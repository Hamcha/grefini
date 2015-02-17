package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/url"
)

var hsurl string = "http://tryhaskell.org/eval"

func hs(msg Message) {
	if msg.Command == MESSAGE {
		if len(msg.Text) < 3 { return }
		if msg.Text[0:3] == "hs " {
			resp, err := http.PostForm(hsurl, url.Values{"exp":{msg.Text[3:]}})
			if err == nil {
				defer resp.Body.Close()
				body, _ := ioutil.ReadAll(resp.Body)
				var data struct{
					Success map[string]string
					Error string
				}
				_ = json.Unmarshal(body, &data)
				out := ""
				if data.Error == "" {
					out = "=> "+data.Success["value"]
				} else {
					out = data.Error
				}
				send(Message{
					Command: MESSAGE,
					Target : msg.Target,
					Text   : out,
				})
			}
		}
	}
}
