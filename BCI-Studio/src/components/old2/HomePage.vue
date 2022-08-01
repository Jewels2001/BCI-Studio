<template>
  <img alt="logo" src="../assets/minion.jpg">
  <!-- <HomePage msg="BCI Wizards: Brain Musical Visualizer!!!!"/> -->
  <div class="hello">
    <h1>BCI Wizards: Brain Musical Visualizer!!!!</h1>
    <p>
      This app uses your very own brain data to create SICK visuals and music. <br>
    </p>
    <button @click="$router.push('recording')"> Click here to start </button>
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
h1 {
  color: black;
}
p {
  color: black;
}
body {
  background-color: rgb(207, 32, 32);
}
a {
  color: #42b983;
}
button {
  background-color: #42b983;
  border: none;
  border-radius: 5px;
  box-shadow: 0 5px #999;
  color: white;
  text-align: center;
  padding: 20px 50px;
  font-size: 20px;
}
button:hover {
  background-color: #2d7a58;
}
button:active {
  background-color: #2d7a58;
  box-shadow: 0 3px #666 ;
  transform: translateY(2px);
}

</style>
