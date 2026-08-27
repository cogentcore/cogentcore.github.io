<script setup lang="ts">
const features = [
  {
    title: 'All platforms',
    description: 'Deploy to web, desktop (macOS, Windows, Linux), and mobile (iOS, Android) from a single Go codebase.',
  },
  {
    title: 'Powerful features',
    description: 'Support for text editors, video players, interactive 3D graphics, customizable data plots, Markdown and HTML rendering, and SVG and canvas vector graphics.',
  },
  {
    title: 'Completely customizable',
    description: 'Allows developers and users to customize apps to fit their needs through a robust styling system and color settings.',
  },
  {
    title: 'Open source',
    description: 'Free and open under the BSD-3 License, allowing you to use Cogent Core for any purpose, commercially or personally.',
  },
]

const products = [
  {
    name: 'Docs',
    href: 'https://cogentcore.org/core',
    description: 'Interactive documentation for the framework, built with Cogent Core and deployed using WebAssembly.',
  },
  {
    name: 'Lab',
    href: 'https://cogentcore.org/lab',
    description: 'An extensible math, data science, and statistics platform.',
  },
  {
    name: 'Apps',
    href: 'https://cogentcore.org/cogent',
    description: 'Apps built with Cogent Core.',
  },
]
</script>

<template>
  <!-- Hero -->
  <section class="flex flex-col items-center justify-center min-h-screen px-6 pt-14 pb-16 text-center">
    <img src="/icon.svg" alt="Cogent Core" class="w-64 h-64 mb-8" />
    <img src="/name.png" alt="Cogent Core" class="w-[min(612px,80vw)] mb-8" />
    <p class="text-[28px] leading-[36px] text-gray-900 dark:text-gray-100 max-w-2xl mb-10">
      A cross-platform framework for building powerful, fast, elegant 2D and 3D apps
    </p>
    <div class="flex flex-wrap gap-3 justify-center">
      <a href="https://cogentcore.org/core" target="_blank"
        class="px-5 py-2.5 rounded-md bg-brand-blue text-white hover:bg-[#004299] dark:bg-brand-blue-muted dark:text-gray-900 dark:hover:bg-[#C5D8FF] font-medium transition-colors">
        Docs
      </a>
      <a href="https://cogentcore.org/core/install" target="_blank"
        class="px-5 py-2.5 rounded-md border border-gray-300 dark:border-gray-600 text-gray-700 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-800 font-medium transition-colors">
        Install
      </a>
    </div>
  </section>

  <!-- Features -->
  <section class="px-6 py-20 border-t border-gray-200 dark:border-gray-800">
    <div class="max-w-5xl mx-auto grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-8">
      <div v-for="feature in features" :key="feature.title">
        <h3 class="font-semibold text-gray-900 dark:text-gray-100 mb-2">{{ feature.title }}</h3>
        <p class="text-sm text-gray-500 dark:text-gray-400 leading-relaxed">{{ feature.description }}</p>
      </div>
    </div>
  </section>

  <!-- Products -->
  <section class="px-6 py-20 border-t border-gray-200 dark:border-gray-800">
    <div class="max-w-5xl mx-auto">
      <h2 class="text-xl font-semibold text-gray-900 dark:text-gray-100 mb-6">The ecosystem</h2>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
        <a v-for="product in products" :key="product.name" :href="product.href" target="_blank"
          rel="noopener noreferrer"
          class="group block p-6 rounded-lg border border-gray-200 dark:border-gray-700 hover:border-brand-blue dark:hover:border-brand-blue-muted hover:bg-gray-50 dark:hover:bg-gray-800/40 transition-colors">
          <h3
            class="font-semibold text-gray-900 dark:text-gray-100 mb-2 group-hover:text-brand-blue dark:group-hover:text-brand-blue-muted transition-colors">
            {{ product.name }} <span class="text-gray-400 dark:text-gray-500">→</span>
          </h3>
          <p class="text-sm text-gray-500 dark:text-gray-400 leading-relaxed">{{ product.description }}</p>
        </a>
      </div>
    </div>
  </section>
</template>
