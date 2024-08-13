package campaign

import (
	localerrors "emailn/internal/local-errors"
	"time"

	"github.com/rs/xid"
)

const (
	Pending  string = "Pending"
	Started  string = "Started"
	Done     string = "Done"
	Canceled string = "Canceled"
)

type Contact struct {
	ID         string `gorm:"size:50"`
	Email      string `validate:"email" gorm:"size:100"`
	CampaignId string `gorm:"size:50"`
}

type Campaign struct {
	ID        string    `validate:"required" gorm:"size:50"`
	Name      string    `validate:"min=5,max=24" gorm:"size:100"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024" gorm:"size:1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string    `gorm:"size:50"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))
	for i, v := range emails {
		contacts[i].ID = xid.New().String()
		contacts[i].Email = v
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
		Status:    Pending,
	}

	err := localerrors.ValidateStruct(campaign)
	if err != nil {
		return nil, err
	}

	return campaign, nil
}

func (c *Campaign) Cancel() {
	c.Status = Canceled
}
