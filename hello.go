package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

type Person struct {
	name string
}

type Consciousness struct {
	name string
	mood string
}

type Thing struct {
	name          string
	thingType     string
	properties    []interface{}
	relationships []interface{}
	memories      []interface{}
}

func main() {
	fmt.Println("Maiiiiin, man")
	//innit()
	think()
	talkPerson()
}

func innit() {
	file, err := os.Open("language.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewReader(file)
	for {
		line, err := scanner.ReadBytes('\n')
		if err == io.EOF {
			//os.Exit(0)
			break
		}
		if err != nil {
			panic(err)
		}
		if match, _ := regexp.Match(`^#`, line); match == true {
			continue
		}
		sline := strings.Split(strings.TrimRight(string(line), "\n"), "|")
		entryType, keyWord, replacement := sline[0], sline[1], stringify(sline[2:])
		storageKey := entryType + ":" + keyWord
		fmt.Println(storageKey, replacement)
		saveKnowledge(storageKey, replacement)
	}
}
