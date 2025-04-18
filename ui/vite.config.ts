import type { UserConfig } from "vite";

export default {
    base: process.env.THEPRESSDEPARTMENT_SERVER_PATH,

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
