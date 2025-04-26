package spotify

import "fmt"


func SpotifyTets() {
  //https://open.spotify.com/playlist/1gmYqSJ68IZZmt15oIU8l5?si=219aa181b3274f3f

  playlist, name, image:= IngestSpotifyPlaylist("https://open.spotify.com/playlist/1gmYqSJ68IZZmt15oIU8l5?si=219aa181b3274f3f")
  //fmt.Println(playlist.Items)
  fmt.Println(playlist.Items[0].Track.Album.Images[0].URL)
  fmt.Println(len(playlist.Items))
  fmt.Println(playlist.Total)
  fmt.Println(name)
  fmt.Println(image)
}