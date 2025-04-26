<script setup>
import { useRoute } from 'vue-router'
import { computed } from 'vue'

const route = useRoute()
const transitionName = computed(() => {
  return route.meta.transition || 'zoom'
})
</script>

<template>
  <div class="app-container">
    <router-view></router-view>
    <!--<transition :name="transitionName" mode="out-in">
      <router-view></router-view>
    </transition>-->
  </div>
</template>