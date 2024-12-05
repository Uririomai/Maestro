package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
)

func (r *Repository) CreateWebsite(ctx context.Context, alias string, adminId int) (*model.Website, error) {
	query := `
	INSERT INTO websites (alias, admin_id) VALUES ($1, $2)
	RETURNING id, alias, admin_id, active`

	website := &model.Website{}
	err := r.conn.GetContext(ctx, website, query, alias, adminId)
	if isSQLError(err, model.UniqueConstraintViolationCode) {
		return nil, model.ErrAliasTaken
	}
	if err != nil {
		return nil, err
	}

	return website, nil
}

func (r *Repository) AdminHaveWebsite(ctx context.Context, adminId int) (bool, error) {
	query := `SELECT COUNT(*) FROM websites WHERE admin_id=$1`

	var count int

	err := r.conn.GetContext(ctx, &count, query, adminId)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *Repository) GetWebsiteByAlias(ctx context.Context, alias string) (*model.Website, error) {
	query := `SELECT id, admin_id, alias, active FROM websites WHERE alias = $1`

	website := &model.Website{}
	err := r.conn.GetContext(ctx, website, query, alias)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, model.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return website, nil
}

func (r *Repository) GetWebsiteByAdminId(ctx context.Context, adminId int) (*model.Website, error) {
	query := `SELECT id, admin_id, alias, active FROM websites WHERE admin_id = $1 LIMIT 1`

	website := &model.Website{}
	err := r.conn.GetContext(ctx, website, query, adminId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, model.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return website, nil
}

func (r *Repository) GetSectionsByWebsiteAlias(ctx context.Context, websiteAlias string) ([]*model.Section, error) {
	querySelectSections := `
	SELECT id, uuid, website_alias, width, full_width, height, full_height
	FROM sections WHERE website_alias = $1 ORDER BY id`
	querySelectBlocks := `
	SELECT id, section_uuid, website_alias, text
	FROM blocks WHERE website_alias = $1 ORDER BY id`

	var sections []*model.Section
	err := r.conn.SelectContext(ctx, &sections, querySelectSections, websiteAlias)
	if err != nil {
		return nil, err
	}

	rows, err := r.conn.QueryContext(ctx, querySelectBlocks, websiteAlias)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	// Мапа списков блоков для секций
	blocks := make(map[string][]*model.Block)

	for rows.Next() {
		block := &model.Block{}
		if err = rows.Scan(
			&block.Id,
			&block.SectionUUID,
			&block.WebsiteAlias,
			&block.Text,
		); err != nil {
			return nil, err
		}

		if _, ok := blocks[block.SectionUUID]; !ok {
			blocks[block.SectionUUID] = make([]*model.Block, 0)
		}

		blocks[block.SectionUUID] = append(blocks[block.SectionUUID], block)
	}

	for i := range sections {
		if _, ok := blocks[sections[i].UUID]; !ok {
			sections[i].Blocks = make([]*model.Block, 0)
			continue
		}
		sections[i].Blocks = blocks[sections[i].UUID]
	}

	return sections, nil
}

func (r *Repository) CreateSections(ctx context.Context, websiteAlias string, sections []*model.Section) error {
	queryDeleteOldBlocks := `DELETE FROM blocks WHERE website_alias = $1`
	queryDeleteOldSections := `DELETE FROM sections WHERE website_alias = $1`

	sqCreateSections := squirrel.
		Insert("sections").
		Columns("uuid, website_alias, width, full_width, height, full_height")
	for _, s := range sections {
		sqCreateSections = sqCreateSections.
			Values(s.UUID, s.WebsiteAlias, s.Width, s.FullWidth, s.Height, s.FullHeight)
	}
	queryCreateSections, argsCreateSections, err := sqCreateSections.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}

	sqCreateBlocks := squirrel.
		Insert("blocks").
		Columns("section_uuid, website_alias, text")
	for _, s := range sections {
		for _, b := range s.Blocks {
			sqCreateBlocks = sqCreateBlocks.
				Values(b.SectionUUID, b.WebsiteAlias, b.Text)
		}
	}
	queryCreateBlocks, argsCreateBlocks, err := sqCreateBlocks.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}

	tx, err := r.conn.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	_, err = tx.ExecContext(ctx, queryDeleteOldBlocks, websiteAlias)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, queryDeleteOldSections, websiteAlias)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, queryCreateSections, argsCreateSections...)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, queryCreateBlocks, argsCreateBlocks...)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
