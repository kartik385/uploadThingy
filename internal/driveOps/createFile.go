package driveops

import (
	"fmt"

	"karti385/uploadThingy/pkg"

	"os"

	"google.golang.org/api/drive/v3"
)

func Upload() error {
    files, err := pkg.GetDirectoryFiles()
    if err != nil {
        return err
    }

    filePath, err := pkg.Questions(files)
    if err != nil {
        return err
    }

    driveService,err:=pkg.NewDriveService()
    if err != nil {
        return fmt.Errorf("failed to start Drive Service: %v", err)
    }

    fileService := drive.NewFilesService(driveService)

    file, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
    if err != nil {
        return fmt.Errorf("failed to open file: %v", err)
    }
    defer file.Close()

    driveFile := drive.File{
        Name: filePath,
    }

    fileInsertCall := fileService.Create(&driveFile).Media(file)

    fileInsertCall.ProgressUpdater(func(current, total int64) {
        fmt.Printf("Uploaded %d bytes out of %d bytes\n", current, total)
    })

    fileUploaded, err := fileInsertCall.Do()
    if err != nil {
        return fmt.Errorf("failed to upload file: %v", err)
    }

	fmt.Printf("%v",fileUploaded);

    fmt.Printf("%v has been successfully uploaded to Drive! ", fileUploaded.Name)

    return nil
}



