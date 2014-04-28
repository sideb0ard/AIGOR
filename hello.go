package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type Conch struct {
  name string
  }

func main() {
        //c:= Conch{}
	fmt.Printf("Hullo. I am AIGOR\n")
	bio := bufio.NewReader(os.Stdin)
	for {
		line, err := bio.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		sline := strings.TrimRight(string(line), "\n")
		// fmt.Printf(sline + "\n")
		reply := getAnswer(sline)
		//fmt.Println("Default is ", c)
		fmt.Println(reply)
	}
}

func getAnswer(q string) string {
	r, _ := regexp.Compile(`^help$`)
	rn, _ := regexp.Compile(`my name is (.*)`)
	//fmt.Printf("In Answer section, i gots " + q)
	if r.MatchString(q) == true {
		return ("I AM AIGOR. NO HELP REQUIRED\n")
	} else if rn.MatchString(q) {
                nom := rn.FindStringSubmatch(q)
                sentence := "i know yer name " + string(nom[1])
                return (sentence)
        } else {
		return "Ima tha answer\n"
	}
}
