// console.log('Hello!');

const audio = document.getElementById('audio-click');

const STARTING_POINTS = 20;

var app = new Vue({
  el: '#app',
  data: {
    message: '',
    clicksLeft: STARTING_POINTS,
    score: 0,
    nextPrize: '?',
    started: false,
    isFetching: false
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
        if (this.score != data.score && url === "/click")
          this.message = `You won ${data.score-this.score} points!`;
        else
          this.message = '';
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
        audio.currentTime = 0;
        audio.play();
        this.fetchData('/click')
      }
    }
  },
  created: function() {
    // TODO: Change this
    this.fetchData('/state');
  }
})
