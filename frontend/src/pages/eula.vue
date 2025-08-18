<template>
    <div class="eula-page">
        <div class="content">
            <div class="eula-container">
                <h1>End User License Agreement</h1>
                <div class="eula-text">
                    <p>By continuing to use SynkRip, you acknowledge that you have read, understood, and agree to be
                        bound by the terms and conditions of the End User License Agreement.</p>
                    <p>The complete EULA can be found on the GitHub repository.</p>
                </div>

                <div class="github-link">
                    <a href="#" @click.prevent="openGitHub" class="github-button">
                        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24"
                            fill="currentColor">
                            <path
                                d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z" />
                        </svg>
                        View EULA on GitHub
                    </a>
                </div>

                <div class="action-buttons">
                    <!--<button @click="decline" class="decline-button">
                        Decline
                    </button>-->
                    <button @click="accept" class="accept-button">
                        Accept & Continue
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
const router = useRouter()

function accept() {
    router.push({ name: 'LandingPage' })
}
function openGitHub() {
    window.runtime.BrowserOpenURL('https://github.com/CunioloRaffaele/synkrip?tab=readme-ov-file#disclaimers')
}
</script>

<style scoped>
.eula-page {
    display: flex;
    align-items: center;
    justify-content: center;
    min-height: 100vh;
    position: relative;
    overflow: hidden;
    color: white;
    text-align: center;
    padding: 2rem;
    box-sizing: border-box;
}

.content {
    position: relative;
    z-index: 1;
    width: 100%;
    max-width: 600px;
}

.eula-container {
    background-color: rgba(255, 255, 255, 0.1);
    border-radius: 16px;
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
    backdrop-filter: blur(10px);
    border: 1px solid rgba(255, 255, 255, 0.2);
    padding: 3rem 2rem;
}

.eula-container h1 {
    font-size: 2.5rem;
    font-weight: bold;
    margin-bottom: 2rem;
    color: #ffffff;
}

.eula-text {
    margin-bottom: 2rem;
}

.eula-text p {
    font-size: 1.1rem;
    line-height: 1.6;
    margin-bottom: 1rem;
    color: rgba(255, 255, 255, 0.9);
}

.github-link {
    margin-bottom: 2.5rem;
}

.github-button {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.75rem 1.5rem;
    background-color: rgba(255, 255, 255, 0.15);
    color: white;
    text-decoration: none;
    border-radius: 8px;
    font-weight: 500;
    transition: all 0.3s ease;
    border: 1px solid rgba(255, 255, 255, 0.2);
}

.github-button:hover {
    background-color: rgba(255, 255, 255, 0.25);
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
}

.action-buttons {
    display: flex;
    gap: 1rem;
    justify-content: center;
}

.decline-button,
.accept-button {
    padding: 0.75rem 2rem;
    font-size: 1rem;
    font-weight: 600;
    border: none;
    border-radius: 8px;
    cursor: pointer;
    transition: all 0.3s ease;
    min-width: 120px;
}

.decline-button {
    background-color: rgba(255, 59, 48, 0.8);
    color: white;
}

.decline-button:hover {
    background-color: rgba(255, 59, 48, 1);
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(255, 59, 48, 0.3);
}

.accept-button {
    background-color: rgba(52, 199, 89, 0.8);
    color: white;
}

.accept-button:hover {
    background-color: rgba(52, 199, 89, 1);
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(52, 199, 89, 0.3);
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
</style>