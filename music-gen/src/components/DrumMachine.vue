<script>
import * as Tone from 'tone'

var kickSynth
var hatSynth
var loop

export default {
    data() {
        return {
            playing: false,
            counter: 0,
        }
    },
    methods: {
        loop(time) {
            if(this.playing) {
                if(this.counter % 4 == 0) {
                    kickSynth.triggerAttackRelease("C1", "8n", time, 1)
                }
                if(this.counter % 4 != 1) {
                    hatSynth.triggerAttackRelease("C4", "32n", time, 0.2)
                }
            }
            this.counter = (this.counter + 1) % 16
        },
    },
    async beforeMount() {
        kickSynth = new Tone.MembraneSynth().toDestination()
        hatSynth = new Tone.MetalSynth({
            "frequency": 250,
            "envelope": {
                "attack": 0.001,
                "decay": 0.1,
                "release": 0.2
            },
            "harmonicity": 3.1,
            "modulationIndex": 16,
            "resonance": 8000,
            "octaves": 0.5
        }).toDestination()
        this.counter = 0
    }
}
</script>

<template>
    <div>
        <h2>DRUM MACHINE</h2>
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