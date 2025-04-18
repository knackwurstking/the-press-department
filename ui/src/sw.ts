const CACHE_NAME = "the-press-department-v1";
const urlsToCache = [
    process.env.SERVER_PATH + "/",
    process.env.SERVER_PATH + "/index.html",
    process.env.SERVER_PATH + "/main.umd.cjs",
    process.env.SERVER_PATH + "/sw.js",
    process.env.SERVER_PATH + "/the-press-department.wasm",
    process.env.SERVER_PATH + "/wasm_exec.js",
];

self.addEventListener("install", (event: any) => {
    event.waitUntil(
        caches.open(CACHE_NAME).then((cache) => {
            console.log("Opened cache");
            return cache.addAll(urlsToCache);
        }),
    );
});

self.addEventListener("fetch", (event: any) => {
    event.respondWith(
        fetch(event.request)
            .then((response) => {
                if (
                    !response ||
                    response.status !== 200 ||
                    response.type !== "basic"
                ) {
                    return response;
                }
                const responseToCache = response.clone();
                caches.open(CACHE_NAME).then((cache) => {
                    cache.put(event.request, responseToCache);
                });
                return response;
            })
            .catch(() => {
                return caches.match(event.request);
            }),
    );
});
