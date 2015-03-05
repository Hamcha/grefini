package main

import (
	"math/rand"
	"strings"
)

var ponies []string = []string{
	"Princess_Luna", "Princess_Celestia", "Twilight", "Applejack", "Rarity",
	"Fluttershy", "Rainbow_Dash", "Pinkie_Pie", "Mi_Amore_Cadenza", "Sweetcream",
	"Aloe", "Lotus", "Spitfire", "Fleetfoot", "Big_Mac", "Scootaloo",
	"Apple_Bloom", "Sweetie_Belle", "Maud", "Lyra", "BonBon", "Nurse_Redheart",
	"Berry_Punch", "Button's_Mon", "Cheerilee", "Queen_Chrysalis", "Colgate",
	"Daring_Do", "Derpy", "Fluffle_Puff", "Gilda", "Littlepip", "Milky_Way",
	"Octavia", "DJ-PON3", "Sunset_Shimmer", "TheGreatAndPowerfulTrixie", "Zecora",
}

var sentences []string = []string{
	"Non piangere, {X}: ci sono qui io",
	"{X}: ti voglio bene e te ne vorrò per sempre...",
	"Non c'è bisogno che ti trattenga, {X}: ci sono qui io ad ascoltarti.",
	"Sfogati pure su di me, {X}, non avere paura.",
	"Qualunque cosa accada, {X}, sarò sempre al tuo fianco.",
	"Coraggio, {X}, domani andrà meglio.",
	"Ti passerà, {X}, vedrai.",
	"Avanti, {X}, abbi fiducia nel futuro!",
	"Sente male, {X}, ma non bisogna perdere la speranza.",
	"Succede a tutti, {X}; ma il tempo guarisce ogni ferita.",
	"Non tenerti tutto dentro, {X}, dimmi cosa ti tormenta.",
	"Non essere così negativo, {X}: dai una possibilità al mondo.",
	"Coraggio, {X}: stasera ci penserò io a consolarti ;)",
}

func tulpa(sid string, msg Message) {
	if msg.Command == MESSAGE {
		if msg.Text == "!tulpa" {
			n := rand.Intn(len(ponies))
			m := rand.Intn(len(sentences))
			sentence := strings.Replace(sentences[m], "{X}", msg.Source.Nickname, 1)
			sentence = customize(sentence, n)
			send(sid, Message{
				Command: MESSAGE,
				Target:  msg.Target,
				Text:    "<" + ponies[n] + ">: " + sentence,
			})
			return
		}
	}
}

func customize(sentence string, i int) string {
	switch ponies[i] {
	case "Maud":
		return "Roccia. Sei una roccia. Grigia. Sei grigia. Come una roccia. Che è quel che sei. Roccia."
	case "Fluffle_Puff":
		return "Prrrrrht."
	case "Gilda":
		return "Non toccarmi con le tue manacce sporche, sfigato!"
	case "Zecora":
		return "Ricordati, anche in questa buia ora / che almeno non vai in giro col fedora."
	default:
		return sentence
	}
}
