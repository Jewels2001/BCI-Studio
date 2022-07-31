<!-- Main Page / Routes -->
<script>
import HomeView from './pages/HomeView.vue'
import AboutView from './pages/AboutView.vue'
import StudioView from './pages/StudioView.vue'
import NotFound from './pages/NotFound.vue'

import HomePage from './components/HomePage.vue'
import SideBar from './components/SideBar.vue'

const routes = {
  '/': HomeView,
  '/about': AboutView,
  '/studio': StudioView
}

export default {
  data() {
    return {
      currentPath: window.location.hash
    }
  },
  computed: {
    currentView() {
      return routes[this.currentPath.slice(1) || '/'] || NotFound
    }
  },
  mounted() {
    window.addEventListener('hashchange', () => {
      this.currentPath = window.location.hash
    })
  }
}
//
// export default {
//   name: 'App',
//   components: {
//     SideBar
// }
// }
</script>

<!-- Without # causes the page to reload;
with it for routing purposes the page does not reload -->

<template>
  <a href="#/"> Home </a> |
  <a href="#/about"> About </a> |
  <a href="/studio"> Studio! </a> |
  <a href="#/non-existent-path">Broken Link</a>
  <component :is="currentView" />
  <!-- <img alt="logo" src="./assets/minion.jpg"> -->
  <!-- <HomePage msg="BCI Wizards: Brain Musical Visualizer!!!!"/> -->
  <!-- <SideBar/>
  <router-view/> -->
</template>

<!-- /* <style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style> */ -->
