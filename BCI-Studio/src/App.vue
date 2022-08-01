<!-- Main Page / Routes -->
<script>
// import { RouterLink, RouterView } from 'vue-router'
import TopBar from './components/TopBar.vue'

export default {
  name: 'App',
  components: {
    TopBar
  },
}
</script>

<template>
  <TopBar />
  <router-view/>
</template>

<style scoped>
@font-face {
  font-family: "Vanitas";
  src: local("Vanitas"),
  url(./fonts/VanitasVanitatum.ttf)format("truetype");
}

@font-face {
  font-family: "HustleBright";
  src: local("HustleBright"),
  url(./fonts/HustleBright.ttf)format("truetype");
}

#app {
  font-family: Vanitas, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #c0dbf7;
}
  /* color: #2c3e50; */

</style>
