// console.log('Hello!');

var app = new Vue({
  el: '#app',
  data: {
    message: 'Hello Vue!',
    clicksLeft: 1,
    total: 2,
    fetched: false
  },
  methods: {
    initData: function(url) {
      const req = new Request(url);
      fetch(req)
      .then((res) => {
        return res.json();
      })
      .then((data) => {
        console.log(data);
        a = Object.keys(data)[0]
        this.clicksLeft = data[a].clicksLeft;;
        console.log(a);
        this.fetched = true;
      });
    },
    handleClick: function() {
      this.initData('http://localhost:3000/inc')
    }
  },
  created: function() {
    this.initData('http://localhost:3000/state');
  }
})
