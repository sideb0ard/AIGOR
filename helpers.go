package main

import (
	"bytes"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"regexp"
	"strings"
)

func stringify(stingz []string) string {
	var unifiedString bytes.Buffer
	for i := range stingz {
		if i != 0 {
			unifiedString.WriteString(" ")
		}
		unifiedString.WriteString(stingz[i])
	}
	return unifiedString.String()
}
func spaceify(dashy string) string {
	spxex, _ := regexp.Compile("-")
	spaceyString := strings.ToLower(spxex.ReplaceAllString(dashy, " "))
	//fmt.Println("SPACEME CALLED! Sending back: " + spacyString)
	return spaceyString
}
func dashify(spacey string) string {
	spxex, _ := regexp.Compile(" ")
	dashyString := strings.ToLower(spxex.ReplaceAllString(spacey, "-"))
	//fmt.Println("DASHME CALLED! Sending back: " + dashyString)
	return dashyString
}
func saveKnowledge(thing string, meaning string) {
	meaning = dashify(meaning)
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()
	key := "aigor:memory:" + thing
	_, err = c.Do("SET", key, meaning)
	if err != nil {
		fmt.Println(err)
	}
}
func getReply(q string) string {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()

	rkey := "aigor:memory:" + q
	r, err := redis.String(c.Do("GET", rkey))
	if err != nil {
		fmt.Println(err)
	}
	if len(r) > 0 {
		return spaceify(r)
	} else {
		return ""
	}
}
