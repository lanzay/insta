package insta

import (
	"encoding/json"
	"insta/models"
	"io/ioutil"
	"log"
	"net/http"
)

func GetShortCodeMedia(shortcode string) *models.Insta {

	u := END_POINT + "/p/" + shortcode + "/"

	res, err := http.Get(u)
	if err != nil || res.StatusCode != 200 {
		log.Panicln("[E] GET tag", res.StatusCode, u, err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panicln("[E] GET.Read tag", res.StatusCode, u, err)
	}
	res.Body.Close()

	jsonBody := getJSONFromBody(body)

	var insta models.Insta
	err = json.Unmarshal(jsonBody, &insta)
	if err != nil {
		log.Panicln("[E] json", err)
	}
	return &insta
}
