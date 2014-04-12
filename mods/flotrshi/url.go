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
var lastLink string
var lastAuthor string
var lastLinkName string
func initurl() {
	fmt.Println("Url Scraper READY!")
	re  = regexp.MustCompile("http\\S*|www.\\S*")
	re2 = regexp.MustCompile("<title>(.*)</title>")
}

func url(msg Message) {
	if msg.Command == MESSAGE {
		if msg.Text == "!last" {
			send(Message{
				Command: MESSAGE,
				Target: msg.Target,
				Text: "Ultimo link (di "+lastAuthor+"): "+lastLink,
			})
			if lastLinkName != "" {
				send(Message{
					Command: MESSAGE,
					Target: msg.Target,
					Text: "Nome pagina: "+lastLinkName,
				})
			}
		}
		if url := re.FindString(msg.Text); url != "" {
			lastLink = url
			lastAuthor = msg.Source.Nickname
			lastLinkName = ""
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
					str = html.UnescapeString(str)
					lastLinkName = str
					send(Message{
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
