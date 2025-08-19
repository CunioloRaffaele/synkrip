package api

import (
	"fmt"
	"log"
)

type UnifiedSongFormat struct {
	ID string
	Name string
	Artists []string
	Album string
}

type UnifiedPlaylistFormat struct {
	Name string
	Songs []UnifiedSongFormat
	Image string 				// base64 encoded image
}

type ApiMusicService interface {
    // FetchPlaylist gets data and converts it to the unified format
    FetchPlaylist(url string) (*UnifiedPlaylistFormat, error)

    // GetServiceName returns the name of the service
    GetServiceName() string
}

// Registry
var serviceRegistry = make(map[string]func() ApiMusicService)

// RegisterService (to register a new music service)
func RegisterService(name string, factory func() ApiMusicService) {
    serviceRegistry[name] = factory
	log.Printf("Registering service: %s\n", name)
}

// GetMusicService (to get a music service by name)
func GetMusicService(serviceName string) (ApiMusicService, error) {
    factory, exists := serviceRegistry[serviceName]
    if !exists {
        return nil, fmt.Errorf("unknown service: %s", serviceName)
    }
    return factory(), nil
}