package dbHandler

type Song struct {
        SONG_NAME        string `json:"song_name"`
        SONG_ALBUM_NAME  string `json:"song_album_name"`
        SONG_ARTIST_NAME string `json:"song_artist_name"`
        SONG_YT_ID       string `json:"song_yt_id"`
        SONG_PLATFORM_ID string `json:"song_platform_id"`
        IS_DOWNLOADED    int    `json:"is_downloaded"`
        DIR_ID          string `json:"dir_id"`
        ADD_DATE         string `json:"add_date"`
}