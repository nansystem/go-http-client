package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	url := "http://localhost:3000/posts"

	resp, err := http.Get(url)
	if err != nil {
		t.Fatal("request failed")
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(byteArray))
}
