package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/infrastructure/database"
	"emailn/internal/infrastructure/mail"
	"fmt"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../../.env")

	db := database.NewDb()
	repository := database.CampaignRepository{
		Db: db,
	}
	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{
			Db: db,
		},
		SendMail: mail.SendMail,
	}

	for {
		campaigns, err := repository.GetStarted()

		if err != nil {
			fmt.Println("error on send email")
			return
		}

		for _, campaign := range campaigns {
			campaignService.SendEmailAndUpdateStatus(&campaign)
		}

		time.Sleep(time.Hour * 4)
	}

}
