package main

import (
	"bufio"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io"
	"math/rand"
	"os"
	//"regexp"
	//"strconv"
	"strings"
)

type Person struct {
	name string
}

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

func main() {
	p := Person{}
	fmt.Printf("Hullo. I am AIGOR\n")

	if len(p.name) == 0 {
		fmt.Println("What is your name?")
		p.name = listen()
		fmt.Printf("YOU ARE %v\n", p.name)
	}

	for {
		line := listen()
		reply := getReply(line)
		if len(reply) == 0 {
			fmt.Println("LAME - no REPLY!")
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

func saveKnowledge(wurd string, meaning string) {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()
	key := "aigor:reply:" + wurd
	r, err := c.Do("SET", key, meaning)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}

func getReply(q string) string {
	tokens := strings.Split(q, " ")
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()

	//reply := ""
	for _, str := range tokens {
		rkey := "aigor:reply:" + str
		r, err := redis.Values(c.Do("SMEMBERS", rkey))
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println("NOM:" + strconv.Itoa(index) + " //VAL: " + str)
		//fmt.Println("REDREPLY IS", len(r), "LONG")
		if len(r) > 0 {
			randReply := r[rand.Intn(len(r))]
			return (string(randReply.([]byte)))
			//fmt.Println("RANDREPLY:", string(randReply.([]byte)))
		}

	}
	return ("")
}
