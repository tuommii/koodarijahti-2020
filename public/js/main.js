
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
	this.isAudio = !this.isAudio;
}

function updateState(data) {
	this.points = data.points;
	this.nextPrize = data.nextPrize;
	this.isFetching = false;
}

function handleClick(e) {
  e.preventDefault();
  if (this.points) {
    playClickAudio(this.isAudio);
    this.isDisabled = true;
    window.setTimeout(() => {
      this.isDisabled = false;
    }, DELAY);
    this.fetchData(CLICK)
  }
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
      console.log(data);
	  this.showMessage(data, url);
	  this.updateState(data);
    });
}

function showMessage(data, url) {
  if (!data.points) {
    if (url === CLICK) {
		playGameOverAudio(this.isAudio)
    }
    this.message = `Game Over!`;
  }
  else if (data.points > this.points && url === CLICK) {
		playPrizeAudio(this.isAudio);
      this.message = `You won ${data.points - this.points + 1} points!`;
  }
  else {
    this.message = '';
  }
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
	isAudio: true,
    points: 0,
  },
  methods: {
    fetchData: fetchData,
    handleClick: handleClick,
	showMessage: showMessage,
	updateState: updateState,
	toggleAudio: toggleAudio,
    reset: reset
  },
  created: function () {
    // TODO: Change this
    this.fetchData('/state');
  },
});
