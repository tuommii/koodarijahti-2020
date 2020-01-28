
const prizeAudio = document.getElementById('audio-prize');
const clickAudio = document.getElementById('audio-click');
const gameOverAudio = document.getElementById('audio-game-over');

const DELAY = 200;
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

function updateState(data) {
	this.points = data.points;
	this.nextPrize = data.nextPrize;
	this.isFetching = false;
}

function handleClick(e) {
	e.preventDefault();
	if (this.points) {
		playClickAudio(this.isAudioEnabled);
		this.isDisabled = true;
		window.setTimeout(() => {
			this.isDisabled = false;
		}, DELAY);
		this.fetchData(CLICK)
	}
}

function checkGameOver(data, url, self) {
	if (!data.points) {
		// Play audio only after click
		if (url === CLICK) {
			playGameOverAudio();
		}
		self.message = `Game Over!`;
		return (1);
	}
	return (0);
}

function checkPrizeWon(data, url, points) {
	let message = `You won ${data.points - points + 1} points!`;
	if (data.points > points && url === CLICK) {
		playPrizeAudio();
		// this.message = `You won ${data.points - this.points + 1} points!`;
	}
	else {
		message = '';
	}
	return message
}

function checkGameState(data, url, self) {
	if (checkGameOver(data, url, self))
		return;
	self.message = checkPrizeWon(data, url, self.points);
}

function fetchData(url) {
	const req = new Request(url);
	this.isFetching = true;
	fetch(req)
		.then((res) => {
			return res.json();
		})
		.then((data) => {
			if (data === undefined)
				return;
			this.checkGameState(data, url, this);
			this.updateState(data);
		})
		.catch((error) => {
			console.log(error);
		});
}

function reset(e) {
	e.preventDefault();
	this.fetchData(RESET);
}

var app = new Vue({
	el: '#app',
	data: {
		message: '',
		nextPrize: '?',
		isFetching: false,
		isDisabled: false,
		isAudioEnabled: true,
		points: 0,
	},
	methods: {
		fetchData: fetchData,
		handleClick: handleClick,
		checkGameState: checkGameState,
		updateState: updateState,
		reset: reset
	},
	created: function () {
		// TODO: Change this
		this.fetchData('/state');
	},
});
