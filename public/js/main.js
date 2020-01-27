// console.log('Hello!');

const prizeAudio = document.getElementById('audio-prize');
const clickAudio = document.getElementById('audio-click');

const STARTING_POINTS = 20;

// TODO: timeout for button, reset counter when >=
var app = new Vue({
  el: '#app',
  data: {
    message: '',
    clicksLeft: STARTING_POINTS,
    score: 0,
    nextPrize: '?',
    started: false,
    isFetching: false,
    isDisabled: false,
  },
  methods: {
    fetchData: function(url) {
      const req = new Request(url);
      this.isFetching = true;
      fetch(req)
      .then((res) => {
        return res.json();
      })
      .then((data) => {
        console.log(data);
        this.clicksLeft = data.clicksLeft;
        this.showMessage(data, url);
        this.score = data.score
        if (this.clicksLeft != STARTING_POINTS) {
          this.nextPrize = data.nextPrize;
        }
        this.started = true;
        this.isFetching = false;
      });
    },
    handleClick: function(e) {
      e.preventDefault();
      // TODO: Change this
      if (this.clicksLeft || !started) {
        clickAudio.currentTime = 0;
        clickAudio.play();
        this.isDisabled = true;
        window.setTimeout(() => {
          console.log(this.isDisabled);
          this.isDisabled = false;
        }, 250);
        this.fetchData('/click')
      }
    },
    showMessage: function(data, url) {
      if (this.score != data.score && url === "/click")
      {
        prizeAudio.currentTime = 0;
        prizeAudio.play();
      }
      if (!this.clicksLeft) {
        this.message = `gg! You scored ${data.score} points!`;
      }
      else if (this.score != data.score && url === "/click") {
        this.message = `You won ${data.score-this.score} points!`;
      }
      else {
        this.message = '';
       }
    },
    reset: function(e) {
      e.preventDefault();
      this.fetchData('/reset');
    }
  },
  created: function() {
    // TODO: Change this
    this.fetchData('/state');
  }
})
