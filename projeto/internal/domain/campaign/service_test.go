package campaign_test

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	localerrors "emailn/internal/local-errors"
	localmock "emailn/internal/test/local-mock"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	service     = campaign.ServiceImp{}
	newCampaign = contract.NewCampaign{
		Name:    "Test Y",
		Content: "Body body body body",
		Emails: []string{
			"email1@email.com",
			"email2@email.com",
			"email3@email.com",
			"email4@email.com",
			"email5@email.com",
			"email6@email.com",
			"email7@email.com",
			"email8@email.com",
		},
	}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(localmock.CampaignRepositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)
	service.Repository = repositoryMock

	id, err := service.Create(newCampaign)

	assert.NoError(err)
	assert.NotEmpty(id)
}

func Test_Create_SaveCampaign(t *testing.T) {
	repositoryMock := new(localmock.CampaignRepositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		if campaign.Name != newCampaign.Name || campaign.Content != newCampaign.Content || len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}
		return true
	})).Return(nil)
	service.Repository = repositoryMock

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	newCampaignCopy := contract.NewCampaign{
		Name:    "",
		Content: newCampaign.Content,
		Emails:  newCampaign.Emails,
	}

	_, err := service.Create(newCampaignCopy)

	assert.EqualError(err, "name requires a minimum of 5")
}

func Test_Create_ValidateRepositoryError(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(localmock.CampaignRepositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(localerrors.ErrInternal)
	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	assert.EqualError(err, localerrors.ErrInternal.Error())
}

func Test_GetById_Success(t *testing.T) {
	assert := assert.New(t)
	campaign := campaign.Campaign{
		ID:      "1",
		Name:    newCampaign.Name,
		Content: newCampaign.Content,
		Status:  campaign.Pending,
	}
	repositoryMock := new(localmock.CampaignRepositoryMock)
	repositoryMock.On("GetBy", mock.MatchedBy(func(id string) bool {
		return id == campaign.ID
	})).Return(&campaign, nil)
	service.Repository = repositoryMock

	resp, err := service.GetBy(campaign.ID)

	assert.Nil(err)
	assert.Equal(resp.ID, campaign.ID)
	assert.Equal(resp.Name, campaign.Name)
	assert.Equal(resp.Content, campaign.Content)
	assert.Equal(resp.Status, campaign.Status)
}

func Test_GetById_Error(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(localmock.CampaignRepositoryMock)
	repositoryMock.On("GetBy", mock.Anything).Return(nil, errors.New("error"))
	service.Repository = repositoryMock

	_, err := service.GetBy("1")

	assert.Error(err)
}
