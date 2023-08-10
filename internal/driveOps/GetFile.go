package driveops

import (
	"fmt"
	"io"
	"karti385/uploadThingy/pkg"
	"os"

	"google.golang.org/api/drive/v3"
)

func DownloadFile(fileId string) error {
	driveService, err := pkg.NewDriveService()
	if err != nil {
		return fmt.Errorf("failed to start Drive Service: %v", err)
	}

	fileService := drive.NewFilesService(driveService)
	file, err := fileService.Get(fileId).Do()
	if err != nil {
		return fmt.Errorf("failed to download the file!: %v", err)
	}

	response, err := fileService.Get(fileId).Download()

	if err != nil {
		return fmt.Errorf("failed to download the file!: %v", err)
	}

	defer response.Body.Close()

	// scanner := bufio.NewScanner(response.Body)

	newFile, err := os.Create(file.Name)

	if err != nil {
		return fmt.Errorf("failed to create the file!: %v", err)
	}

	defer newFile.Close()

	_, err = io.Copy(newFile, response.Body)

	// for scanner.Scan() {
	// 	writer.Write([]byte(scanner.Text()))
	// }

	// err = writer.Flush()

	if err != nil {
		return fmt.Errorf("failed to copy the file!: %v", err)
	}

	fmt.Printf("%v has been successfully downloaded! ", file.Name)

	return nil
}
