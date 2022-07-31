<script>
// @ is an alias to /src
import HomePage from '@/components/HomePage.vue'

export default {
  name: 'HomeView',
  components: {
    HomePage
  }
}
</script>

<template>
  <img alt="logo" src="../assets/minion.jpg">
  <div class="hello">
    <h1>{{ msg }}</h1>
    <p>
      This app uses your very own brain data to create SICK visuals and music. <br>
    </p>
    <button v-on:click="alert('hi!')" >
    Click here to start
    </button>
  </div>
</template>
