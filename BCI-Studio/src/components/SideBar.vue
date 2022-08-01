
<template>
  <div id="sidebar">
    <h1>Options</h1>
    <!-- <span> Size: {{ number }} </span> -->
    <!-- <SlideBar v-model="number" :min="0" :max="99" class="slider" /> -->
    <div class="option-box">
      <ShapeOptions/>
    </div>
    <!-- <button @click:></button> -->
    <!-- <ColourPicker/> -->
    <div class="option-box">
      <BrainPicker :options="opts"/>

    </div>
    <div>
      <Music/>
    </div>
  </div>

</template>

<script>
import ShapeOptions from './ShapeOptions.vue'
// import SlideBar from './SlideBar.vue'
import ColourPicker from './ColourPicker.vue';
import BrainPicker from './BrainPicker.vue';
import Music from './music/Music.vue'

export default {
  name: "SideBar",
  components: {
    ShapeOptions,
    // SlideBar,
    ColourPicker,
    BrainPicker,
    Music,
},
  data() {
    return {
      opts: [
      { name: "Color", value: 0 },
      { name: "Speed", value: 1 },
      { name: "Shape", value: 2 },
      { name: "Size", value: 3 }
      ]
    };
  },
  methods: {

  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.option-box {
  height: 250px;
}

#sidebar {
    background-color: beige;
    height: 100%;
    width: 300px;
    position: absolute;
    /* clear: both;  */
    /* float: right; */
    /* right: 0px;
    top: 0px; */
    border-left: 5px solid #d2b48c;
    /* align-items: center; */
    /* opacity: 50%; */
}
h1 {
  font-weight: lighter;
  font-family: "Vanitas";
  font-size: 80px;
  color: #41b3a3;
  margin: 10px;
}
span {
  font-family: "Bodoni Moda";
  font-size: 30px;
  color: #d2b48c
}

</style>
