var CACHE_NAME = "the-press-department-v1";
var urlsToCache = ["/", "/the-press-department.wasm", "/wasm_exec.js"];
self.addEventListener("install", function (event) {
    console.warn(event);
    event.waitUntil(caches.open(CACHE_NAME).then(function (cache) {
        console.log("Opened cache");
        return cache.addAll(urlsToCache);
    }));
});
self.addEventListener("fetch", function (event) {
    console.warn(event);
    event.respondWith(fetch(event.request)
        .then(function (response) {
        if (!response ||
            response.status !== 200 ||
            response.type !== "basic") {
            return response;
        }
        var responseToCache = response.clone();
        caches.open(CACHE_NAME).then(function (cache) {
            cache.put(event.request, responseToCache);
        });
        return response;
    })
        .catch(function () {
        return caches.match(event.request);
    }));
});
