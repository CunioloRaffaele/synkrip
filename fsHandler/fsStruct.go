package fsHandler


type Song struct {
	Name string `json:"name"`
}

type Playlist struct {
	PlaylistName string `json:"playlist_name"`
    File         []Song `json:"files"`
}

type FileSystem struct {
	Directories []Playlist `json:"directories"`
}