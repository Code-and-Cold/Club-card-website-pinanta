import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import { playwright } from '@vitest/browser-playwright'
import vue from '@vitejs/plugin-vue'

// https://vite.dev/config/
export default defineConfig({
  base: "/Club-card-website-pinanta/",
  plugins: [vue()],

  test: {
    browser: {
      enabled: true,
      provider: playwright(),
      instances: [
        { browser: 'chromium' },
      ],
    },

    environment: 'jsdom',
    globals: true,
    include: ['test/**/*.{test,spec}.{js,ts,jsx,tsx}'],
    exclude: ['node_modules', 'dist', '.idea', '.git', '.cache'],
    testTimeout: 10000,
  },

  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
})
