package youtube

import (
	"strings"
	"synkrip/api"
)

// YoutubeMusicService implementa api.ApiMusicService
type YoutubeMusicService struct{}

func (s *YoutubeMusicService) FetchPlaylist(url string) (*api.UnifiedPlaylistFormat, error) {

	
	id := ""
	if strings.Contains(url, "list=") {
		id = strings.Split(url, "list=")[1]
		if strings.Contains(id, "&") {
			id = strings.Split(id, "&")[0]
		}
	}
	name, playlist, image := IngestYoutubePlaylist(id)

    // Converte nel formato unificato
    var songs []api.UnifiedSongFormat
    for _, item := range playlist.Contents.TwoColumnBrowseResultsRenderer.Tabs[0].TabRenderer.Content.SectionListRenderer.Contents[0].ItemSectionRenderer.Contents[0].PlaylistVideoListRenderer.Contents {
        var artists []string
        for _, artist := range item.PlaylistVideoRenderer.ShortBylineText.Runs {
            artists = append(artists, artist.Text)
        }
        
        songs = append(songs, api.UnifiedSongFormat{
            ID:      item.PlaylistVideoRenderer.VideoID,
            Name:    item.PlaylistVideoRenderer.Title.Runs[0].Text,
            Artists: artists,
            Album:   "",
        })
    }

	if (name == "") {
		name = "Unknown YouTube Playlist"
	}
	
    return &api.UnifiedPlaylistFormat{
		Name:    name,
        Songs:   songs,
        Image:   image,
    }, nil
}

func (s *YoutubeMusicService) GetServiceName() string {
    return "YouTube"
}

// Automatic registration on package import
func init() {
    api.RegisterService("YouTube", func() api.ApiMusicService {
        return &YoutubeMusicService{}
    })
}