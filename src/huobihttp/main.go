package main

import "net/http"
//f52a9b23-5e3ab8a3-f6d1584f-a5962
//5ec99983-02d1ed81-0937e76b-0ec14
func main() {

	res, err := http.Get("http://www.baidu.com")
	if err != nil {
		print(err)
	}
	defer res.Body.Close()
	println(res.Body)
}
