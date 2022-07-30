<script setup>
defineProps({
  ch: {
    type: Object,
    required: true
  }
})
</script>

<template>
    <div v-for="(label, i) in ch.labels">
        <h3>{{label}}</h3>
        <p>{{ch.datasets[0].data[i]}}</p>
    </div>
</template>