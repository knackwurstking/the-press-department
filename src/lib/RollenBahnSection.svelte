<script>
  import Data from "../js/data";

  /** @type {string} */
  export let name;
  /** @type {import("../js/data").Engine} */
  export let engine;
  $: {
    const newEngineParts = [];
    for (let x = 1; x <= engine.count; x++) {
      newEngineParts.push(x);
    }
    engineParts = newEngineParts;
  }

  /** @type {number[]} */
  let engineParts = [];
</script>

<section
  class={name}
  style={`
    width: ${
      engine.count * (Data.ROLLEN_BAHN_LENGTH / Data.ROLLEN_BAHN_MAX_COUNT)
    }px;
  `}
>
  <div class="gestell">
    {#each engineParts as part}
      <div class={`engine-part engine-part-${part}`} />
    {/each}
  </div>
</section>

<style>
  div.gestell {
    width: 100%;
    height: 100%;
    border: 3px solid peru;
    border-top-width: 10px;
    border-bottom-width: 10px;
    display: flex;
    flex-direction: row-reverse;
    justify-content: space-evenly;
  }

  section:last-child div.gestell {
    border-left: none;
  }

  section:first-child div.gestell {
    border-right: none;
  }

  .engine-part {
    width: 4px;
    height: 100%;
    animation: rollen-10hz 0.55s infinite;
  }

  @keyframes rollen-10hz {
    0% {
      background-image: linear-gradient(90deg, gray 25%, dimgray, darkgray 25%);
    }
    50% {
      background-image: linear-gradient(90deg, dimgray, darkgray 25%, gray 25%);
    }
    100% {
      background-image: linear-gradient(90deg, darkgray 25%, gray 25%, dimgray);
    }
  }
</style>
