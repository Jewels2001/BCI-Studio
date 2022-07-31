<!-- Main Page / Routes -->
<script>
import TopBar from './components/TopBar.vue'

// import HomeView from './pages/HomeView.vue'
// import AboutView from './pages/AboutView.vue'
// import StudioView from './pages/StudioView.vue'
// import NotFound from './pages/NotFound.vue'

import HomePage from './components/HomePage.vue'
// import SideBar from './components/SideBar.vue'

// const routes = {
//   '/': HomeView,
//   '/about': AboutView,
//   '/studio': StudioView
// }
//
// export default {
//   data() {
//     return {
//       currentPath: window.location.hash
//     }
//   },
//   computed: {
//     currentView() {
//       return routes[this.currentPath.slice(1) || '/'] || NotFound
//     }
//   },
//   mounted() {
//     window.addEventListener('hashchange', () => {
//       this.currentPath = window.location.hash
//     })
//   }
// }
export default {
  components: {
    TopBar
  }
}
</script>

<!-- Without # causes the page to reload;
with it for routing purposes the page does not reload -->
<template>
    <TopBar />

  <!-- <HomePage msg="BCI Wizards: Brain Musical Visualizer!!!!"/> -->
<!-- <SideBar/> -->
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
  color: #2c3e50;
}
</style>
