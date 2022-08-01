<script>
import DrumMachine from './DrumMachine.vue'
import HarmonyGen from './HarmonyGen.vue'
import MelodyGen from './MelodyGen.vue'

import * as Tone from 'tone'

export default {
    data() {
        return {
            playing: false,
            chord: [],
            key: [],
            cur_tempo: 100,
            loop: null,
        };
    },
    methods: {
        toggleMusic() {
            if(this.playing) {
                this.playing = false
                this.loop.stop()
            }
            else {
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
            this.$refs.drumMachine.loop(time)
            this.$refs.harmonyGen.loop(time)
            this.$refs.melodyGen.loop(time)
        },
    },
    async beforeMount() {
        await Tone.start()
    },
    components: { HarmonyGen, DrumMachine, MelodyGen},
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
        </div>
    </div>
</template>

<style>

.music-components {
    display: flex;
}

.instrument {
    padding: 5%;
}

</style>