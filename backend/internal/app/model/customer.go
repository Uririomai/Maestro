package model

type CustomerAliasEmailPasswordRequest struct {
	Alias    string `json:"alias"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CustomerTokenResponse struct {
	Token string `json:"token"`
}

type Customer struct {
	Id                   int    `db:"id"`
	WebsiteAlias         string `db:"website_alias"`
	Email                string `db:"email"`
	FirstName            string `db:"first_name"`
	LastName             string `db:"last_name"`
	FatherName           string `db:"father_name"`
	Phone                string `db:"phone"`
	Telegram             string `db:"telegram"`
	DeliveryType         string `db:"delivery_type"`
	PaymentType          string `db:"payment_type"`
	TelegramNotification bool   `db:"telegram_notification"`
	EmailNotification    bool   `db:"email_notification"`
}

type CustomerDTO struct {
	Id                   int    `json:"id"`
	WebsiteAlias         string `json:"website_alias"`
	Email                string `json:"email"`
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	FatherName           string `json:"father_name"`
	Phone                string `json:"phone"`
	Telegram             string `json:"telegram"`
	DeliveryType         string `json:"delivery_type"`
	PaymentType          string `json:"payment_type"`
	TelegramNotification bool   `json:"telegram_notification"`
	EmailNotification    bool   `json:"email_notification"`
}
