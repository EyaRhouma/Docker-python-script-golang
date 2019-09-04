package main

import "fmt"
import "net/http"
import "io/ioutil"

func main() {
	resp, err := http.Get("http://localhost:8000/solve")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}
