<template>
    <div class="detail-page-container">
        <div class="left-pane">
            <PlaylistCover v-if="playlist" :image="playlist.image" :title="playlist.dir_id"
                :songs="playlist.songs.map(song => ({ name: song.song_name, downloaded: song.is_downloaded === 1 }))"
                :needsSync="false" @sync-clicked="handleSync" :service="playlist.service"
                class="sticky-cover" 
                layout="compact"/>
            
            <div v-if="playlist" class="playlist-dates">
                <div class="date-item">
                    <span class="date-label">Created:</span>
                    <span class="date-value">{{ formattedCreationDate }}</span>
                </div>
                <div class="date-item">
                    <span class="date-label">Synced:</span>
                    <span class="date-value">{{ formattedLastModifiedDate }}</span>
                </div>
            </div>

            <!-- Button Group -->
            <div class="button-group">
                <button v-if="playlist && playlist.url" @click="openInPlatform" class="action-button platform-button">
                    Open in {{ playlist.service }}
                </button>
                <button @click="goBack" class="action-button back-button">Back to Playlists</button>
            </div>
        </div>
        <div class="right-pane">
            <h2 v-if="playlist">{{ playlist.dir_id }}</h2>
            <div v-if="playlist" class="song-list">
                <SongListItem v-for="song in playlist.songs" :key="song.song_id" :song-name="song.song_name" :song-artist="song.song_artist_name" :is-downloaded="song.is_downloaded === 1" />
            </div>
            <div v-else>
                <p>Loading playlist details...</p>
            </div>
        </div>
    </div>
</template>

<script>
import PlaylistCover from '../components/playlistCover.vue';
import SongListItem from '../components/songListItem.vue';
import Header from '../components/header.vue';
import { GetDB, SyncPlaylist } from '../../wailsjs/go/main/App.js';

export default {
  name: 'PlaylistDetailPage',
  components: {
    PlaylistCover,
    SongListItem,
    Header
  },
  props: {
    id: {
      type: String,
      required: true,
    },
  },
  data() {
    return {
      playlist: null,
      currentTime: new Date(), // Add current time for reactivity
      interval: null,
    };
  },
  computed: {
    formattedCreationDate() {
      if (!this.playlist || !this.playlist.add_date) return 'N/A';
      const options = { year: 'numeric', month: 'long', day: 'numeric' };
      return new Date(this.playlist.add_date).toLocaleDateString(undefined, options);
    },
    formattedLastModifiedDate() {
      if (!this.playlist || !this.playlist.last_modified) return 'Never';
      // This makes the computed property reactive to the interval timer
      const now = this.currentTime; 
      return this.formatTimeAgo(this.playlist.last_modified);
    },
  },
  methods: {
    formatTimeAgo(dateString) {
      const date = new Date(dateString);
      const now = new Date();
      const seconds = Math.round((now - date) / 1000);

      const intervals = {
        year: 31536000,
        month: 2592000,
        week: 604800,
        day: 86400,
        hour: 3600,
        minute: 60
      };

      if (seconds < 30) return 'just now';

      for (const [unit, secondsInUnit] of Object.entries(intervals)) {
        const value = Math.floor(seconds / secondsInUnit);
        if (value >= 1) {
          return `${value} ${unit}${value > 1 ? 's' : ''} ago`;
        }
      }
      return `${Math.floor(seconds)} seconds ago`;
    },
    fetchPlaylistDetails() {
      GetDB()
        .then((result) => {
          const allPlaylists = JSON.parse(result);
          this.playlist = allPlaylists.find(p => p.dir_id === this.id);
          if (!this.playlist) {
            console.error("Playlist not found:", this.id);
            this.$router.push({ name: 'MainPage' });
          }
        })
        .catch((error) => {
          console.error('Error fetching playlist details:', error);
        });
    },
    handleSync(playlistId) {
      console.log("Syncing playlist:", playlistId);
      SyncPlaylist(playlistId)
        .then(() => {
          this.fetchPlaylistDetails(); // Re-fetch details to update the view
        })
        .catch((error) => {
          console.error('Error syncing playlist:', error);
        });
    },
    goBack() {
      this.$router.back();
    },
    openInPlatform() {
      if (this.playlist && this.playlist.url) {
        window.runtime.BrowserOpenURL(this.playlist.url);
      }
    }
  },
  created() {
    this.fetchPlaylistDetails();
    // Set up an interval to update the 'currentTime' every minute
    this.interval = setInterval(() => {
      this.currentTime = new Date();
    }, 60000); // 60000 ms = 1 minute
  },
  beforeUnmount() {
    // Clear the interval when the component is destroyed to prevent memory leaks
    clearInterval(this.interval);
  },
};
</script>

<style scoped>
.detail-page-container {
  display: flex;
  height: 100vh;
  width: 100%;
  overflow: hidden; /* CRITICAL: Prevents the page itself from scrolling */
}

.left-pane {
  flex: 0 0 300px;
  padding: 20px;
  display: flex;
  flex-direction: column;
  align-items: center;

  margin: 20px;

  /* CRITICAL: Hides any content that might overflow on the left pane */
  overflow: hidden; 

  /* --- Frosted Glass Effect --- */
  background: rgba(40, 40, 40, 0.6);
  backdrop-filter: blur(12px);
  -webkit-backdrop-filter: blur(12px);
  border-radius: 16px;
  border: 1px solid rgba(255, 255, 255, 0.1);
}

.sticky-cover {
  position: sticky;
  top: 20px;
}

.right-pane {
  flex: 1;
  padding: 20px;
  overflow-y: auto;
}

.right-pane h2 {
  margin-top: 0;
  border-bottom: 1px solid #333;
  padding-bottom: 10px;
  margin-bottom: 20px;
}

.song-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.back-button {
  margin-top: 20px;
  padding: 10px 20px;
  background-color: #333;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.back-button:hover {
  background-color: #555;
}

.playlist-dates {
  margin-top: 20px;
  width: 100%;
  text-align: left;
}

.date-item {
  display: flex;
  justify-content: space-between;
  margin-bottom: 10px;
}

.date-label {
  font-weight: bold;
}

.date-value {
  color: #666;
}

.button-group {
  display: flex;
  flex-direction: column;
  width: 100%;
  margin-top: 20px;
}

.action-button {
  padding: 10px 20px;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  margin-bottom: 10px;
}

.platform-button {
  background-color: #007bff;
}

.platform-button:hover {
  background-color: #0056b3;
}
</style>