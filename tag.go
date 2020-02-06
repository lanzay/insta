package insta

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"insta/models"
	"log"
)

// https://www.instagram.com/explore/tags/aaa/?__a=1
func GetPostsByTag(tag string) {

	queryHash, insta := getFirstTagPage(tag)

	gql := insta.EntryData.TagPage[0].Graphql
	src := gql.Hashtag
	getTagPageByScroll(queryHash, src, 0, 7000)
}

func getFirstTagPage(tag string) (string, *models.Insta) {

	u := END_POINT + "explore/tags/" + tag + "/"
	code, body, err := fasthttp.Get(nil, u)
	if err != nil || code != 200 {
		log.Panicln("[E]", code, err)
	}
	queryHash := getQueryHash(body, Tag)

	jsonBody := getJSONFromBody(body)
	var insta models.Insta
	err = json.Unmarshal(jsonBody, &insta)
	if err != nil {
		log.Panicln("[E]", err)
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
		next := GetNextScroll(queryHash, "tag_name", hashTag.Name, count, *src.PageInfo.EndCursor)
		getTagPageByScroll(queryHash, next.Data.Hashtag, count, total)
	}
}
