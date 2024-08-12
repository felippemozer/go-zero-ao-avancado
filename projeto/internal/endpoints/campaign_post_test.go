package endpoints

import (
	"bytes"
	"emailn/internal/contract"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type serviceMock struct {
	mock.Mock
}

func (s *serviceMock) Create(newCampaign contract.NewCampaign) (string, error) {
	args := s.Called(newCampaign)
	return args.String(0), args.Error(1)
}

func Test_CampaignPost_SaveNewCampaign(t *testing.T) {
	assert := assert.New(t)
	campaign := contract.NewCampaign{
		Name:    "Teste",
		Content: "Conteúdo",
		Emails: []string{
			"teste@teste.com",
		},
	}
	service := new(serviceMock)
	service.On("Create", mock.MatchedBy(func(c contract.NewCampaign) bool {
		if c.Name == campaign.Name && c.Content == campaign.Content {
			return true
		}
		return false
	})).Return("123", nil)
	handler := Handler{
		CampaignService: service,
	}
	var body bytes.Buffer
	json.NewEncoder(&body).Encode(campaign)
	req, _ := http.NewRequest("POST", "/", &body)
	res := httptest.NewRecorder()

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
	}
	service := new(serviceMock)
	service.On("Create", mock.Anything).Return("", errors.New("error"))
	handler := Handler{
		CampaignService: service,
	}
	var body bytes.Buffer
	json.NewEncoder(&body).Encode(campaign)
	req, _ := http.NewRequest("POST", "/", &body)
	res := httptest.NewRecorder()

	_, _, err := handler.CampaignPost(res, req)

	assert.Error(err)
}
