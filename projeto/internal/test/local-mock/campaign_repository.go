package mock

import (
	"emailn/internal/domain/campaign"

	"github.com/stretchr/testify/mock"
)

type CampaignRepositoryMock struct {
	mock.Mock
}

func (r *CampaignRepositoryMock) Create(c *campaign.Campaign) error {
	args := r.Called(c)
	return args.Error(0)
}

func (r *CampaignRepositoryMock) Get() ([]campaign.Campaign, error) {
	args := r.Called()
	return args.Get(0).([]campaign.Campaign), nil
}

func (r *CampaignRepositoryMock) GetStarted() ([]campaign.Campaign, error) {
	args := r.Called()
	return args.Get(0).([]campaign.Campaign), nil
}

func (r *CampaignRepositoryMock) GetBy(id string) (*campaign.Campaign, error) {
	args := r.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*campaign.Campaign), nil
}

func (r *CampaignRepositoryMock) Update(c *campaign.Campaign) error {
	args := r.Called(c)
	return args.Error(0)
}

func (r *CampaignRepositoryMock) Delete(c *campaign.Campaign) error {
	args := r.Called(c)
	return args.Error(0)
}
