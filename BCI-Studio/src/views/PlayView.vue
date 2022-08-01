<template>
    <div class="main">
    </div>
    <div class="side">
      <SideBar/>
    </div>
</template>

<script>
import SideBar from '../components/SideBar.vue'

export default {
  name: 'App',
  components: {
    SideBar
  }
}
</script>

<style>

.side {
  height: 100%;
}

</style>
