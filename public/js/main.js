// console.log('Hello!');

const audio = document.getElementById('audio-click');

var app = new Vue({
  el: '#app',
  data: {
    message: 'Hello Vue!',
    clicksLeft: 0,
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
        if (this.fetched)
          this.nextPrize = data.nextPrize;
        this.fetched = true;
      });
    },
    handleClick: function() {
      audio.currentTime = 0;
      audio.play();
      // TODO: Change this
      this.fetchData('http://192.168.43.68:3000/action')
    }
  },
  created: function() {
    // TODO: Change this
    this.fetchData('http://192.168.43.68:3000/state');
  }
})
