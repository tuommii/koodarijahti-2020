
const prizeAudio = document.getElementById('audio-prize');
const clickAudio = document.getElementById('audio-click');
const gameOverAudio = document.getElementById('audio-game-over');

const DELAY = 50;
const CLICK = '/click';
const STATE = '/state';
const RESET = '/reset';

function playClickAudio(enabled) {
	if (enabled) {
		clickAudio.currentTime = 0;
		clickAudio.play();
	}
}

function playPrizeAudio(enabled) {
	if (enabled) {
		prizeAudio.currentTime = 0;
		prizeAudio.play();
	}
}

function playGameOverAudio(enabled) {
	if (enabled) {
		gameOverAudio.currentTime = 0;
		gameOverAudio.play();
	}
}

function toggleAudio() {
	this.isAudio = !this.isAudio
}

function updateState(data) {
	this.points = data.points;
	this.firstTry = data.firstTry;
	// Show clicks to next prize only if button has been clicked
	this.nextPrize = '?';
	if (!this.firstTry) {
		this.nextPrize = data.nextPrize;
	}
	if (!this.points) {
		this.isClickDisabled = true;
	} else {
		setTimeout(() => {
			this.isClickDisabled = false;
		}, DELAY);
	}
}

function handleClick(e) {
	e.preventDefault();
	if (this.points) {
		this.isClickDisabled = true;
		playClickAudio(this.isAudio);
		this.fetchData(CLICK);
	}
}

function fetchData(url, cb) {
	if (url === CLICK)
		this.isClickDisabled = true;
	const req = new Request(url);
	fetch(req)
		.then((res) => {
			return res.json();
		})
		.then((data) => {
			if (data === undefined)
				return;
			console.log(data);
			this.checkGameState(data, url);
			this.updateState(data);
		});
}

function checkGameState(data, url) {
	// GameOver
	if (!data.points) {
		if (url === CLICK) {
			playGameOverAudio(this.isAudio);
		}
		this.message = `Game Over!`;
	}
	// Prize Won
	if (data.points > this.points && url === CLICK) {
		playPrizeAudio(this.isAudio);
		this.message = `You won ${data.points - this.points + 1} points!`;
	}
	// No prize
	else {
		this.message = '';
	}
}

function reset(e) {
	e.preventDefault();
	this.fetchData(RESET);
}


// TODO: timeout for button, reset counter when >=
var app = new Vue({
	el: '#app',
	data: {
		message: '',
		nextPrize: '?',
		firstTry: true,
		isClickDisabled: false,
		isAudio: true,
		points: 0,
	},
	methods: {
		fetchData: fetchData,
		handleClick: handleClick,
		checkGameState: checkGameState,
		updateState: updateState,
		toggleAudio: toggleAudio,
		reset: reset
	},
	mounted: function () {
		this.fetchData('/state');
	},
});
