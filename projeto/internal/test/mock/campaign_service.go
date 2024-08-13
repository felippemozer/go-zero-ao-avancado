package mock

import (
	"emailn/internal/contract"

	"github.com/stretchr/testify/mock"
)

type CampaignServiceMock struct {
	mock.Mock
}

func (s *CampaignServiceMock) Create(newCampaign contract.NewCampaign) (string, error) {
	args := s.Called(newCampaign)
	return args.String(0), args.Error(1)
}

func (s *CampaignServiceMock) GetBy(campaignID string) (*contract.GetCampaignByIdResponse, error) {
	args := s.Called(campaignID)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*contract.GetCampaignByIdResponse), nil
}
