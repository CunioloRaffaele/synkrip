<template>
    <div class="os-libs">
        <div class="header">
            <button @click="goBack" class="back-button">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                    stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                    class="arrow-icon">
                    <path d="M19 12H5"></path>
                    <path d="M12 19l-7-7 7-7"></path>
                </svg>
                <span>Back</span>
            </button>
        </div>
        <ul>
            <li v-for="lib in libraries" :key="lib.name">
                <strong>{{ lib.name }}</strong><br>{{ lib.description }}
            </li>
        </ul>
    </div>
</template>

<script>
export default {
    name: "OsLibs",
    data() {
        return {
            libraries: [
                {
                    name: "Go",
                    description: "The Go programming language.",
                },
                {
                    name: "Wails",
                    description: "A framework for building desktop applications using Go and web technologies.",
                },
                {
                    name: "GJSON",
                    description: "A Go package that provides a fast and simple way to get values from a JSON document.",
                },
                {
                    name: "yt-dlp",
                    description: "A feature-rich command-line audio/video downloader ",
                },
                {
                    name: "FFmpeg",
                    description: "A complete, cross-platform solution to record, convert and stream audio and video.",
                },
                {
                    name: "mattn/go-sqlite3",
                    description: "A Go SQLite3 driver for the database/sql package.",
                },
                {
                    name: "Vue.js",
                    description: "The progressive JavaScript framework.",
        
                },
                {
                    name: "Tailwind CSS",
                    description: "A utility-first CSS framework for rapidly building custom designs.",
                },
            ],
        };
    },

    methods: {
            goBack() {
                this.$router.go(-1); // Navigate to previous page in history
            }
        }
};
</script>

<style scoped>
ul {
    list-style-type: none;
    padding: 0;
}

li {
    margin: 10px 0;
}

.header {
    display: flex;
    align-items: center;
    margin-bottom: 20px;
}

.back-button {
    display: flex;
    align-items: center;
    background: none;
    border: none;
    cursor: pointer;
    color: #4a90e2;
    padding: 8px 12px;
    margin-right: 16px;
    border-radius: 4px;
    transition: background-color 0.5s;
    font-size: 16px;
}

.back-button:hover {
    background-color: rgba(74, 144, 226, 0.1);
}

.back-button span {
    margin-left: 4px;
}

.arrow-icon {
    width: 20px;
    height: 20px;
}

</style>