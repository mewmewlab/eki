import tailwindcss from "@tailwindcss/vite";
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [tailwindcss(), sveltekit()],
	server: {
		proxy: {
			'^/(api|_|-)/': {
				target: 'http://localhost:8090',
				changeOrigin: true,
			}
		}
	}
});
