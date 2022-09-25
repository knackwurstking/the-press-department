<script>
  import RollenBahnSection from "./lib/RollenBahnSection.svelte";

  import Data from "./js/data";
</script>

<main class="game-screen">
  <div class="trockner-container" />
  <div class="rollenbahn-container">
    <div class="rollenbahn">
      {#each Data.rollenBahn as section}
        <RollenBahnSection
          name={section.name}
          engine={section.engine}
        />
      {/each}
    </div>
  </div>
  <div class="presse-container" />
</main>

<style>
  :global(body) {
    display: flex;
    justify-content: center;
    place-items: center;
  }

  * {
    border: 1px solid red;
  }

  .game-screen {
    /* main screen contains the game elements; should be scalable to fit the screen */
    width: 1800px;
    height: 500px;
    position: relative;
    overflow: hidden;
    display: flex;
    position: relative;
  }

  .trockner-container {
    /* end of screen (left side) */
    height: 100%;
    width: 75px;
    position: relative;
  }

  .rollenbahn-container {
    /* between '.trockner' and '.presse' */
    height: 100%;
    width: 1700px;
    display: flex;
    place-items: center;
  }

  .rollenbahn {
    height: 150px;
    width: 100%;
    border: none;
    display: flex;
    flex-direction: row-reverse;
  }

  .presse-container {
    /* start of screen (right screen) */
    height: 100%;
    width: 25px;
    position: relative;
  }
</style>
