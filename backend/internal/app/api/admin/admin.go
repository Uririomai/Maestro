package admin

type Service interface {
	Foo()
}

type Admin struct {
	srv Service
}

func NewAPI(srv Service) *Admin {
	return &Admin{
		srv: srv,
	}
}
