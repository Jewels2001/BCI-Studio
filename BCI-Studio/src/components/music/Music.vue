<script>
import DrumMachine from './DrumMachine.vue'
import HarmonyGen from './HarmonyGen.vue'
import MelodyGen from './MelodyGen.vue'
import BassGen from './BassGen.vue'

import Queue from '../util/queue.js'

import * as Tone from 'tone'

export default {
    data() {
        return {
            playing: false,
            chord: [],
            key: [],
            cur_tempo: 100,
            loop: null,
            avg_hist: new Queue(),
            avg_activity: 0
        };
    },
    methods: {
        async toggleMusic() {
            console.log("DATA", this.$store.getters.getData)
            if(this.playing) {
                this.playing = false
                this.loop.stop()
            }
            else {
                await Tone.start()
                this.playing = true
                this.startLoops()
            }
        },

        startLoops() {
            this.loop = new Tone.Loop(this.triggerComponents, '16n')

            Tone.Transport.bpm.value = this.cur_tempo
            Tone.Transport.start()
            console.log("Starting Music Generation")
            this.loop.start(0)
        },

        triggerComponents(time) {
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

            this.$refs.drumMachine.loop(time, this.avg_activity)
            this.$refs.harmonyGen.loop(time)
            this.$refs.melodyGen.loop(time, this.avg_activity)
            this.$refs.bassGen.loop(time)
        },
    },
    // async beforeMount() {
        
    // },
    components: { HarmonyGen, DrumMachine, MelodyGen, BassGen },
}
</script>

<template>
    <div>
        <h1>MUSIC</h1>
        <button @click="this.toggleMusic">{{this.playing ? "Stop" : "Start"}}</button>
        <div class="music-components">
            <HarmonyGen
                class="instrument"
                ref="harmonyGen"
                @changeKey="this.key = $event"
                @changeChord="this.chord = $event"
            />
            <DrumMachine
                class="instrument"
                ref="drumMachine"
                @changeTempo="this.cur_tempo = $event"
            />
            <MelodyGen
                class="instrument"
                ref="melodyGen"
                :keyT="this.key"
                :chord="this.chord"
            />
            <BassGen
                class="instrument"
                ref="bassGen"
                :chord="this.chord"
            />
        </div>
    </div>
</template>

<style scoped>
h1{
    color: #41b3a3;
    letter-spacing: 5px;
}
button {
  padding: 0.8rem 2.5rem;
  background: transparent;
  text-align: center;
  font-size: 20px;
  font-family: "Bodoni Moda";
  text-transform: uppercase;
  cursor: pointer;
  border: 1px solid #e27d60;
  color: rgb(31, 31, 31);
  font-weight: bold;
  letter-spacing: 1px;
  position:relative;
}
button:hover {
  background-color: #c38d9e;
}
button:active {
  background-color: #c38d9e;
  transform: translateY(2px);
}
.music-components {
  width: 290px;
  margin: 10px;
  float: center;
  text-align: left;
}

.instrument {
    letter-spacing: 5px;
    float: left;
    font-weight: lighter;
    font-size: 13px;
    text-align: left;
    padding: 10px;
    right: 5px;
    color: #e27d60;
    /* background-color: #41b3a3; */
}

</style>