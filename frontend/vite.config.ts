import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import Icons from 'unplugin-icons/vite'
import { FileSystemIconLoader } from 'unplugin-icons/loaders'
import tailwindcss from '@tailwindcss/vite'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    tailwindcss(),
    Icons({
      compiler: 'vue3',
      customCollections: {
        'mdi-outlined': FileSystemIconLoader(
          './node_modules/@material-design-icons/svg/outlined',
          (svg) => svg.replace(/^<svg /, '<svg fill="currentColor" '),
        ),
        'mdi-filled': FileSystemIconLoader(
          './node_modules/@material-design-icons/svg/filled',
          (svg) => svg.replace(/^<svg /, '<svg fill="currentColor" '),
        ),
      },
      iconCustomizer(collection, icon, props) {
        props.width = '2em'
        props.height = '2em'
      },
    }),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
  server: {
    host: '0.0.0.0',
    port: 5173,
  },
})
