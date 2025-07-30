<template>
    <div class="landing-page" :class="{ 'exit-animation': isExiting }">
        <div class="animated-background" :class="{ 'fade-out': isExiting }"></div>
        <div class="content" :class="{ 'content-exit': isExiting }">
            <h1 :class="{ 'fade-out': isExiting }">Welcome to SynkRip</h1>
            <p :class="{ 'fade-out': isExiting }">Select your library or create a new one</p>
        </div>

        <!-- Recently Used Libraries Box -->
        <transition-group name="fade" tag="div" class="recent-libraries" v-if="recentLibraries.length > 0">
            <h2>Recently Used Libraries</h2>
            <ul>
                <li v-for="(library, index) in recentLibraries" :key="library" class="library-item">
                    <button @click="openLibrary(library)">{{ library.split('/').pop() }}</button>
                </li>
            </ul>
        </transition-group>

        <div class="footer" :class="{ 'fade-out': isExiting }">
            <p>Created with ❤️ by Cuniolo Raffaele</p>
            <a href="https://github.com/CunioloRaffaele" target="_blank">
                Open sourced project at GitHub
            </a>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { EventsOn } from '../../wailsjs/runtime/runtime'
import { GetSettings, OpenLibrary } from '../../wailsjs/go/main/App'

const router = useRouter()
const animate = ref(false)
const isExiting = ref(false)
const recentLibraries = ref([]) // Make recentLibraries reactive

fetchRecentLibraries()

EventsOn("LibSelected", () => {
    console.log("LibSelected")
    router.push({ name: 'mainPage' })
})

function startExitAnimation() {
    if (isExiting.value) return;

    animate.value = true
    isExiting.value = true

    // Navigate after animation completes
    setTimeout(() => {
        router.push({ name: 'mainPage' })
    }, 1500) // Timing matches our exit animations
}

function fetchRecentLibraries() {
    // Fetch playlists from the database
    return GetSettings()
        .then((result) => {
            const settings = JSON.parse(result)
            console.log("Fetched settings:", settings)
            recentLibraries.value = settings.recentLibraries || [] // Update reactive variable
            console.log("Recent libraries:", recentLibraries.value)
        })
        .catch((error) => {
            console.error('Error fetching playlists:', error)
        })
}

function openLibrary(library) {
    console.log("Opening library:", library)
    return OpenLibrary(library)
        .then(() => {
            console.log("Library opened successfully:", library)
        })
        .catch((error) => {
            console.error('Error opening library:', error)
        })
}
</script>

<style scoped>
/* General layout */
.landing-page {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100vh;
    position: relative;
    overflow: hidden;
    color: white;
    text-align: center;
    padding: 2rem;
    box-sizing: border-box;
}

/* Animated background */
.animated-background {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    /*background: linear-gradient(135deg, #667eea 0%, #764ba2 25%, #6B8DD6 50%, #8E37D7 75%, #667eea 100%);
    background-size: 400% 400%;
    animation: gradientAnimation 15s ease infinite;
    z-index: -1;*/
}

/* Content styling */
.content {
    position: relative;
    z-index: 1;
    margin-bottom: 2rem;
}

.content h1 {
    font-size: 3rem;
    font-weight: bold;
    margin-bottom: 1rem;
}

.content p {
    font-size: 1.2rem;
    margin-bottom: 2rem;
}

/* Recently Used Libraries */
.recent-libraries {
    background-color: rgba(255, 255, 255, 0.1);
    border-radius: 12px;
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2);
    padding: 1.5rem;
    width: 100%;
    max-width: 600px;
    text-align: center;
    margin-bottom: 2rem;
}

.recent-libraries h2 {
    font-size: 1.8rem;
    margin-bottom: 1rem;
    color: #ffffff;
}

.recent-libraries ul {
    list-style: none;
    padding: 0;
    margin: 0;
}

.library-item {
    margin: 0.5rem 0;
    animation: fadeIn 0.5s ease forwards;
    animation-delay: calc(var(--index) * 0.2s);
    opacity: 0;
}

.recent-libraries button {
    padding: 0.8rem 1.5rem;
    font-size: 1rem;
    border: none;
    border-radius: 8px;
    cursor: pointer;
    background-color: #000000;
    color: rgba(255, 255, 255, 0.827);
    transition: background-color 0.3s ease, transform 0.2s ease;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.recent-libraries button:hover {
    box-shadow: 0 0 10px 3px rgba(255, 255, 255, 0.1);
    filter: brightness(1.15);
    transform: scale(1.08);
    color: white;
}

.recent-libraries button:active {
    filter: brightness(0.9);
    transform: scale(0.98);
    transition: all 0.2s ease-in-out;
}

/* Footer styling */
.footer {
    text-align: center;
    font-size: 0.9rem;
    color: rgba(255, 255, 255, 0.8);
    font-style: italic;
    text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3);
}

.footer a {
    color: rgba(255, 255, 255, 0.8);
    text-decoration: none;
    font-weight: bold;
    transition: color 0.3s ease;
}

.footer a:hover {
    color: #feb47b;
}

/* Animations */
@keyframes gradientAnimation {
    0% {
        background-position: 0% 50%;
    }
    50% {
        background-position: 100% 50%;
    }
    100% {
        background-position: 0% 50%;
    }
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