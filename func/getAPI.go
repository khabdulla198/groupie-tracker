package groupie

import (
	"encoding/json"
	"io"
	"net/http"
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate string   `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDate  string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func getData() ([]Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data []Artist

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil

}
