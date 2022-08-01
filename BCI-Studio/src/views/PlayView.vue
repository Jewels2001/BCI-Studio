<template>
  <div class="container">
    <div class="main">
      <Viz2 />

      <!-- <canvas id="test1" class="test">
      </canvas> -->
    </div>
    <div class="side">
      <SideBar/>
    </div>
  </div>
</template>

<script>
import SideBar from '../components/SideBar.vue'
import Viz from '../components/Viz.vue'
import Viz2 from '../components/Viz2.vue'

export default {
  name: 'App',
  components: {
    Viz,
    Viz2,
    SideBar
  }
}
</script>

<style>

.container {
  height: 95%;
  margin: 0 0 0 0;
}

.side {
  height: 100%;
  width: 300px;
  right: 0px;
  float: right;
  position: absolute;
  display: inline-flex;
  flex-direction: row;
}

.main {
  height: 100%;
  max-width: 100%;
  padding-top: 5px;
  left: 0px;
  float: left;
  margin-right:300px;
  display: inline-flex;
  flex-direction: row;
}

/* .test {
  position: absolute;
  height: 100%;
  width: 100%;
} */

canvas {
  height: 100%;
  width: 100%;
}

</style>
