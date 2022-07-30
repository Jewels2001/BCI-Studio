<script>
import { zipSamples, channelNames, MuseClient } from 'muse-js';
import { powerByBand, epoch, fft } from "@neurosity/pipes";
import DataBandDisplay from './DataBandDisplay.vue';

export default {
    data() {
        return {
            connected: false,
            ch0: {
                labels: ["Delta", "Theta", "Alpha", "Beta", "Gamma"],
                datasets: [
                    {
                        fillColor: "rgba(220,220,220,0.2)",
                        strokeColor: "rgba(220,220,220,1)",
                        pointColor: "rgba(220,220,220,1)",
                        pointStrokeColor: "#fff",
                        pointHighlightFill: "#fff",
                        pointHighlightStroke: "rgba(220,220,220,1)",
                        data: [0, 0, 0, 0, 0]
                    }
                ]
            },
            ch1: {
                labels: ["Delta", "Theta", "Alpha", "Beta", "Gamma"],
                datasets: [
                    {
                        fillColor: "rgba(220,220,220,0.2)",
                        strokeColor: "rgba(220,220,220,1)",
                        pointColor: "rgba(220,220,220,1)",
                        pointStrokeColor: "#fff",
                        pointHighlightFill: "#fff",
                        pointHighlightStroke: "rgba(220,220,220,1)",
                        data: [0, 0, 0, 0, 0]
                    }
                ]
            },
            ch2: {
                labels: ["Delta", "Theta", "Alpha", "Beta", "Gamma"],
                datasets: [
                    {
                        fillColor: "rgba(220,220,220,0.2)",
                        strokeColor: "rgba(220,220,220,1)",
                        pointColor: "rgba(220,220,220,1)",
                        pointStrokeColor: "#fff",
                        pointHighlightFill: "#fff",
                        pointHighlightStroke: "rgba(220,220,220,1)",
                        data: [0, 0, 0, 0, 0]
                    }
                ]
            },
            ch3: {
                labels: ["Delta", "Theta", "Alpha", "Beta", "Gamma"],
                datasets: [
                    {
                        fillColor: "rgba(220,220,220,0.2)",
                        strokeColor: "rgba(220,220,220,1)",
                        pointColor: "rgba(220,220,220,1)",
                        pointStrokeColor: "#fff",
                        pointHighlightFill: "#fff",
                        pointHighlightStroke: "rgba(220,220,220,1)",
                        data: [0, 0, 0, 0, 0]
                    }
                ]
            }
        };
    },
    methods: {
        async connect(event) {
            const client = new MuseClient();
            client.connectionStatus.subscribe((status) => {
                this.connected = status ? "Connected!" : "Disconnected",
                    console.log(status ? "Connected!" : "Disconnected");
            });
            try {
                await client.connect();
                await client.start();
                zipSamples(client.eegReadings).pipe(epoch({ duration: 1024, interval: 100, samplingRate: 256 }), fft({ bins: 256 }), powerByBand()).subscribe((data) => {
                    this.ch0.datasets[0].data = [data.delta[0], data.theta[0], data.alpha[0], data.beta[0], data.gamma[0]];
                    this.ch1.datasets[0].data = [data.delta[1], data.theta[1], data.alpha[1], data.beta[1], data.gamma[1]];
                    this.ch2.datasets[0].data = [data.delta[2], data.theta[2], data.alpha[2], data.beta[2], data.gamma[2]];
                    this.ch3.datasets[0].data = [data.delta[3], data.theta[3], data.alpha[3], data.beta[3], data.gamma[3]];
                    console.log(data);
                });
            }
            catch (err) {
                console.error("Connection failed", err);
            }
        }
    },
    components: { DataBandDisplay }
}
</script>

<template>
    <h3>{{this.connected ? "Connected!" : "Not Connected :("}}</h3>
    <div>
	    <button @click="connect">Connect Muse</button>
    </div>
    <div>
        <DataBandDisplay :ch=this.ch0></DataBandDisplay>
        <DataBandDisplay :ch=this.ch1></DataBandDisplay>
        <DataBandDisplay :ch=this.ch2></DataBandDisplay>
        <DataBandDisplay :ch=this.ch3></DataBandDisplay>
    </div>
</template>