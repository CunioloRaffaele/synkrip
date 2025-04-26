<template>
    <div class="myLanding" :class="{ 'exit-animation': isExiting }">
        <div class="animated-background" :class="{ 'fade-out': isExiting }"></div>
        <div class="content" :class="{ 'content-exit': isExiting }">
            <h1 :class="{ 'fade-out': isExiting }">Welcome to SynkRip</h1>
            <p :class="{ 'fade-out': isExiting }">Select you library or create new one</p>
            <!--<button @click="startExitAnimation" :class="{ 'animate': animate, 'disabled': isExiting }">Enter</button>-->
        </div>
        <div class="myFooter" :class="{ 'fade-out': isExiting }">
            Created with ❤️ by Cuniolo Raffaele<br>
            <a href="https://github.com/CunioloRaffaele" target="_blank"
                style="color: rgba(255, 255, 255, 0.8); text-decoration: none;">
                Open sourced project at GitHub
            </a>
        </div>
        
    </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { EventsOn } from '../../wailsjs/runtime/runtime'

const router = useRouter()
const animate = ref(false)
const isExiting = ref(false)

EventsOn("LibSelected", () => {
    console.log("LibSelected")
    //startExitAnimation()
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
</script>

<style scoped>
.myLanding {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100vh;
    position: relative;
    overflow: hidden;
    color: white;
    text-align: center;
    transition: all 1.5s ease;
}

.exit-animation {
    transform: scale(1.1);
    opacity: 0;
}

.animated-background {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: linear-gradient(135deg,
            #667eea 0%,
            #764ba2 25%,
            #6B8DD6 50%,
            #8E37D7 75%,
            #667eea 100%);
    background-size: 400% 400%;
    animation: gradientAnimation 15s ease infinite;
    z-index: -1;
    transition: opacity 1.5s ease;
}

.fade-out {
    opacity: 0;
}

.content {
    position: relative;
    z-index: 1;
    transition: all 1.5s ease;
}

.content-exit {
    transform: translateY(-50px);
    opacity: 0;
}

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

.myLanding h1 {
    font-size: 3rem;
    margin-bottom: 2rem;
    transition: opacity 0.8s ease;
}

.myLanding button {
    padding: 1rem 2rem;
    font-size: 1.2rem;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    background-color: #ff7e5f;
    color: white;
    transition: background-color 0.3s ease;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.myLanding button:hover:not(.disabled) {
    background-color: #feb47b;
}

.myLanding button.disabled {
    cursor: default;
    pointer-events: none;
}

.myLanding button.animate {
    animation: buttonAnimation 0.8s forwards;
}

@keyframes buttonAnimation {
    0% {
        transform: scale(1);
    }

    50% {
        transform: scale(1.2);
    }

    100% {
        transform: scale(0);
        opacity: 0;
    }
}

.myFooter {
    position: absolute;
    bottom: 20px;
    width: 100%;
    text-align: center;
    font-size: 0.9rem;
    color: rgba(255, 255, 255, 0.8);
    z-index: 1;
    font-style: italic;
    text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3);
    transition: opacity 0.8s ease;
}
</style>