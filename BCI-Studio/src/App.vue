<!-- Main Page / Routes -->
<script>
// import { RouterLink, RouterView } from 'vue-router'
import TopBar from './components/TopBar.vue'

export default {
  name: 'App',
  components: {
    TopBar
  }
}
</script>

<template>
  <TopBar />
  <p> hi </p>
  <router-view/>
</template>

<style>
  body {
    margin: 0 0 0 0;
  }

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #c0dbf7;
}

/* <!--
<style>
  body {
    margin: 0 0 0 0;
  }
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
--> */

</style>
