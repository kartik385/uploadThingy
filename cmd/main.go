package main

import (
	"fmt"
	"karti385/uploadThingy/internal"
	readerwriters "karti385/uploadThingy/pkg/readerWriters"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "UploadThingy",
		Usage: "Upload.....",
		Commands:[]*cli.Command{
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
					token,err:=readerwriters.GetToken()
					if err!=nil {
						panic(err)
					}
					fmt.Printf("Token is %v",token.AccessToken);
					return nil
				},
			},
		},
		
		
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
