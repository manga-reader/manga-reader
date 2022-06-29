package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func GetPageLatestVolumeAndDate(id string) (string, *time.Time, error) {
	url := fmt.Sprintf("https://www.comicabc.com/html/%s.html", id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", nil, fmt.Errorf("failed to create http.Request: %w", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", nil, fmt.Errorf("failed to do request to %v: %w", url, err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil, fmt.Errorf("failed to read response body: %w", err)
	}
	latestChapterTag := "id=lch>"
	latestChapterRaws := strings.Split(string(html), latestChapterTag)
	latestChapterRaws = strings.Split(latestChapterRaws[1], "</font>")

	updatedAtRaws := strings.Split(string(html), "更新：")
	updatedAtRaws = strings.Split(updatedAtRaws[1], "</span>")
	updatedAtRaw := updatedAtRaws[0][len(updatedAtRaws[0])-10:]

	updatedAtRaw = updatedAtRaw + "T00:00:00Z"
	t, err := time.Parse(time.RFC3339, updatedAtRaw)
	if err != nil {
		return "", nil, fmt.Errorf("failed to parse updated at time: %s: %w", updatedAtRaw, err)
	}

	return latestChapterRaws[0], &t, nil
}
