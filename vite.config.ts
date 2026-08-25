import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  // vite-ssg reads this custom field
  // @ts-ignore
  ssgOptions: {
    dirStyle: 'nested',
  },
})
