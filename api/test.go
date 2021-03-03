package handler

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.URL.String())
	fmt.Println("full:>", r.URL.String())

	spliturl := strings.Split(r.URL.String(), "?")
	webhook := strings.Split(spliturl[1], "~")

	fullurl := "https://discordapp.com/api/webhooks/" + webhook[0]
	fmt.Println("URL:>", fullurl)

	webhook[1] = strings.Replace(webhook[1], "%E2%99%A1", "♡", -1)
	webhook[1] = strings.Replace(webhook[1], "%C3%A2%E2%84%A2%C2%A1", "♡", -1)
	webhook[1] = strings.Replace(webhook[1], "%C3%83%C2%A2%C3%A2%E2%80%9E%C2%A2%C3%82%C2%A1", "♡", -1)
	
	message, failure := url.QueryUnescape(webhook[1])
	if failure != nil {
		return
	}
	
	message = strings.Replace(message, "@", "@ ", -1)
	if strings.HasPrefix(message, "_____") {
		message = strings.Replace(message, "_", " ", -1)
		if r.Method != http.MethodHead {
			return	
		}
	}

	var jsonStr = []byte(`{"content":"[` + hash(r.RemoteAddr) + "] " + strings.TrimSpace(message) + `"}`)
	req, err := http.NewRequest("POST", fullurl, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		panic(err)
	}

	if r.Header.Get("If-Unmodified-Since") == "" {
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
	}
}

func hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	rs := fmt.Sprintf("%X", bs)
	return rs[0:6]
}
