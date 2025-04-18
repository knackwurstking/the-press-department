import { KeepAwake } from "@capacitor-community/keep-awake";

if ("serviceWorker" in navigator) {
    window.addEventListener("load", async () => {
        const registration = await navigator.serviceWorker.register(
            process.env.SERVER_PATH + "/sw.js",
        );

        try {
            console.log(
                "ServiceWorker registration successful with scope: ",
                registration.scope,
            );
        } catch (error) {
            console.log("ServiceWorker registration failed: ", error);
        }
    });
}

KeepAwake.isSupported().then((result) => {
    result.isSupported && KeepAwake.keepAwake();
});
