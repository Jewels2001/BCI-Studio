<script>
// var canvas = document.getElementById('c');
// var ctx = canvas.getContext('2d');

// var randomHeight = Math.round(Math.random() * 100);
// var randomWidth = Math.round(Math.random() * 100);

var W = 400;
var H = 200;

var colors = ['orange', 'blue', 'green', 'teal', 'purple'];
var balls = [];

for(let i=0; i<50; i++) {
  balls[i] = {
    x: Math.floor(Math.random() * 500),
    y: Math.floor(Math.random() * 500),
    vx: Math.floor(Math.random() * 20) + (-10),
    vy: Math.floor(Math.random() * 20) + (-10),
    radius: 5,
    color: colors[Math.floor(Math.random() * 5)],
  };
}

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
  },
  data() {
    return {
      c: null,
      intervalId1: '',
    };
  },
  mounted() {
    var canvas = document.getElementById('c1');
    var ctx = canvas.getContext('2d');
    this.c = ctx;
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
      this.c.fillStyle = 'rgba(255,255,255, 0.15)';
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

    startDrawing() {
      console.log('HERE');
      let self = this;
      this.intervalId1 = setInterval(
        function () {
          console.log(balls.length);
          for(let i=0; i<balls.length; i++) {
            self.moveBall(balls[i]);
          }
        }
      , 100);
    },
  },
};
</script>

<template>
  <canvas id="c1"></canvas>
</template>

<style>
#c1 {
  border: 1px solid grey;
  /*Fill canvas with black by default*/
  background: #000;
  /*  #fa0 is orange  */
}
</style>
