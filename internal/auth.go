package internal

import (
	"context"
	"fmt"
	"io"

	"log"
	"os"

	"net/http"

	readerwriters "karti385/uploadThingy/pkg/readerWriters"

	"karti385/uploadThingy/internal/domain"

	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
)

func Auth() {
	
	config,err:=domain.NewConfig()
	if err!=nil {
		log.Fatal(err)
	}

	url := config.AuthCodeURL("x-state")


	open.Run(url)
	server(config)

}

func server(config *oauth2.Config) {
	mux := http.NewServeMux()

	urlchannel := make(chan string)

	mux.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {

		code := r.URL.Query().Get("code")
		fmt.Println(code);
		showThanks(w)
		getToken(code, config)
		urlchannel <- r.UserAgent()
		

	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go showUrl(urlchannel, server)

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error starting server: %v", err)
	}
}

func showUrl(ch chan string, server *http.Server) {
	<-ch

	err := server.Shutdown(context.Background())
	if err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}
}

func getToken(code string, config *oauth2.Config) {
	token, err := config.Exchange(context.Background(), code)
	if !token.Valid() || err != nil {
		log.Fatal("Error in token exchange")
	}

	err=readerwriters.SaveToken(token);
	if err!=nil {
		log.Fatal(err)
	}
	
}
func showThanks(w http.ResponseWriter) {
	file, err := os.OpenFile("index.html", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(content))
}
