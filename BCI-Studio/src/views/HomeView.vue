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
  <HomePage/>
</template>


<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1 {
  color: #e27d60;
  font-family: "Vanitas";
  font-weight: lighter;
  font-size: 200px;
  margin: 20px;
  margin-bottom: 10px;
}

p {
  color: black;
  font-family: "Bodoni Moda";
  font-size: 30px;
}
img {
  height: 400px;
}

button {
  /* background-color: #42b983;
  border: none;
  border-radius: 5px;
  box-shadow: 0 5px #999;
  color: white;
  text-align: center;
  padding: 20px 50px;
  font-size: 20px; */
  padding: 0.8rem 2.5rem;
  background: transparent;
  text-align: center;
  font-size: 30px;
  font-family: "Bodoni Moda";
  text-transform: uppercase;
  cursor: pointer;
  border: 1px solid #e27d60;
  color: rgb(31, 31, 31);
  font-weight: bold;
  letter-spacing: 1px;
  transition: 0.3s linear;
  position: relative;
}
/* button:hover {
  background-color: #2d7a58;
}
button:active {
  background-color: #2d7a58;
  box-shadow: 0 3px #666 ;
  transform: translateY(2px);
} */

button::after {
  content: "";
  position: absolute;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  z-index: -1;
  transform: scalex(0);
  background: #c38d9e;
  transform-origin: right;
  transition: transform 200ms ease-in;
}

button:hover::after {
  transform: scalex(1);
  transform-origin: left;
}

</style>
