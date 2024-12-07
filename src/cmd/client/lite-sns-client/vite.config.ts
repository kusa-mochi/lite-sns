import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'
import Pages from "vite-plugin-pages"

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    react(),
    Pages({
      dirs: [
        // basic
        { dir: 'src/pages', baseRoute: '' },
        // admin pages
        { dir: 'src/admin/pages', baseRoute: 'admin' },
      ],
    }),
  ],
  server: {
    host: '0.0.0.0',
    port: 8080,
  },
})
