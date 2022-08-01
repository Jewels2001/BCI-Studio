<script>

export default {
  props: {
    options: { type: Array, required: false, default: () => [{name: "None", value: 0}] }
  },
  data() {
    return {
      labels: [
        {
          f: "Delta"
        },
        {
          f: "Theta"
        },
        {
          f: "Alpha"
        },
        {
          f: "Beta"
        },
        {
          f: "Gamma"
        }
      ],
      selectedOptions: [],
    }
  },

  methods: {
  }
}
</script>

<template>
  <div class="box-option" v-for="item in options">
      <h3> {{ item.name }}:
      <select>
        <option v-for="(label) in labels">{{label.f}}</option>
      </select>
      </h3>
  </div>

</template>

<style>
.box-option {
  height: 50px;
  font-family: "Bodoni Moda";
  font-size: 30px;
  color: #d2b48c
}

</style>
