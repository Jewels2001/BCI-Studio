<script>
import DrumMachine from './DrumMachine.vue'
import HarmonyGen from './HarmonyGen.vue'
import MelodyGen from './MelodyGen.vue'
import BassGen from './BassGen.vue'

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
            this.$refs.drumMachine.loop(time)
            this.$refs.harmonyGen.loop(time)
            this.$refs.melodyGen.loop(time)
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

<style>

.music-components {
    display: flex;    
}

.instrument {
    padding: 5%;
    margin-left: 10px;
    text-align: center;
    background-color: rgb(73, 73, 73);
}

</style>