package campaign

import (
	"emailn/internal/contract"
	localerrors "emailn/internal/local-errors"
)

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	GetBy(campaignID string) (*contract.GetCampaignByIdResponse, error)
}

type ServiceImp struct {
	Repository Repository
}

func (s *ServiceImp) Create(newCampaign contract.NewCampaign) (string, error) {
	c, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	if err != nil {
		return "", err
	}

	err = s.Repository.Save(c)

	if err != nil {
		return "", localerrors.ErrInternal
	}

	return c.ID, nil
}

func (s *ServiceImp) GetBy(campaignID string) (*contract.GetCampaignByIdResponse, error) {
	campaign, err := s.Repository.GetBy(campaignID)

	if err != nil {
		return nil, localerrors.ErrInternal
	}

	return &contract.GetCampaignByIdResponse{
		ID:      campaign.ID,
		Name:    campaign.Name,
		Content: campaign.Content,
		Status:  campaign.Status,
	}, nil
}
