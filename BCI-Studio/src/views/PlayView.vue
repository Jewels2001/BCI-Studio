<template>
    <div class="main">
      <p> Hi </p>
    </div>
    <div class="side">
      <SideBar/>
    </div>
</template>

<script>
import SideBar from '../components/SideBar.vue'
import Viz from '../components/Viz.vue'

export default {
  name: 'App',
  components: {
    SideBar
  }
}
</script>

<style>

.side {
  height: 100%;
  width: 50%;
  float: right;
}

.main {
  width: 50%;
  float: left;
  margin-top: 5px;
}


</style>
