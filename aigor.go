package main

import (
	"bufio"
	"fmt"
	//"github.com/garyburd/redigo/redis"
	"io"
	"math/rand"
	"os"
	//"regexp"
	//"strconv"
	"strings"
)

func listen() string {
	bio := bufio.NewReader(os.Stdin)
	line, err := bio.ReadBytes('\n')
	if err == io.EOF {
		os.Exit(0)
	}
	if err != nil {
		panic(err)
	}
	sline := strings.TrimRight(string(line), "\n")
	return sline
}

func talkPerson() {
	p := Person{}
	i := Consciousness{}
	i.name = "AIGOR"
	i.mood = "HAPPY!"

	fmt.Printf("AIGOr:: Hullo. I am " + i.name + "\n" + "Me am " + i.mood + "\n")

	if len(p.name) == 0 {
		fmt.Println("AIGORr:: What is your name?")
		p.name = listen()
		fmt.Printf("AIGOr:: YOU ARE %v\n", p.name)
	}

	for {
		line := listen()
		reply := getReply(line)
		if len(reply) == 0 {
			//fmt.Println("LAME - no REPLY!")
			fmt.Printf("Sorry, i don't know what %v means - can you tell me?\n", line)
			explanation := listen()
			fmt.Printf("Thanks, so \"%v\" means \"%v\" - got it (i think!!)\n", line, explanation)
			saveKnowledge(line, explanation)
		}
		sayName := rand.Intn(4)
		if sayName == 0 {
			reply = p.name + ", " + reply
		}
		fmt.Println(reply)
	}
}

func think() {
	question := "What am I?"
	answer := "you are a process."
	saveKnowledge(question, answer)

}
