package endpoints

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignGetByID(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	campaign, err := h.CampaignService.GetBy(id)

	if campaign == nil {
		return nil, http.StatusNotFound, err
	}

	return campaign, http.StatusOK, err
}
