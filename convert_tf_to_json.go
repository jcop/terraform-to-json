package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing mandatory param: terraformVars file")
		os.Exit(-1)
	}

	fileContent, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}

	tfJSONString, err := getJSONStr(fileContent)
	if err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}

	fmt.Println(tfJSONString)
}

func getJSONStr(fileContent []byte) (string, error) {

	var out interface{}
	err := hcl.Decode(&out, string(fileContent))
	//fmt.Println(out)

	outData, err := json.Marshal(out)
	if err != nil {
		return "", err
	}
	return string(outData), nil
}
