package model

type CreateWebsiteRequest struct {
	Alias string `json:"alias"`
}

type Website struct {
	Id      int    `db:"id"`
	AdminId int    `db:"admin_id"`
	Alias   string `db:"alias"`
}

type WebsiteDTO struct {
	Id    int    `json:"id"`
	Alias string `json:"alias"`
}

func FromWebsiteToDTO(website *Website) *WebsiteDTO {
	return &WebsiteDTO{
		Id:    website.Id,
		Alias: website.Alias,
	}
}
