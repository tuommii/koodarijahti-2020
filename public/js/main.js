
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

function toggleAudio() {
	this.isAudio = !this.isAudio
}

function updateState(data) {
	this.points = data.points;
	this.firstTry = data.firstTry;
	// Show clicks to next prize only if button is clicked
	if (!this.firstTry) {
		this.nextPrize = data.nextPrize;
	} else {
		this.nextPrize = '?';
	}
	if (this.points === 0)
		this.isDisabled = true;
	else
		this.isDisabled = false;
}

function handleClick(e) {
	e.preventDefault();
	if (this.points) {
		this.isDisabled = true;
		playClickAudio(this.isAudio);
		window.setTimeout(() => {
			this.isDisabled = false;
		}, DELAY);
		this.fetchData(CLICK);
	}
}

function fetchData(url) {
	if (url === CLICK)
		this.isDisabled = true;
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

function checkGameOver(data, url)
{
	if (!data.points) {
		if (url === CLICK) {
			playGameOverAudio(this.isAudio);
		}
		this.message = `Game Over!`;
		return true;
	}
	return false;
}

function checkGameState(data, url) {
	// if (!data.points) {
	// 	if (url === CLICK) {
	// 		playGameOverAudio(this.isAudio)
	// 	}
	// 	this.message = `Game Over!`;
	// }
	if (checkGameOver(data, url))
		return ;
	if (data.points > this.points && url === CLICK) {
		playPrizeAudio(this.isAudio);
		this.message = `You won ${data.points - this.points + 1} points!`;
	}
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
		isDisabled: false,
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
