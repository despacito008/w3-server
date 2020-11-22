package ctrl

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

func Substring(source string, start int, end int) string {
	var substring = ""
	var pos = 0
	if end > len(source) {
		end = len(source)
	}
	for _, c := range source {
		if pos < start {
			pos++
			continue
		}
		if pos >= end {
			break
		}
		pos++
		substring += string(c)
	}

	return substring
}

func ReverseList(s []string) []string {
	// sort.Reverse(sort.StringSlice(stringList))
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func GenerateRandStr(len int) string {
	str := ""
	cList := "abcdefghjkmnpqrstuvwxyz123456789"
	for i := 0; i < len; i++ {
		str += string(cList[rand.Intn(32)])
	}
	return str
}

// 判断obj是否在target中
func Contain(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}

	return false
}

func Uint64ToBytes(i uint64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToUint64(buf []byte) uint64 {
	return binary.BigEndian.Uint64(buf)
}

func httpPostJson(targetUrl string, vals string) string {
	jsonStr := []byte(vals)
	req, err := http.NewRequest("POST", targetUrl, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)

}

func httpPostForm(targetUrl string, vals url.Values) string {
	// resp, err := http.PostForm(url, vals)

	// proxyUrl := "http://127.0.0.1:8118"
	// proxy, _ := url.Parse(proxyUrl)
	tr := &http.Transport{
		// Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	timeout := time.Duration(30000 * time.Millisecond) //超时时间50秒
	client := &http.Client{
		Transport: tr,
		Timeout:   timeout,
	}

	request, err := http.NewRequest("POST", targetUrl, strings.NewReader(vals.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(request)

	if err != nil {
		return ""
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	return string(body)

}

func httpGet(auth string, targetUrl string) []byte {
	client := &http.Client{}

	reqest, err := http.NewRequest("GET", targetUrl, nil)

	if auth != "" {
		reqest.Header.Add("Authorization", "token "+auth)
	}

	if err != nil {
		panic(err)
	}
	resp, _ := client.Do(reqest)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return nil
	}

	return body
}
