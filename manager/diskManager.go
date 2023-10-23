package manager

import (
	"fmt"
	"main/constant"
	"main/helper"
	"os"
	"time"
)

func StoreInDisk(url string, content string) {
	filePath := helper.UrlToFilename(url)
	dirPath := helper.UrlToDirectoryName(url)

	fmt.Println("Filepath is ", filePath)

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error in storing data in the disc")
	}

	go DeleteFromDisc(url)
}

func DeleteFromDisc(url string) {

	sleepDuration := constant.EXPIRATION_TIMEOUT * int(time.Second)
	time.Sleep(time.Duration(sleepDuration))

	err := os.RemoveAll(helper.UrlToDirectoryName(url))
	if err != nil {
		fmt.Println("Unable to remove from disc")
	}
}

func IsStoredInDisk(url string) bool {
	filePath := helper.UrlToFilename(url)
	_, err := os.Stat(filePath)
	return err == nil
}

func GetStoredPage(url string) (string, error) {
	filePath := helper.UrlToFilename(url)
	data, err := os.ReadFile(filePath)
	
	if err != nil {
		return "", err
	}

	println("Found data stored in cache")

	return string(data), nil
}