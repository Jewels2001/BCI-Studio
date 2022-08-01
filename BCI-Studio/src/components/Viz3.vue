<script>
const PI = Math.PI;
const TAU = PI * 2;

function rand(min, max) {
	if (!max) {
		var max = min;
		min = 0;
	}
	return Math.random() * (max - min) + min;
};

// $.events = () => {
// 	window.addEventListener( 'resize', $.reset.bind( this ) );	
// };

export default {
    data() {
        return {
            canvas: null,
            context: null,
            simplex: null,
            dpr: null,
            w: null,
            h: null,
            cx: null,
            cy: null,
            count: null,
            xoff: null,
            xinc: null,
            yoff: null,
            goff: null,
            ginc: null,
            y: null,
            length: null,
            amp: null,
        }
    },
    methods: {
        reset() {
            this.dpr = window.devicePixelRatio;
            this.w = window.innerWidth;
            this.h = window.innerHeight;
            this.cx = this.w / 2;
            this.cy = this.h / 2;
            this.canvas.width = this.w * this.dpr;
            this.canvas.height = this.h * this.dpr;
            this.canvas.style.width = `this{this.w}px`;
            this.canvas.style.height = `this{this.h}px`;
            this.context.scale(this.dpr, this.dpr);
            
            this.count = Math.floor( this.w / 50 );
            this.xoff = 0;
            this.xinc = 0.05;	
            this.yoff = 0;
            this.yinc = 0.003;
            this.goff = 0;
            this.ginc = 0.003;
            this.y = this.h * 0.66;
            this.length = this.w + 10;
            this.amp = 40;
        },

        wave() {
            this.context.beginPath();
            let sway = this.simplex.noise2D( this.goff, 0 ) * this.amp;
            for( let i = 0; i <= this.count; i++ ) {
                this.xoff += this.xinc;
                let x = this.cx - this.length / 2 + ( this.length / this.count ) * i;
                let y = this.y + this.simplex.noise2D( this.xoff, this.yoff ) * this.amp + sway;
                this.context[ i === 0 ? 'moveTo' : 'lineTo' ]( x, y );
            }
            this.context.lineTo( this.w, this.h );
            this.context.lineTo( 0, this.h );
            this.context.closePath();
            this.context.fillStyle = 'hsla(210, 90%, 50%, 0.2)';
            this.context.fill();
        },

        loop() {
            requestAnimationFrame( this.loop );
            this.context.clearRect( 0, 0, this.w, this.h );
            this.xoff = 0;
            this.wave();
            this.wave();
            this.wave();
            this.wave();
            this.yoff += this.yinc;
            this.goff += this.ginc;
        },
    },
    mounted() {
        this.canvas = this.$refs.cw;
        this.context = this.canvas.getContext('2d');
        this.simplex = new SimplexNoise();
        // this.events();
        this.reset();
        this.loop();
    }
}

</script>

<template>
    <canvas ref="cw" @resize="this.reset()"></canvas>
</template>

<style>
html,
body {
	height: 100%;	
}

body {
	background: radial-gradient( circle cover, hsl( 210, 90%, 90% ), hsl( 210, 90%, 70% ) );
	overflow: hidden;
}

canvas {
	display: block;	
}
</style>