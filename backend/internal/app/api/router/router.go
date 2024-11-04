package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/Nikita-Kolbin/Maestro/docs"
	"github.com/Nikita-Kolbin/Maestro/internal/app/api/admin"
)

type service interface {
	admin.Service
}

func New(ctx context.Context, srv service, address string) http.Handler {
	router := chi.NewRouter()

	// middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	//authMiddleware := authMW.Auth

	// CORS
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*", "https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
	}))

	// swagger
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", address)),
	))

	// APIs
	adminAPI := admin.NewAPI(srv)

	// handlers
	router.Post("/api/admin/sign-up", adminAPI.SignUp)
	router.Post("/api/admin/sign-in", adminAPI.SignIn)

	return router
}
