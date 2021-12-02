package main

import (
	"fmt"
	"os"

	"github.com/niuee/webcrawlerc/interhttp"
)

func main() {
	url := os.Args[1]

	data := interhttp.GetWebContent(url)

	fmt.Println(data)

}
