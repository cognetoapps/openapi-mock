package data

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const LOREMBASEURL = "https://loripsum.net/api/1"
var LOREMPATHS = map[string]string{
	"plain": "plaintext",
	"full":  "decorate/links/ul/ol/dl/dq/headers",
}
const CACHESIZE = 10

var cache = make(map[string][]string)
var initialized = false

func loripsum(length, typePath string) string {
	cacheKey := fmt.Sprintf("%s-%s", length, typePath)
	if len(cache[cacheKey]) >= CACHESIZE {
		return cachedResponse(cacheKey)
	} else {
		lorem := get(length, typePath)
		cache[cacheKey] = append(cache[cacheKey], lorem)
		return lorem
	}
}

func get(length, typeName string) string {
	url := fmt.Sprintf("%s/%s/%s", LOREMBASEURL, length, LOREMPATHS[typeName])
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		panic(errRead)
	}
	trimmed := strings.Trim(strings.TrimSpace(string(body)), "\n")
	return strings.ReplaceAll(trimmed , "\t", "")
}

func cachedResponse(cacheKey string) string {
	if !initialized {
		rand.Seed(time.Now().UnixNano())
		initialized = true
	}
	return cache[cacheKey][rand.Intn(CACHESIZE)]
}
