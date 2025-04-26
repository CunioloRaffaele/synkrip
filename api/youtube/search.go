package youtube

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func GetYTid(title string) (string, error) {

	url := searchEndpoint
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(searchPayload, title))

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Println(err)
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer res.Body.Close()
	
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		log.Printf("Error: received status code %d\n", res.StatusCode)
		return "", fmt.Errorf("failed to get token from Youtube Api, status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	//fmt.Println("Youtube API search response: ", string(body))

	var response searchStruct
	err = json.Unmarshal(body, &response)
	if err != nil {	
		log.Println(err)
		return "", err
	}

	// The video id could be found in one of the response.Contents.TabbedSearchResultsRenderer.Tabs[0].TabRenderer.Content.SectionListRenderer.Contents[] array location
	for _, contentsArray := range response.Contents.TabbedSearchResultsRenderer.Tabs[0].TabRenderer.Content.SectionListRenderer.Contents {
		var videoIdResponse string
		if contentsArray.MusicCardShelfRenderer != nil {
			videoIdResponse = contentsArray.MusicCardShelfRenderer.Title.Runs[0].NavigationEndpoint.WatchEndpoint.VideoID
			//log.Println("Youtube API search response items: ", videoIdResponse)
			//return videoIdResponse, nil
		}else if contentsArray.MusicShelfRenderer != nil {
			videoIdResponse = contentsArray.MusicShelfRenderer.Contents[0].MusicResponsiveListItemRenderer.Overlay.MusicItemThumbnailOverlayRenderer.Content.MusicPlayButtonRenderer.PlayNavigationEndpoint.WatchEndpoint.VideoID
			//log.Println("Youtube API search response items: ", videoIdResponse)
			//return videoIdResponse, nil
		}
		if (videoIdResponse != "") {
			log.Println("Youtube API search response items: ", videoIdResponse)
			return videoIdResponse, nil
		}
	}

	//videoIdResponse:=response.Contents.TabbedSearchResultsRenderer.Tabs[0].TabRenderer.Content.SectionListRenderer.Contents[1].MusicCardShelfRenderer.Title.Runs[0].NavigationEndpoint.WatchEndpoint.VideoID
	//log.Println("Youtube API search response items: ", videoIdResponse)
	return "", err
}