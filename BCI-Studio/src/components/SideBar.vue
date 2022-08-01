
<template>
  <div id="sidebar">
    <h1>Options</h1>
  <span> Size: {{ number }} </span>
  <SlideBar v-model="number" :min="0" :max="99" class="slider" />
  <ShapeOptions/>
  <ColourPicker/>
  </div>

</template>

<script>
import ShapeOptions from './ShapeOptions.vue'
import SlideBar from './SlideBar.vue'
import ColourPicker from './ColourPicker.vue';

export default {
  name: "SideBar",
  components: {
    ShapeOptions,
    SlideBar,
    ColourPicker
},
  data() {
    return {
      number: 0,
    };
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#sidebar {
    background-color: beige;
    height: 100%;
    width: 300px;
    /* position: relative;
    clear: both; */
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
