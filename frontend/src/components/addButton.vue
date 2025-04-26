<template>
    <button @click="toggleButton" class="fab-button" :class="{ 'is-active': isActive }">
        <span class="fab-icon"></span>
    </button>
</template>

<script>
export default {
    name: 'FloatingActionButton',
    props: {
        initialActive: {
            type: Boolean,
            default: false
        },
        color: {
            type: String,
            default: '#ff4136' // Default red color
        }
    },
    data() {
        return {
            isActive: this.initialActive
        }
    },
    methods: {
        toggleButton() {
            this.isActive = !this.isActive;
            // Emit an event that parent components can listen to
            this.$emit('toggle', this.isActive);
        }
    },
    watch: {
        // Watch for changes in the initialActive prop
        initialActive(newValue) {
            this.isActive = newValue;
        }
    }
}
</script>

<style scoped>
/* Floating action button */
.fab-button {
    position: fixed;
    bottom: 2rem;
    right: 2rem;
    width: 4rem;
    height: 4rem;
    border-radius: 50%;
    background-color: v-bind('color');
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2);
    border: none;
    cursor: pointer;
    z-index: 999;
    transition: all 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275);
    display: flex;
    align-items: center;
    justify-content: center;
}

.fab-button:hover {
    transform: scale(1.05);
    box-shadow: 0 6px 14px rgba(0, 0, 0, 0.25);
}

.fab-button:active {
    transform: scale(0.95);
}

/* Plus/X icon */
.fab-icon {
    position: relative;
    width: 50%;
    height: 2px;
    background-color: white;
    transition: all 0.3s ease;
}

.fab-icon::before,
.fab-icon::after {
    content: '';
    position: absolute;
    width: 100%;
    height: 100%;
    background-color: white;
    transition: all 0.3s ease;
    left: 0;
}

.fab-icon::before {
    transform: rotate(90deg);
}

/* Animation for X */
.fab-button.is-active .fab-icon {
    background-color: transparent;
}

.fab-button.is-active .fab-icon::before {
    transform: rotate(45deg);
}

.fab-button.is-active .fab-icon::after {
    transform: rotate(-45deg);
}
</style>