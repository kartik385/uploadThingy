package token

import (
	"context"
	"encoding/gob"
	"karti385/uploadThingy/internal/domain"
	"os"

	"golang.org/x/oauth2"
)

func SaveToken(token *oauth2.Token)(error) {
	file,err:=os.OpenFile("token.gob",os.O_CREATE | os.O_APPEND  | os.O_TRUNC ,0644);
	if err!=nil {
		return err
		
	}
	

	defer file.Close()

	// err = file.Truncate(0)
    // if err != nil {
    //     return err
    // }

	encoder := gob.NewEncoder(file);
	err=encoder.Encode(token);
	if err!=nil {
		return err
	}
	return nil

}

func GetToken()(*oauth2.Token,error) {
	file,err:=os.OpenFile("token.gob",os.O_RDONLY,0644);
	if err!=nil {
		return nil, err
		
	}

	defer file.Close()

	
	var token oauth2.Token
	decoder:=gob.NewDecoder(file);
	decoder.Decode(&token);
	if err!=nil {
		return nil,err
	}
	return &token,nil
}

func RefetchToken()(*oauth2.Token,error) {
	token,err:=GetToken();
	if err!= nil {
		return nil,err;
		
	}
	config,err:=domain.NewConfig()
	if err!= nil {
		return nil,err;
		
	}
	tokenSource:=config.TokenSource(context.Background(),token);
	freshToken,err:=tokenSource.Token()
	if err!=nil {
		return nil,err;
		
	}
	return freshToken,nil;
}