package main

import (
	"math/rand"
	"fmt"
)

var replies []string = []string { "Direi di sì", "Nah.jpg", "Ne dubitavi?" , "Ahahahahah no.", "Che domande, Certo!", "Ma neanche lontanamente!", "Ma spero veramente di no!", "Ci puoi giurare!" }

func initball() {
	fmt.Println("8ball ready! (trigger: secondo te)")
}

func ball(msg Message) {
	if msg.Command == MESSAGE {
		if len(msg.Text) < 10 { return }
		if msg.Text[0:10] == "secondo te" {
			n := rand.Intn(len(replies)-1)
			send(Message{
				Command: MESSAGE,
				Target : msg.Target,
				Text   : replies[n],
			})
			return
		}
	}
}
