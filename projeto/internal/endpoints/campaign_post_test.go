package endpoints

import (
	"bytes"
	"context"
	"emailn/internal/contract"
	localmock "emailn/internal/test/local-mock"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setup(campaign contract.NewCampaign) (*http.Request, *httptest.ResponseRecorder) {
	var body bytes.Buffer
	json.NewEncoder(&body).Encode(campaign)
	req, _ := http.NewRequest("POST", "/", &body)
	ctx := context.WithValue(req.Context(), "email", campaign.CreatedBy)
	req = req.WithContext(ctx)
	res := httptest.NewRecorder()

	return req, res
}

func Test_CampaignPost_SaveNewCampaign(t *testing.T) {
	assert := assert.New(t)
	campaign := contract.NewCampaign{
		Name:    "Teste",
		Content: "Conteúdo",
		Emails: []string{
			"teste@teste.com",
		},
		CreatedBy: "teste@teste.com.br",
	}
	service := new(localmock.CampaignServiceMock)
	service.On("Create", mock.MatchedBy(func(c contract.NewCampaign) bool {
		if c.Name == campaign.Name && c.Content == campaign.Content && c.CreatedBy == campaign.CreatedBy {
			return true
		}
		return false
	})).Return("123", nil)
	handler := Handler{
		CampaignService: service,
	}
	req, res := setup(campaign)

	_, status, err := handler.CampaignPost(res, req)

	assert.Equal(http.StatusCreated, status)
	assert.Nil(err)
}

func Test_CampaignPost_Error(t *testing.T) {
	assert := assert.New(t)
	campaign := contract.NewCampaign{
		Name:    "Teste",
		Content: "Conteúdo",
		Emails: []string{
			"teste@teste.com",
		},
		CreatedBy: "teste@teste.com",
	}
	service := new(localmock.CampaignServiceMock)
	service.On("Create", mock.Anything).Return("", errors.New("error"))
	handler := Handler{
		CampaignService: service,
	}
	req, res := setup(campaign)

	_, _, err := handler.CampaignPost(res, req)

	assert.Error(err)
}
