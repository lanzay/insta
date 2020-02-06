package insta

import (
	"bytes"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"insta/models"
	"io/ioutil"
	"log"
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
	gql := insta.EntryData.ProfilePage[0].Graphql
	src := gql.User
	getUserPageByScroll(queryHash, src, 0, 7000)
}

func getFirstUserPage(user string) (string, *models.Insta) {

	u := END_POINT + user
	code, body, err := fasthttp.Get(nil, u)
	if err != nil || code != 200 {
		log.Panicln("[E]", code, err)
	}
	queryHash := getQueryHash(body, Page)

	jsonBody := getJSONFromBody(body)
	var insta models.Insta
	err = json.Unmarshal(jsonBody, &insta)
	if err != nil {
		log.Panicln("[E]", err)
	}
	return queryHash, &insta
}

func getUserPageByScroll(queryHash string, o models.User, count, total int) {

	src := o.EdgeOwnerToTimelineMedia
	{ // doing something
		for _, v := range src.Edges {
			n := v.Node
			hook(n)
		}
	}

	//NextScroll
	if src.PageInfo.HasNextPage && (total == -1 || (count+12) < total) {
		count += 50
		next := GetNextScroll(queryHash, "id", src.Edges[0].Node.Owner.ID, count, *src.PageInfo.EndCursor)
		getUserPageByScroll(queryHash, next.Data.User, count, total)
	}
}
