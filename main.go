package insta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"insta/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

const (
	END_POINT = "https://www.instagram.com/"
)

var (
	DATA_DIR = "data"
	WG       sync.WaitGroup
	count    int64 = 0
)

func init() {
	os.MkdirAll(DATA_DIR, 0777)
}

func hook(n models.PurpleNode) {

	count = atomic.AddInt64(&count, 1)
	if count == 1 || count%500 == 0 {
		log.Printf("[I] %d) https://instagramm/%s UserID:%s IMG:%s https://instagramm/p/%s\n", count, n.Owner.Username, n.Owner.ID, n.ID, n.Shortcode)
	}

	// TODO 2020-03-13 Anton
	//img := n.DisplayURL
	//go func() {
	//	WG.Add(1)
	//	getIMGMust(n.Owner.Username, n.Owner.ID, n.ID, img)
	//	WG.Done()
	//}()

	body, _ := json.Marshal(n)

	//f, _ := os.OpenFile("insta_detail.json", os.O_APPEND|os.O_CREATE, 0666)
	//defer f.Close()
	//f.Write(body)
	//f.Write([]byte("\n"))
	//f.Close()
	////log.Println(string(body))

	if webHooks := viper.GetStringSlice("webhooks"); len(webHooks) != 0 {
	m1:
		for _, webHook := range webHooks {
			for try := 1; ; try++ {
				res, err := http.Post(webHook, "application/json", bytes.NewReader(body))
				res.Body.Close()
				if err != nil {
					log.Println("[E] POST hook ERR", try, err)
					runtime.Gosched()
					time.Sleep(500 * time.Millisecond)
					continue m1
				}
				if err == nil && res.StatusCode != 200 {
					log.Println("[E] POST hook !200", res.StatusCode, err, string(body))
				}
				break m1
			}
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
	var QueryHash string
	if s = bytes.Index(bodyJS, targetHash) + len(targetHash); s > 0 {
		if e = bytes.Index(bodyJS[s:], []byte("\"")); e > 0 {
			QueryHash = string(bodyJS[s : s+e])
		}
	}

	//log.Println("============================",string(targetHash))
	//log.Println(string(bodyJS))

	return QueryHash
}

func GetNextScroll(query_hash, p1, v1 string, count int, after string, try int) *models.InstaNext {

	try++
	variables := fmt.Sprintf("{\"%s\":\"%s\",\"first\":%d,\"after\":\"%s\"}", p1, v1, count, after)
	u := fmt.Sprintf(`https://www.instagram.com/graphql/query/?query_hash=%s&variables=%s`, query_hash, variables)

	res, err := http.Get(u)
	if err != nil {
		log.Println(u)
		log.Println("[E] E001", res.StatusCode, err)
		return nil
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panicln("[E] GET.Read next scroll", res.StatusCode, u, err)
	}
	defer res.Body.Close()

	if res.StatusCode == 429 {
		if bytes.Contains(body, []byte(`{"message": "rate limited", "status": "fail"}`)) {
			log.Println("[I] Rate limit 5000 pet hour. Wait 1 min...")
			time.Sleep(1 * time.Minute)
			return GetNextScroll(query_hash, p1, v1, count, after, try)
			return nil
		} else {
			log.Println("[E] StatusCode: 429 But any problem", u)
			log.Println(string(body))
			return nil
		}
	}

	if res.StatusCode != 200 {
		log.Println("[E] E005", string(body))
		log.Println("[E] E005", res.StatusCode, err)
		log.Println("[E] E005", query_hash, p1, v1, count, after, try)
		log.Println("[E] E005 url:", u, try)
		if try <= 3 {
			log.Println("[D] wait 1 mimutes...")
			time.Sleep(1 * time.Minute)
			return GetNextScroll(query_hash, p1, v1, count, after, try)
		}
		return nil
	}

	// JSON
	if len(body) < 1000 {
		log.Println("[E] E002", string(body))
		return nil
	}
	jsonBody := body

	var insta models.InstaNext
	err = json.Unmarshal(jsonBody, &insta)
	if err != nil {
		log.Println("[E] E003", err)
		return nil
	}
	return &insta
}
func getIMGMust(userName, userID, imgID, url string) {

	for try := 0; try <= 3; try++ {
		if err := getIMG(userName, userID, imgID, url); err == nil {
			break
		}
	}
}

func getIMG(userName, userID, imgID, url string) error {

	folder := userName
	if len(folder) == 0 {
		folder = "__by_id"
	}
	os.MkdirAll(DATA_DIR+"/"+folder, 0777)

	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		log.Println("[E] E004", res.StatusCode, err)
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("[E] GET.Read.IMG", res.StatusCode, url, err)
		return err
	}
	defer res.Body.Close()

	f, _ := os.Create(DATA_DIR + "/" + folder + "/" + userName + "-[" + userID + "]-" + imgID + ".jpg")
	defer f.Close()
	f.Write(body)
	f.Close()
	return nil
}
