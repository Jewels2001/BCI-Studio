
<template>
  <div class="sidebar">
  <h1>Shapes</h1>
  <ShapeOptions/>
  </div>
  
</template>

<script>
import ShapeOptions from './ShapeOptions.vue'

export default {
  name: 'SideBar',
  components: {
    ShapeOptions
  }
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.sidebar {
    background-color: #42b983;
    height: 100%;
    width: 300px;
    position: fixed;
    right: 0px;
    top: 0px;
    align-items: center;
}

</style>
