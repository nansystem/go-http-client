package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestGet(t *testing.T) {
	u := "http://localhost:3000/post"

	resp, err := http.Get(u)
	if err != nil {
		t.Fatal("request failed")
	}
	// 関数を抜ける際に必ずresponseをcloseするようにdeferでcloseを呼ぶ
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(byteArray))
}

func TestGetQueries(t *testing.T) {
	u := "http://localhost:3000/posts"
	v := url.Values{}
	v.Add("title", "json-server")
	v.Add("author", "typicode")

	resp, err := http.Get(u + "?" + v.Encode())
	if err != nil {
		t.Fatal("request failed")
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(byteArray))
}
