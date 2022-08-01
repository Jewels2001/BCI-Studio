<script>
import * as Tone from 'tone'

var melodySynth

export default {
    data() {
        return {
            playing: false,
            counter: 0,
            melody_note: 0,
            melody_note_chance: 50,
            key_copy: this.keyT,
            chord_copy: this.chord,
        }
    },
    props: {
        keyT: Array,
        chord: Array,
    },
    watch: {
        keyT: {
            handler(newKey, oldKey) {
                this.key_copy = newKey
            },
            deep: true
        },
        chord: {
            handler(newChord, oldChord) {
                this.chord_copy = newChord
            },
            deep: true
        },
    },
    methods: {
        loop(time) {
            if(this.counter % 2 == 0) {
                if(chance(this.melody_note_chance)) {
                    this.genMelodyNote()
                    if(this.playing) {
                        // console.log("Melody", this.melody_note, numToNote(this.melody_note))
                        melodySynth.triggerAttackRelease(numToNote(this.melody_note), "8n", time)
                    }
                }
            }
            this.counter = (this.counter + 1) % 16
        },

        genMelodyNote() {
            let offset = rand(4) * 2
            this.melody_note = this.key_copy[(this.key_copy.indexOf(this.chord_copy[0]) + offset) % 7] + 72
            if(this.melody_note > 84) {
                this.melody_note -= 12
            }
        },
    },
    beforeMount() {
        melodySynth = new Tone.MonoSynth().toDestination()
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
        <h2>MELODY</h2>
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