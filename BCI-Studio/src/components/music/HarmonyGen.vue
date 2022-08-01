<script>
import * as Tone from 'tone'
import {rand, chance, numToNote} from '../util/music.js'

var harmonySynth
const KEYS =   [[0, 2, 4, 5, 7, 9, 11],     // C
                [7, 9, 11, 0, 2, 4, 6],     // G
                [2, 4, 6, 7, 9, 11, 1],     // D
                [9, 11, 1, 2, 4, 6, 8],     // A
                [4, 6, 8, 9, 11, 1, 3],     // E
                [11, 1, 3, 4, 6, 8, 10],    // B
                [6, 8, 10, 11, 1, 3, 5],    // Gb
                [1, 3, 5, 6, 8, 10, 0],     // Db
                [8, 10, 0, 1, 3, 5, 7],     // Ab
                [3, 5, 7, 8, 10, 0, 2],     // Eb
                [10, 0, 2, 3, 5, 7, 9],     // Bb
                [5, 7, 9, 10, 0, 2, 4]]     // F

export default {
    data() {
        return {
            playing: false,
            counter: 0,
            chord: [],
            voicing: [],
            key: [],
            cur_key: -1,
            key_change_chance: 0,
        }
    },
    methods: {
        loop(time) {
            if(this.counter == 0) {
                // Try to change key
                if(chance(Math.floor(this.key_change_chance))) {
                    this.genKey()
                    this.key_change_chance = 0
                }
                else {
                    this.key_change_chance += 22
                }

                this.genChord()
                this.genVoicing()

                if(this.playing) {
                    console.log([numToNote(this.voicing[0]), numToNote(this.voicing[1]), numToNote(this.voicing[2])])
                    harmonySynth.triggerAttackRelease([numToNote(this.voicing[0]), numToNote(this.voicing[1]), numToNote(this.voicing[2])], "1n", time, 0.8)
                }
            }

            this.counter = (this.counter + 1) % 16
        },

        // Generator functions
        genKey() {
            let i = 0
            if(this.cur_key == -1) {
                i = rand(12)
            }
            else {
                i = ((i + (rand(5)-3)) + 12) % 12
            }
            console.log("KEY", i)
            this.key = KEYS[i]
            this.cur_key = i
            this.$emit('changeKey', this.key)
        },

        genChord() {
            let i = rand(7)
            let root = this.key[i]
            let third = this.key[(i+2) % 7]
            let fifth = this.key[(i+4) % 7]
            
            this.chord = [root, third, fifth]
            this.$emit('changeChord', this.chord)
        },

        genVoicing() {
            this.voicing = this.chord.slice()
            for(let i = 0; i < this.voicing.length; i++) {
                this.voicing[i] += 48
                if(this.voicing[i] > 60) {
                    this.voicing[i] -= 12
                }
            }

            this.voicing.sort()
            console.log("Voicing", this.voicing)
        },
    },
    beforeMount() {
        harmonySynth = new Tone.PolySynth(Tone.AMSynth).toDestination()
        this.genKey()
        this.genChord()
        this.genVoicing()
        this.counter = 0
    }
}
</script>

<template>
    <div>
        <h2>HARMONY</h2>
        <div>
            <label class="CircleSwitch">
                <input type="checkbox" @change="this.playing = !this.playing">
                <span class="slider round"></span>
            </label>
        </div>
    </div>
</template>

<style>

.CircleSwitch {
    position: relative;
    display: inline-block;
    width: 60px;
    height: 34px;
    /* margin-top: 116px; */
  
}
    
.CircleSwitch input {
  opacity: 0;
  width: 0;
  height: 0;
}

/* slider */
.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #ccc;
  -webkit-transition: .4s;
  transition: .4s;
}

.slider:before {
  position: absolute;
  content: "";
  height: 26px;
  width: 26px;
  left: 4px;
  bottom: 4px;
  background-color: white;
  -webkit-transition: .4s;
  transition: .4s;
}

input:checked + .slider {
  background-color: #2196F3;
}

input:focus + .slider {
  box-shadow: 0 0 1px #2196F3;
}

input:checked + .slider:before {
  -webkit-transform: translateX(26px);
  -ms-transform: translateX(26px);
  transform: translateX(26px);
}

.slider.round {
  border-radius: 34px;
}

.slider.round:before {
  border-radius: 50%;
}

</style>