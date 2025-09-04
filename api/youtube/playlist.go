package youtube

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func imageToBase64 (playlistImage string) string {
    // Convert the image URL to base64
    if playlistImage != "" {
        resp, err := http.Get(playlistImage)
        if err != nil {
            log.Println("Error fetching playlist image:", err)
        } else {
            defer resp.Body.Close()
            imageData, err := io.ReadAll(resp.Body)
            if err != nil {
                log.Println("Error reading image data:", err)
            } else {
                playlistImage = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(imageData)
            }
        }
    }
    return playlistImage
}

func extractImageFromHtml (url string) string {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	// Regex per <meta property="og:image" content="...">
	re := regexp.MustCompile(`<meta property="og:image" content="([^"]+)"`)
	match := re.FindSubmatch(body)

	if len(match) > 1 {
		log.Println("Cover URL:", string(match[1]))
	} else {
		log.Println("Meta og:image non trovato")
	}
	return string(match[1])
}

// get playlist information from playlist id
func IngestYoutubePlaylist(id string) (name string , playlist playlistStruct, image string) {
 	url := playlistEndpoint
	log.Println("YT Fetching playlist from:", url)
  	method := "POST"
 	payload := strings.NewReader(fmt.Sprintf(playlistPayload, id))
	//log.Println("Payload:", payload)

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Println("Error: ", res.StatusCode)
		return
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var response playlistStruct
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println(err)
		return
	}

	var playlistYtUrlHtml = "https://music.youtube.com/playlist?list=" + id
	var playlistImage = extractImageFromHtml(playlistYtUrlHtml)
	var playlistImageBase64 = imageToBase64(playlistImage)

	/*if (response.OnResponseReceivedActions != nil) {
		fmt.Println ("playlist no second part")
	} else {
		fmt.Println ("playlist has second part")
	}*/
	
	//fmt.Println("Playlist Name: ", response.Header.PageHeaderRenderer.Content.PageHeaderViewModel.Title.DynamicTextViewModel.Text)
	//fmt.Println("Playlist first song: ", response.Contents.TwoColumnBrowseResultsRenderer.Tabs[0].TabRenderer.Content.SectionListRenderer.Contents[0].ItemSectionRenderer.Contents[0].PlaylistVideoListRenderer.Contents[0].PlaylistVideoRenderer.Title.Runs[0].Text)

	return response.Header.PageHeaderRenderer.Content.PageHeaderViewModel.Title.DynamicTextViewModel.Text.Content, 
	response, 
	playlistImageBase64
}