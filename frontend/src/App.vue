<template>
  <div id="app-container">
    <router-view></router-view>
    <FloatingDownloadStatus 
      :is-visible="isVisible"
      :download-text="downloadText"
      :current-item="currentItem"
      :total-items="totalItems"
      :show-details="totalItems > 0"
    />
  </div>
</template>

<script setup>
import FloatingDownloadStatus from './components/floatingDownloadStatus.vue';
import { useDownloadStatus } from './downloadStatus.js';
import { useRoute } from 'vue-router'
import { computed } from 'vue'

// Get the reactive state from the store
const { isVisible, downloadText, currentItem, totalItems } = useDownloadStatus();

const route = useRoute()
const transitionName = computed(() => {
  return route.meta.transition || 'zoom'
})
</script>

<style>
/* Optional: Ensure the container fills the viewport */
#app-container {
  height: 100vh;
  width: 100vw;
}
</style>