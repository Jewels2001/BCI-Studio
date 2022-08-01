<template>
  <div class="page">
        <div class="bg" :style="{ background: color }">
            <div class="cover">
                <ColorPicker
                :theme="light"
                :color="color"
                :sucker-hide="false"
                :sucker-canvas="suckerCanvas"
                :sucker-area="suckerArea"
                @changeColor="changeColor"
                @openSucker="openSucker"
                />
            </div>
        </div>
  </div>
</template>

<script>
import { ColorPicker } from 'vue-color-kit'
import 'vue-color-kit/dist/vue-color-kit.css'

export default {
  name: 'ColourPicker',
  components:{ColorPicker},
  data() {
      return {
        color: '#59c7f9',
        suckerCanvas: null,
        suckerArea: [],
        isSucking: false,
      }
    },
  methods: {
        changeColor(color) {
            const { r, g, b, a } = color.rgba
            this.color = `rgba(${r}, ${g}, ${b}, ${a})`
        }
    },
    createCanvas(cover, coverRect) {
      const canvas = document.createElement('canvas')
      const ctx = canvas.getContext('2d')
      canvas.width = coverRect.width
      canvas.height = coverRect.height
      ctx.drawImage(cover, 0, 0, coverRect.width, coverRect.height)
      Object.assign(canvas.style, {
        position: 'absolute',
        left: coverRect.left + 'px',
        top: coverRect.top + 'px',
        opacity: 0,
      })
      return canvas
    },
}
</script>

<style scoped>
html, body {
  height: 100%;
}

body {
  margin: 0;
}

.page {
  height: 100;
}

.bg {
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    align-items: center;
}
.cover {
    /* display: flex; */
    justify-content: space-around;
    align-items: center;
    width: 100%;
}
</style>
