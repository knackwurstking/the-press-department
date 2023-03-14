<script lang="ts">
  import { onMount } from "svelte";
  import { KeepAwake } from "@capacitor-community/keep-awake";

  onMount(async () => {
    if ((await KeepAwake.isSupported()).isSupported) {
      KeepAwake.keepAwake();
    }
  });
</script>

<iframe src="main.html" title="The Press Department" scrolling="no"></iframe>

<style>
  iframe {
    width: 100vw;
    height: 100vh;
    border: none;
  }
</style>
