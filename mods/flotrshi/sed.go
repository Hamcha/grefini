package main

import (
	"regexp"
	"fmt"
	"strings"
)

var rep *regexp.Regexp
var last map[string]string

func initsed() {
	fmt.Println("sed Module ready!")
	rep = regexp.MustCompile("s/([^/]*)/([^/]*)/")
	last = make(map[string]string)
}

func sed(sid string, msg Message) {
	if msg.Command == MESSAGE {
		if len(msg.Text) > 2 && msg.Text[0:2] == "s/" {
			matches := rep.FindAllStringSubmatch(msg.Text,-1)
			if len(matches) > 0 && len(matches[0])>2 {
				// Compile regexp
				rex, err := regexp.Compile(matches[0][1])
				if err != nil {
					send(sid, Message{
						Command: MESSAGE,
						Target: msg.Target,
						Text: "Regexp error: "+err.Error(),
					})
					return
				}
				// Get target message
				parts := strings.Split(msg.Text," ")
				var lastMsg string
				var nick string
				if val, ok := last[parts[len(parts)-1]]; ok {
					lastMsg = val
					nick = parts[len(parts)-1]
				} else {
					lastMsg = last[msg.Source.Nickname]
					nick = msg.Source.Nickname
				}
				out := rex.ReplaceAllString(lastMsg,matches[0][2])
				send(sid, Message{
					Command: MESSAGE,
					Target: msg.Target,
					Text: "<"+nick+"> "+out,
				})
			}
		} else {
			last[msg.Source.Nickname] = msg.Text
		}
	}
}
