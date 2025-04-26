<template>
  <transition name="fade">
    <div 
      v-if="isVisible" 
      class="floating-download-status"
      @mouseenter="isHovered = true"
      @mouseleave="isHovered = false"
    >
      <div class="download-content">
        <div class="loading-spinner" :class="{ 'hovered': isHovered }">
          <template v-if="!isHovered">
            <svg class="spinner" viewBox="0 0 50 50">
              <circle class="path" cx="25" cy="25" r="20" fill="none" stroke-width="5"></circle>
            </svg>
          </template>
          <template v-else>
            <!-- X button when hovered -->
            <button @click="stopDownload" class="cancel-button">
              <svg viewBox="0 0 24 24" width="24" height="24">
                <path d="M19 6.41L17.59 5 12 10.59 6.41 5 5 6.41 10.59 12 5 17.59 6.41 19 12 13.41 17.59 19 19 17.59 13.41 12z"/>
              </svg>
            </button>
          </template>
        </div>
        <div class="download-info">
          <div class="download-text">
            {{ isHovered ? 'Click X to cancel' : downloadText }}
          </div>
          <div class="progress-bar-container">
            <div class="progress-bar" :class="{ 'hovered': isHovered }">
              <div 
                class="progress-container"
                :class="{ 'hovered': isHovered }"
                :style="{ width: progressPercentage + '%' }"
              >
                <div class="progress-bar-animation" :class="{ 'hovered': isHovered }"></div>
              </div>
            </div>
          </div>
          <div class="download-details" v-if="showDetails">
            {{ currentItem }} of {{ totalItems }} items
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>

<script>
export default {
  name: 'FloatingDownloadStatusAndCancel',
  props: {
    isVisible: {
      type: Boolean,
      default: false
    },
    downloadText: {
      type: String,
      default: 'Downloading...'
    },
    showDetails: {
      type: Boolean,
      default: false
    },
    currentItem: {
      type: Number,
      default: 0
    },
    totalItems: {
      type: Number,
      default: 0
    }
  },
  data() {
    return {
      isHovered: false
    }
  },
  computed: {
    progressPercentage() {
      if (this.totalItems <= 0) return 0;
      return Math.min(100, Math.floor((this.currentItem / this.totalItems) * 100));
    }
  },
  methods: {
    stopDownload() {
      // Emit an event to be caught by the parent component
      this.$emit('cancel-download');
    }
  }
}
</script>

<style scoped>
.floating-download-status {
  position: fixed;
  bottom: 20px;
  left: 20px;
  background-color: rgba(32, 32, 32, 0.9);
  border-radius: 8px;
  padding: 15px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.3);
  color: white;
  z-index: 9999;
  display: flex;
  align-items: center;
  min-width: 280px;
  max-width: 320px;
  backdrop-filter: blur(5px);
  border: 1px solid rgba(75, 75, 75, 0.5);
  transition: all 0.3s ease;
}

.floating-download-status:hover {
  background-color: rgba(32, 32, 32, 0.95);
  box-shadow: 0 4px 15px rgba(255, 0, 0, 0.2);
  border-color: rgba(255, 100, 100, 0.6);
  backdrop-filter: blur(10px);
}

.download-content {
  display: flex;
  align-items: center;
  width: 100%;
}

.loading-spinner {
  margin-right: 15px;
  flex-shrink: 0;
  transition: all 0.3s ease;
}

.loading-spinner.hovered {
  transform: scale(1.1);
}

.spinner {
  animation: rotate 2s linear infinite;
  height: 30px;
  width: 30px;
}

.path {
  stroke: #4CAF50;
  stroke-linecap: round;
  animation: dash 1.5s ease-in-out infinite;
}

.cancel-button {
  background: none;
  border: none;
  cursor: pointer;
  padding: 5px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background-color: rgba(255, 80, 80, 0.9);
  width: 30px;
  height: 30px;
  transition: all 0.2s ease;
}

.cancel-button:hover {
  background-color: rgba(255, 50, 50, 1);
  transform: scale(1.1);
}

.cancel-button svg {
  fill: white;
}

.cancel-button:active {
  transform: scale(0.95);
}

.download-info {
  flex-grow: 1;
}

.download-text {
  font-size: 14px;
  margin-bottom: 8px;
  font-weight: bold;
  transition: color 0.3s ease;
}

.floating-download-status:hover .download-text {
  color: #ff7070;
}

.progress-bar-container {
  width: 100%;
  margin-bottom: 5px;
}

.progress-bar {
  height: 8px;
  background-color: rgba(120, 120, 120, 0.3);
  border-radius: 4px;
  position: relative;
  overflow: hidden;
  transition: all 0.3s ease;
}

.progress-bar.hovered {
  background-color: rgba(150, 120, 120, 0.3);
}

.progress-container {
  position: absolute;
  height: 100%;
  left: 0;
  background-color: rgba(46, 204, 113, 0.5);
  border-radius: 4px;
  transition: all 0.3s ease;
  overflow: hidden;
}

.progress-container.hovered {
  background-color: rgba(255, 80, 80, 0.5);
}

.progress-bar-animation {
  position: absolute;
  height: 100%;
  width: 100%;
  background: linear-gradient(90deg, 
    rgba(46, 204, 113, 0.1) 0%, 
    rgba(46, 204, 113, 0.9) 50%,
    rgba(46, 204, 113, 0.1) 100%);
  border-radius: 4px;
  animation: progressAnim 1.5s infinite;
  background-size: 200% 100%;
  transition: background 0.3s ease;
}

.progress-bar-animation.hovered {
  background: linear-gradient(90deg, 
    rgba(255, 80, 80, 0.1) 0%, 
    rgba(255, 80, 80, 0.9) 50%,
    rgba(255, 0, 0, 0.1) 100%);
}

.download-details {
  font-size: 12px;
  color: #ccc;
  transition: color 0.3s ease;
}

.floating-download-status:hover .download-details {
  color: #ff9999;
}

/* Animations */
@keyframes rotate {
  100% {
    transform: rotate(360deg);
  }
}

@keyframes dash {
  0% {
    stroke-dasharray: 1, 150;
    stroke-dashoffset: 0;
  }
  50% {
    stroke-dasharray: 90, 150;
    stroke-dashoffset: -35;
  }
  100% {
    stroke-dasharray: 90, 150;
    stroke-dashoffset: -124;
  }
}

@keyframes progressAnim {
  0% {
    left: -100%;
  }
  100% {
    left: 100%;
  }
}

/* Fade transition */
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s, transform 0.3s;
}

.fade-enter-from, .fade-leave-to {
  opacity: 0;
  transform: translateY(20px);
}
</style>