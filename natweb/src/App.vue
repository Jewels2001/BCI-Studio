<template>
  <div id = "nav">
    <router-link to="/">Home</router-link>
    <router-link to="/recording">Record Brain</router-link>
    <router-link to="/play">Make ART</router-link>
  </div>
  <!-- <SideBar/> -->
  <router-view/>
</template>

<script>
// import SideBar from './components/SideBar.vue'

export default {
  name: 'App',
  components: {
    // SideBar

}
}

</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #c0dbf7;
  margin-top: 60px;
}
#nav {
  background-color: lightblue;
  margin-top: 0px;
  
}
nav li:hover,
nav li:active {
  background-color: indianred;
  cursor: pointer;
}

</style>
