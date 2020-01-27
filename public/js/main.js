// console.log('Hello!');

const prizeAudio = document.getElementById('audio-prize');
const clickAudio = document.getElementById('audio-click');
const gameOverAudio = document.getElementById('audio-game-over');

const STARTING_POINTS = 20;

function handleClick(e) {
  e.preventDefault();
  // TODO: Change this
  if (this.points) {
    clickAudio.currentTime = 0;
    clickAudio.play();
    this.isDisabled = true;
    window.setTimeout(() => {
      this.isDisabled = false;
    }, 250);
    this.fetchData('/click')
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
      if (this.started) {
        this.nextPrize = data.nextPrize;
      }
      this.started = true;
      this.isFetching = false;
    });
}

function showMessage(data, url) {
  if (!data.points) {
    if (url === "/click") {
      gameOverAudio.currentTime = 0;
      gameOverAudio.play();
    }
    this.message = `Game Over!`;
  }
  else if (data.points > this.points) {
    if (url === "/click") {
      prizeAudio.currentTime = 0;
      prizeAudio.play();
    }
    this.message = `You won ${data.points - this.points + 1} points!`;
  }
  else {
    this.message = '';
  }
}

function reset(e) {
  e.preventDefault();
  this.started = false;
  this.fetchData('/reset');
}


// TODO: timeout for button, reset counter when >=
var app = new Vue({
  el: '#app',
  data: {
    message: '',
    nextPrize: '?',
    started: false,
    isFetching: false,
    isDisabled: false,
    points: STARTING_POINTS,
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
  }
})
