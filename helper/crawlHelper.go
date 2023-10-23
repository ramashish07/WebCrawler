package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"main/constant"
)

func Sha256Hash(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}

func UrlToFilename(url string) string {
	hashUrl := string(Sha256Hash(url))
	return constant.CACHE_PATH + hashUrl + constant.FILE_PATH
}

func UrlToDirectoryName(url string) string {
	hashUrl := string(Sha256Hash(url))
	return constant.CACHE_PATH + hashUrl
}