package main

import (
	"fmt"
	"strings"
	"os/exec"
)

func initmarkov() {
	fmt.Println("Markov pronto! (!talk)")
}

func markov(msg Message) {
	if msg.Command == MESSAGE {
		if len(msg.Text) < 8 { return }
		if msg.Text[0:6] == "!talk " {
			nick := msg.Text[6:]
			if strings.ContainsAny(nick, "; & . : > <") {
				send(Message{
					Target: msg.Target,
					Command: MESSAGE,
					Text: "Nice try",
				})
				return
			}
			out, err := exec.Command("zsh", "-c", "cat /usr/home/znc-admin/.znc/users/Hamcha/moddata/log/logs/ponychat_* | grep -i \\<\\*"+nick+"\\> | cut -d \" \" -f3- | sort -R | ./markov -words 10").Output()
			if err != nil {
				fmt.Println(err.Error())
				send(Message{
					Command: MESSAGE,
					Target:  msg.Target,
					Text:    "Something's not right..",
				})
				return
			}
			send(Message{
				Command: MESSAGE,
				Target : msg.Target,
				Text   : "<"+msg.Text[6:]+"> " + string(out),
			})
			return
		}
	}
}
