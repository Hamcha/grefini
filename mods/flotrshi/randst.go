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
	err := json.Unmarshal(bytes,&randstX)
	if err != nil {
		fmt.Println("RANDST ERROR: "+err.Error())
	} else {
		fmt.Println("Randst loaded!")
	}
}

const PROBOF = 2
const PROBIN = 50

func randst(msg Message) {
	if (msg.Command == MESSAGE) {
		if rand.Intn(PROBIN) == PROBOF {
			var msgs []string
			if msgx, ok := randstX[msg.Source.Nickname]; ok && rand.Intn(3) < 2 {
				msgs = msgx
			} else {
				msgs = randstX["*"]
			}
			idx := rand.Intn(len(msgs)-1)
			msx := msgs[idx]
			send(Message{
				Command: MESSAGE,
				Target:  msg.Target,
				Text:    msx,
			})
		}
	}
}
