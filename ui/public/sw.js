const s = "the-press-department-v1", n = [
  "/test/",
  "/test/index.html",
  "/test/main.umd.cjs",
  "/test/style.css",
  "/test/sw.js",
  "/test/the-press-department.wasm",
  "/test/wasm_exec.js"
];
self.addEventListener("install", (t) => {
  console.warn(t), t.waitUntil(
    caches.open(s).then((e) => (console.log("Opened cache"), e.addAll(n)))
  );
});
self.addEventListener("fetch", (t) => {
  console.warn(t), t.respondWith(
    fetch(t.request).then((e) => {
      if (!e || e.status !== 200 || e.type !== "basic")
        return e;
      const c = e.clone();
      return caches.open(s).then((a) => {
        a.put(t.request, c);
      }), e;
    }).catch(() => caches.match(t.request))
  );
});
