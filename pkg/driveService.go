package pkg

import (
	"context"
	"karti385/uploadThingy/internal/domain"
	tokenM "karti385/uploadThingy/pkg/token"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func NewDriveService() (*drive.Service, error) {
	config, err := domain.NewConfig()
    if err != nil {
        return nil,err
    }

    token, err := tokenM.GetToken()
    if err != nil {
        return nil,err
	}
    driveService, err := drive.NewService(context.Background(), option.WithTokenSource(config.TokenSource(context.Background(), token)))
    if err != nil {
        return nil,err
    }
	return driveService,nil
}