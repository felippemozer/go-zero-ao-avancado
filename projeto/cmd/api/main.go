package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/endpoints"
	"emailn/internal/infrastructure/database"
	"emailn/internal/infrastructure/mail"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

type Product struct {
	ID   string
	Name string
}

func main() {
	godotenv.Load("../../.env")
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
		SendMail: mail.SendMail,
	}
	handler := endpoints.Handler{
		CampaignService: &campaignService,
	}
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	r.Route("/campaign", func(r chi.Router) {
		r.Use(endpoints.Auth)
		r.Post("/", endpoints.HandlerError(handler.CampaignPost))
		r.Get("/{id}", endpoints.HandlerError(handler.CampaignGetByID))
		r.Patch("/cancel/{id}", endpoints.HandlerError(handler.CampaignCancelPatch))
		r.Patch("/start/{id}", endpoints.HandlerError(handler.CampaignStart))
		r.Delete("/{id}", endpoints.HandlerError(handler.CampaignDelete))

	})

	http.ListenAndServe(":3000", r)
}
