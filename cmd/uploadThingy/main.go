package main

import (
	"fmt"
	"karti385/uploadThingy/internal"
	driveOps "karti385/uploadThingy/internal/driveOps"
	tokenUtil "karti385/uploadThingy/pkg/token"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "UploadThingy",
		Usage: "Upload.....",
		Commands: []*cli.Command{
			{
				Name: "login",
				Aliases: []string{
					"l",
				},
				Usage: "Command for login & auth",
				Action: func(*cli.Context) error {
					internal.Auth()
					return nil
				},
			},
			{
				Name: "getToken",
				Aliases: []string{
					"gt",
				},
				Usage: "Command for login & auth",
				Action: func(*cli.Context) error {
					token, err := tokenUtil.GetToken()
					if err != nil {
						panic(err)
					}
					fmt.Printf("Token is %v", token.RefreshToken)
					return nil
				},
			},
			{
				Name: "upload",
				Aliases: []string{
					"ud",
				},
				Usage: "command For uploading files",
				Action: func(*cli.Context) error {
					err := driveOps.Upload()
					if err != nil {
						panic(err)

					}

					return nil
				},
			},
			{
				Name: "listFiles",
				Aliases: []string{
					"lf",
				},
				Usage: "command For Listing Drive files",
				Action: func(*cli.Context) error {
					err := driveOps.ListFiles()
					if err != nil {
						panic(err)

					}

					return nil
				},
			},
			{
				Name: "download",
				Aliases: []string{
					"d",
				},
				Usage: "command For downloading Drive files",
				Action: func(ctx *cli.Context) error {
					fileId := ctx.Args().Get(0)
					if fileId == "" {
						fmt.Println("Please provide correct fleId")
					} else {
						err := driveOps.DownloadFile(fileId)
						if err != nil {
							panic(err)

						}
					}

					return nil
				},
			},
			{
				Name: "refreshToken",
				Aliases: []string{
					"rt",
				},
				Usage: "Refresh the token",
				Action: func(*cli.Context) error {
					token, err := tokenUtil.RefetchToken()
					if err != nil {
						panic(err)
					}
					fmt.Printf("Token is %v", token)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
