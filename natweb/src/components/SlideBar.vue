<template>
  <div>
    <div class="slider-component">
      <div class="slidecontainer">
        <input
          class="slider"
          type="range"
          ref="input"
          v-model="currentValue"
          :min="min"
          :max="max"
          @input="onInput"
        >
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'SlideBar',
  props: {
    value: {
      type: Number,
      required: true
    },
    min: {
      type: Number,
      required: true
    },
    max: {
      type: Number,
      required: true
    }
  },
  data(){
    return {
      currentValue: this.value
    };
  },
  methods: {
    onInput() {
      // this.currentValue is a string because HTML is weird
      this.$emit('input', parseInt(this.currentValue));
    }
  }
}
</script>

<style scoped>
.slider-component .slidecontainer {
	width: 100%;
}

.slider-component .slidecontainer .slider {
	-webkit-appearance: none;
	appearance: none;
	width: 100%;
	height: 4px;
	border-radius: 2px;
	background: #c2c2c2;
	outline: none;
	opacity: 0.7;
	-webkit-transition: .2s;
	transition: opacity .2s;
}

.slider-component .slidecontainer .slider:hover {
	opacity: 1;
}

.slider-component .slidecontainer .slider::-webkit-slider-thumb {
	-webkit-appearance: none;
	appearance: none;
	width: 18px;
	height: 18px;
	background: #83b8e9;
	cursor: pointer;
	border-radius: 50%;
}

.slider-component .slidecontainer .slider::-moz-range-thumb {
	width: 18px;
	height: 18px;
	background: #7caee7;
	cursor: pointer;
	border-radius: 50%;
}
</style>