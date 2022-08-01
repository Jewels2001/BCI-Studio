<script>
//https://wesbos.com/javascript/15-final-round-of-exercise/85-audio-visualization

var W = 400;
var H = 200;

function ribbon() {
  //location on the canvas
  this.location = {
    x: Math.random() * W,
    y: Math.random() * H,
  };
  //radius - lets make this 0
  this.radius = 10;
  //speed
  this.speed = 3;
  //steering angle in degrees range = 0 to 360
  this.angle = Math.random() * 360;
  //colors
  var r = Math.round(Math.random() * 255);
  var g = Math.round(Math.random() * 255);
  var b = Math.round(Math.random() * 255);
  var a = Math.random();
  this.rgba = 'rgba(' + r + ', ' + g + ', ' + b + ', ' + a + ')';
}

let particles = [];

class Particle {
  constructor(x, y, xSpeed, ySpeed, pColor, size) {
    this.x = x;
    this.y = y;
    this.xSpeed = xSpeed;
    this.ySpeed = ySpeed;
    this.color = pColor;
    this.size = size;
    this.isAlive = true;
    this.trail = [];
    this.trailIndex = 0;
  }

  step() {
    this.trail[this.trailIndex] = createVector(this.x, this.y);
    this.trailIndex++;
    if (this.trailIndex > 10) {
      this.trailIndex = 0;
    }
    this.x += this.xSpeed;
    this.y += this.ySpeed;

    this.ySpeed += gravity;

    if (this.y > height) {
      this.isAlive = false;
    }
  }

  draw() {
    this.drawTrail();
    fill(this.color);
    noStroke();
    rect(this.x, this.y, this.size, this.size);
  }

  drawTrail() {
    let index = 0;

    for (let i = this.trailIndex - 1; i >= 0; i--) {
      const tColor = lerpColor(
        color(this.color),
        endColor,
        index / this.trail.length
      );
      fill(tColor);
      noStroke();
      rect(this.trail[i].x, this.trail[i].y, this.size, this.size);
      index++;
    }

    for (let i = this.trail.length - 1; i >= this.trailIndex; i--) {
      const tColor = lerpColor(
        color(this.color),
        endColor,
        index / this.trail.length
      );
      fill(tColor);
      noStroke();
      rect(this.trail[i].x, this.trail[i].y, this.size, this.size);
      index++;
    }
  }
}

class Firework extends Particle {
  constructor(x, y) {
    super(x, y, random(-2, 2), random(-10, -15), random(colors), 10);
    this.countdown = random(30, 60);
  }

  step() {
    super.step();

    this.countdown--;
    if (this.countdown <= 0) {
      const explosionSize = random(20, 50);
      for (let i = 0; i < explosionSize; i++) {
        const speed = random(5, 10);
        const angle = random(TWO_PI);
        const xSpeed = cos(angle) * speed;
        const ySpeed = sin(angle) * speed;

        particles.push(
          new Particle(this.x, this.y, xSpeed, ySpeed, this.color, 5)
        );
      }
      this.isAlive = false;
    }
  }
}

export default {
  props: {
    drawInterval: { default: 3000, type: Number },
  },
  data() {
    return {
      c: null,
      rectWidth: 50,
      intervalId1: '',
    };
  },
  mounted() {
    var canvas = document.getElementById('canvas1');
    var ctx = canvas.getContext('2d');
    this.c = ctx;
    canvas.width = W;
    canvas.height = H;
    this.startDrawing();
  },
  methods: {
    drawRect() {
      this.c.clearRect(0, 0, 400, 200);
      this.c.beginPath();
      this.c.rect(10, 10, this.rectWidth, 100);
      this.c.stroke();
    },
    addWidth() {
      console.log('HI');
      this.rectWidth += 10;
      this.drawRect();
    },
    subWidth() {
      console.log('SUB');
      this.rectWidth -= 10;
      this.drawRect();
    },
    drawSquare() {
      // ctx.fillStyle = "rgba(0, 0, 0, 0.02)";
      // ctx.fillStyle = "white";
      // ctx.fillRect(0, 0, W, H);
      var a = new ribbon();
      this.c.fillRect(a.location.x, a.location.y, a.radius, a.radius);
      this.c.fillStyle = a.rgba;
      console.log('X', a.location.x);
      console.log('Y', a.location.y);
      console.log('RAD', a.radius);
    },
    drawCircle() {
      var a = new ribbon();
      this.c.beginPath();
      this.c.arc(a.location.x, a.location.y, a.radius, 0, 2 * Math.PI);
      this.c.fillRect(a.location.x, a.location.y, a.radius, a.radius);
      this.c.fillStyle = a.rgba;
      this.c.stroke();
    },
    drawExplode() {
      particles.forEach((p) => {
        p.step();
        p.draw();
      });
      particles = particles.filter((p) => p.isAlive);
    },
    mousePressed() {
      particles.push(new Firework(100, 200));
    },
    startDrawing() {
      console.log('HERE');
      let self = this;
      this.intervalId1 = setInterval(function () {
        self.drawSquare();
      }, 3000);
      this.intervalId1 = setInterval(function () {
        self.drawCircle();
      }, 3000);
    },
  },
};
</script>

<template>
  <p>Hi!</p>
  <canvas id="canvas1"> </canvas>
  <button @click="draw">Draw</button>
  <button @click="mousePressed">Firework</button>
</template>

<style>
#canvas1 {
  border: 1px solid grey;
  /*Fill canvas with black by default*/
  background: #fff;
  /*  #fa0 is orange  */
}
</style>
