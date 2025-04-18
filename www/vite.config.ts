import type { UserConfig } from "vite";

export default {
    build: {
        emptyOutDir: false,
        outDir: "public/",
        lib: {
            name: "main",
            entry: "src/main.ts",
            fileName: "main",
        },
    },
} satisfies UserConfig;
