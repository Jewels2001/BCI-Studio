<template>
  <!-- <img alt="logo" src="./assets/minion.jpg"> -->
  <!-- <HomePage msg="BCI Wizards: Brain Musical Visualizer!!!!"/> -->
  <SideBar/>
  <router-view/>
</template>

<script>
// import HomePage from './components/HomePage.vue'
import SideBar from './components/SideBar.vue'

export default {
  name: 'App',
  components: {
    SideBar
}
}

</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
