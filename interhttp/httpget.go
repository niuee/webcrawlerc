package interhttp

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetWebContent(url string) string {
	fmt.Println()
	r, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
