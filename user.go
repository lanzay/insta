package insta

import (
	"bytes"
	"encoding/json"
	"insta/models"
	"io/ioutil"
	"log"
	"net/http"
)

func GetPostsByUserList(list string) {

	body, err := ioutil.ReadFile(list)
	if err != nil {
		log.Println("[E]", err)
		return
	}
	lines := bytes.Split(body, []byte("\r\n"))
	if len(lines) == 1 {
		lines = bytes.Split(body, []byte("\n"))
	}

	for _, v := range lines {
		v = bytes.Trim(v, "@")
		v = bytes.TrimSpace(v)
		GetPostsByUser(string(v))
	}
}

func GetPostsByUser(user string) {

	queryHash, insta := getFirstUserPage(user)

	if insta == nil || len(insta.EntryData.TagPage) == 0 || len(queryHash) == 0 {
		return
	}

	gql := insta.EntryData.ProfilePage[0].Graphql
	src := gql.User
	getUserPageByScroll(queryHash, src, 0, 7000)
}

func getFirstUserPage(user string) (string, *models.Insta) {

	u := END_POINT + user
	res, err := http.Get(u)
	if err != nil || res.StatusCode != 200 {
		log.Panicln("[E] GET user", res.StatusCode, u, err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panicln("[E] GET.Read user", res.StatusCode, u, err)
	}
	defer res.Body.Close()

	queryHash := getQueryHash(body, Page)

	jsonBody := getJSONFromBody(body)
	var insta models.Insta
	err = json.Unmarshal(jsonBody, &insta)
	if err != nil {
		log.Panicln("[E] JSON", err)
	}
	return queryHash, &insta
}

func getUserPageByScroll(queryHash string, o models.User, count, total int) {

	src := o.EdgeOwnerToTimelineMedia
	{ // doing something
		for _, v := range src.Edges {
			hook(v.Node)
		}
	}

	//NextScroll
	if src.PageInfo.HasNextPage && (total == -1 || (count+12) < total) {
		count += 50
		next := GetNextScroll(queryHash, "id", src.Edges[0].Node.Owner.ID, count, *src.PageInfo.EndCursor)
		getUserPageByScroll(queryHash, next.Data.User, count, total)
	}
}
