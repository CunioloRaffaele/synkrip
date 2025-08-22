import { ref, computed } from 'vue';
import { EventsOn } from '../wailsjs/runtime/runtime';

// --- Reactive State ---
const isVisible = ref(false);
const downloadText = ref('Preparing download...');
const currentItem = ref(0);
const totalItems = ref(0);
const needsRefresh = ref(false); // New state to signal a data refresh

// --- Wails Event Listener (Singleton) ---
// This runs once when the module is imported and listens for events globally.
EventsOn("updateDownloadProgress", (data) => {
  console.log("Global store received download progress:", data);
  if (data && typeof data === 'object') {
    downloadText.value = data.title || "Loading...";
    isVisible.value = data.isVisible !== undefined ? data.isVisible : true;
    currentItem.value = data.currentItem || 0;
    totalItems.value = data.totalItems || 0;

    // If the download is finished (isVisible is false), trigger a refresh.
    if (!isVisible.value) {
      needsRefresh.value = true;
      // Optional: Hide the bar after a delay to show the final message
      setTimeout(() => {
        // Reset state if needed, or just keep it hidden
      }, 2000);
    }
  }
});


// --- Composable Function ---
// Components will call this to get access to the state.
export function useDownloadStatus() {
  
  // Function to reset the refresh flag after a component has handled it
  const resetRefreshFlag = () => {
    needsRefresh.value = false;
  };

  const progressPercentage = computed(() => {
    if (totalItems.value <= 0) return 0;
    return Math.min(100, Math.floor((currentItem.value / totalItems.value) * 100));
  });

  return {
    isVisible,
    downloadText,
    currentItem,
    totalItems,
    progressPercentage,
    needsRefresh,
    resetRefreshFlag,
  };
}