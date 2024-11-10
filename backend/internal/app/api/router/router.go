package router

import (
	"context"
	"fmt"
	"github.com/Nikita-Kolbin/Maestro/internal/app/api/customer"
	"github.com/Nikita-Kolbin/Maestro/internal/app/api/product"
	"github.com/Nikita-Kolbin/Maestro/internal/app/api/website"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/Nikita-Kolbin/Maestro/docs"
	"github.com/Nikita-Kolbin/Maestro/internal/app/api/admin"
	authMW "github.com/Nikita-Kolbin/Maestro/internal/pkg/middleware"
)

type service interface {
	admin.Service
	website.Service
	customer.Service
	product.Service
}

func New(ctx context.Context, srv service, address string) http.Handler {
	router := chi.NewRouter()

	// middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	authMiddleware := authMW.Auth(srv.GetJWTSecret())

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
	websiteAPI := website.NewAPI(srv)
	customerAPI := customer.NewAPI(srv)
	productAPI := product.NewAPI(srv)

	// handlers
	router.Post("/api/admin/sign-up", adminAPI.AdminSignUp)
	router.Post("/api/admin/sign-in", adminAPI.AdminSignIn)

	router.Post("/api/website/create", authMiddleware(websiteAPI.CreateWebsite))
	router.Get("/api/website/get-my-website", authMiddleware(websiteAPI.GetMyWebsite))

	router.Post("/api/customer/sign-up", customerAPI.CustomerSignUp)
	router.Post("/api/customer/sign-in", customerAPI.CustomerSignIn)

	router.Post("/api/product/create", authMiddleware(productAPI.CreateProduct))
	router.Get("/api/product/get-active-by-alias", productAPI.GetActiveProductByAlias)

	return router
}
