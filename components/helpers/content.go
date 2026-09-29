package helpers

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var ContentPath string
var ContentDir string
var MainUrl string

func GetImageUrl(image string) string {
	return ContentPath + "/" + image
}

func RandSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func FormatDate(date string, fromFormat string, toFormat string) string {
	formattedDate, _ := time.Parse(fromFormat, date)
	date = formattedDate.Format(toFormat)

	return date
}

func SaveFileToContent(uploadedFile io.Reader, inFolder string, fileName string) (bool, string) {
	var docPath string
	var contentPath = ContentDir + inFolder + "/"

	if _, err := os.Stat(contentPath); os.IsNotExist(err) {
		err := os.Mkdir(contentPath, 0777)
		if err != nil {
			log.Println(err)
		}
	}

	targetFile, err := os.OpenFile(contentPath+fileName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Println(err)
		return false, docPath
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		log.Println(err)
		return false, docPath
	}

	docPath = "/media/" + inFolder + "/" + fileName
	return true, docPath
}

func ArrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
