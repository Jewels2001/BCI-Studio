<template>
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

<script>
export default {
  name: 'HomePage',
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
button {
  background-color: #42b983;
  border: none;
  color: white;
  text-align: center;
  padding: 20px 50px;
  font-size: 20px;
}
</style>
