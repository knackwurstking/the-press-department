import type { UserConfig } from "vite";

export default {
    build: {
        emptyOutDir: false,
        outDir: "public/",
        copyPublicDir: false,
        lib: {
            name: "main",
            entry: "src/main.ts",
            fileName: "main",
            formats: ["umd"],
        },
    },
} satisfies UserConfig;
