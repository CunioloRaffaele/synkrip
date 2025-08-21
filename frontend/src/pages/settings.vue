<template>
    <div class="settings-page">

        <header>
            <Header title="Settings" />
        </header>

        <main class="settings-content">
            <!--<section class="settings-section">
                <h2 class="section-title">Download Options</h2>
                <div class="setting-item">
                    <label for="download-quality">Download file format</label>
                    <p class="setting-description">Select one or more preferred formats.</p>
                    <select id="download-quality" v-model="selectedQualities" multiple class="multi-select-box">
                        <option value="m4a">m4a</option>
                    </select>
                </div>
            </section>

            <section class="settings-section">
                <h2 class="section-title">Download order</h2>
                <div class="setting-item checkbox-item">
                    <label for="enable-ordered">Name file after playlist specific order</label>
                    <div class="checkbox-wrapper">
                        <input type="checkbox" id="enable-ordered" v-model="enableOrdered"
                            class="custom-checkbox" />
                        <span class="checkbox-custom-display" aria-hidden="true"></span>
                    </div>
                </div>
            </section>

            <section class="settings-section">
                <h2 class="section-title">Cache size</h2>
                <div class="setting-item">
                    <label for="cache-size">mb</label>
                </div>
            </section>-->
        </main>

        <!-- User identifier in bottom right -->
        <div class="user-identifier" v-if="userIdentifier">
            <span>{{ userIdentifier }}</span>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import Header from '../components/header.vue';
import { GetSettings } from '../../wailsjs/go/main/App';

const router = useRouter();

// Reactive variables for all settings
const selectedQualities = ref(['m4a']);
const enableOrdered = ref(true);
const cacheSize = ref(500);
const userIdentifier = ref('userId');

// Fetch settings on component mount
onMounted(() => {
    fetchSettings();
});

function fetchSettings() {
    return GetSettings()
        .then((result) => {
            const settings = JSON.parse(result);
            console.log("Fetched settings:", settings);
            userIdentifier.value = settings.userIdentifier;
        })
        .catch((error) => {
            console.error('Error fetching settings:', error);
        });
}
</script>

<style scoped>
.settings-page {
    padding: 2rem 3rem;
    color: #e0e0e0;
    min-height: 100vh;
}

.settings-section {
    background-color: rgba(255, 255, 255, 0.1);
        border-radius: 12px;
        box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2);
        padding: 1.5rem;
    margin-bottom: 2rem;
}

.section-title {
    font-size: 1.5rem;
    margin-top: 0;
    margin-bottom: 1.5rem;
    color: #ffffff;
    padding-bottom: 0.8rem;
}

.setting-item {
    display: flex;
    flex-direction: column;
}

.setting-item label {
    font-size: 1.1rem;
    font-weight: 600;
    margin-bottom: 0.5rem;
}

.setting-description {
    font-size: 0.9rem;
    color: #a0a0c0;
    margin-top: 0;
    margin-bottom: 1rem;
}

/* Multiple Selection Box */
.multi-select-box {
    background-color: rgba(255, 255, 255, 0.1);
    color: #e0e0e0;
    border-radius: 8px;
    padding: 0.8rem;
    font-size: 1rem;
    min-height: 120px; /* Ensure it's visible */
}

.multi-select-box option {
    padding: 0.8rem;
}

/* Checkbox */
.checkbox-item {
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
}

.checkbox-wrapper {
    position: relative;
    display: inline-block;
}

.custom-checkbox {
    opacity: 0;
    width: 0;
    height: 0;
    position: absolute;
}

.checkbox-custom-display {
    position: relative;
    display: inline-block;
    width: 50px;
    height: 28px;
    background-color: #000000;
    border-radius: 14px;
    cursor: pointer;
    transition: background-color 0.3s ease;
}

.checkbox-custom-display::before {
    content: '';
    position: absolute;
    width: 22px;
    height: 22px;
    border-radius: 50%;
    background-color: white;
    top: 3px;
    left: 4px;
    transition: transform 0.3s ease;
}

.custom-checkbox:checked+.checkbox-custom-display {
    background-color: #000000;
}

.custom-checkbox:checked+.checkbox-custom-display::before {
    transform: translateX(21px);
}

.user-identifier {
    font-family: Arial, Helvetica, sans-serif;
    padding: 0.5rem 1rem;
    font-size: 0.9rem;
    color: #ffffff;
}

.user-identifier span {
    font-style: italic;
}
</style>