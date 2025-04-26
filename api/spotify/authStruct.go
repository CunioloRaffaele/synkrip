package spotify

type authResponse struct {
    AccessToken 	string `json:"access_token"`
    TokenType   	string `json:"token_type"`
    ExpiresIn   	int    `json:"expires_in"`
	ExpiresTime     int64 
}