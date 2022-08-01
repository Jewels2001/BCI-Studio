<template>
    <div class="SlideBar">
    <input type="range" min="1" max="100" value="50" class="slider" id="myRange">
    
    </div>

</template>

<script>
export default {
    name: 'SlideBar'
}
</script>

<style scoped>
.slidecontainer {
  width: 100%; 
}

.slider {
  -webkit-appearance: none;  
  appearance: none;
  width: 100%; 
  height: 25px; 
  background: #d3d3d3; 
  outline: none; 
  opacity: 0.7; 
  -webkit-transition: .2s; 
  transition: opacity .2s;
}
.slider:hover {
  opacity: 1; 
}


.slider::-webkit-slider-thumb {
  -webkit-appearance: none;
  appearance: none;
  width: 25px; 
  height: 25px; 
  background: #04AA6D; 
  border-radius: 50%;
  cursor: pointer;
}

.slider::-moz-range-thumb {
  width: 25px; 
  height: 25px; 
  background: #04AA6D; 
  cursor: pointer;
}
</style>