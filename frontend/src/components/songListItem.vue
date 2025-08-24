<template>
  <div class="song-item" :class="{ downloaded: isDownloaded }">
    <div class="song-details">
      <span class="song-artist">{{ songArtist }}</span>
      <span class="separator">-</span>
      <span class="song-name">{{ songName }}</span>
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

.song-item:not(.downloaded) .status-indicator {
  background-color: #dc3545;
}
</style>