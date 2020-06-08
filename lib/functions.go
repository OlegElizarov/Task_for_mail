package lib

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

func NumberInFile(path string) (int, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, err
	}
	re := regexp.MustCompile(`([Gg])o`)
	//re := regexp.MustCompile(`(^| |\n|\W)Go(\n| |$|\W)`) in order if you want only "Go" string
	return len(re.FindAll(dat, -1)), nil
	//return strings.Count(string(body), "Go")+strings.Count(string(body), "go"), nil   bad way
}

func NumberInURL(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	re := regexp.MustCompile(`([Gg])o`)
	//re := regexp.MustCompile(`(^| |\n|\W)Go(\n| |$|\W)`) in order if you want only "Go" string
	return len(re.FindAll(body, -1)), nil
	//return strings.Count(string(body), "Go")+strings.Count(string(body), "go"), nil   bad way
}
