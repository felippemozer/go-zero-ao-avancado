package endpoints

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	localmock "emailn/internal/test/mock"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignGetByID_Success(t *testing.T) {
	assert := assert.New(t)
	campaign := contract.GetCampaignByIdResponse{
		ID:      "1",
		Name:    "Teste",
		Content: "Teste",
		Status:  campaign.Pending,
	}
	service := new(localmock.CampaignServiceMock)
	service.On("GetBy", mock.Anything).Return(&campaign, nil)
	handler := Handler{
		CampaignService: service,
	}
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	response, status, _ := handler.CampaignGetByID(res, req)

	assert.Equal(http.StatusOK, status)
	assert.Equal(campaign.ID, response.(*contract.GetCampaignByIdResponse).ID)
}

func Test_CampaignGetByID_Error(t *testing.T) {
	assert := assert.New(t)
	service := new(localmock.CampaignServiceMock)
	service.On("GetBy", mock.Anything).Return(nil, errors.New("error"))
	handler := Handler{
		CampaignService: service,
	}
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	_, _, err := handler.CampaignGetByID(res, req)

	assert.Error(err)
}
