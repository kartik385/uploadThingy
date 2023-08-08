package domain

import (
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v2"
)

func NewConfig() (*oauth2.Config,error){
	err := godotenv.Load()
	if err != nil {
		return nil,err
	}

	gcid := os.Getenv("GCID")
	gcs := os.Getenv("GCS")

	config := oauth2.Config{
		ClientID:     gcid,
		ClientSecret: gcs,
		Endpoint:     google.Endpoint,
		Scopes: []string{
			drive.DriveScope,
		},
		RedirectURL: "http://localhost:8080/redirect",
	}

	return &config,nil
}