<template>
  <div class="song-item" :class="{ downloaded: isDownloaded }">
    <span class="song-name">{{ songName }}</span>
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
    isDownloaded: Boolean,
  },
};
</script>

<style scoped>
.song-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 15px;
  background-color: #2c2c2c;
  border-radius: 5px;
}

.song-name {
  font-size: 16px;
}

.status-indicator {
  font-size: 12px;
  padding: 3px 8px;
  border-radius: 10px;
  color: white;
}

.song-item.downloaded .status-indicator {
  background-color: #28a745;
}

.song-item:not(.downloaded) .status-indicator {
  background-color: #dc3545;
}
</style>