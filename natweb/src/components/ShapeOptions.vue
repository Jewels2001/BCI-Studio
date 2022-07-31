<template>
    <div class="shapes">
    <div id ="square">
        <label class="SquareSwitch">
        <input type="checkbox">
        <span class="slider round"></span>
    </label>
    </div>
    
    <div id="circle">
        <label class="CircleSwitch">
        <input type="checkbox">
        <span class="slider round"></span>
    </label>
    </div>
    
    </div>

</template>

<script>

export default {
    name: 'ShapeOptions'
}
</script>

<style scoped>
/* shapes */
#square {
  width: 100px;
  height: 100px;
  background: white;
  float: right;
  margin: 25px;
}
#circle {
  width: 100px;
  height: 100px;
  background: white;
  border-radius: 50%;
  float: right;
  margin: 25px 20px;
}

/* switches */
.SquareSwitch {
    position: relative;
    display: inline-block;
    width: 60px;
    height: 34px;
    margin-top: 115px;

    
}
.SquareSwitch input {
  opacity: 0;
  width: 0;
  height: 0;
}

.CircleSwitch {
    position: relative;
    display: inline-block;
    width: 60px;
    height: 34px;
    margin-top: 116px;
  
}
    
.CircleSwitch input {
  opacity: 0;
  width: 0;
  height: 0;
}

/* slider */
.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #ccc;
  -webkit-transition: .4s;
  transition: .4s;
}

.slider:before {
  position: absolute;
  content: "";
  height: 26px;
  width: 26px;
  left: 4px;
  bottom: 4px;
  background-color: white;
  -webkit-transition: .4s;
  transition: .4s;
}

input:checked + .slider {
  background-color: #2196F3;
}

input:focus + .slider {
  box-shadow: 0 0 1px #2196F3;
}

input:checked + .slider:before {
  -webkit-transform: translateX(26px);
  -ms-transform: translateX(26px);
  transform: translateX(26px);
}

.slider.round {
  border-radius: 34px;
}

.slider.round:before {
  border-radius: 50%;
}
</style>