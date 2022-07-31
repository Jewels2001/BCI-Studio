
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
    background-color: rgb(0, 0, 0);
    height: 100%;
    width: 300px;
    position: fixed;
    right: 0px;
    top: 0px;
    /* align-items: center; */
    opacity: 50%;
}

</style>
