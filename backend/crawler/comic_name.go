package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetComicName(id string) (string, error) {
	url := fmt.Sprintf("https://www.comicabc.com/html/%s.html", id)

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

	nameTag := "<h3 class=\"item_name\" title=\""
	nameRaws := strings.Split(string(html), nameTag)
	nameRaws = strings.Split(nameRaws[1], "\">")

	return nameRaws[0], nil
}
