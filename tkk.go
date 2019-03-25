package tk

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

var homeURL string = "https://translate.google.cn"

func GetTKK() (tkk string, err error) {
	resp, err := http.Get(homeURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	data := string(body)

	r := regexp.MustCompile(`tkk:'(\d+\.\d+)'`)
	if r.MatchString(data) {
		tkk = r.FindStringSubmatch(data)[1]
	}
	return
}
