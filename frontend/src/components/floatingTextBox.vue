<template>
    <div>
        <!-- Semi-transparent overlay -->
        <transition name="fade">
            <div v-if="isVisible" class="overlay" @click="cancel"></div>
        </transition>

        <!-- Floating text box -->
        <transition name="fade-scale">
            <div v-if="isVisible" class="floating-text-box">
                <!-- Selection bar -->
                <div class="selection-bar">
                    <div
                        v-for="option in options"
                        :key="option"
                        :class="['selection-option', { active: selectedOption === option }]"
                        @click="selectOption(option), console.log('Selected option:', option)"
                    >
                        {{ option }}
                    </div>
                </div>

                <div class="input-container">
                    <input 
                        type="text" 
                        :placeholder="placeholder" 
                        v-model="inputText" 
                        class="text-input" 
                        ref="inputField"
                        @keyup.enter="submit"
                    >
                </div>
                <button class="action-button submit" @click="submit">{{ submitButtonText }}</button>
            </div>
        </transition>
    </div>
</template>

<script>
export default {
    name: 'FloatingTextBox',
    props: {
        isVisible: {
            type: Boolean,
            default: false
        },
        title: {
            type: String,
            default: 'Add New Item'
        },
        placeholder: {
            type: String,
            default: 'Enter text here...'
        },
        submitButtonText: {
            type: String,
            default: 'Add'
        }
    },
    data() {
        return {
            inputText: '',
            options: ['Spotify'/*, 'SoundCloud', 'YouTube'*/],
            selectedOption: 'Spotify' // Default selection
        }
    },
    watch: {
        isVisible(newVal) {
            if (newVal) {
                // When box becomes visible, focus the input field and clear previous input
                this.$nextTick(() => {
                    this.inputText = '';
                    if (this.$refs.inputField) {
                        this.$refs.inputField.focus();
                    }
                });
            }
        }
    },
    methods: {
        selectOption(option) {
            this.selectedOption = option;
        },
        submit() {
            if (this.inputText.trim()) {
                this.$emit('submit', { text: this.inputText, source: this.selectedOption });
                this.inputText = '';
            }
        },
        cancel() {
            this.$emit('cancel');
        }
    }
}
</script>

<style scoped>
/* Floating text box styles */
.floating-text-box {
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 80%;
    max-width: 400px;
    padding: 2rem;
    border-radius: 16px;
    background: rgba(40, 40, 40, 0.85);
    backdrop-filter: blur(12px);
    -webkit-backdrop-filter: blur(12px);
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.3);
    z-index: 1000;
    display: flex;
    flex-direction: column;
    gap: 1.2rem;
    border: 1px solid rgba(255, 255, 255, 0.1);
}

/* Selection bar styles */
.selection-bar {
    display: flex;
    justify-content: space-around;
    margin-bottom: 1rem;
    background: rgba(255, 255, 255, 0.1);
    border-radius: 12px;
    padding: 0.5rem;
}

.selection-option {
    flex: 1;
    text-align: center;
    padding: 0.5rem 1rem;
    font-size: 1rem;
    font-weight: 500;
    color: rgba(255, 255, 255, 0.7);
    cursor: pointer;
    transition: all 0.3s ease;
    border-radius: 8px;
}

.selection-option:hover {
    background: rgba(255, 255, 255, 0.1);
    color: white;
}

.selection-option.active {
    background: rgba(255, 255, 255, 0.3);
    color: white;
    font-weight: bold;
}

/* Other styles remain unchanged */
.floating-text-box h3 {
    color: white;
    margin: 0;
    font-size: 1.5rem;
    font-weight: 500;
    text-shadow: 0 1px 2px rgba(0, 0, 0, 0.2);
}

.input-container {
    position: relative;
    width: 100%;
}

.text-input {
    width: 90%;
    padding: 14px 16px;
    font-size: 1.1rem;
    border: none;
    border-radius: 12px;
    background: rgba(255, 255, 255, 0.9);
    box-shadow: 
        inset 0 2px 4px rgba(0, 0, 0, 0.05),
        0 1px 2px rgba(255, 255, 255, 0.05);
    color: #333;
    transition: all 0.2s ease;
    outline: none;
}

.text-input:focus {
    background: white;
    box-shadow: 
        0 0 0 3px rgba(255, 255, 255, 0.2),
        inset 0 2px 4px rgba(0, 0, 0, 0.05);
}

.text-input::placeholder {
    color: #999;
    font-weight: 300;
}

.action-button {
    padding: 12px 20px;
    border: none;
    border-radius: 12px;
    font-weight: 600;
    font-size: 1rem;
    cursor: pointer;
    transition: all 0.3s ease;
    background-color: white;
    color: #333;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
    letter-spacing: 0.5px;
}

.submit:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
    background-color: #ffffff;
}

.submit:active {
    transform: translateY(0);
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    background-color: #f5f5f5;
}

/* Semi-transparent overlay */
.overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(0, 0, 0, 0.6);
    z-index: 888;
    backdrop-filter: blur(5px);
    -webkit-backdrop-filter: blur(5px);
}

/* Animations */
.fade-enter-active,
.fade-leave-active {
    transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
    opacity: 0;
}

.fade-scale-enter-active {
    animation: fade-scale-in 0.3s ease-out forwards;
}

.fade-scale-leave-active {
    animation: fade-scale-out 0.2s ease-in forwards;
}

@keyframes fade-scale-in {
    0% {
        opacity: 0;
        transform: translate(-50%, -50%) scale(0.9);
    }

    100% {
        opacity: 1;
        transform: translate(-50%, -50%) scale(1);
    }
}

@keyframes fade-scale-out {
    0% {
        opacity: 1;
        transform: translate(-50%, -50%) scale(1);
    }

    100% {
        opacity: 0;
        transform: translate(-50%, -50%) scale(0.9);
    }
}
</style>