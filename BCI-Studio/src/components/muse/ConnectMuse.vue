<script>
import { zipSamples, channelNames, MuseClient } from 'muse-js';
import { powerByBand, epoch, fft } from "@neurosity/pipes";
import DataBandDisplay from './DataBandDisplay.vue';

export default {
    data() {
        return {
            connected: false,
            recording: false,
            recorded: false,
            samples: 0,
            client: null,
            files: [],
            selectedFile: null,
            ch0: {
                labels: ["Delta", "Theta", "Alpha", "Beta", "Gamma"],
                data: [0, 0, 0, 0, 0],
                hist: [[]]
            },
            ch1: {
                labels: ["Delta", "Theta", "Alpha", "Beta", "Gamma"],
                data: [0, 0, 0, 0, 0],
                hist: [[]]
            },
            ch2: {
                labels: ["Delta", "Theta", "Alpha", "Beta", "Gamma"],
                data: [0, 0, 0, 0, 0],
                hist: [[]]
            },
            ch3: {
                labels: ["Delta", "Theta", "Alpha", "Beta", "Gamma"],
                data: [0, 0, 0, 0, 0],
                hist: [[]]
            }
        };
    },
    methods: {
        async connect(event) {
            this.client = new MuseClient();
            this.client.connectionStatus.subscribe((status) => {
                this.connected = status ? "Connected!" : "Disconnected",
                    console.log(status ? "Connected!" : "Disconnected");
            });
            try {
                await this.client.connect();
                await this.client.start();
                zipSamples(this.client.eegReadings).pipe(
                    epoch({ duration: 1024, interval: 100, samplingRate: 256 }),
                    fft({ bins: 256 }),
                    powerByBand()
                ).subscribe((data) => {
                    this.ch0.data = [data.delta[0], data.theta[0], data.alpha[0], data.beta[0], data.gamma[0]];
                    this.ch1.data = [data.delta[1], data.theta[1], data.alpha[1], data.beta[1], data.gamma[1]];
                    this.ch2.data = [data.delta[2], data.theta[2], data.alpha[2], data.beta[2], data.gamma[2]];
                    this.ch3.data = [data.delta[3], data.theta[3], data.alpha[3], data.beta[3], data.gamma[3]];
                    // console.log(data);

                    this.$store.commit('updateData', {ch0: this.ch0, ch1: this.ch1, ch2: this.ch2, ch3: this.ch3})

                    if(this.recording) {
                        this.ch0.hist.push(this.ch0.data)
                        this.ch1.hist.push(this.ch1.data)
                        this.ch2.hist.push(this.ch2.data)
                        this.ch3.hist.push(this.ch3.data)
                        this.samples++
                    }
                });
            }
            catch (err) {
                console.error("Connection failed", err);
            }
        },

        toggleRecord(event) {
            if(this.client != null) {
                if(this.recording) {
                    this.recording = false
                    this.recorded = true
                }
                else {
                    this.ch0.hist = [[]]
                    this.ch1.hist = [[]]
                    this.ch2.hist = [[]]
                    this.ch3.hist = [[]]
                    this.recording = true
                    this.samples = 0
                }
            }
            else {
                console.log("Recording failed, no connection")
            }
        },

        async exportRecording(event) {
            let histObj = {
                ch0: {
                    hist: this.ch0.hist
                },
                ch1: {
                    hist: this.ch1.hist
                },
                ch2: {
                    hist: this.ch2.hist
                },
                ch3: {
                    hist: this.ch3.hist
                },
                samples: this.samples
            }

            fetch("http://localhost:8080/upload", {
                method: "POST",
                headers: {'Content-Type': 'application/json'}, 
                body: JSON.stringify(histObj)
            }).then(res => {
                console.log("Request complete! response:", res);
                this.getFiles()
            });

            // console.log("CH0", this.ch0.hist)
            // console.log("CH1", this.ch1.hist)
            // console.log("CH2", this.ch2.hist)
            // console.log("CH3", this.ch3.hist)
        },

        async testREST(event) {
            fetch("http://localhost:8080", {
                method: "POST",
                headers: {'Content-Type': 'application/json'}
            }).then(res => res.json())
            .then(data => {
                console.log("Request complete! response:", data);
            });
        },

        async testFiles(event) {
            this.getFiles()
            console.log("Current files:", this.files)
        },

        async getFiles() {
            fetch("http://localhost:8080/files", {
                method: "GET",
                headers: {'Content-Type': 'application/json'}
            }).then(res => res.json())
            .then(data => {
                console.log("Request complete! response:", data);
                this.files = data.files
            });
        },

        async getFileFromName(filename) {
            console.log("Requesting file", filename)
            fetch("http://localhost:8080/download?" + new URLSearchParams({
                filename: this.selectedFile,
            }),{
                method: "GET",
                headers: {'Content-Type': 'application/json'}
            }).then(res => res.json())
            .then(data => {
                console.log("Request complete! response:", data);
            });
        }
    },
    components: { DataBandDisplay },
    beforeMount() {
        this.getFiles()
    }
}
</script>

<template>
    <div class="buttons">
        <h3>{{this.connected ? "Connected!" : "Not Connected :("}}</h3>
	    <button @click="connect">Connect Muse</button>
        <button v-if="this.connected" @click="toggleRecord">{{this.recording ? "Stop Recording" : "Record"}}</button>
        <button v-if="this.connected && this.recorded" :disabled="this.recording" @click="exportRecording">Export Recording</button>
        <button @click="testREST">Test API Connection</button>
        <button @click="testFiles">Test GET Filenames</button>
        <div id="dropdown">
            <select v-model="this.selectedFile">
                <option disabled value=null>Please Select</option>
                <option v-for="file in this.files" :value="file">{{file}}</option>
            </select>
            <button :disabled="this.selectedFile == null" @click="getFileFromName(this.selectedFile)">Select</button>
        </div>
        <span style="padding-left:5%">Your Choice is: {{this.selectedFile}}</span>
    </div>
    <div class="data">
        <DataBandDisplay :ch=this.ch0 name='ch0'></DataBandDisplay>
        <DataBandDisplay :ch=this.ch1 name='ch1'></DataBandDisplay>
        <DataBandDisplay :ch=this.ch2 name='ch2'></DataBandDisplay>
        <DataBandDisplay :ch=this.ch3 name='ch3'></DataBandDisplay>
    </div>
</template>

<style>

.data {
    display: flex;
    align-content: space-between;
}

.buttons {
    display: flex;
    flex-direction: column;
    padding: 10px;
}

</style>