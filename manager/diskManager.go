package manager

import (
	"fmt"
	"main/constant"
	"main/helper"
	"os"
	"time"
)

// Function to store crawled content on disk
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

	//Starting separate go-routine to delete the file after timeout
	go DeleteFromDisc(url)
}

// Delete the file from disk
func DeleteFromDisc(url string) {

	sleepDuration := constant.EXPIRATION_TIMEOUT * int(time.Minute)
	time.Sleep(time.Duration(sleepDuration))

	err := os.RemoveAll(helper.UrlToDirectoryName(url))
	if err != nil {
		fmt.Println("Unable to remove from disc")
	}
}

// Check if the result is stored in disk
func IsStoredInDisk(url string) bool {
	filePath := helper.UrlToFilename(url)
	_, err := os.Stat(filePath)
	return err == nil
}

// Extract the result from the disk
func GetStoredPage(url string) (string, error) {
	filePath := helper.UrlToFilename(url)
	data, err := os.ReadFile(filePath)
	
	if err != nil {
		return "", err
	}

	println("Found data stored in cache")

	return string(data), nil
}