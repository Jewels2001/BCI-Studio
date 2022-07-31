
<script>
import HomeView from '../pages/HomeView.vue'
import AboutView from '../pages/AboutView.vue'
import StudioView from '../pages/StudioView.vue'
import NotFound from '../pages/NotFound.vue'

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
    // Once we add rounting, can delete all of query selecting to add 'active status'
    document.querySelector('a[href="' + this.currentPath + '"]').classList.add('active');
    window.addEventListener('hashchange', () => {
      document.querySelector('a[href="' + this.currentPath + '"]').classList.remove('active');
      console.log(document.querySelector('a[href="' + this.currentPath + '"]'))
      this.currentPath = window.location.hash
      document.querySelector('a[href="' + this.currentPath + '"]').classList.add('active');
    })
  }
}

</script>

<template>
  <div class="top">
    <a href="#/"> Home </a> |
    <a href="#/about"> About </a> |
    <a href="#/studio"> Studio! </a> |
    <a href="#/non-existent-path">Broken Link</a>
  </div>
  <br>
  <component :is="currentView" />
</template>

<style>
.top {
  background-color: #333;
  overflow: hidden;
  display: block;
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
  color: white;
}

</style>
