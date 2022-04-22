package main

import (
	"fmt"
	"os"
	"io/ioutil"
	
	"gopkg.in/yaml.v3"
)


func main() {
	var p map[string]interface{}

	file, err := ioutil.ReadFile("./dict.yaml")

	err = yaml.Unmarshal(file, &p)
	if err != nil {
		fmt.Println(err)
		
	}
	if len(os.Args) <= 1 {
		fmt.Println(`---no arguments provided---`)
		os.Exit(1)
	}
	if val, ok := p[os.Args[1]]; ok {
		fmt.Println(val)
	} else {
		fmt.Println(`---repo does not exist ---`)
	}
	
	
}
