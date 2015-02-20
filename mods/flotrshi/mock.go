package main

import (
	"os/exec"
	"strings"
	"fmt"
)

func mock(sid string, msg Message) {
	if (msg.Command == MESSAGE) {
		parts := strings.Split(msg.Text, " ")
		if parts[0] == "!mock" {
			if len(parts) < 2 {
				send(sid, Message{
					Command:MESSAGE,
					Target:	msg.Target,
					Text:	"fuk u "+msg.Source.Nickname,
				})
				return
			}
			nick := parts[1]
			origin := msg.Source.Nickname
			if len(parts) > 2 {
				origin = parts[1]
				nick = parts[2]
			}
			if nick == msg.Source.Nickname {
				send(sid, Message{
					Command:MESSAGE,
					Target:msg.Target,
					Text: msg.Source.Nickname+": Can't do yourself!",
				})
				return
			}
			if strings.ContainsAny(nick, " ;&,.:<>`#%$()[]{}\"\\'/") {
				send(sid, Message{
					Command:MESSAGE,
					Target:msg.Target,
					Text: "Nice try",
				})
				return
			}
			bout, err := exec.Command("zsh", "-c", "grep -i \\<"+nick+"\\> /znc-logs/*spernet_* | grep -i "+origin+": | sort -R | head -n 1").Output()
			if err != nil {
				fmt.Println(err.Error())
				send(sid, Message{
					Command: MESSAGE,
					Target: msg.Target,
					Text: "Something's not right..",
				})
				return
			}
			if len(bout) < 1 {
				text := msg.Source.Nickname+": "+nick+" never really fancied "+origin+" apparently"
				if origin == msg.Source.Nickname {
					text = msg.Source.Nickname+": "+nick+" never really fancied you apparently"
				}
				send(sid, Message{
					Command: MESSAGE,
					Target: msg.Target,
					Text: text,
				})
				return
			}
			out := string(bout)
			out = out[strings.Index(out,origin + ":")+len(origin+":"):]
			extra := ""
			if len(parts) > 2 {
				extra = " (to " + origin + ")"
			}
			send(sid, Message{
				Command:MESSAGE,
				Target: msg.Target,
				Text:	msg.Source.Nickname + ": " + nick + " be like \"" + strings.TrimSpace(out) +"\"" + extra,
			})
			return
		}
	}
}
