package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"sort"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var (
	playListID = flag.String("playListID", "", "youtube playlist id to analyse")
	apiKey     = flag.String("apiKey", "", "apiKey for youtube api")
)

type Video struct {
	Title     string `json:"title,omitempty"`
	ViewCount uint64 `json:"viewCount,omitempty"`
	ID        string `json:"id,omitempty"`
	URL       string `json:"url,omitempty"`
}

// List all videos from a given playlist
func lsPlaylistVideos(yts *youtube.Service, pid string) error {
	q := yts.PlaylistItems.List([]string{"snippet,contentDetails"})

	var pt = ""
	var result []Video
	for {
		xs, err := q.Do(
			googleapi.QueryParameter("playlistId", pid),
			googleapi.QueryParameter("maxResults", "20"),
			googleapi.QueryParameter("pageToken", pt),
		)
		if err != nil {
			return err
		}

		for _, x := range xs.Items {
			videoResponse, _ := yts.Videos.List([]string{"statistics", "snippet", "contentDetails"}).Id(x.Snippet.ResourceId.VideoId).Do()
			if len(videoResponse.Items) > 0 {
				video := videoResponse.Items[0]
				result = append(result, Video{
					Title:     video.Snippet.Title,
					ViewCount: video.Statistics.ViewCount,
					ID:        video.Id,
					URL:       fmt.Sprintf("https://www.youtube.com/watch?v=%s", video.Id),
				})
			}
		}
		pt = xs.NextPageToken

		if pt == "" {
			break
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].ViewCount < result[j].ViewCount
	})

	resultJSON, _ := json.Marshal(result)
	fmt.Println(string(resultJSON))
	return nil
}

func main() {
	flag.Parse()
	service, err := youtube.NewService(context.Background(), option.WithAPIKey(*apiKey))
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	err = lsPlaylistVideos(service, *playListID)
	if err != nil {
		log.Fatal(err)
	}
}
