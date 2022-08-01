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
        drumMachineTrigger() {
            if(this.playing) {
                this.playing = false
                loop.stop()
            }
            else {
                this.playing = true
                this.start()
            }
        },

        start() {
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
            loop = new Tone.Loop(this.song, '16n')

            Tone.Transport.bpm.value = 120
            Tone.Transport.start()
            console.log("Starting Drum Machine")
            loop.start(0)
        },

        song(time) {
            if(this.counter % 4 == 0) {
                kickSynth.triggerAttackRelease("C1", "8n", time, 1)
            }
            if(this.counter % 4 != 1) {
                hatSynth.triggerAttackRelease("C4", "32n", time, 0.2)
            }

            this.counter = (this.counter + 1) % 16
        }
    },
    async beforeMount() {
        await Tone.start()
    }
}
</script>

<template>
    <div>
        <h1>DRUM MACHINE</h1>
        <button @click="drumMachineTrigger">{{this.playing ? "Stop" : "Start"}}</button>
        <!-- <button @click="playTestTone">Play Test Tone</button> -->
    </div>
</template>

<style>
</style>