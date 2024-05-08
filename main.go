package main

import (
	"XiaoMi/crossplane"
	"encoding/json"
	"fmt"
)

func main() {
	path := "./nginx.conf"
	//path := []string{"./nginx.conf", "./nginx.conf"}

	payload, err := crossplane.Parse(path, &crossplane.ParseOptions{})

	if err != nil {
		fmt.Println(err)
	}

	b, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))
}



