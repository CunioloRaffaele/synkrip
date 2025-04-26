<template>
  <transition name="fade">
    <div v-if="isVisible" class="floating-download-status">
      <div class="download-content">
        <div class="loading-spinner">
          <svg class="spinner" viewBox="0 0 50 50">
            <circle class="path" cx="25" cy="25" r="20" fill="none" stroke-width="5"></circle>
          </svg>
        </div>
        <div class="download-info">
          <div class="download-text">{{ downloadText }}</div>
          <div class="progress-bar-container">
            <div class="progress-bar">
              <div 
                class="progress-container"
                :style="{ width: progressPercentage + '%' }"
              >
                <div class="progress-bar-animation"></div>
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
  name: 'FloatingDownloadStatus',
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
  computed: {
    progressPercentage() {
      if (this.totalItems <= 0) return 0;
      return Math.min(100, Math.floor((this.currentItem / this.totalItems) * 100));
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
}

.download-content {
  display: flex;
  align-items: center;
  width: 100%;
}

.loading-spinner {
  margin-right: 15px;
  flex-shrink: 0;
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

.download-info {
  flex-grow: 1;
}

.download-text {
  font-size: 14px;
  margin-bottom: 8px;
  font-weight: bold;
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
}

.progress-container {
  position: absolute;
  height: 100%;
  left: 0;
  background-color: rgba(46, 204, 113, 0.5);
  border-radius: 4px;
  transition: width 0.3s ease;
  overflow: hidden;
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
}

.download-details {
  font-size: 12px;
  color: #ccc;
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