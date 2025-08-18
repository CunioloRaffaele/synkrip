<template>
    <div class="os-libs">
        <header>
            <Header title="OS Libraries" />
        </header>
        <ul>
            <li v-for="lib in libraries" :key="lib.name">
                <strong>{{ lib.name }}</strong><br>{{ lib.description }}
            </li>
        </ul>
    </div>
</template>

<script setup>
import Header from '../components/header.vue';
import { ref } from 'vue';
import { useRouter } from 'vue-router';
const router = useRouter();

const libraries = ref([
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
]);

</script>

<style scoped>
.os-libs {
    padding: 2rem 3rem;
    color: #e0e0e0;
    min-height: 100vh;
}
ul {
    list-style-type: none;
    padding: 0;
}

li {
    margin: 10px 0;
}

</style>