import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import tailwindcss from 'tailwindcss'
// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    host: true,
    port: 8091
},
preview: {
    host: true,
    port: 8091
} ,
css: {
  postcss: {
    plugins: [tailwindcss()],
  },
}
  
})
