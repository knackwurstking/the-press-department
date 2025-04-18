import "./sw-register";

import { KeepAwake } from "@capacitor-community/keep-awake";

KeepAwake.isSupported().then((result) => {
    result.isSupported && KeepAwake.keepAwake();
});

if ("serviceWorker" in navigator) {
    window.addEventListener("load", async () => {
        const registration = await navigator.serviceWorker.register("/sw.js");

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
