package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name    = "Campaign X"
	content = "Body body body body"
	emails  = []string{
		"email1@email.com",
		"email2@email.com",
		"email3@email.com",
		"email4@email.com",
		"email5@email.com",
		"email6@email.com",
		"email7@email.com",
		"email8@email.com",
	}
	fake = faker.New()
)

func Test_NewCampaign(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	createdOn := time.Now()
	// Act
	c, _ := NewCampaign(name, content, emails)
	// Assert
	assert.NotNil(c.ID)
	assert.WithinDuration(c.CreatedOn, createdOn, time.Minute)
	assert.Equal(c.Name, name)
	assert.Equal(c.Content, content)
	assert.Len(c.Contacts, len(emails))
}

func Test_NewCampaign_MustValidateNameMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, emails)

	assert.EqualError(err, "name requires a minimum of 5")
}

func Test_NewCampaign_MustValidateNameMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(fake.Lorem().Text(40), content, emails)

	assert.EqualError(err, "name requires a maximum of 24")
}

func Test_NewCampaign_MustValidateContentMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", emails)

	assert.EqualError(err, "content requires a minimum of 5")
}

func Test_NewCampaign_MustValidateContentMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, fake.Lorem().Text(1100), emails)

	assert.EqualError(err, "content requires a maximum of 1024")
}

func Test_NewCampaign_MustValidateContactsMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assert.EqualError(err, "contacts requires a minimum of 1")
}

func Test_NewCampaign_MustValidateContactsEmailPattern(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{
		"invalid_email",
	})

	assert.EqualError(err, "email is not a valid email")

}
