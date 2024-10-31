package customer

type Service interface {
	Bar()
}

type Admin struct {
	srv Service
}

func NewAPI(srv Service) *Admin {
	return &Admin{
		srv: srv,
	}
}
