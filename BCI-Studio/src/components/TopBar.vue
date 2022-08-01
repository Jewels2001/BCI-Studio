
<script>
import HomeView from '../views/HomeView.vue'
// import AboutView from '../views/AboutView.vue'
// import StudioView from '../views/StudioView.vue'
import NotFound from '../views/NotFound.vue'
//
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
//     // Once we add rounting, can delete all of query selecting to add 'active status'
//     document.querySelector('a[href="' + this.currentPath + '"]').classList.add('active');
//     window.addEventListener('hashchange', () => {
//       document.querySelector('a[href="' + this.currentPath + '"]').classList.remove('active');
//       console.log(document.querySelector('a[href="' + this.currentPath + '"]'))
//       this.currentPath = window.location.hash
//       document.querySelector('a[href="' + this.currentPath + '"]').classList.add('active');
//     })
//   }
// }

</script>

<!-- Without # causes the page to reload;
with it for routing purposes the page does not reload -->
<template>
  <!-- <div class="top">
    <a href="#/"> Home </a> |
    <a href="#/about"> About </a> |
    <a href="#/studio"> Studio! </a> |
    <a href="#/non-existent-path">Broken Link</a>
  </div>
  <br>
  <component :is="currentView" /> -->
  <div class="top">
    <router-link to="/">Home</router-link>
    <router-link to="/recording">Record Brain</router-link>
    <router-link to="/play">Make ART</router-link>
  </div>
  <br>
</template>

<style>
.top {
  background-color: lightblue;
  overflow: hidden;
  display: block;
}

.top.router-link-active {
  background-color: indianred;
  color: red;
}


.top a {
  float: left;
  color: #f2f2f2;
  text-align: center;
  padding: 14px 16px;
  text-decoration: none;
  font-size: 14px;
}

.top a:hover {
  background-color: #ddd;
  color: black;
}

/* https://stackoverflow.com/questions/46083220/how-to-vuejs-router-link-active-style */
.top a.active {
  background-color: #04AA6D;
  color: black;
}

nav li:hover,
nav li.active,
nav li.router-link-exact-active {
  background-color: indianred;
  cursor: pointer;
}

</style>
