<template>
  <footer class="border-t border-gray-200 dark:border-gray-800 py-8 px-4">
    <div
      class="max-w-5xl mx-auto flex flex-col sm:flex-row items-center justify-between gap-4 text-sm text-gray-500 dark:text-gray-500">
      <p>© {{ new Date().getFullYear() }} Cogent Core</p>
      <div class="flex items-center gap-5">
        <a href="https://github.com/cogentcore" target="_blank" rel="noopener noreferrer"
          class="hover:text-gray-900 dark:hover:text-gray-100 transition-colors">GitHub</a>
        <router-link to="/community" class="hover:text-gray-900 dark:hover:text-gray-100 transition-colors">
          Community
        </router-link>
        <router-link to="/faq" class="hover:text-gray-900 dark:hover:text-gray-100 transition-colors">
          FAQ
        </router-link>
        <a href="https://github.com/sponsors/cogentcore" target="_blank" rel="noopener noreferrer"
          class="hover:text-gray-900 dark:hover:text-gray-100 transition-colors">Sponsor</a>
      </div>
    </div>
  </footer>
</template>
