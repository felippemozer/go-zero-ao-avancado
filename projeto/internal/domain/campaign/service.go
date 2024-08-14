package campaign

import (
	"emailn/internal/contract"
	localerrors "emailn/internal/local-errors"
	"errors"
)

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	GetBy(campaignID string) (*contract.GetCampaignByIdResponse, error)
	Cancel(campaignID string) error
	Delete(campaignID string) error
	Start(campaignID string) error
}

type ServiceImp struct {
	Repository Repository
	SendMail   func(campaign *Campaign) error
}

func (s *ServiceImp) Create(newCampaign contract.NewCampaign) (string, error) {
	c, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

	if err != nil {
		return "", err
	}

	err = s.Repository.Create(c)

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

	if campaign == nil {
		return nil, localerrors.ErrNotFound
	}

	return &contract.GetCampaignByIdResponse{
		ID:                   campaign.ID,
		Name:                 campaign.Name,
		Content:              campaign.Content,
		Status:               campaign.Status,
		CreatedBy:            campaign.CreatedBy,
		AmountOfEmailsToSend: len(campaign.Contacts),
	}, nil
}

func (s *ServiceImp) Cancel(campaignID string) error {
	campaign, err := s.Repository.GetBy(campaignID)

	if err != nil {
		return localerrors.ErrInternal
	}

	if campaign.Status == StatusCanceled {
		return nil
	}

	if campaign.Status != StatusPending {
		return errors.New("Campaign status invalid to cancel")
	}

	campaign.Cancel()
	err = s.Repository.Update(campaign)
	if err != nil {
		return localerrors.ErrInternal
	}

	return nil
}

func (s *ServiceImp) Delete(campaignID string) error {
	campaign, err := s.Repository.GetBy(campaignID)

	if err != nil {
		return localerrors.ErrInternal
	}

	if campaign.Status == StatusStarted {
		return errors.New("Campaign status has started and has not finished")
	}

	campaign.Delete()
	err = s.Repository.Delete(campaign)
	if err != nil {
		return localerrors.ErrInternal
	}

	return nil
}

func (s *ServiceImp) Start(campaignID string) error {
	campaign, err := s.Repository.GetBy(campaignID)

	if err != nil {
		return localerrors.ErrInternal
	}

	if campaign.Status != StatusPending {
		return errors.New("Campaign is not pending start")
	}

	go func() {
		err = s.SendMail(campaign)
		if err != nil {
			campaign.Fail()
		} else {
			campaign.Done()
		}
		s.Repository.Update(campaign)
	}()

	campaign.Start()
	err = s.Repository.Update(campaign)
	if err != nil {
		return localerrors.ErrInternal
	}

	return nil
}
