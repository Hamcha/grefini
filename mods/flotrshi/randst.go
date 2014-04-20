package main

import (
	"math/rand"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

var randstX map[string][]string

func initrandst() {
	bytes, _ := ioutil.ReadFile("s.json")
	json.Unmarshal(bytes,&randstX)

	fmt.Println("Randst loaded!")
}

const PROBOF = 2
const PROBIN = 10

func randst(msg Message) {
	if (msg.Command == MESSAGE) {
		if msgs, ok := randstX[msg.Source.Nickname]; ok {
			if rand.Intn(PROBIN) == PROBOF {
				idx := rand.Intn(len(msgs))
				msx := msgs[idx]
				send(Message{
					Command: MESSAGE,
					Target:  msg.Target,
					Text:    msx,
				})
			}
		}
	}
}
