package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func myGetRequest() {
	respGet, err := http.Get("http://localhost:8082/?nick=marker007")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer respGet.Body.Close()

	bodyGet, err := ioutil.ReadAll(respGet.Body)
	if err != nil {
		fmt.Println(err)
	}
	log.Println(string(bodyGet))
}

func myPostFormRequest(c *http.Client, data *url.Values) {
	respPost, err := http.PostForm("http://localhost:8082/", *data)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer respPost.Body.Close()

	bodyPost, err := ioutil.ReadAll(respPost.Body)
	if err != nil {
		fmt.Println(err)
	}
	log.Println(string(bodyPost))
}

func main() {
	client := &http.Client{}
	formData := url.Values{
		"nickname":           {"marker008"},
		"workingHoursPerDay": {"7"},
		"salary":             {"234.1"},
	}
	myGetRequest()
	myPostFormRequest(client, &formData)

}
