package insta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"insta/models"
	"log"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

const (
	END_POINT = "https://www.instagram.com/"
)

var (
	DATA_DIR       = "data"
	count    int64 = 0
)

func init() {
	os.MkdirAll(DATA_DIR, 0666)
}

func hook(n models.PurpleNode) {

	count = atomic.AddInt64(&count, 1)
	img := n.DisplayURL
	log.Printf("[I] %d) %s UserID: %s IMG: %s https://instagramm/p/%s\n", count, n.Owner.Username, n.Owner.ID, n.ID, n.Shortcode) //, img)
	getIMG(n.Owner.Username, n.Owner.ID, n.ID, img)

	f, _ := os.OpenFile("insta_detail.json", os.O_APPEND|os.O_CREATE, 0666)
	defer f.Close()
	body, _ := json.Marshal(n)
	f.Write(body)
	f.Write([]byte("\n"))
	//log.Println(string(body))

	if webHooks := viper.GetStringSlice("webhooks"); len(webHooks) != 0 {
		for _, webHook := range webHooks {
			http.Post(webHook, "application/json", bytes.NewReader(body))
		}
	}
}

func getJSONFromBody(body []byte) []byte {
	s := bytes.Index(body, []byte("window._sharedData = ")) + len("window._sharedData = ")
	e := bytes.Index(body[s:], []byte(";</script>"))
	return body[s : s+e]
}

type TargetType int

const (
	Page TargetType = iota
	Tag
)

// query_hash
// 0 = UserPage
// 1 = Tag
func getQueryHash(body []byte, targetType TargetType) string {

	// exapmle
	// https://www.instagram.com/static/bundles/es6/TagPageContainer.js/4aa59b65e888.js

	var targetCode, targetHash []byte
	switch targetType {
	case Page: // UserProfile->NextScroll
		targetCode = []byte("ProfilePageContainer.js/")
		targetHash = []byte("l.pagination},queryId:\"")
	case Tag: // Tag->NextScroll
		targetCode = []byte("TagPageContainer.js/")
		targetHash = []byte("t.tagMedia.byTagName.get(n)).pagination},queryId:\"")
	case -1:
		// targetHash = []byte("t.comments.byPostId.get(n).pagination},queryId:\"")
		// targetHash = []byte("t.savedPosts.byUserId.get(n).pagination},queryId:\"")
		// targetHash = []byte("o.pagination},queryId:\"")
	}

	s := bytes.Index(body, targetCode) + len(targetCode)
	e := bytes.Index(body[s:], []byte(".js"))
	queryHashURL := END_POINT + "/static/bundles/metro/" + string(targetCode) + string(body[s:s+e]) + ".js"

	//log.Println("============================",string(targetCode))
	//log.Println(string(body))

	_, bodyJS, _ := fasthttp.Get(nil, queryHashURL)
	s = bytes.Index(bodyJS, targetHash) + len(targetHash)
	e = bytes.Index(bodyJS[s:], []byte("\""))

	//log.Println("============================",string(targetHash))
	//log.Println(string(bodyJS))

	QueryHash := string(bodyJS[s : s+e])
	return QueryHash
}

func GetNextScroll(query_hash, p1, v1 string, count int, after string) *models.InstaNext {

	variables := fmt.Sprintf("{\"%s\":\"%s\",\"first\":%d,\"after\":\"%s\"}", p1, v1, count, after)
	u := fmt.Sprintf(`https://www.instagram.com/graphql/query/?query_hash=%s&variables=%s`, query_hash, variables)

	code, body, err := fasthttp.Get(nil, u)
	if code == 429 {
		log.Println(u)
		log.Println("[I] Rate limit 5000 pet hour. Wait 5 min...")
		time.Sleep(5 * time.Minute)
		GetNextScroll(query_hash, p1, v1, count, after)
		return nil
	}
	if err != nil || code != 200 {
		log.Println(u)
		log.Println("[E]", code, err, string(body))
		return nil
	}

	// JSON
	if len(body) < 1000 {
		log.Println("[E]", string(body))
		return nil
	}
	jsonBody := body

	var insta models.InstaNext
	err = json.Unmarshal(jsonBody, &insta)
	if err != nil {
		log.Println("[E]", err)
		return nil
	}
	return &insta
}

func getIMG(userName, userID, imgID, url string) {

	folder := userName
	if len(folder) == 0 {
		folder = "__by_id"
	}
	os.MkdirAll(DATA_DIR+"/"+folder, 0666)

	code, body, err := fasthttp.Get(nil, url)
	if err != nil || code != 200 {
		log.Panicln("[E]", code, err)
	}
	f, _ := os.Create(DATA_DIR + "/" + folder + "/" + userName + "-[" + userID + "]-" + imgID + ".jpg")
	defer f.Close()
	f.Write(body)
	f.Close()
}
