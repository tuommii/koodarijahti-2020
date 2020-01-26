// console.log('Hello!');

const audio = document.getElementById('audio-click');

const STARTING_POINTS = 20;

var app = new Vue({
  el: '#app',
  data: {
    message: 'Hello Vue!',
    clicksLeft: STARTING_POINTS,
    score: 0,
    nextPrize: '?',
    fetched: false
  },
  methods: {
    fetchData: function(url) {
      const req = new Request(url);
      fetch(req)
      .then((res) => {
        return res.json();
      })
      .then((data) => {
        console.log(data);
        this.clicksLeft = data.clicksLeft;
        this.score = data.score
        if (this.clicksLeft != STARTING_POINTS)
          this.nextPrize = data.nextPrize;
        this.fetched = true;
      });
    },
    handleClick: function() {
      audio.currentTime = 0;
      audio.play();
      // TODO: Change this
      this.fetchData('/action')
    }
  },
  created: function() {
    // TODO: Change this
    this.fetchData('/state');
  }
})
