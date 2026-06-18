<script setup lang="ts">
import { ref } from 'vue'

const menuOpen = ref(false)

const externalLinks = [
  { label: 'Blog', href: 'https://cogentcore.org/blog' },
  { label: 'Videos', href: 'https://youtube.com/@CogentCore' },
  { label: 'GitHub', href: 'https://github.com/cogentcore' },
]
</script>

<template>
  <nav class="fixed top-0 inset-x-0 z-50 bg-white/90 dark:bg-[#121316]/90 backdrop-blur-md border-b border-gray-200 dark:border-gray-800">
    <div class="max-w-5xl mx-auto px-4 h-14 flex items-center justify-between">
      <a
        href="/"
        class="flex items-center gap-2 font-semibold text-gray-900 dark:text-gray-100 hover:opacity-75 transition-opacity"
      >
        <img src="/icon.svg" alt="Cogent Core" class="w-7 h-7" />
        Cogent Core
      </a>

      <!-- Desktop nav -->
      <div class="hidden md:flex items-center gap-1 text-sm font-medium">
        <a
          v-for="link in externalLinks"
          :key="link.href"
          :href="link.href"
          target="_blank"
          rel="noopener noreferrer"
          class="px-3 py-1.5 rounded text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-100 hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
        >
          {{ link.label }}
        </a>
        <router-link
          to="/community"
          class="px-3 py-1.5 rounded text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-100 hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
        >
          Community
        </router-link>
        <a
          href="https://github.com/sponsors/cogentcore"
          target="_blank"
          rel="noopener noreferrer"
          class="ml-1 px-3 py-1.5 rounded text-sm border border-brand-blue dark:border-brand-blue-muted text-brand-blue dark:text-brand-blue-muted hover:bg-brand-blue hover:text-white dark:hover:bg-brand-blue-muted dark:hover:text-gray-900 transition-colors"
        >
          Sponsor
        </a>
      </div>

      <!-- Mobile hamburger -->
      <button
        class="md:hidden p-2 rounded text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-100 hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
        :aria-expanded="menuOpen"
        aria-label="Toggle menu"
        @click="menuOpen = !menuOpen"
      >
        <svg v-if="!menuOpen" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
        </svg>
        <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>
    </div>

    <!-- Mobile menu -->
    <div v-if="menuOpen" class="md:hidden border-t border-gray-200 dark:border-gray-800 px-4 pb-3">
      <div class="flex flex-col gap-1 pt-2 text-sm font-medium">
        <a
          v-for="link in externalLinks"
          :key="link.href"
          :href="link.href"
          target="_blank"
          rel="noopener noreferrer"
          class="px-3 py-2 rounded text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-100 hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
          @click="menuOpen = false"
        >
          {{ link.label }}
        </a>
        <router-link
          to="/community"
          class="px-3 py-2 rounded text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-100 hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
          @click="menuOpen = false"
        >
          Community
        </router-link>
        <a
          href="https://github.com/sponsors/cogentcore"
          target="_blank"
          rel="noopener noreferrer"
          class="px-3 py-2 rounded text-brand-blue dark:text-brand-blue-muted hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
          @click="menuOpen = false"
        >
          Sponsor
        </a>
      </div>
    </div>
  </nav>
</template>
