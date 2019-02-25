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

const nodeJSONPath = "./config.json"
const testcasesJSONPath = "./testcast.json"

type CfgParser struct {
    TargetNodes []Ethnode
	TestCases []TestCase
}

type Ethnode struct {
    Host string    `json:"host"`
    User string    `json:"user"`
    Passwd string  `json:"passwd"`
}

type TestCase struct {
	Number string        `json:"number"`
	Category string      `json:"category"`
	Title string         `json:"title"`
	Description string 	 `json:"desc"`
	Commands []string	 `json:"command"`
}

type Command struct {
	String string       `json:"command"`
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


var testTestCases = []TestCase {
	TestCase {
		Number:"1",
		Category:"cate1",
		Title:"title",
		Description:"desc",
		Commands:{
			{Command{"ls"}},
			// Command {String:"ls"},
			// Command {String:"pwd"},
		},
	},
	// TestCase {
	// number:"1",
	// Category:"cate1",
	// Title:"title",
	// Description:"desc",
	// Commands:{
	// "ls",
	// "pwd",
	// {String:"ls"},
			// {String:"pwd"},
		// },
// },
}

func (parser *CfgParser) loadTargetNodes() error {
	
	log.Printf("Loading the targetNodes... [path : %s]\n", nodeJSONPath)
	
    _, err := os.Stat(nodeJSONPath)
    if err == nil {
        log.Printf("file was exsiting [%s] \n", nodeJSONPath)
    } else if os.IsNotExist(err) {
        log.Printf("file not exists\n", nodeJSONPath)
    } else {
        log.Printf("file %s stat error %v", nodeJSONPath, err)
	}

	content, err := ioutil.ReadFile(nodeJSONPath)
	if err != nil {
		log.Fatal(err)
	}

	jsonBytes, err := json.Marshal(testNodes)
	if err != nil {
		panic(err)
	}

	jsonString := string(jsonBytes)
	fmt.Println(jsonString)

	err = json.Unmarshal([]byte(content), &parser.TargetNodes)
	
	if err != nil {
		fmt.Println("Couldn't unmarshal the config...(",err,")")
	}

	log.Println("result", parser.TargetNodes)
	
	
	return nil
}



func (parser *CfgParser) loadTestCases() error {

	log.Println("Loading the test cases... [path : %s]\n", testcasesJSONPath)

	_, err := os.Stat(testcasesJSONPath)
    if err == nil {
        log.Printf("file was exsiting [%s] \n", testcasesJSONPath)
    } else if os.IsNotExist(err) {
        log.Printf("file not exists\n", testcasesJSONPath)
    } else {
        log.Printf("file %s stat error %v", testcasesJSONPath, err)
	}

	content, err := ioutil.ReadFile(testcasesJSONPath)
	if err != nil {
		log.Fatal(err)
	}

	jsonBytes, err := json.Marshal(test)
	if err != nil {
		panic(err)
	}

	jsonString := string(jsonBytes)
	fmt.Println(jsonString)

	err = json.Unmarshal([]byte(content), &parser.TargetNodes)
	
	if err != nil {
		fmt.Println("Couldn't unmarshal the config...(",err,")")
	}

	log.Println("result", parser.TargetNodes)
	
	
	return nil
}


// Load() is
func (parser *CfgParser) Load() bool {
	
	err := parser.loadTargetNodes()
	if err != nil {
		log.Println(err)
		return false
	}

	err = parser.loadTestCases()
	if err != nil {
		log.Println(err)
		return false
	}
	
    return true
}




