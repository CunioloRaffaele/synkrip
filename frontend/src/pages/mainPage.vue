<template>
  <FloatingDownloadStatusAndCancel :downloadText=downloadTxt :isVisible="isDownloadVisible" :showDetails=true
    :currentItem="downloadCurrentItem" :totalItems="downloadTotalItems" @cancel-download="StopSync()" />
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
import FloatingActionButton from '../components/addButton.vue';  
import FloatingTextBox from '../components/floatingTextBox.vue';
import PlaylistCover from '../components/playlistCover.vue';
import FloatingDownloadStatusAndCancel from '../components/floatingDownloadStatus.vue';

import { GetDB, AddEntry, SyncPlaylist, StopSync } from '../../wailsjs/go/main/App.js';
import { EventsOn } from '../../wailsjs/runtime/runtime'

export default {
  name: 'mainPage',
  components: {
    FloatingActionButton,
    FloatingTextBox,
    PlaylistCover,
    FloatingDownloadStatusAndCancel
  },
  data() {
    return {
      fabActive: false,
      fabColor: '#000000',
      downloadTxt: '',
      isDownloadVisible: false,
      downloadCurrentItem: 0,
      downloadTotalItems: 0,
      playlists: []
    }
  },
  methods: {
    viewPlaylistDetails(playlistId) {
      this.$router.push({ name: 'PlaylistDetailPage', params: { id: playlistId } });
    },
    handleFabToggle(isActive) {
      this.fabActive = isActive;
      console.log('FAB is now:', isActive ? 'active' : 'inactive');
      this.fabColor = isActive ? '#ff4136' : '#000000';
    },
    handleSubmit(obj) {
      const text = obj.text;
      const source = obj.source;
      // Check if the text is empty
      if (!text) {
        return;
      }
      // Process the submitted text
      AddEntry(text, source)
        .then((result) => {
        console.log("Result of AddEntry:", result);
        this.fetchPlaylists();
      }).catch((error) => {
        console.error('Error adding entry:', error);
      });
      console.log('Submitted:', text, 'from', source);
      // Then close the input box
      this.fabActive = false;
      this.fabColor = '#000000';
    },
    handleSync(playlistId) {
      console.log("Syncing playlist:", playlistId);
      SyncPlaylist(playlistId)
        .then((result) => {
          console.log("Result of SyncPlaylist:", result);
          this.fetchPlaylists();
        })
        .catch((error) => {
          console.error('Error syncing playlist:', error);
        });
    },
    StopSync() {
      StopSync() // Call the imported StopSync function
        .then(() => {
          console.log("Sync process stopped successfully.");
        })
        .catch((error) => {
          console.error("Error stopping sync process:", error);
        });
    },
    updateDownloadStatus(text, isVisible, currentItem, totalItems) {
      this.downloadTxt = text;
      this.isDownloadVisible = isVisible;
      this.downloadCurrentItem = currentItem;
      this.downloadTotalItems = totalItems;
    },
    fetchPlaylists() {
      // Fetch playlists from the database
      GetDB()
        .then((result) => {
          const files = JSON.parse(result);
          console.log("Db:", files);
          this.playlists = files; // Update the playlists with the new data
        })
        .catch((error) => {
          console.error('Error fetching playlists:', error);
        });
    },
    // sync status
    getPlaylistSyncStatus(playlist) {
      // if download status is visible, hide toBeSynced
      if (this.isDownloadVisible) {
        return false;
      }
      // Otherwise show original value
      return playlist.toBeSynced;
    }
  },
  mounted() {
    this.fetchPlaylists();
    EventsOn("updateDownloadProgress", (data) => {
      console.log("Update download progress:", data);
      if (data && typeof data === 'object') {
        const text = data.title || "Downloading...";
        const isVisible = data.isVisible !== undefined ? data.isVisible : true;
        const currentItem = data.currentItem || 0;
        const totalItems = data.totalItems || 0;

        this.updateDownloadStatus(text, isVisible, currentItem, totalItems);
        this.fetchPlaylists();
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

