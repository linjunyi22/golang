package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpTest(url string) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
	}

	cookie := &http.Cookie{Name: "test", Value: "test"}
	req.AddCookie(cookie)

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Charset", "utf-8;q=0.7,*;q=0.3")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	code := resp.StatusCode
	fmt.Println(code)
	if code == 200 {
		r, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(err)
		}

		ioutil.WriteFile("test.html", r, 0644)
		fmt.Println(string(r))
	}

}

func main() {
	url := "http://www.douban.com"
	httpTest(url)
}

