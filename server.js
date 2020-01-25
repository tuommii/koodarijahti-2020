const express = require('express');
const app = express();
var http = require('http').createServer(app);
var io = require('socket.io')(http);

app.use(express.static('public'));

// app.get('/', (req, res) => {
//     res.send('An alligator approaches!');
// });


let clicks = 0;
let players = {

};


io.on('connection', function(socket) {
	console.log(socket.handshake.address ,'connected');
	if (socket.handshake.address === "127.0.0.1")
		return ;
	if (players[socket.handshake.address] === undefined)
	{
		let stats = {
			clicks: clicks,
			clicksLeft: 20,
			points: 0,
		};
		players[socket.handshake.address] = stats;
	}

	socket.on('disconnect', () => {
		console.log('user disconnected', socket.handshake.address);
		console.log(players[socket.handshake.address] === undefined);

	});

	socket.on('clicked', (val) => {
		clicks++;
		players[socket.handshake.address].clicks = clicks;
		players[socket.handshake.address].points = 0;
		io.emit('answer',players[socket.handshake.address]);
		players[socket.handshake.address].clicksLeft--;
		socket.emit('clicksLeft', players[socket.handshake.address].clicksLeft);
	});

});


app.get('/init', function(req, res) {
	res.header("Access-Control-Allow-Origin", "*");
	res.send(players);
});

http.listen(3000, "0.0.0.0", () => console.log('App listening on port 3000!'));
