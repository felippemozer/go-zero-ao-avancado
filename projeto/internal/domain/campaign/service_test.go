package campaign

import (
	"emailn/internal/contract"
	localerrors "emailn/internal/local-errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	service     = ServiceImp{}
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

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(c *Campaign) error {
	args := r.Called(c)
	return args.Error(0)
}

func (r *repositoryMock) Get() ([]Campaign, error) {
	// args := r.Called(c)
	return nil, nil
}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)
	service.Repository = repositoryMock

	id, err := service.Create(newCampaign)

	assert.NoError(err)
	assert.NotEmpty(id)
}

func Test_Create_SaveCampaign(t *testing.T) {
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name || campaign.Content != newCampaign.Content || len(campaign.Contacts) != len(emails) {
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
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(localerrors.ErrInternal)
	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	assert.EqualError(err, localerrors.ErrInternal.Error())
}
