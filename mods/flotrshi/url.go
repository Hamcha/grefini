package main

import (
	"html"
	"net/http"
	"fmt"
	"regexp"
	"strings"
	"io/ioutil"
)

var re  *regexp.Regexp
var re2 *regexp.Regexp
var lastLink map[string]string
var lastAuthor map[string]string
var lastLinkName map[string]string
func initurl() {
	fmt.Println("Url Scraper READY!")
	re  = regexp.MustCompile("(?i)http\\S*|www.\\S*")
	re2 = regexp.MustCompile("(?i)<title>(.*)</title>")
	lastLink = make(map[string]string)
	lastAuthor = make(map[string]string)
	lastLinkName = make(map[string]string)
}

func urldo(sid string, msg Message) {
	if msg.Command == MESSAGE {
		if msg.Text == "!last" {
			if lastLink[sid] == "" {
				send(sid, Message{
					Command: MESSAGE,
					Target: msg.Target,
					Text: "No links have been posted since I joined",
				})
				return
			}
			send(sid, Message{
				Command: MESSAGE,
				Target: msg.Target,
				Text: "Last posted link (by "+lastAuthor[sid]+"): "+lastLink[sid],
			})
			if lastLinkName[sid] != "" {
				send(sid, Message{
					Command: MESSAGE,
					Target: msg.Target,
					Text: "Link name: "+lastLinkName[sid],
				})
			}
			return
		}
		if url := re.FindString(msg.Text); url != "" {
			lastLink[sid] = url
			lastAuthor[sid] = msg.Source.Nickname
			lastLinkName[sid] = ""
			res, err := http.Get(url)
			if err != nil {	return }
			if head,ok := res.Header["Content-Type"]; ok {
				if len(head) < 1 { return; }
				if strings.Index(head[0],"text/html") >= 0 {
					page, err := ioutil.ReadAll(res.Body)
					res.Body.Close()
					if err != nil { return }
					str := re2.FindString(string(page))
					if len(str) < 5 { return }
					str = str[7:]
					str = strings.Replace(str,"</title>","",-1)
					str = strings.Replace(str,"</TITLE>","",-1)
					str = html.UnescapeString(str)
					lastLinkName[sid] = str
					send(sid, Message{
						Command: MESSAGE,
						Target:  msg.Target,
						Text  :  "^ " + str,
					})
					return
				}
			}
		}
	}
}
