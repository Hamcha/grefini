package main

import (
	"net/http"
	"fmt"
	"regexp"
	"io/ioutil"
	"strconv"
	"encoding/json"
)

var dogebtc string = "http://www.cryptocoincharts.info/v2/api/tradingPair/doge_btc"
var dogefordollars string = "https://www.dogefordollars.com/start.php"
var reg *regexp.Regexp

func initdoge() {
	reg = regexp.MustCompile("<h2>.* DOGE coins")
	fmt.Println("Doge mod loaded! (!doge)")
}

func doge(msg Message) {
	if msg.Command == MESSAGE {
		if msg.Text == "!doge" {
			btcresp, err := http.Get(dogebtc)
			if err == nil { 
				defer btcresp.Body.Close()
				body, _ := ioutil.ReadAll(btcresp.Body)
				var outjson map[string]string
				json.Unmarshal(body, &outjson)
				price   , _ := strconv.ParseFloat(outjson["price"],64)
				price24h, _ := strconv.ParseFloat(outjson["price_before_24h"],64)
				diff := (price/price24h)*100.0-100.0
				sign := ""
				if diff > 0 { sign = "+" }
				send(Message{
					Command: MESSAGE,
					Target : msg.Target,
					Text   : "DOGE/BTC : "+
							 strconv.FormatFloat(price,'f',8,64)+
							 " ("+sign+strconv.FormatFloat(diff,'f',0,64)+"%%%%%%%% since yesterday)",
				})
			}
			dogedol, err := http.Get(dogefordollars)
			if err == nil {
				defer dogedol.Body.Close()
				body, _ := ioutil.ReadAll(dogedol.Body)
				send(Message{
					Command: MESSAGE,
					Target : msg.Target,
					Text   : "Price at dogefordollars.com : "+reg.FindString(string(body))[4:],
				})
			}
		}
	}
}
