package model

import "github.com/google/uuid"

type CreateWebsiteRequest struct {
	Alias string `json:"alias"`
}

type Website struct {
	Id         int    `db:"id"`
	AdminId    int    `db:"admin_id"`
	Alias      string `db:"alias"`
	Active     bool   `db:"active"`
	SectionIds []int  `db:"section_ids"`
}

type WebsiteDTO struct {
	Id     int    `json:"id"`
	Alias  string `json:"alias"`
	Active bool   `json:"active"`
}

type Section struct {
	Id           int    `db:"id"`
	UUID         string `db:"uuid"`
	WebsiteAlias string `db:"website_alias"`
	Width        int    `db:"width"`
	FullWidth    bool   `db:"full_width"`
	Height       int    `db:"height"`
	FullHeight   bool   `db:"full_height"`
	Blocks       []*Block
}

type Block struct {
	Id           int    `db:"id"`
	WebsiteAlias string `db:"website_alias"`
	SectionUUID  string `db:"section_uuid"`
	Text         string `db:"text"`
}

type WebsiteStylesDTO struct {
	Sections []*SectionDTO `json:"sections"`
}

type SectionDTO struct {
	Id           int         `json:"id"`
	UUID         string      `json:"uuid"`
	WebsiteAlias string      `json:"website_alias"`
	Width        int         `json:"width"`
	FullWidth    bool        `json:"full_width"`
	Height       int         `json:"height"`
	FullHeight   bool        `json:"full_height"`
	Blocks       []*BlockDTO `json:"blocks"`
}

type BlockDTO struct {
	Id           int    `json:"id"`
	WebsiteAlias string `json:"website_alias"`
	SectionUUID  string `json:"section_uuid"`
	Text         string `json:"text"`
}

type SetWebsiteStyleRequest struct {
	WebsiteAlias string               `json:"website_alias"`
	Sections     []*SetSectionRequest `json:"sections"`
}

type SetSectionRequest struct {
	Width      int                `json:"width"`
	FullWidth  bool               `json:"full_width"`
	Height     int                `json:"height"`
	FullHeight bool               `json:"full_height"`
	Blocks     []*SetBlockRequest `json:"blocks"`
}

type SetBlockRequest struct {
	Text string `json:"text"`
}

func FromWebsiteToDTO(website *Website) *WebsiteDTO {
	return &WebsiteDTO{
		Id:     website.Id,
		Alias:  website.Alias,
		Active: website.Active,
	}
}

func FromSetWebsiteStyleRequestToSections(req *SetWebsiteStyleRequest) []*Section {
	sections := make([]*Section, 0, len(req.Sections))

	for _, s := range req.Sections {
		id := uuid.New().String()

		section := &Section{
			UUID:         id,
			WebsiteAlias: req.WebsiteAlias,
			Width:        s.Width,
			FullWidth:    s.FullWidth,
			Height:       s.Height,
			FullHeight:   s.FullHeight,
			Blocks:       make([]*Block, 0, len(s.Blocks)),
		}

		for _, b := range s.Blocks {
			block := &Block{
				SectionUUID:  id,
				WebsiteAlias: req.WebsiteAlias,
				Text:         b.Text,
			}
			section.Blocks = append(section.Blocks, block)
		}

		sections = append(sections, section)
	}

	return sections
}

func FromSectionsToDTO(sections []*Section) *WebsiteStylesDTO {
	sectionsDTO := make([]*SectionDTO, 0, len(sections))
	for _, s := range sections {
		sectionDTO := &SectionDTO{
			Id:           s.Id,
			UUID:         s.UUID,
			WebsiteAlias: s.WebsiteAlias,
			Width:        s.Width,
			FullWidth:    s.FullWidth,
			Height:       s.Height,
			FullHeight:   s.FullHeight,
			Blocks:       make([]*BlockDTO, 0, len(s.Blocks)),
		}

		for _, b := range s.Blocks {
			blockDTO := &BlockDTO{
				Id:           b.Id,
				WebsiteAlias: b.WebsiteAlias,
				SectionUUID:  b.SectionUUID,
				Text:         b.Text,
			}
			sectionDTO.Blocks = append(sectionDTO.Blocks, blockDTO)
		}

		sectionsDTO = append(sectionsDTO, sectionDTO)
	}

	return &WebsiteStylesDTO{Sections: sectionsDTO}
}
