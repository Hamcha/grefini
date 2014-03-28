from chatterbotapi import ChatterBotFactory, ChatterBotType
import json
import socket

# Create Cleverbot instance
factory = ChatterBotFactory()
bot = factory.create(ChatterBotType.CLEVERBOT)
botsession = bot.create_session()

# Create TCP connection

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect(("127.0.0.1",4314))
print "Connected to proxy!"
while (1):
	data = s.recv(1024)
	if not data: break
	msg = ""
	try:
		msg = json.loads(data.decode('utf-8'))
	except:
		continue
	if msg["Command"] == "PRIVMSG":
		if msg["Text"].lower().startswith("grefini:"):
			if len(msg["Text"]) < 9: 
				continue
			res = botsession.think(msg["Text"][8:].strip())
			msg["Text"] = msg["Source"]["Nickname"] + ": " + res
			out = json.dumps(msg)
			s.sendall(out.encode('utf-8')+"\r\n")
s.close()
