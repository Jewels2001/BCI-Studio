<script>
import Queue from './util/queue.js'

export default {
    data() {
        return {
            canvas: null,
            context: null,
            cursor: null,
            particlesArray: [],
            avg_hist: new Queue(),
            avg_activity: 0
        }
    },
    methods: {
        generateParticles(amount) {
            for (let i = 0; i < amount; i++) {
                this.particlesArray[i] = this.Particle(
                    innerWidth / 2,
                    innerHeight / 2,
                    4,
                    this.generateColor(),
                    0.02
                );
            }
        },

        generateColor() {
            let hexSet = "0123456789ABCDEF";
            let finalHexString = "#";
            for (let i = 0; i < 6; i++) {
                finalHexString += hexSet[Math.ceil(Math.random() * 15)];
            }
            return finalHexString;
        },

        setSize() {
            this.canvas.height = innerHeight;
            this.canvas.width = innerWidth;
        },

        Particle(x, y, particleTrailWidth, strokeColor, rotateSpeed) {
            var p = {}
            p.x = x;
            p.y = y;
            p.particleTrailWidth = particleTrailWidth;
            p.strokeColor = strokeColor;
            p.theta = Math.random() * Math.PI * 2;
            p.rotateSpeed = rotateSpeed;
            p.t = Math.random() * 150;

            p.rotate = () => {
                const ls = {
                x: p.x,
                y: p.y,
                };
                p.theta += p.rotateSpeed;
                p.x = this.cursor.x + Math.cos(p.theta) * p.t;
                p.y = this.cursor.y + Math.sin(p.theta) * p.t;
                this.context.beginPath();
                this.context.lineWidth = p.particleTrailWidth;
                this.context.strokeStyle = p.strokeColor;
                this.context.moveTo(ls.x, ls.y);
                this.context.lineTo(p.x, p.y);
                this.context.stroke();
            };

            return p
        },

        anim() {
            window.requestAnimationFrame(this.anim);
            this.triggerComponents();


            this.context.fillStyle = "rgba(0,0,0,0.05)";
            this.context.fillRect(0, 0, this.canvas.width, this.canvas.height);

            this.particlesArray.forEach((particle) => {
              particle.rotateSpeed = this.avg_activity;
              particle.rotate()
            });
        },

        cursorUpdate(event) {
          this.cursor.x = event.clientX;
          this.cursor.y = event.clientY;
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
        // updateData() {
        //   let self = this;
        //   setInterval(
        //     function () {
        //       self.triggerComponents();
        //     },
        //   100);
        // }
    },
    mounted() {
        this.canvas = this.$refs.cw
        this.context = this.canvas.getContext("2d");
        this.context.globalAlpha = 0.5;
        this.cursor = {
            x: innerWidth / 2,
            y: innerHeight / 2,
        };
        this.generateParticles(101);
        this.setSize();
        this.anim();
        // this.updateData();

    }
}

</script>

<template>
    <canvas ref="cw" @mousemove="this.cursorUpdate"></canvas>
</template>

<style>

#cw {
  position: fixed;
  z-index: -1;
}

body {
  margin: 0;
  padding: 0;
  background-color: rgba(0, 0, 0, 0.05);
}

</style>
