// console.log('Hello!');

const audio = document.getElementById('audio-click');

var app = new Vue({
  el: '#app',
  data: {
    message: 'Hello Vue!',
    clicksLeft: 0,
    score: 0,
    nextPrize: 0,
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
        name = Object.keys(data)[0]
        this.clicksLeft = data[name].clicksLeft;
        this.score = data[name].score
        this.nextPrize = data[name].nextPrize;
        console.log(name);
        this.fetched = true;
      });
    },
    handleClick: function() {
      audio.currentTime = 0;
      audio.play();
      this.fetchData('http://localhost:3000/action')
    }
  },
  created: function() {
    this.fetchData('http://localhost:3000/state');
  }
})
