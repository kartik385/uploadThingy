package pkg

import (
	"os"
)

func GetDirectoryFiles()([]string,error) {
	files,err:= os.ReadDir(".");
	var fileNames []string
	if err!=nil {
		return fileNames,err
	}

	

	
	for _,v := range files {
		fileNames = append(fileNames, v.Name())
	}
	
	return fileNames,nil

}