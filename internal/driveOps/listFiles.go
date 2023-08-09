package driveops

import (
	"fmt"
	"karti385/uploadThingy/pkg"

	"google.golang.org/api/drive/v3"
)

func ListFiles() error {
	driveService, err := pkg.NewDriveService()
	if err != nil {
		return fmt.Errorf("failed to start Drive Service: %v", err)
	}

	fileService := drive.NewFilesService(driveService)

	fileList,err:=fileService.List().OrderBy("name").Do()

	if err!=nil {
		return fmt.Errorf("failed to List Files from Drive: %v", err)
	}

	

	for _,v := range fileList.Files {
		fmt.Println(v.Name);
	}

	if fileList.NextPageToken != "" {
		nxtToken := fileList.NextPageToken
		for {
			fileList,err = fileService.List().OrderBy("name").PageToken(nxtToken).Do();
			if err!=nil {
				return fmt.Errorf("failed to List Files from Drive: %v", err)
			}
			for _,v := range fileList.Files {
				fmt.Println(v.Name);
			}
			if fileList.NextPageToken == "" {
				break
			} else {
				nxtToken = fileList.NextPageToken
			}
			

		}
	}

	return nil
}
