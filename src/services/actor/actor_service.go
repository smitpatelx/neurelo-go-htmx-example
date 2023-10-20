package actor

import (
	"fmt"
	"io"
	"log"
	"net/url"
	"strings"

	"github.com/goccy/go-json"
	"github.com/smitpatelx/neurelo-go-htmx-example/src/lib"
)

/* Actor Service */

func ReadAllActorSvc(req GetAllActorRequest) *[]Actor {
	request_url := "/rest/actor/?take=12"
	if req.Page > 1 {
		request_url = fmt.Sprintf("/rest/actor/?skip=%d&take=12", (req.Page-1)*12)
	}

	string_filter := fmt.Sprintf(`{"first_name":"%s"}`, req.Search)
	if len(strings.TrimSpace(req.Search)) > 0 {
		request_url = fmt.Sprintf("%s&filter=%s", request_url, url.QueryEscape(string_filter))
	}

	res, err := lib.Call(request_url)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	var actors *Client_GetAllActorResponse

	json_decoder := json.NewDecoder(io.Reader(res.Body))
	err1 := json_decoder.Decode(&actors)
	if err1 != nil {
		log.Fatal(err1.Error())
		return nil
	}

	defer res.Body.Close()

	data := actors.Data
	for i := 0; i < len(data); i++ {
		// Make date more readable
		data[i].LastUpdate = data[i].LastUpdate[:10]
	}

	return &actors.Data
}

func GetTotalActorsSvc(req GetAllActorRequest) *int {
	string_select := `{"_count":["actor_id"]}`

	request_url := fmt.Sprintf("/rest/actor/__aggregate?select=%s", url.QueryEscape(string_select))

	string_filter := fmt.Sprintf(`{"first_name":"%s"}`, req.Search)
	if strings.TrimSpace(req.Search) != "" {
		request_url = fmt.Sprintf("%s&filter=%s", request_url, url.QueryEscape(string_filter))
	}

	res, err := lib.Call(request_url)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	var actor_count_res *Client_GetTotalActorResponse

	json_decoder := json.NewDecoder(io.Reader(res.Body))
	err1 := json_decoder.Decode(&actor_count_res)
	if err1 != nil {
		log.Fatal(err1.Error())
		return nil
	}

	defer res.Body.Close()

	return &actor_count_res.Data.Count.ActorId
}

/* Film Service */

func ReadAllFilmSvc(req GetAllFilmRequest) *[]Film {
	request_url := "/rest/film/?take=12"
	if req.Page > 1 {
		request_url = fmt.Sprintf("/rest/film/?skip=%d&take=12", (req.Page-1)*12)
	}

	string_filter := fmt.Sprintf(`{"title":"%s"}`, req.Search)
	if strings.TrimSpace(req.Search) != "" {
		request_url = fmt.Sprintf("%s&filter=%s", request_url, url.QueryEscape(string_filter))
	}

	res, err := lib.Call(request_url)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	var films *Client_GetAllFilmResponse

	json_decoder := json.NewDecoder(io.Reader(res.Body))
	err1 := json_decoder.Decode(&films)
	if err1 != nil {
		log.Fatal(err1.Error())
		return nil
	}

	defer res.Body.Close()

	data := films.Data
	for i := 0; i < len(data); i++ {
		// Make date more readable
		data[i].LastUpdate = data[i].LastUpdate[:10]
	}

	return &films.Data
}

func GetTotalFilmSvc(req GetAllFilmRequest) *int {
	string_select := `{"_count":["film_id"]}`
	request_url := fmt.Sprintf("/rest/film/__aggregate?select=%s", url.QueryEscape(string_select))

	string_filter := fmt.Sprintf(`{"title":"%s"}`, req.Search)
	if strings.TrimSpace(req.Search) != "" {
		request_url = fmt.Sprintf("%s&filter=%s", request_url, url.QueryEscape(string_filter))
	}

	res, err := lib.Call(request_url)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	var film_count_res *Client_GetTotalFilmResponse

	json_decoder := json.NewDecoder(io.Reader(res.Body))
	err1 := json_decoder.Decode(&film_count_res)
	if err1 != nil {
		log.Fatal(err1.Error())
		return nil
	}

	defer res.Body.Close()

	return &film_count_res.Data.Count.FilmId
}
