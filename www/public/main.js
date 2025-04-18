var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
import { KeepAwake } from "@capacitor-community/keep-awake";
KeepAwake.isSupported().then((result) => {
    result.isSupported && KeepAwake.keepAwake();
});
if ("serviceWorker" in navigator) {
    window.addEventListener("load", () => __awaiter(void 0, void 0, void 0, function* () {
        const registration = yield navigator.serviceWorker.register("/service-worker.js");
        try {
            console.log("ServiceWorker registration successful with scope: ", registration.scope);
        }
        catch (error) {
            console.log("ServiceWorker registration failed: ", error);
        }
    }));
}
