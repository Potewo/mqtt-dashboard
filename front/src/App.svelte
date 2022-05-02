<script lang="ts">
  import Card from "./lib/Card.svelte";
  import { Styles } from "sveltestrap";
</script>

<Styles />
<main>
  <div class="my-container">
    <Card type="temperature" value={25} />
    <Card type="humidity" value={65} />
  </div>
</main>

<style>
  main {
    width: 100vw;
  }

  .my-container {
    display: flex;
  }
</style>
