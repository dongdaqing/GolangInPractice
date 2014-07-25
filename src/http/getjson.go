package main

import (
	"../github.com/widuu/gojson"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, _ := http.Get("http://test.api.3g.youku.com/shows/z06464d7c434611e18195/next_series?pid=69b81504767483cf&vid=XNDcwNjUxNzcy&guid=9c553730ef5b6c8c542bfd31b5e25b69")
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
	c := gojson.Json(string(body)).Get("status").Tostring()
	c1 := gojson.Json(string(body)).Get("result").Get("showid").Tostring()
	fmt.Println(c)
	fmt.Println(c1)
}
