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
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDate  string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Relations struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func GetData() ([]Artist, error) {
	//get the data from API
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}

	//close the response body
	defer resp.Body.Close()

	//read the response body which is the data in the API
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data []Artist

	//convert the data
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func GetRelations() ([]Relations, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data map[string][]Relations

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	// Extract data from the "index" key
	relations, exist := data["index"]
	if !exist {
		return nil, err
	} 

	return relations, nil

}
