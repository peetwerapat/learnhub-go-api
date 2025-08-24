package oembed

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type OEmbedDto struct {
	AuthorName   string `json:"author_name"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnail_url"`
	Title        string `json:"title"`
}

type OEmbedErrorDto struct {
	Error string `json:"error"`
}

func GetOEmbedInfo(videoUrl string) (*OEmbedDto, error) {
	endpoint := fmt.Sprintf("https://noembed.com/embed?url=%s", url.QueryEscape(videoUrl))

	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch oEmbed data: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read oEmbed response: %w", err)
	}

	var errDto OEmbedErrorDto
	if err := json.Unmarshal(body, &errDto); err == nil && errDto.Error != "" {
		return nil, fmt.Errorf("invalid video URL: %s", errDto.Error)
	}

	var data OEmbedDto
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to parse oEmbed data: %w", err)
	}

	return &data, nil
}
