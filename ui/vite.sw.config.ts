import type { UserConfig } from "vite";

export default {
    base: process.env.THEPRESSDEPARTMENT_SERVER_PATH,

    build: {
        emptyOutDir: false,
        outDir: "public/",
        copyPublicDir: false,
        lib: {
            name: "main",
            entry: "src/sw.ts",
            fileName: "sw",
            formats: ["es"],
        },
    },

    define: {
        "process.env.SERVER_PATH": JSON.stringify(
            process.env.THEPRESSDEPARTMENT_SERVER_PATH,
        ),
    },
} satisfies UserConfig;
