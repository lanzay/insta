package insta

import (
	"encoding/json"
	"insta/models"
	"io/ioutil"
	"log"
	"net/http"
)

// https://www.instagram.com/explore/tags/aaa/?__a=1
func GetPostsByTag(tag string) {

	queryHash, insta := getFirstTagPage(tag)

	if insta == nil || len(insta.EntryData.TagPage) == 0 || len(queryHash) == 0 {
		return
	}

	gql := insta.EntryData.TagPage[0].Graphql
	src := gql.Hashtag
	getTagPageByScroll(queryHash, src, 0, 7000)
}

func getFirstTagPage(tag string) (string, *models.Insta) {

	u := END_POINT + "explore/tags/" + tag + "/"
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
	queryHash := getQueryHash(body, Tag)

	jsonBody := getJSONFromBody(body)
	var insta models.Insta
	err = json.Unmarshal(jsonBody, &insta)
	if err != nil {
		log.Panicln("[E] firs page", err)
	}
	return queryHash, &insta
}

func getTagPageByScroll(queryHash string, hashTag models.Hashtag, count, total int) {

	src := hashTag.EdgeHashtagToMedia
	{ // doing something
		for _, v := range src.Edges {
			n := v.Node
			hook(n)
		}
	}
	//NextScroll
	if src.PageInfo.HasNextPage && (total == -1 || (count+70) < total) {
		count += 50
		next := GetNextScroll(queryHash, "tag_name", hashTag.Name, count, *src.PageInfo.EndCursor, 1)
		if next == nil {
			log.Println("[E] E006", queryHash, next, count, total)
			return
		}
		getTagPageByScroll(queryHash, next.Data.Hashtag, count, total)

	}
}
