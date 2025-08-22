<template>
  <FloatingActionButton :color="fabColor" :initial-active="fabActive" @toggle="handleFabToggle" />
  <FloatingTextBox :is-visible="fabActive" title="Add New Playlist" placeholder="Enter public playlist url"
    submit-button-text="Add" @submit="handleSubmit" />
  <transition-group name="playlist" tag="div">
    <div v-for="(playlist, index) in playlists" :key="playlist.dir_id" class="playlist-item"
      :style="{ '--index': index }">
      <PlaylistCover :image="playlist.image" :title="playlist.dir_id"
        :songs="playlist.songs.map(song => ({ name: song.song_name, downloaded: song.is_downloaded === 1 }))"
        :needsSync="getPlaylistSyncStatus(playlist)" @sync-clicked="handleSync" 
        @cover-clicked="viewPlaylistDetails"
        :service="playlist.service" 
        layout="full"/>
    </div>
  </transition-group>
</template>

<script>
import { watch } from 'vue'; // Import watch from Vue
import FloatingActionButton from '../components/addButton.vue';  
import FloatingTextBox from '../components/floatingTextBox.vue';
import PlaylistCover from '../components/playlistCover.vue';

import { GetDB, AddEntry, SyncPlaylist, StopSync } from '../../wailsjs/go/main/App.js';
import { useDownloadStatus } from '../downloadStatus.js'; // Import the global store

export default {
  name: 'mainPage',
  components: {
    FloatingActionButton,
    FloatingTextBox,
    PlaylistCover,
  },
  // Use the setup() function to integrate the Composition API store
  setup() {
    // Get all the reactive state and functions from the global store
    const { isVisible, needsRefresh, resetRefreshFlag } = useDownloadStatus();
    
    // Expose them to the rest of the component (template, methods, etc.)
    return {
      isVisible,
      needsRefresh,
      resetRefreshFlag,
    };
  },
  data() {
    return {
      fabActive: false,
      fabColor: '#000000',
      // The old local download state is now completely removed
      playlists: []
    }
  },
  methods: {
    viewPlaylistDetails(playlistId) {
      this.$router.push({ name: 'PlaylistDetailPage', params: { id: playlistId } });
    },
    handleFabToggle(isActive) {
      this.fabActive = isActive;
      this.fabColor = isActive ? '#ff4136' : '#000000';
    },
    handleSubmit(obj) {
      const text = obj.text;
      const source = obj.source;
      if (!text) return;
      AddEntry(text, source)
        .then(() => this.fetchPlaylists())
        .catch(error => console.error('Error adding entry:', error));
      this.fabActive = false;
      this.fabColor = '#000000';
    },
    handleSync(playlistId) {
      console.log("Syncing playlist:", playlistId);
      SyncPlaylist(playlistId).catch(error => console.error('Error syncing playlist:', error));
    },
    StopSync() {
      StopSync().catch(error => console.error("Error stopping sync process:", error));
    },
    fetchPlaylists() {
      GetDB()
        .then((result) => {
          this.playlists = JSON.parse(result);
        })
        .catch((error) => {
          console.error('Error fetching playlists:', error);
        });
    },
    getPlaylistSyncStatus(playlist) {
      // Use the 'isVisible' from the global store (made available via setup())
      if (this.isVisible) {
        return false;
      }
      return playlist.toBeSynced;
    }
  },
  mounted() {
    this.fetchPlaylists();
    
    // This now works correctly because 'needsRefresh' and 'resetRefreshFlag'
    // are properly exposed from the setup() function.
    watch(() => this.needsRefresh, (newValue) => {
      if (newValue) {
        console.log("Download finished. Refreshing playlists...");
        this.fetchPlaylists();
        this.resetRefreshFlag(); // Reset the flag in the store
      }
    });
  }
}
</script>

<style scoped>
p {
  font-size: 20px;
}

/* Transition for playlist items */
.playlist-enter-active, .playlist-leave-active {
  transition: all 0.5s ease;
}
.playlist-enter-from {
  opacity: 0;
  transform: translateY(-20px);
}
.playlist-enter-to {
  opacity: 1;
  transform: translateY(0);
}
.playlist-leave-from {
  opacity: 1;
  transform: translateY(0);
}
.playlist-leave-to {
  opacity: 0;
  transform: translateY(-20px);
}

/* Staggered animation */
.playlist-item {
  animation: fadeIn 0.5s ease forwards;
  animation-delay: calc(var(--index) * 0.2s); /* Delay based on index */
  opacity: 0;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(-20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>

