package spotify

import "synkrip/api"

// SpotifyService implementa api.ApiMusicService
type SpotifyService struct{}

func (s *SpotifyService) FetchPlaylist(url string) (*api.UnifiedPlaylistFormat, error) {
    
	playlist, name, image := IngestSpotifyPlaylist(url)
    
    // Converte nel formato unificato
    var songs []api.UnifiedSongFormat
    for _, item := range playlist.Items {
        var artists []string
        for _, artist := range item.Track.Artists {
            artists = append(artists, artist.Name)
        }
        
        songs = append(songs, api.UnifiedSongFormat{
            ID:      item.Track.ID,
            Name:    item.Track.Name,
            Artists: artists,
            Album:   item.Track.Album.Name,
        })
    }

    return &api.UnifiedPlaylistFormat{
		Name:    name,
        Songs:   songs,
        Image:   image,
    }, nil
}

func (s *SpotifyService) GetServiceName() string {
    return "Spotify"
}

// Automatic registration on package import
func init() {
    api.RegisterService("Spotify", func() api.ApiMusicService {
        return &SpotifyService{}
    })
}