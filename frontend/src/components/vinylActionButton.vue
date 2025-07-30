<template>
  <button @click="toggleButton" class="vinyl-fab" :class="{ 'is-active': isActive }">
    <img
      src="../assets/images/Vinyl-Icon.png"
      alt="Vinyl Disk"
      class="vinyl-img"
      :class="{ spinning: isActive }"
    />
  </button>
</template>

<script>
export default {
  name: 'VinylActionButton',
  props: {
    initialActive: {
      type: Boolean,
      default: true
    },
    color: {
      type: String,
      default: '#222' // Button background color
    }
  },
  data() {
    return {
      isActive: this.initialActive
    }
  },
  methods: {
    toggleButton() {
      this.isActive = !this.isActive;
      this.$emit('toggle', this.isActive);
    }
  },
  watch: {
    initialActive(newValue) {
      this.isActive = newValue;
    }
  }
}
</script>

<style scoped>
.vinyl-fab {
  position: fixed;
  bottom: 4rem;
  right: 2rem;
  width: 4rem;
  height: 4rem;
  border-radius: 50%;
  background-color: v-bind('color');
  box-shadow: 0 4px 10px rgba(0,0,0,0.2);
  border: none;
  cursor: pointer;
  z-index: 999;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s cubic-bezier(0.175,0.885,0.32,1.275);
  overflow: hidden;
}

.vinyl-fab:hover {
  transform: scale(1.05);
  box-shadow: 0 6px 14px rgba(0,0,0,0.25);
}

.vinyl-fab:active {
  transform: scale(0.95);
}

.vinyl-img {
  width: 60%;
  height: 60%;
  border-radius: 50%;
  transition: box-shadow 0.3s;
  box-shadow: 0 0 8px rgba(0,0,0,0.15);
}

.spinning {
  animation: spin 1.2s linear infinite;
}

@keyframes spin {
  100% { transform: rotate(360deg); }
}
</style>