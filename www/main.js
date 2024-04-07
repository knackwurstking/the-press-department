import { KeepAwake } from "@capacitor-community/keep-awake";

KeepAwake.isSupported().then(
  (result) => result.isSupported && KeepAwake.keepAwake(),
);
