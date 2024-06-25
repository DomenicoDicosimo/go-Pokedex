package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListEncounters(locationName string) (RespEncounters, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		encountersResp := RespEncounters{}
		err := json.Unmarshal(val, &encountersResp)
		if err != nil {
			return RespEncounters{}, err
		}

		return encountersResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespEncounters{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespEncounters{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespEncounters{}, err
	}

	encountersResp := RespEncounters{}
	err = json.Unmarshal(dat, &encountersResp)
	if err != nil {
		return RespEncounters{}, err
	}

	return encountersResp, nil
}
