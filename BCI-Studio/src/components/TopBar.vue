
<script>
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
  <div class="top">
    <div>
    <router-link to="/">Home</router-link>
    <router-link to="/recording">Record Brain</router-link>
    <router-link to="/play">Make ART</router-link>
    <router-link to="/studio">STUDIO</router-link>
    </div>
    <div class="top-right">
      <router-link to="/about">About</router-link>
      <router-link to="/contact">Contact</router-link>
    </div>
  </div>
</template>

<style>
.top {
  background-color: lightblue;
  overflow: hidden;
  display: block;
  clear: both;
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
  font-family: 'HustleBright';
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

.top-right {
  float: right;
}


nav li:hover,
nav li.active,
nav li.router-link-exact-active {
  background-color: indianred;
  cursor: pointer;
}

</style>
