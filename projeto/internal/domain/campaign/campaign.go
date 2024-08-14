package campaign

import (
	localerrors "emailn/internal/local-errors"
	"time"

	"github.com/rs/xid"
)

const (
	StatusPending  string = "Pending"
	StatusStarted  string = "Started"
	StatusDone     string = "Done"
	StatusCanceled string = "Canceled"
	StatusDeleted  string = "Deleted"
	StatusFailed   string = "Failed"
)

type Contact struct {
	ID         string `gorm:"size:50"`
	Email      string `validate:"email" gorm:"size:100"`
	CampaignId string `gorm:"size:50"`
}

type Campaign struct {
	ID        string    `validate:"required" gorm:"size:50;not null"`
	Name      string    `validate:"min=5,max=24" gorm:"size:100;not null"`
	CreatedOn time.Time `validate:"required" gorm:"not null"`
	UpdatedOn time.Time
	CreatedBy string    `validate:"email" gorm:"size:50;not null"`
	Content   string    `validate:"min=5,max=1024" gorm:"size:1024;not null"`
	Contacts  []Contact `validate:"min=1,dive" gorm:"not null"`
	Status    string    `gorm:"size:50;not null"`
}

func NewCampaign(name string, content string, emails []string, createdBy string) (*Campaign, error) {
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
		Status:    StatusPending,
		CreatedBy: createdBy,
	}

	err := localerrors.ValidateStruct(campaign)
	if err != nil {
		return nil, err
	}

	return campaign, nil
}

func (c *Campaign) Cancel() {
	c.Status = StatusCanceled
	c.UpdatedOn = time.Now()
}

func (c *Campaign) Delete() {
	c.Status = StatusDeleted
	c.UpdatedOn = time.Now()
}

func (c *Campaign) Start() {
	c.Status = StatusStarted
	c.UpdatedOn = time.Now()
}

func (c *Campaign) Done() {
	c.Status = StatusDone
	c.UpdatedOn = time.Now()
}

func (c *Campaign) Fail() {
	c.Status = StatusFailed
	c.UpdatedOn = time.Now()
}
