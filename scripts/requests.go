package GIFBot

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"net/url"
)

// Remplace cette clé par la clé API Giphy obtenue
const GIPHY_API_KEY string = "wFlZrH69PNMzVclU281h9LSIxtB9lJd1"

type GiphyResponse struct {
	Data []struct {
		URL string `json:"url"`
	} `json:"data"`
}

func GifRequest(search string) string {

	reqURL, err := url.Parse("https://api.giphy.com/v1/gifs/search")
	if err != nil {
		fmt.Println("Error at creating request : ", err)
		return ""
	}

	params := url.Values{}
	params.Add("q", search)
	params.Add("api_key", GIPHY_API_KEY)
	params.Add("limit", "5")
	reqURL.RawQuery = params.Encode()

	resp, err := http.Get(reqURL.String())
	if err != nil {
		fmt.Println("Error at get method for your request : ", err)
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Http Code Error : ", resp.StatusCode)
		return ""
	}

	var giphyResponse GiphyResponse

	if err := json.NewDecoder(resp.Body).Decode(&giphyResponse); err != nil {
		fmt.Println("Error at parse response : ", err)
		return ""
	}

	if len(giphyResponse.Data) > 0 {
		return giphyResponse.Data[rand.IntN(len(giphyResponse.Data))].URL
	} else {
		fmt.Println(giphyResponse)
		return "https://giphy.com/gifs/ThisIsMashed-among-us-sus-amongus-XwsIv9WXc0FUUroO2H"
	}
}
