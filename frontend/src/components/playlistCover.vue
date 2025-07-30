<template>
    <div class="playlist-entry" :style="{ backgroundImage: `url(${image})` }" :class="{ 'needs-sync': needsSync }"
        @click="toggleSelection">
        <!-- Frosted Glass Overlay -->
        <div v-if="isSelected" class="frosted-overlay"></div>

        <div class="blur-overlay"></div>
        <div class="sync-indicator" v-if="needsSync"></div>
        <div class="content-wrapper">
            <div class="playlist-image-container">
                <img :src="image" alt="Playlist Cover" class="playlist-image" />
                <!-- Service image overlay -->
                <img 
                    :src="(`/src/assets/images/${service}.png`)" 
                    alt="Service Icon" 
                    class="service-icon" 
                />
                <!-- Add sync icon overlay -->
                <div class="sync-icon" v-if="needsSync" @click="sync">
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24">
                        <path fill="#000000"
                            d="M12 4V1L8 5l4 4V6c3.31 0 6 2.69 6 6 0 1.01-.25 1.97-.7 2.8l1.46 1.46C19.54 15.03 20 13.57 20 12c0-4.42-3.58-8-8-8zm0 14c-3.31 0-6-2.69-6-6 0-1.01.25-1.97.7-2.8L5.24 7.74C4.46 8.97 4 10.43 4 12c0 4.42 3.58 8 8 8v3l4-4-4-4v3z" />
                    </svg>
                </div>
            </div>
            <div class="playlist-info">
                <div class="playlist-title">{{ title }}</div>
                <div class="song-list-container">
                    <ul class="song-list">
                        <li v-for="(song, index) in songs" :key="index" class="song-item"
                            :class="{ ' not-downloaded': !song.downloaded }">
                {{ song.name }}
                </li>
                </ul>
            </div>
        </div>
    </div>
    </div>
</template>

<script>
export default {
    name: "PlaylistCover",
    props: {
        title: {
            type: String,
            required: true,
        },
        image: {
            type: String,
            required: true,
        },
        songs: {
            type: Array,
            required: true,
            validator(value) {
                // Ensure each song is an object with `name` and `downloaded` properties
                return value.every(song => 'name' in song && 'downloaded' in song);
            },
        },
        needsSync: {
            type: Boolean,
            required: true,
        },
        service: {
            type: String,
            default: 'unknown', // Default to 'unknown' if not provided
        },
    },
    data() {
        return {
            isSelected: false,
        };
    },
    mounted() {
        //console.log("needsSync value:", this.needsSync);
    },
    methods: {
        sync() {
            // Emit an event to the parent with the playlist title or ID
            this.$emit('sync-clicked', this.title);
        },
        toggleSelection() {
            this.isSelected = !this.isSelected;
            console.log("Selected:", this.isSelected);
        },
    },
};
</script>

<style scoped>
.playlist-entry {
    position: relative;
    display: flex;
    align-items: center;
    border-radius: 16px;
    overflow: hidden;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    padding: 16px;
    width: 100%;
    max-width: 800px;
    margin: 16px auto;
    height: 180px;
    background-size: cover;
    background-position: center;
    transition: all 0.3s ease-in-out;
}

.playlist-entry:hover {
    box-shadow: 0 0 10px 3px rgba(255, 255, 255, 0.1);
    filter: brightness(1.15);
    transform: scale(1.03);
}

.playlist-entry:active {
    filter: brightness(0.9);
    transform: scale(0.98);
    transition: all 0.2s ease-in-out;
}

.needs-sync::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    border-radius: 16px;
    box-shadow: 0 0 0 2px rgba(255, 0, 0, 0);
    z-index: 4; /* Above all other elements */
    pointer-events: none;
}

.needs-sync::before {
    animation: border-trace 4s linear infinite;
}

@keyframes border-trace {
    0% {
        background-color: rgba(255, 0, 0, 0.0);
    }
    25% {
        background-color: rgba(255, 0, 0.4, 0.2);
    }
    50% {
        background-color: rgba(255, 0, 0.4, 0.3);
    }
    75% {
        background-color: rgba(255, 0, 0.4, 0.2);
    }
    100% {
        background-color: rgba(255, 0, 0, 0.0);
    }
}

.blur-overlay {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    backdrop-filter: blur(20px) brightness(0.7);
    -webkit-backdrop-filter: blur(20px) brightness(0.7);
    z-index: 1;
}

.frosted-overlay {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    backdrop-filter: blur(20px) brightness(1.2);
    -webkit-backdrop-filter: blur(20px) brightness(1.2);
    z-index: 1000;
}

.sync-indicator {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(255, 0, 0, 0.1);
    pointer-events: none;
}

.content-wrapper {
    position: relative;
    z-index: 5; /* Above the blur overlay */
    display: flex;
    width: 100%;
    height: 100%;
    align-items: center;
}

.playlist-image-container {
    position: relative;
    width: 150px;
    height: 150px;
    margin-right: 16px;
    flex-shrink: 0;
    transition: all 0.3s ease-in-out;
}

.playlist-image-container:hover {
    transform: scale(1.05);
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.3);
}

.playlist-image {
    width: 100%;
    height: 100%;
    object-fit: cover;
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.2);
}

.service-icon {
    position: absolute;
    top: 50%;
    left: 50%;
    width: 60px;
    height: 60px;
    transform: translate(-50%, -50%);
    z-index: 10;
    border-radius: 50%;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.8);
}

.sync-icon {
    cursor: pointer;
    position: absolute;
    top: -12px;
    right: -12px;
    width: 40px;
    height: 40px;
    background-color: #ff3b30;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.3);
    z-index: 5;
    animation: sync-rotate 2s linear infinite;
    transition: all 0.3s ease-in-out;
}

.sync-icon:hover {
    scale: 1.15;
}

.playlist-info {
    flex: 1;
    display: grid;
    grid-template-columns: 1fr 1fr;
    /* Create two equal columns */
    grid-gap: 16px;
    height: 150px;
    color: white;
}

.playlist-title {
    display: flex;
    align-items: center;
    font-size: 1.8rem;
    font-family: 'Avenir Next';
    font-weight: bold;
    white-space: nowrap;
    /* Prevent text from wrapping */
    overflow: hidden;
    text-overflow: ellipsis;
    justify-content: center;
}

.song-list-container {
    position: relative;
    overflow-y: auto;
    background: linear-gradient(rgba(0, 0, 0, 0.2), rgba(0, 0, 0, 0.4));
    backdrop-filter: blur(5px);
    box-shadow: inset 0 0 10px rgba(0, 0, 0, 0.2);
    border: 1px solid rgba(255, 255, 255, 0.1);
    border-radius: 8px;
    padding: 8px;
}

.song-list-container::-webkit-scrollbar {
    width: 8px;
}

.song-list-container::-webkit-scrollbar-track {
    background: rgba(255, 255, 255, 0);
    border-radius: 8px;
}

.song-list-container::-webkit-scrollbar-thumb {
    background: rgba(255, 255, 255, 0.3);
    border-radius: 8px;
}

.song-list-container::-webkit-scrollbar-thumb:hover {
    background: rgba(255, 255, 255, 0.5);
}

.song-list {
    list-style: none;
    padding: 0 0 16px 0;
    margin: 0;
}

.song-item {
    font-size: 1rem;
    margin-bottom: 8px;
    text-shadow: 0 1px 2px rgba(0, 0, 0, 0.6);
}

.song-item.not-downloaded {
    color: rgb(226, 56, 56);
}
</style>