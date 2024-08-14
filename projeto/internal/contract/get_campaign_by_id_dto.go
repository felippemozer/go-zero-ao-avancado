package contract

type GetCampaignByIdResponse struct {
	ID                   string
	Name                 string
	Content              string
	Status               string
	CreatedBy            string
	AmountOfEmailsToSend int
}
