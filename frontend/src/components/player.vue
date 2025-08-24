<template>
  <transition name="slide-up">
    <div v-if="song" class="music-player">
      <div class="player-content">
        <!-- Album Art -->
        <div class="album-art">
          <img :src="song.albumArtUrl || song.image" alt="Album Art" />
        </div>

        <!-- Info and Controls -->
        <div class="player-info">
          <div class="song-details">
            <span class="song-title">{{ song.song_name || song.title }}</span>
            <span class="song-artist">{{ song.song_artist_name || song.artist }}</span>
          </div>

          <!-- Progress Bar -->
          <div class="progress-section">
            <div class="time-display">{{ formattedCurrentTime }}</div>
            <div class="progress-bar-container" @click="seek">
              <div class="progress-bar">
                <div class="progress-fill" :style="{ width: progressPercentage + '%' }"></div>
              </div>
            </div>
            <div class="time-display">{{ formattedDuration }}</div>
          </div>

          <!-- Player Controls -->
          <div class="player-controls">
            <button @click="$emit('prev')" class="control-button">
              <svg viewBox="0 0 24 24" width="24" height="24"><path d="M6 6h2v12H6zm3.5 6 8.5 6V6z"/></svg>
            </button>
            <button @click="$emit('toggle-play')" class="control-button play-button">
              <!-- Show Pause icon if playing, Play icon if paused -->
              <svg v-if="isPlaying" viewBox="0 0 24 24" width="32" height="32"><path d="M6 19h4V5H6v14zm8-14v14h4V5h-4z"/></svg>
              <svg v-else viewBox="0 0 24 24" width="32" height="32"><path d="M8 5v14l11-7z"/></svg>
            </button>
            <button @click="$emit('next')" class="control-button">
              <svg viewBox="0 0 24 24" width="24" height="24"><path d="M16 6h2v12h-2zm-3.5 6 8.5 6V6z" transform="scale(-1, 1) translate(-24, 0)"/></svg>
            </button>
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>

<script>
export default {
  name: 'MusicPlayer',
  props: {
    song: { type: Object, default: null },
    isPlaying: { type: Boolean, default: false },
    playbackTime: { type: Number, default: 0 },
  },
  emits: ['toggle-play', 'seek', 'next', 'prev'],
  computed: {
    progressPercentage() {
      if (!this.song || !this.song.duration) return 0;
      return (this.playbackTime / this.song.duration) * 100;
    },
    formattedCurrentTime() {
      return this.formatTime(this.playbackTime);
    },
    formattedDuration() {
      if (!this.song || !this.song.duration) return '0:00';
      return this.formatTime(this.song.duration);
    }
  },
  methods: {
    formatTime(seconds) {
      if (isNaN(seconds)) return '0:00';
      const mins = Math.floor(seconds / 60);
      const secs = Math.floor(seconds % 60).toString().padStart(2, '0');
      return `${mins}:${secs}`;
    },
    seek(event) {
      if (!this.song || !this.song.duration) return;
      const bar = event.currentTarget;
      const clickPosition = event.clientX - bar.getBoundingClientRect().left;
      const barWidth = bar.offsetWidth;
      const seekPercentage = clickPosition / barWidth;
      const newTime = this.song.duration * seekPercentage;
      this.$emit('seek', newTime);
    }
  }
};
</script>

<style scoped>
/* --- Main Container - Styled like the download status --- */
.music-player {
  position: fixed;
  bottom: 20px;
  left: 50%;
  transform: translateX(-50%);
  background-color: rgba(32, 32, 32, 0.9);
  border-radius: 12px;
  padding: 15px;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.4);
  color: white;
  z-index: 9998;
  display: flex;
  align-items: center;
  width: 90%;
  max-width: 500px;
  backdrop-filter: blur(8px);
  -webkit-backdrop-filter: blur(8px);
  border: 1px solid rgba(75, 75, 75, 0.5);
  transition: all 0.4s ease;
}

.player-content {
  display: flex;
  align-items: center;
  width: 100%;
  gap: 15px;
}

/* --- Album Art --- */
.album-art {
  width: 60px;
  height: 60px;
  flex-shrink: 0;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 2px 8px rgba(0,0,0,0.3);
}

.album-art img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

/* --- Info & Controls Column --- */
.player-info {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  gap: 8px;
  min-width: 0; /* Prevents flexbox overflow issues */
}

.song-details {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
}

.song-title {
  font-size: 15px;
  font-weight: bold;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  max-width: 100%;
}

.song-artist {
  font-size: 12px;
  color: #b3b3b3;
}

/* --- Progress Bar Section --- */
.progress-section {
  display: flex;
  align-items: center;
  gap: 10px;
  width: 100%;
}

.time-display {
  font-size: 11px;
  color: #ccc;
  font-family: monospace;
}

.progress-bar-container {
  flex-grow: 1;
  height: 8px;
  cursor: pointer;
}

.progress-bar {
  height: 100%;
  background-color: rgba(120, 120, 120, 0.3);
  border-radius: 4px;
  position: relative;
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background-color: #1DB954; /* Spotify Green */
  border-radius: 4px;
  transition: width 0.1s linear;
}

/* --- Player Controls --- */
.player-controls {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 15px;
}

.control-button {
  background: none;
  border: none;
  cursor: pointer;
  padding: 5px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  fill: #b3b3b3;
  transition: all 0.2s ease;
}

.control-button:hover {
  fill: #ffffff;
  transform: scale(1.1);
}

.play-button {
  background-color: #ffffff;
  fill: #121212;
  width: 40px;
  height: 40px;
}

.play-button:hover {
  background-color: #1DB954;
  fill: #ffffff;
}

/* --- Transitions --- */
.slide-up-enter-active, .slide-up-leave-active {
  transition: opacity 0.4s, transform 0.4s;
}

.slide-up-enter-from, .slide-up-leave-to {
  opacity: 0;
  transform: translate(-50%, 20px);
}
</style>