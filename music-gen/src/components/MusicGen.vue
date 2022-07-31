<script>
import * as Tone from 'tone'

const LETTERS = "ABCDEFG"
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

const TEMPO = 80
const SUBDIVISION = 2
const WAIT_TIME = 1.0 / ((TEMPO / 60) * SUBDIVISION)

// Helper functions
function rand(max) {
    return Math.floor(Math.random() * max)
}

function chance(pct) {
    return rand(100) <= pct
}

const sleep = s => new Promise(r => setTimeout(r, s * 1000));

function numToNote(num) {
    if(num % 12 == 0) {
        
    }
}

export default {
    data() {
        return {
            cur_key: -1,
            playing: false,
            chord: [],
            voicing: [],
            key: [],
            key_change_chance: 0,
            melody_note: 0,
            melody_note_chance: 25
        }
    },
    methods: {
        togglePlay() {
            if(this.playing) {
                this.playing = false
            }
            else {
                this.playing = true
                this.generateMusic()
            }
        },

        playTestTone() {
            try {
				const synth = new Tone.PolySynth().toDestination();
                const synth2 = new Tone.PolySynth().toDestination();
                let octave = "4"
                let note = LETTERS.charAt(Math.floor(Math.random()*7))
                let note2 = LETTERS.charAt(Math.floor(Math.random()*7))
                console.log(octave, note, note2)
                synth.triggerAttackRelease(note+octave, "8n")
                synth.triggerAttackRelease(note2+octave, "8n")
            }
            catch(err) {
                console.log(err.message, err.name, err.code)
            }
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
            return KEYS[i], i
        },

        genChord(key) {
            let i = rand(7)
            let root = key[i]
            let third = key[(i+2) % 7]
            let fifth = key[(i+4) % 7]
            
            return [root, third, fifth]
        },

        genMelodyNote(key, chord) {
            let offset = rand(4) * 2
            let note = key[(key.index(chord[0]) + offset) % 7] + 72
            if(note > 84) {
                note -= 12
            }
            return note
        },

        genVoicing(chord) {
            let voicing = chord
            for(let i = 0; i < length(voicing); i++) {
                voicing[i] += 48
                if(voicing[i] > 60) {
                    voicing[i] -= 12
                }
            }

            voicing.sort()
            return voicing
        },

        async generateMusic() {
            const chordSynth = new Tone.PolySynth().toDestination();
            const bassSynth = new Tone.PolySynth().toDestination();
            const melodySynth = new Tone.PolySynth().toDestination();
            while(this.playing) {
                // START OF BAR

                // Try to change key
                if(chance(this.key_change_chance)) {
                    this.key, this.cur_key = genKey()
                    this.key_change_chance = 0
                }
                else {
                    this.key_change_chance += 2
                }
                
                // Generate current chord
                this.chord = genChord(this.key)
                this.voicing = genVoicing(this.chord)
                // print([x % 12 for x in voicing], voicing, chord)

                let bass_note = this.chord[0] + 36
                let bass_vel = 20 + rand(20)
                // midiout.send_message([144, bass_note, bass_vel])
                bassSynth.triggerAttackRelease(numToNote(bass_note), "8n")

                // Play chord
                // midiout.send_message([144, voicing[0], 40 + rand(10)])
                // midiout.send_message([144, voicing[1], 40 + rand(10)])
                // midiout.send_message([144, voicing[2], 40 + rand(10)])
                chordSynth.triggerAttackRelease([numToNote(this.voicing[0]), numToNote(this.voicing[1]), numToNote(this.voicing[2])], "8n")

                for(let beat = 0; beat < 4; beat++) {
                    for(let sub = 0; sub < SUBDIVISION; sub++) {

                        // Try to play melody note
                        if(this.melody_note == 0 || chance(this.melody_note_chance)) {
                            // midiout.send_message([128, melody_note, 0])
                            this.melody_note = genMelodyNote(this.key, this.chord)
                            // print("MELODY", melody_note % 12)
                            // midiout.send_message([144, melody_note, 50 + rand(20)])
                            melodySynth.triggerAttackRelease(numToNote(this.melody_note))
                        }

                        // Change chance of next melody note occuring
                        this.melody_note_chance + (rand(40)-20)
                        this.melody_note_chance = max(this.melody_note_chance, 0)
                        this.melody_note_chance = min(this.melody_note_chance, 100)

                        // Wait for next tick
                        // time.sleep(WAIT_TIME)
                        await sleep(WAIT_TIME)
                        // midiout.send_message([176, 64, 127])
                    }
                }

                // Stop playing chord
                // midiout.send_message([128, voicing[0], 0])
                // midiout.send_message([128, voicing[1], 0])
                // midiout.send_message([128, voicing[2], 0])
                // midiout.send_message([128, bass_note, 0])
                // midiout.send_message([176, 64, 0])
            }
        }
    },
    async beforeMount() {
        await Tone.start()
    }
}




</script>

<template>
    <div>
        <h1>MUSIC</h1>
        <button @click="togglePlay">{{this.playing ? "Stop" : "Start"}}</button>
        <button @click="playTestTone">Play Test Tone</button>
    </div>
</template>

<style>
</style>