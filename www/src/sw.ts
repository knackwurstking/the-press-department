const CACHE_NAME = "the-press-department-v1";
const urlsToCache = ["/", "/the-press-department.wasm", "/wasm_exec.js"];

self.addEventListener("install", (event: any) => {
    console.warn(event);
    event.waitUntil(
        caches.open(CACHE_NAME).then((cache) => {
            console.log("Opened cache");
            return cache.addAll(urlsToCache);
        }),
    );
});

self.addEventListener("fetch", (event: any) => {
    console.warn(event);
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
