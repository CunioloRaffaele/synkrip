package dbHandler

type Playlist struct {
        DIR_ID          string `json:"dir_id"`
        SERVICE         string `json:"service"`
        URL             string `json:"url"`
        IMAGE           string `json:"image"`
        ADD_DATE        string `json:"add_date"`
        LAST_MODIFIED   string `json:"last_modified"`
        Songs           []Song `json:"songs"`
        ToBeSynced      bool   `json:"toBeSynced"`
        DuplicateCount int    `json:"duplicate_count"`
        _songSet        map[string]struct{} `json:"-"` // Internal set to track duplicates
}