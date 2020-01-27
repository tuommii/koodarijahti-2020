// console.log('Hello!');

const prizeAudio = document.getElementById('audio-prize');
const clickAudio = document.getElementById('audio-click');

const STARTING_POINTS = 20;

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
    fetchData: function(url) {
      const req = new Request(url);
      this.isFetching = true;
      fetch(req)
      .then((res) => {
        return res.json();
      })
      .then((data) => {
        console.log(data);
        this.showMessage(data, url);
        this.points = data.points;
        // if (this.clicksLeft != STARTING_POINTS) {
        this.nextPrize = data.nextPrize;
        // }
        this.started = true;
        this.isFetching = false;
      });
    },
    handleClick: function(e) {
      e.preventDefault();
      // TODO: Change this
      if (this.points) {
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
      if (this.points != data.points && url === "/click")
      {
      }
      if (!this.points) {
        this.message = `gg!`;
      }
      else if (data.points > this.points && url === "/click") {
        prizeAudio.currentTime = 0;
        prizeAudio.play();
        this.message = `You won ${data.points-this.points+1} points!`;
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
