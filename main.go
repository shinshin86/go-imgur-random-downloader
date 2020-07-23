package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

const rsLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {
	n := flag.Int("n", 1, "Number of files")
	imgPath := flag.String("path", ".", "Image download path")
	flag.Parse()

	for range make([]int, *n) {
		download(getUrl(), *imgPath)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getUrl() string {
	b := make([]byte, 5)
	for i := range b {
		b[i] = rsLetters[rand.Intn(len(rsLetters))]
	}
	extArr := []string{".jpg", ".png"}
	ext := extArr[rand.Intn(len(extArr))]

	return "https://i.imgur.com/" + string(b) + ext
}

func download(url string, imgPath string) {
	fmt.Println("Download URL:" + url)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("status:", resp.Status)

	if resp.Request.URL.String() == "https://i.imgur.com/removed.png" {
		fmt.Println("Couldn't downloaded. This is removed image")
		return
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, filename := path.Split(url)

	fmt.Println("Download Path:" + filepath.Join(imgPath, filename))

	file, err := os.OpenFile(filepath.Join(imgPath, filename), os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		file.Close()
	}()

	file.Write(body)
}
