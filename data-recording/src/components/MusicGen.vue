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

const TEMPO = 60
const SUBDIVISION = 4
const WAIT_TIME = 1.0 / ((TEMPO / 60) * SUBDIVISION)

export default {
    data() {
        return {
            cur_key: -1,
            playing: false,
            chord: [],
            chord_note: 0,
            chord_note_chance: 80,
            voicing: [],
            key: [],
            key_change_chance: 0,
            melody_note: 0,
            melody_note_chance: 20
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
                i = this.rand(12)
            }
            else {
                i = ((i + (this.rand(5)-3)) + 12) % 12
            }
            console.log("KEY", i)
            this.key = KEYS[i]
            this.cur_key = i
        },

        genChord(key) {
            let i = this.rand(7)
            let root = this.key[i]
            let third = this.key[(i+2) % 7]
            let fifth = this.key[(i+4) % 7]
            
            return [root, third, fifth]
        },

        genMelodyNote(key, chord) {
            let offset = this.rand(4) * 2
            let note = key[(key.indexOf(chord[0]) + offset) % 7] + 72
            if(note > 84) {
                note -= 12
            }
            return note
        },

        genVoicing(chord) {
            let voicing = chord.slice()
            for(let i = 0; i < voicing.length; i++) {
                voicing[i] += 48
                if(voicing[i] > 60) {
                    voicing[i] -= 12
                }
            }

            voicing.sort()
            return voicing
        },

        // Helper functions
        rand(max) {
            return Math.floor(Math.random() * max)
        },

        chance(pct) {
            return this.rand(100) <= pct
        },

        sleep(s) {
            return new Promise(resolve => setTimeout(resolve, s*1000));
        },

        numToNote(num) {
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
        },

        async generateMusic() {
            const chordSynth = new Tone.PolySynth(Tone.AMSynth).toDestination();
            const bassSynth = new Tone.MonoSynth().toDestination();
            const melodySynth = new Tone.FMSynth().toDestination();
            this.genKey()
            console.log("Key", this.key)
            while(this.playing) {
                // START OF BAR

                // Try to change key
                if(this.chance(this.key_change_chance)) {
                    this.genKey()
                    this.key_change_chance = 0
                }
                else {
                    this.key_change_chance += 2
                }
                
                // Generate current chord
                this.chord = this.genChord(this.key)
                console.log("Chord", this.chord)
                this.voicing = this.genVoicing(this.chord)
                console.log("Voicing", this.voicing)

                let bass_note = this.chord[0] + 36
                console.log("Bass", bass_note)
                // let bass_vel = 20 + this.rand(20)
                bassSynth.triggerAttackRelease(this.numToNote(bass_note), "8n")

                // Play chord
                console.log("Voicing", [this.numToNote(this.voicing[0]), this.numToNote(this.voicing[1]), this.numToNote(this.voicing[2])])
                // chordSynth.triggerAttackRelease([this.numToNote(this.voicing[0]), this.numToNote(this.voicing[1]), this.numToNote(this.voicing[2])], "8n")
                
                for(let beat = 0; beat < 4; beat++) {
                    for(let sub = 0; sub < SUBDIVISION; sub++) {
                        if(!this.playing) {
                            return
                        }

                        // Try to play melody note
                        if(this.melody_note == 0 || this.chance(this.melody_note_chance)) {
                            this.melody_note = this.genMelodyNote(this.key, this.chord)
                            console.log("Melody", this.melody_note)
                            melodySynth.triggerAttackRelease(this.numToNote(this.melody_note), "8n")
                        }

                        if(this.chance(this.chord_note_chance)) {
                            if(sub == 3) {
                                this.chord_note = 1
                            }
                            else {
                                this.chord_note = sub
                            }
                            chordSynth.triggerAttackRelease(this.numToNote(this.voicing[this.chord_note]), "8n")
                            // midiout.send_message([144, voicing[chord_note], 30 + random.randrange(10)])
                        }

                        // Change chance of next melody note occuring
                        this.melody_note_chance + (this.rand(40)-20)
                        this.melody_note_chance = Math.max(this.melody_note_chance, 0)
                        this.melody_note_chance = Math.min(this.melody_note_chance, 100)

                        // Wait for next tick
                        await this.sleep(WAIT_TIME)
                    }
                }
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