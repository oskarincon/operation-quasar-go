import react from '@vitejs/plugin-react';
import path from 'path';
import reactRefresh from '@vitejs/plugin-react-refresh'
import { defineConfig, loadEnv } from 'vite';

export default ({ mode }: { mode: string }) => {
  return defineConfig({
    plugins: [react(), reactRefresh()],
    resolve: {
      alias: {
        '@': path.resolve(__dirname, './src')
      }
    },
    define: { 'process.env': { ...loadEnv(mode, process.cwd()) } },
    server: {
      host: '0.0.0.0',
      port: 3000,
    }
  });
};
