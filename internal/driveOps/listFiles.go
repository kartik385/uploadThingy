package driveops

import (
	"fmt"
	"karti385/uploadThingy/pkg"
	"os"

	"github.com/olekukonko/tablewriter"
	"google.golang.org/api/drive/v3"
)

func ListFiles() error {
	driveService, err := pkg.NewDriveService()
	if err != nil {
		return fmt.Errorf("failed to start Drive Service: %v", err)
	}

	fileService := drive.NewFilesService(driveService)

	fileList, err := fileService.List().OrderBy("name").Do()

	if err != nil {
		return fmt.Errorf("failed to List Files from Drive: %v", err)
	}

	values := [][]string{}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Type"})

	for _, v := range fileList.Files {
		var fileType string
		if v.MimeType == "application/vnd.google-apps.folder" {
			fileType = "Folder"
		} else {
			fileType = "File"
		}

		values = append(values, []string{
			v.Id,
			v.Name,
			fileType,
		})
	}

	if fileList.NextPageToken != "" {
		nxtToken := fileList.NextPageToken
		for {
			fileList, err = fileService.List().OrderBy("name").PageToken(nxtToken).Do()
			if err != nil {
				return fmt.Errorf("failed to List Files from Drive: %v", err)
			}
			for _, v := range fileList.Files {
				var fileType string
				if v.MimeType == "application/vnd.google-apps.folder" {
					fileType = "Folder"
				} else {
					fileType = "File"
				}
				values = append(values, []string{
					v.Id,
					v.Name,
					fileType,
				})
			}
			if fileList.NextPageToken == "" {
				break
			} else {
				nxtToken = fileList.NextPageToken
			}

		}
	}

	for _, v := range values {
		table.Append(v)
	}

	table.Render()

	return nil
}
