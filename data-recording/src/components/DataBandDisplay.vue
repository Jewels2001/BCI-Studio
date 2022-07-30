<script setup>
defineProps({
  ch: {
    type: Object,
    required: true
  },
  name: {
    type: String,
    required: true
  }
})
</script>

<template>
    <div class="band">
      <h1>{{name}}</h1>
      <div v-for="(label, i) in ch.labels">
          <h2>{{label}}</h2>
          <p>{{ch.data[i]}}</p>
      </div>
    </div>
</template>

<style>
.band {
  display: flex;
  flex-direction: column;
  padding: 50px
}
</style>