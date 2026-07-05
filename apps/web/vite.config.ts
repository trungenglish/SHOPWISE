import tailwindcss from "@tailwindcss/vite";
import { tanstackRouter } from "@tanstack/router-plugin/vite";
import viteReact from "@vitejs/plugin-react";
import { defineConfig, loadEnv } from "vite";

const defaultServerUrl = "http://localhost:18080";

export default defineConfig(({ mode }) => {
  const environment = loadEnv(mode, process.cwd(), "");
  const configuredServerUrl = environment.VITE_SERVER_URL?.trim();
  const serverUrl =
    configuredServerUrl !== undefined && configuredServerUrl.length > 0
      ? configuredServerUrl
      : defaultServerUrl;

  return {
    plugins: [
      tailwindcss(),
      tanstackRouter({
        autoCodeSplitting: true,
        target: "react",
      }),
      viteReact(),
    ],
    resolve: {
      tsconfigPaths: true,
    },
    server: {
      port: 3001,
      proxy: {
        "/api": {
          changeOrigin: true,
          target: serverUrl,
        },
        "/health": {
          changeOrigin: true,
          target: serverUrl,
        },
        "/swagger": {
          changeOrigin: true,
          target: serverUrl,
        },
      },
    },
  };
});
