package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
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

func TestGetPath(t *testing.T) {
	u := "http://localhost:3000"
	p := path.Join(u, "users", "1", "comments")

	resp, err := http.Get(u + "/" + p)
	if err != nil {
		t.Fatalf("request failed err:%#v", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(byteArray))
}

func TestGetURLParseQuery(t *testing.T) {
	u, err := url.Parse("http://localhost:3000/users")
	if err != nil {
		t.Fatalf("invalid url")
	}
	q := u.Query()
	q.Add("title", "json-server")
	q.Add("author", "typicode")
	u.RawQuery = q.Encode()

	us := u.String()
	t.Logf(us)
	resp, err := http.Get(us)
	if err != nil {
		t.Fatalf("request failed err:%#v", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(byteArray))
}

func TestGetWithHeader(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:3000/posts", nil)
	if err != nil {
		t.Fatalf("invalid request")
	}
	req.Header.Set("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request failed err:%#v", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(byteArray))
}

type Post struct {
	ID     uint
	Title  string
	Author string
}

type Response struct {
	Posts []Post
}

func TestGetJson(t *testing.T) {
	u := "http://localhost:3000/posts"

	resp, err := http.Get(u)
	if err != nil {
		t.Fatal("request failed")
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	posts := make([]Post, 0)
	err = json.Unmarshal(byteArray, &posts)
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("posts: %#v", posts)
}

func TestGetHeader(t *testing.T) {
	u := "http://localhost:3000/posts?_start=20&_end=30"

	resp, err := http.Get(u)
	if err != nil {
		t.Fatal("request failed")
	}
	defer resp.Body.Close()
	t.Logf("resp.Header: %#v", resp.Header.Get("X-Total-Count"))
}

func TestPostJSONByBytes(t *testing.T) {
	u := "http://localhost:3000/posts"

	var reqByteArray = []byte(`{"title":"Refactoring", "author": "some"}`)
	resp, err := http.Post(u, "application/json", bytes.NewBuffer(reqByteArray))
	if err != nil {
		t.Fatal("request failed")
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(byteArray))
}

func TestPostJSONByMap(t *testing.T) {
	u := "http://localhost:3000/posts"

	jsonMap := map[string]string{"title": "Clean Architecture", "author": "fuga"}
	reqByteArray, _ := json.Marshal(jsonMap)
	resp, err := http.Post(u, "application/json", bytes.NewBuffer(reqByteArray))
	if err != nil {
		t.Fatal("request failed")
	}
	defer resp.Body.Close()
	t.Logf("resp.Header: %#v", resp.Header.Get("X-Total-Count"))
}

func TestPostJSONByString(t *testing.T) {
	// TODO
}
