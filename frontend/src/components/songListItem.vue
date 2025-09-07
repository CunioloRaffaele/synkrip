<template>
  <div class="song-item" :class="{ downloaded: isDownloaded, duplicate: isDuplicate, 'has-error': hasError }">
    <div class="song-details">
      <span class="song-artist">{{ songArtist }}</span>
      <span class="separator">-</span>
      <span class="song-name">{{ songName }}</span>
      <!-- Show ERROR badge if YT ID is missing -->
      <span v-if="hasError" class="error-badge" title="Missing YouTube ID. Song cannot be downloaded.">ERROR</span>
      <!-- Show DUPLICATE badge only if there is no error -->
      <span v-else-if="isDuplicate" class="duplicate-badge">Duplicate</span>
    </div>
    <span class="status-indicator">
      {{ isDownloaded ? 'Downloaded' : 'Not Downloaded' }}
    </span>
  </div>
</template>

<script>
export default {
  name: 'SongListItem',
  props: {
    songName: String,
    songArtist: String,
    isDownloaded: Boolean,
    isDuplicate: Boolean,
    songYtId: String,
  },
  computed: {
    // Computed property to check for the error condition
    hasError() {
      return !this.songYtId || this.songYtId.trim() === '';
    },
  },
};
</script>

<style scoped>
.song-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 15px;
  background-color: #2c2c2c;
  border-radius: 8px;
  gap: 16px;
}

.song-details {
  display: flex;
  align-items: baseline;
  gap: 8px;
  min-width: 0;
}

.song-artist {
  font-size: 14px;
  color: #b3b3b3;
  font-weight: 500;
  flex-shrink: 0;
}

.separator {
  color: #666;
}

.song-name {
  font-size: 16px;
  color: #ffffff;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.error-badge {
  font-size: 10px;
  font-weight: 700;
  color: #fff;
  background-color: #dc3545;
  padding: 2px 5px;
  border-radius: 4px;
  margin-left: 10px;
  flex-shrink: 0;
  line-height: 1;
  cursor: help; 
}

.duplicate-badge {
  font-size: 10px;
  font-weight: 700;
  color: #111;
  background-color: #ffc107;
  padding: 2px 5px;
  border-radius: 4px;
  margin-left: 10px;
  flex-shrink: 0;
  line-height: 1;
}

.status-indicator {
  font-size: 12px;
  padding: 3px 8px;
  border-radius: 10px;
  color: white;
  flex-shrink: 0; /* Prevents indicator from shrinking */
}

.song-item.downloaded .status-indicator {
  background-color: #28a745;
}

/* De-emphasize the entire duplicate item */
.song-item.duplicate {
  opacity: 0.6;
  transition: opacity 0.2s ease-in-out;
}

.song-item:not(.downloaded) .status-indicator {
  background-color: #dc3545;
}
</style>