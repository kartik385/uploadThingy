package internal

import (
	"context"
	"fmt"
	"io"

	"log"
	"os"

	"net/http"

	tokenUtil "karti385/uploadThingy/pkg/token"

	"karti385/uploadThingy/internal/domain"

	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
)

func Auth() {

	
	_,err:=tokenUtil.GetToken();

	if(err!=nil) {
		config,err:=domain.NewConfig()
	if err!=nil {
		log.Fatal(err)
	}

	url := config.AuthCodeURL("state",oauth2.AccessTypeOffline)
	open.Run(url)
	server(config)
	} else {
		fmt.Println("Already Logged In")
	}


	
	

}

func server(config *oauth2.Config) {
	mux := http.NewServeMux()

	urlchannel := make(chan string)

	mux.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
	
		showThanks(w)
		getToken(code, config,r)
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

func getToken(code string, config *oauth2.Config,r *http.Request) {
	token, err := config.Exchange(r.Context(), code)
	if !token.Valid() || err != nil {
		log.Fatal("Error in token exchange")
	}
	fmt.Printf("Token is %v",token);

	err=tokenUtil.SaveToken(token);
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
