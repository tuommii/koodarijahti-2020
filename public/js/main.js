

console.log('Hello!');

var socket = io('http://10.12.5.13:3000');


let state = {
	counter: 0,
	points: 20,
	next: 100
};

const req = new Request('http://10.12.5.13:3000/init');
fetch(req)
.then((res) => {
	return res.json();
})
.then((data) => {
	console.log(data);
	state.points = Object.keys(data)[0].clicksLeft;
	state.counter = Object.keys(data)[0].clicks;
	state.points = Object.keys(data)[0].points;
});

socket.on('answer', (e) => {
	console.log(e);
	state.counter = e.clicks;
	state.points = e.clicksLeft;
});

socket.on('clicksLeft', (e) => {
	console.log(e);
	state.points = e;
});

new Vue({
	el: '#app',
	data: state,
	methods: {
		handleClick: function(e) {
			e.preventDefault();
			socket.emit('clicked');
		}
	},
	created: function() {
		const req = new Request('http://10.12.5.13:3000/init');
			fetch(req)
			.then((res) => {
				return res.json();
			})
			.then((data) => {
				console.log(data);
				state.points = Object.keys(data)[0].clicksLeft;
				state.counter = Object.keys(data)[0].clicks;
				state.points = Object.keys(data)[0].points;
			});
	}
});

// new Vue({
// 	el: '#app',
// 	data: {
// 		counter: 0,
// 		message: 'Hello Vue.js!',
// 	},
// 	methods: {
// 		fetchData: function(url) {
// 			const req = new Request(url);
// 			fetch(req)
// 			.then((res) => {
// 				return res.json();
// 			})
// 			.then((data) => {
// 				console.log(this.message);
// 				this.counter = data["Counter"];
// 				// this.count = data.Counter.
// 					// self.data = data;
// 					// self.Counter++;
// 					// a = data["Counter"];
// 					// console.log(this.Counter);
// 					// self.count = data["Counter"];
// 					// this.count++;
// 				});
// 		},
// 		handleClick: function() {
// 			// this.fetchData('http://10.12.5.13:3000/inc');
// 			// socket.send("PASKA!");
// 		}
// 	},
// 	mounted() {
// 		// this.fetchData('http://10.12.5.13:3000/get');
// 	}
// })
