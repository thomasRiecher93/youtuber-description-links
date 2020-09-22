package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type PlaylistItemsResponse struct {
	Items []PlaylistItem `json:"items"`
}

type PlaylistItem struct {
	Snippet VideoSnippet `json:"snippet"`
}

type VideoSnippet struct {
	ResourceId ResourceId `json:"resourceId"`
	Description string `json:"description"`
}

type ResourceId struct {
	VideoId string `json:"videoId"`
}

func GetVideosService(playlistID string) PlaylistItemsResponse {

	// Open our jsonFile
	jsonFile, err := os.Open("playlist-" + playlistID + ".json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var playlistItemsResponse PlaylistItemsResponse
	errGJson := json.Unmarshal(byteValue, &playlistItemsResponse)
	if errGJson != nil {
		fmt.Println(err)
	}

	return playlistItemsResponse
}
