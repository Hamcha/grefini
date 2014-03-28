fs = require "fs"
net = require "net"

markov = require "markov"

client = net.connect {port:4314}, () ->
	console.log("Connected")

client.on 'data', (data) ->
	try
		msg = JSON.parse(data.toString())
	catch e
		return
	if msg.Command is "PRIVMSG"
		parts = msg.Text.split(" ")
		if parts[0] == "!markov" or parts[0] == "!ask"
			if !parts[1]?
				msg.Text = "fuk u" + msg.Source.Nickname
				client.write(JSON.stringify(msg)+"\r\n")
				return
			if fs.existsSync(__dirname + "/logs/" + parts[1]) == false
				msg.Text = "Non ho log di quell'utente :("
				client.write(JSON.stringify(msg)+"\r\n")
				return
			s = fs.createReadStream(__dirname + "/logs/" + parts[1])
			limit = 80
			grd = 1
			if parts[2]?
				grd = parseInt(parts[2])
				if isNaN(grd) then grd = 1
			m = markov grd
			m.seed s, () ->
				wordn = 0
				tries = 0
				while wordn < 6 and tries < 10
					tries += 1
					if parts[0] == "!ask"
						out = m.respond(parts.splice(2).join " ").join " "
					else
						key = m.pick()
						out = m.backward(key,limit).join(" ")+" "+m.forward(key,limit).join(" ")
					wordn = out.split(" ").length
				#out = m.fill(key,limit).join(" ")
				msg.Text = "<"+parts[1]+"> "+out
				client.write(JSON.stringify(msg)+"\r\n")
