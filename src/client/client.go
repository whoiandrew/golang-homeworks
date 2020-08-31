package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	loginPwdURL   = "http://localhost:8082/loginFromServer"
	loginTokenURL = "http://localhost:8082/loginWithTokenFromServer"
)

func myPostFormRequest(c *http.Client, data *url.Values, loginURL string) {
	respPost, err := http.PostForm(loginURL, *data)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer respPost.Body.Close()

	bodyPost, err := ioutil.ReadAll(respPost.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println(string(bodyPost))
}

func loginPwd(l *string, p *string) {
	fmt.Println("Nickname: ")
	fmt.Scan(l)
	fmt.Println("Password: ")
	fmt.Scan(p)
}

func loginToken(t *string){
	fmt.Println("Token: ")
	fmt.Scan(t)
}



func main() {
	var nick, pwd string
	loginPwd(&nick, &pwd)
	client := &http.Client{}
	formData := url.Values{
		"nickname": {nick},
		"pwd":      {pwd},
	}


	myPostFormRequest(client, &formData, loginPwdURL)
	myPostFormRequest(client, &formData, loginPwdURL)

	var token, choise string

	tokenData := url.Values{
		"token": {token},
	}

	fmt.Println("Would you like to login with token? (y/n):")
	fmt.Scan(&choise)
	switch choise {
	case "y":
		loginToken(&token)
		myPostFormRequest(client, &tokenData, loginTokenURL)
	case "n":
		loginPwd(&nick, &pwd)
		myPostFormRequest(client, &formData, loginPwdURL)
	default:
		fmt.Println("Wrong command")
	}

}
