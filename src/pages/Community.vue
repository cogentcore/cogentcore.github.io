<script setup lang="ts">
const channels = [
  {
    title: 'Slack',
    description: '#cogentcore channel on the Gophers Slack',
    links: [
      { label: '#cogentcore', href: 'https://gophers.slack.com/archives/C07ENRTB2F7' },
      { label: 'Get invite', href: 'https://invite.slack.golangbridge.org' },
    ],
  },
  {
    title: 'GitHub',
    description: 'Issues and discussions in the repositories',
    links: [
      { label: 'github.com/cogentcore', href: 'https://github.com/cogentcore' },
    ],
  },
  {
    title: 'Mailing list',
    description: 'Announcements and discussion via Google Groups',
    links: [
      { label: 'groups.google.com/g/cogentcore', href: 'https://groups.google.com/g/cogentcore' },
    ],
  },
  {
    title: 'General email',
    description: 'For private inquiries and consulting',
    links: [
      { label: 'contact@cogentcore.org', href: 'mailto:contact@cogentcore.org' },
    ],
  },
  {
    title: 'Security',
    description: 'For reporting security vulnerabilities',
    links: [
      { label: 'security@cogentcore.org', href: 'mailto:security@cogentcore.org' },
    ],
  },
]
</script>

<template>
  <div class="pt-28 pb-20 px-6">
    <div class="max-w-2xl mx-auto">
      <h1 class="text-3xl font-bold text-gray-900 dark:text-gray-50 mb-3">Community</h1>
      <p class="text-gray-500 dark:text-gray-400 mb-10">
        Connect with the Cogent Core community through any of the following channels.
      </p>
      <div class="space-y-3">
        <div
          v-for="channel in channels"
          :key="channel.title"
          class="flex flex-col sm:flex-row sm:items-start sm:justify-between gap-3 p-5 rounded-lg border border-gray-200 dark:border-gray-700"
        >
          <div>
            <h3 class="font-semibold text-gray-900 dark:text-gray-100 mb-0.5">{{ channel.title }}</h3>
            <p class="text-sm text-gray-500 dark:text-gray-400">{{ channel.description }}</p>
          </div>
          <div class="flex flex-col items-start sm:items-end gap-1 shrink-0">
            <a
              v-for="link in channel.links"
              :key="link.href"
              :href="link.href"
              target="_blank"
              rel="noopener noreferrer"
              class="text-sm text-brand-blue dark:text-brand-blue-muted hover:underline"
            >
              {{ link.label }}
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
