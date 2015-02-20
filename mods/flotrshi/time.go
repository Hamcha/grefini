package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"time"
	"strings"
)

var timezones map[string]map[string]string

func inittime(sid string) {
	if timezones == nil {
		timezones = make(map[string]map[string]string)
	}
	bytes, _ := ioutil.ReadFile("time."+sid+".json")
	var timezoneFile map[string]string
	json.Unmarshal(bytes,&timezoneFile)
	timezones[sid] = timezoneFile

	fmt.Println("Loaded timezones for '"+sid+"'")
}

func dotime(sid string, msg Message) {
	if (msg.Command == MESSAGE) {
		parts := strings.Split(msg.Text, " ")
		if parts[0] == "!time" {
			if len(parts) < 2 {
				send(sid, Message{
					Command:MESSAGE, 
					Target:	msg.Target,
					Text:	"fuk u "+msg.Source.Nickname,
				})
				return
			}
			var out string
			var ok bool
			parts[1] = strings.ToLower(parts[1])
			if out, ok = timezones[sid][parts[1]]; !ok {
				out = "No timezone set for that user"
			} else {
				loc, err := time.LoadLocation(out)
				if err != nil {
					out = "The location set for " + parts[1] + "is broken, please re-set"
				} else {
					out = time.Now().In(loc).String()
				}
			}

			send(sid, Message{
				Command:MESSAGE,
				Target: msg.Target,
				Text:	msg.Source.Nickname+": "+out,
			})
			return
		}
		if parts[0] == "!settime" {
			if len(parts) < 3 {
				send(sid, Message{
					Command:MESSAGE,
					Target:	msg.Target,
					Text:	"fuk u "+msg.Source.Nickname,
				})
				return
			}
			name := strings.ToLower(parts[1])
			value := parts[2]
			_, err := time.LoadLocation(value)
			if err != nil {
				send(sid, Message{
					Command:MESSAGE,
					Target: msg.Target,
					Text:   "Invalid timezone name, use a valid IANA TZ from http://bit.ly/1BKGEkq",
				})
				return
			}
			timezones[sid][name] = value
			send(sid, Message{
				Command:MESSAGE,
				Target:	msg.Target,
				Text:	"Timezone set for "+name+"!",
			})
			savetime(sid)
			return
		}
	}
}

func savetime(sid string) {
	bytes,err := json.Marshal(timezones[sid])
	if err != nil { 
		fmt.Printf("CAN'T SAVE TIMEZONES: %s\r\n",err.Error())
		return
	}
	ioutil.WriteFile("time."+sid+".json",bytes, 0777)
}

