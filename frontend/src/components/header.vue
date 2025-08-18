<template>
    <header class="app-header">
        <button @click="goBack" class="back-button" aria-label="Go back">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M19 12H5"></path>
                <path d="M12 19l-7-7 7-7"></path>
            </svg>
            <span>Back</span>
        </button>
        <h1>{{ title }}</h1>
    </header>
</template>

<script>
export default {
    name: 'header',
    props: {
        title: {
            type: String,
            default: 'Synkrip'
        }
    },
    methods: {
        goBack() {
            this.$router.go(-1);
        }
    }
}
</script>

<style scoped>
.app-header {
    display: flex;
    align-items: center;
    margin-bottom: 2.5rem;
    padding-bottom: 1.5rem;
}

.app-header h1 {
    font-size: 2.5rem;
    font-weight: 700;
    margin: 0;
    color: #ffffff;
}

.back-button {
    display: flex;
    align-items: center;
    background: black;
    border: black;
    color: #ffffff;
    padding: 0.6rem 1.2rem;
    border-radius: 8px;
    cursor: pointer;
    font-size: 1rem;
    margin-right: 1.5rem;
    transition: all 0.2s ease-in-out;
}

.back-button:hover {
    box-shadow: 0 0 10px 3px rgba(255, 255, 255, 0.1);
    filter: brightness(1.15);
    transform: scale(1.03);
}

.back-button:active {
    filter: brightness(0.9);
    transform: scale(0.98);
    transition: all 0.2s ease-in-out;
}

.back-button svg {
    margin-right: 0.5rem;
    transition: transform 0.2s ease-in-out;
}

.back-button:active svg {
    transform: translateX(-8px);
}
</style>