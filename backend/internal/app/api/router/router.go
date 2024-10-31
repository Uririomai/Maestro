package router

import (
	"net/http"
)

type service interface {
}

func New(srv service) http.Handler {
	return nil
}
