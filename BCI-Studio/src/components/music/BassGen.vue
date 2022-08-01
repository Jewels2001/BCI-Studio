<script>
import * as Tone from 'tone'

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

function rand(max) {
    return Math.floor(Math.random() * max)
}

function chance(pct) {
    return rand(100) <= pct
}

function numToNote(num) {
    let note = ""
    if(num % 12 == 0) {
        note = "C"
    } else if(num % 12 == 1) {
        note = 'C#'
    } else if(num % 12 == 2) {
        note = 'D'
    } else if(num % 12 == 3) {
        note = 'D#'
    } else if(num % 12 == 4) {
        note = 'E'
    } else if(num % 12 == 5) {
        note = 'F'
    } else if(num % 12 == 6) {
        note = 'F#'
    } else if(num % 12 == 7) {
        note = 'G'
    } else if(num % 12 == 8) {
        note = 'G#'
    } else if(num % 12 == 9) {
        note = 'A'
    } else if(num % 12 == 10) {
        note = 'A#'
    } else if(num % 12 == 11) {
        note = 'B'
    }

    let octave = Math.floor(num / 12) - 1

    return note + octave
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
    position: relative;
    display: inline-block;
    width: 60px;
    height: 34px;
    margin-top: 116px;
  
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