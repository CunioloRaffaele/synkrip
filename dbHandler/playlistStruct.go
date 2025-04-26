package dbHandler

type Playlist struct {
        DIR_ID          string `json:"dir_id"`
        SERVICE         string `json:"service"`
        URL             string `json:"url"`
        IMAGE           string `json:"image"`
        Songs           []Song `json:"songs"`
        ToBeSynced      bool   `json:"toBeSynced"`
}