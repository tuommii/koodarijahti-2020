
const prizeAudio = document.getElementById('audio-prize');
const clickAudio = document.getElementById('audio-click');
const gameOverAudio = document.getElementById('audio-game-over');

// const 0 = 20;
const DELAY = 200;
const CLICK = '/click';
const STATE = '/state';
const RESET = '/reset';

function handleClick(e) {
  e.preventDefault();
  if (this.points) {
    clickAudio.currentTime = 0;
    clickAudio.play();
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
      this.points = data.points;
      this.nextPrize = data.nextPrize;
      this.isFetching = false;
    });
}

function showMessage(data, url) {
  if (!data.points) {
    if (url === CLICK) {
      gameOverAudio.currentTime = 0;
      gameOverAudio.play();
    }
    this.message = `Game Over!`;
  }
  else if (data.points > this.points && url === CLICK) {
      prizeAudio.currentTime = 0;
      prizeAudio.play();
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
    isFetching: false,
    isDisabled: false,
    points: 0,
  },
  methods: {
    fetchData: fetchData,
    handleClick: handleClick,
    showMessage: showMessage,
    reset: reset
  },
  created: function () {
    // TODO: Change this
    this.fetchData('/state');
  },
});
