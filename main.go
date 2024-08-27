package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/valyala/fasthttp"
)

const siteInfoRegex = `window\.siteInfo\s*=\s*({[^}]+})`

func main() {
	args := getArgs()
	urlString := args[1]

	validateURL(urlString)

	status, resp, err := fasthttp.Get(nil, urlString)
	checkError(err, status)

	siteInfo, err := generateSiteInfo(resp)
	checkError(err, fasthttp.StatusOK)

	uid := siteInfo["uid"].(string)
	host := siteInfo["host"].(string)

	reqUrl := "https://" + host + "/cache/" + uid
	status, resp, err = fasthttp.Get(nil, reqUrl)
	checkError(err, status)

	cacheData := make(map[string]interface{})
	err = json.Unmarshal(resp, &cacheData)
	checkError(err, fasthttp.StatusOK)

	var wg sync.WaitGroup
	total := len(cacheData)
	progress := 0
	progressMutex := &sync.Mutex{}

	for i := range cacheData {
		urlString = "https://" + host + "/access/" + uid + "/" + i
		path := filepath.Join(args[2], i)
		createParentFolder(path)
		wg.Add(1)
		go downloadAndSave(urlString, path, &wg, &progress, total, progressMutex)
	}
	wg.Wait()
}

func validateURL(urlString string) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		log.Fatal("Invalid URL")
	}
}

func generateSiteInfo(resp []byte) (map[string]interface{}, error) {
	re := regexp.MustCompile(siteInfoRegex)
	matches := re.FindStringSubmatch(string(resp))
	if len(matches) < 2 {
		log.Fatal("Unable to extract siteInfo")
	}
	siteInfo := make(map[string]interface{})
	err := json.Unmarshal([]byte(matches[1]), &siteInfo)
	return siteInfo, err
}

func downloadAndSave(url string, path string, wg *sync.WaitGroup, progress *int, total int, progressMutex *sync.Mutex) {
	defer wg.Done()
	status, resp, err := fasthttp.Get(nil, url)
	checkError(err, status)

	out, err := os.Create(path)
	checkError(err, fasthttp.StatusOK)
	defer func(out *os.File) {
		err = out.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(out)

	_, err = io.WriteString(out, string(resp))
	checkError(err, fasthttp.StatusOK)

	progressMutex.Lock()
	*progress++
	printProgressBar(*progress, total)
	progressMutex.Unlock()
}

func printProgressBar(progress int, total int) {
	percent := float64(progress) / float64(total) * 100
	barLength := 50
	progressLength := int(float64(barLength) * percent / 100)
	bar := fmt.Sprintf("[%s%s] %d%%",
		string(repeat('#', progressLength)),
		string(repeat(' ', barLength-progressLength)),
		int(percent))
	fmt.Printf("\r%s", bar)
	if progress == total {
		fmt.Println()
	}
}

func repeat(char rune, count int) []rune {
	result := make([]rune, count)
	for i := 0; i < count; i++ {
		result[i] = char
	}
	return result
}

func getArgs() []string {
	args := os.Args
	if len(args) < 3 {
		printUsageExiting()
	}
	return args
}

func printUsageExiting() {
	fmt.Println("\033[31mUsage: go run main.go <url> <directory>\033[0m")
	os.Exit(1)
}

func createParentFolder(path string) {
	parentFolder := filepath.Dir(path)
	if _, err := os.Stat(parentFolder); os.IsNotExist(err) {
		err = os.MkdirAll(parentFolder, os.ModePerm)
		checkError(err, fasthttp.StatusOK)
	}
}

func checkError(err error, status int) {
	if err != nil || status != fasthttp.StatusOK {
		log.Fatal(err)
	}
}
