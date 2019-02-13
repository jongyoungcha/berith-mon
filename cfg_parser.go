package main

import (
    "fmt"
	"encoding/json"
	"io/ioutil"
    // "io"
    "log"
    "os"
    // "strings"
)

const cfgFilePath = "./config.json"
// const sampleData = string(`"{EthNode : {"host" : "192.168.56.102:22", "user" : "root", "passwd" : "jongyoungcha" }`)

type CfgParser struct {
    Nodes []Ethnode
}

type Ethnode struct {
    Host string    `json:"host"`
    User string    `json:"user"`
    Passwd string  `json:"passwd"`
}


var testNodes = []Ethnode {
	Ethnode {
		Host:"192.168.56.102:22",
		User:"root",
		Passwd:"jongyoungcha",
	},
	Ethnode {
		Host:"192.168.56.103:22",
		User:"root",
		Passwd:"jongyoungcha",
	},
}


 // Load() is
func (parser *CfgParser) Load() bool {
	
    _, err := os.Stat(cfgFilePath)
    if err == nil {
        log.Printf("file was exsiting\n", cfgFilePath)
    } else if os.IsNotExist(err) {
        log.Printf("file not exists\n", cfgFilePath)
    } else {
        log.Printf("file %s stat error %v", cfgFilePath, err)
	}

	content, err := ioutil.ReadFile(cfgFilePath)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("File contents: %s", content)

	err = json.Unmarshal([]byte(content), &parser.Nodes)
	
	// err = json.Unmarshal([]byte(sampleData), &ethnodes)
	if err != nil {
		fmt.Println("Couldn't unmarshal the config...(",err,")")
	}

	log.Println("result", parser.Nodes)
	
    return err == nil
}



