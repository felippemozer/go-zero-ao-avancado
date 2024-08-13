package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/endpoints"
	"emailn/internal/infrastructure/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Product struct {
	ID   string
	Name string
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := database.NewDb()
	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{
			Db: db,
		},
	}
	handler := endpoints.Handler{
		CampaignService: &campaignService,
	}
	r.Post("/campaign", endpoints.HandlerError(handler.CampaignPost))
	r.Get("/campaign/{id}", endpoints.HandlerError(handler.CampaignGetByID))
	r.Patch("/campaign/cancel/{id}", endpoints.HandlerError(handler.CampaignCancelPatch))
	r.Delete("/campaign/{id}", endpoints.HandlerError(handler.CampaignDelete))

	http.ListenAndServe(":3000", r)
}
