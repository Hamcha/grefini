package main

import (
	"fmt"
	"strings"
	"os/exec"
)

var sidtab map[string]string
func initmarkov() {
	sidtab = make(map[string]string)
	sidtab["ponychat"] = "ponychat"
	sidtab["espernet"] = "Espernet"
	sidtab["azzurra"]  = "azzurra"
}

func markov(sid string, msg Message) {
	if msg.Command == MESSAGE {
		if len(msg.Text) < 8 { return }
		if msg.Text[0:6] == "!talk " {
			nick := msg.Text[6:]
			if strings.ContainsAny(nick, "; & , . : > < ` # % $ ( ) [ ] { } \" \\ /") {
				send(sid, Message{
					Target: msg.Target,
					Command: MESSAGE,
					Text: "Nice try",
				})
				return
			}
			out, err := exec.Command("zsh", "-c", "cat /znc-logs/"+sidtab[sid]+"_* | grep -i \"] <\"\\*"+nick+"\\> | cut -d \" \" -f3- | sort -R | ./markov -words 10").Output()
			if err != nil {
				fmt.Println(err.Error())
				send(sid, Message{
					Command: MESSAGE,
					Target:  msg.Target,
					Text:    "Something's not right..",
				})
				return
			}
			send(sid, Message{
				Command: MESSAGE,
				Target : msg.Target,
				Text   : "<"+msg.Text[6:]+"> " + string(out),
			})
			return
		}
	}
}
