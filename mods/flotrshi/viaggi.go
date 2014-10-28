package main

import (
	"net/http"
	"fmt"
	"regexp"
	"io/ioutil"
	"net/url"
	"strconv"
	"encoding/json"
)

var viaggiurl string = "http://free.rome2rio.com/api/1.2/json/Search?key=X5JMLHNc&languageCode=IT&currencyCode=EUR"
var reg *regexp.Regexp

func initviaggi() {
	reg = regexp.MustCompile("!go ([^-]+) -> (.+)")
	fmt.Println("Rome2rio loaded! (!go X -> Y)")
}

func viaggi(msg Message) {
	if msg.Command == MESSAGE {
		msgs := reg.FindStringSubmatch(msg.Text)
		if len(msgs) > 2 {
			src := url.QueryEscape(msgs[1])
			dst := url.QueryEscape(msgs[2])
			url := viaggiurl + "&oName="+src+"&dName="+dst
			resp, err := http.Get(url)
			if err == nil {
				defer resp.Body.Close()
				body, _ := ioutil.ReadAll(resp.Body)
				if len(body) < 10 {
					send(Message{
						Command: MESSAGE,
						Target: msg.Target,
						Text: "fak u",
					})
				}
				var outjson Romejson
				json.Unmarshal(body, &outjson)
				var moreeco Romeroute
				var lesstim Romeroute
				if len(outjson.Routes) < 1 {
					send(Message{
						Command: MESSAGE,
						Target: msg.Target,
						Text: "wtf something's not right",
					})
				}
				// Calculate cheapest and fastest
				moreeco = outjson.Routes[0]
				lesstim = outjson.Routes[0]
				for _,v := range outjson.Routes {
					if v.IndicativePrice.Price < moreeco.IndicativePrice.Price {
						moreeco = v
					}
					if v.Duration < lesstim.Duration {
						lesstim = v
					}
				}
				send(Message{
					Command: MESSAGE,
					Target : msg.Target,
					Text   : "Viaggio da "+
							 outjson.Places[0].Name+
							 " a "+outjson.Places[1].Name,
				})
				send(Message{
					Command: MESSAGE,
					Target : msg.Target,
					Text   : "Piu economico: " + moreeco.Name + " (" + parseData(moreeco) + ")",
				})
				send(Message{
					Command: MESSAGE,
					Target : msg.Target,
					Text   : "Piu veloce: " + lesstim.Name + " (" + parseData(lesstim) + ")",
				})
				send(Message{
					Command: MESSAGE,
					Target : msg.Target,
					Text   : "Info: http://www.rome2rio.com/it/s/" + src + "/" + dst,
				})
			}
		}
	}
}

func parseData(route Romeroute) string {
	// Get time
	minutes := int(route.Duration)
	hours := minutes/60
	minutes -= hours*60
	days := hours/24
	hours -= days*24
	timestamp := ""
	if days > 0 {
		timestamp += strconv.Itoa(days)+"d "
	}
	if hours > 0 {
		timestamp += strconv.Itoa(hours)+"h "
	}
	if minutes > 0 {
		timestamp += strconv.Itoa(minutes)+"m"
	}

	return strconv.Itoa(int(route.IndicativePrice.Price)) + " " + route.IndicativePrice.Currency + " - " + strconv.Itoa(int(route.Distance)) + " Km - " + timestamp
}

type Romeplace struct {
	Name string
}

type Romeprice struct {
	Price float64
	Currency string
}

type Romeroute struct {
	Name string
	Distance float64
	Duration float64
	IndicativePrice Romeprice
}

type Romejson struct {
	Places []Romeplace
	Routes []Romeroute
}
