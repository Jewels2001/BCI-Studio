<script>
import Queue from './util/queue.js'
// var canvas = document.getElementById('c');
// var ctx = canvas.getContext('2d');

// var randomHeight = Math.round(Math.random() * 100);
// var randomWidth = Math.round(Math.random() * 100);

var W = 400;
var H = 200;

var colors = ['orange', 'blue', 'green', 'red', 'purple'];
var balls = [];
// Math.floor(Math.random() * 20) + (-10)

for(let i=0; i<5; i++) {
  balls.push({
    x: Math.floor(Math.random() * 500) + 1,
    y: Math.floor(Math.random() * 500) + 1,
    vx: Math.floor(Math.random() * 20) + (-10),
    vy: Math.floor(Math.random() * 20) + (-10),
    radius: 5,
    color: colors[Math.floor(Math.random() * 5)],
  });
}
console.log(balls)

// var ball = {
//   x: 100,
//   y: 100,
//   vx: 5,
//   vy: 2,
//   radius: 5,
//   color: 'orange',
// };

// function startBall() {
//   ctx.fillStyle = 'rgba(255,255,255, 0.15)';
//   ctx.fillRect(0, 0, canvas.width, canvas.height);
//   ball.drawBall();
//   ball.x += ball.vx;
//   ball.y += ball.vy;

//   if(ball.y + ball.vy > canvas.height || ball.y + ball.vy < 0) {
//     ball.vy = -ball.vy;
//   }
//   if(ball.x + ball.vx > canvas.width || ball.x + ball.vx < 0) {
//     ball.vx = -ball.vx;
//   }
// }

// ball.drawBall();
// setInterval(startBall, 100); -->

export default {
  props: {
    drawInterval: { default: 3000, type: Number },
    circleBool: false,
    squareBool: false,
  },
  data() {
    return {
      c: null,
      intervalId1: '',
      intervalId2: '',
      avg_hist: new Queue(),
      avg_activity: 0
    };
  },
  mounted() {
    var canvas = document.getElementById('c1');
    var ctx = canvas.getContext('2d');
    this.c = ctx;
    console.log(canvas.height);
    console.log(canvas.width);
    canvas.width = W;
    canvas.height = H;
    this.startDrawing();
  },
  methods: {
    drawBall(abc) {
      this.c.beginPath();
      this.c.arc(abc.x, abc.y, abc.radius, 0, Math.PI * 2, true);
      this.c.closePath();
      this.c.fillStyle = abc.color;
      this.c.fill();
    },
    moveBall(ball) {
      this.c.fillStyle = 'rgba(255,255,255, 0.01)';
      // console.log(this.c.fillStyle);
      this.c.fillRect(0, 0, W, H);
      this.drawBall(ball);
      ball.x += ball.vx;
      ball.y += ball.vy;

      if (ball.y + ball.vy > H || ball.y + ball.vy < 0) {
        ball.vy = -ball.vy;
      }
      if (ball.x + ball.vx > W || ball.x + ball.vx < 0) {
        ball.vx = -ball.vx;
      }
    },
    triggerComponents() {
        if(this.avg_hist.length == 20) {
            this.avg_hist.pop()
        }

        let data = this.$store.getters.getData
        let sum = 0
        for(let i=0; i<5; i++) {
            sum += data.ch0.data[i]
        }
        for(let i=0; i<5; i++) {
            sum += data.ch1.data[i]
        }
        for(let i=0; i<5; i++) {
            sum += data.ch2.data[i]
        }
        for(let i=0; i<5; i++) {
            sum += data.ch3.data[i]
        }
        sum /= 20
        this.avg_hist.push(sum)

        if(this.avg_hist.length == 20) {
            sum = 0
            for(let i = this.avg_hist.head; i<this.avg_hist.head+20; i++) {
                sum += this.avg_hist.elements[i]
            }
            sum /= 20
            this.avg_activity = sum
        }
        // console.log("AVG Activity", this.avg_activity)
        // console.log("AVG Hist", this.avg_hist.elements)
    },

    startDrawing() {
      console.log('HERE');
      let self = this;
      this.intervalId2 = setInterval(
        function () {
          self.triggerComponents();
          balls.push({
            x: Math.floor(Math.random() * 500),
            y: Math.floor(Math.random() * 500),
            vx: this.avg_activity,
            vy: Math.floor(Math.random() * 20) + (-10),
            radius: 5,
            color: colors[Math.floor(Math.random() * 5)],
          });
          // balls.shift();
        },
      5000);
      this.intervalId1 = setInterval(
        function () {
          // console.log(balls.length);
          for(let j=0; j<balls.length; j++) {
            self.moveBall(balls[j]);
          }
        },
      100);
    },
  },
};
</script>

<template>
  <canvas id="c1" class="c1"></canvas>
</template>

<style>
.c1 {
  /* border: 1px solid grey; */
  /*Fill canvas with black by default*/
  background-color: #fff;
  /*  #fa0 is orange  */
}
</style>
