package model

type AdminEmailPasswordRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminTokenResponse struct {
	Token string `json:"token"`
}

type Admin struct {
	Id                   int    `db:"id"`
	Email                string `db:"email"`
	FirstName            string `db:"first_name"`
	LastName             string `db:"last_name"`
	FatherName           string `db:"father_name"`
	City                 string `db:"city"`
	Telegram             string `db:"telegram"`
	ImageId              int    `db:"image_id"`
	EmailNotification    bool   `db:"email_notification"`
	TelegramNotification bool   `db:"telegram_notification"`
}

type AdminDTO struct {
	Id                   int    `json:"id"`
	Email                string `json:"email"`
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	FatherName           string `json:"father_name"`
	City                 string `json:"city"`
	Telegram             string `json:"telegram"`
	ImageId              int    `json:"image_id"`
	TelegramNotification int    `json:"telegram_notification"`
	EmailNotification    int    `json:"email_notification"`
}
