package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetPageLatestVolume(id string) (string, error) {
	url := fmt.Sprintf("https://www.comicabc.com/html/%s.html", id)
	latestChapterTag := "id=lch>"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create http.Request: %w", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to do request to %v: %w", url, err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}
	raws := strings.Split(string(html), latestChapterTag)
	raws = strings.Split(raws[1], "</font>")
	return raws[0], nil
}
