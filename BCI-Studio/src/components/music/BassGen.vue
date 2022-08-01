<script>
import * as Tone from 'tone'
import {rand, chance, numToNote} from '../util/music.js'

var bassSynth

export default {
    data() {
        return {
            playing: false,
            counter: 0,
            bass_note: 0,
            bass_note_chance: 25,
            chord_copy: this.chord,
        }
    },
    props: {
        chord: Array,
    },
    watch: {
        chord: {
            handler(newChord, oldChord) {
                this.chord_copy = newChord
            },
            deep: true
        },
    },
    methods: {
        loop(time) {
            if(this.counter % 4 == 0) {
                if(this.counter == 0) {
                    if(this.playing) {
                        this.bass_note = this.chord_copy[0] + 36
                        bassSynth.triggerAttackRelease(numToNote(this.bass_note), "8n", time, 0.6)
                    }
                } else if(chance(this.bass_note_chance)) {
                    if(this.playing) {
                        this.bass_note = this.chord_copy[2] + 36
                        bassSynth.triggerAttackRelease(numToNote(this.bass_note), "8n", time, 0.6)
                    }
                }
            }
            this.counter = (this.counter + 1) % 16
        },
    },
    beforeMount() {
        bassSynth = new Tone.MonoSynth().toDestination()
        this.counter = 0
    }
}
</script>

<template>
    <div>
        <h2>BASS</h2>
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
    /* position: relative; */
    width: 60px;
    height: 34px;
    margin: 0;
    float: right;
}
    
.CircleSwitch input {
  opacity: 0;
  width: 0;
  height: 0;
}

/* slider */
.slider {
  /* position: absolute; */
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
  /* position: absolute; */
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
  background-color: #c38d9e;
}

input:focus + .slider {
  box-shadow: 0 0 1px #c38d9e;
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