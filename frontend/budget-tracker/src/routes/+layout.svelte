<script>
  import { onMount } from 'svelte';
  import Sidenav from "../components/Sidenav.svelte";
  import FilterHeader from "../components/FilterHeader.svelte";
  import '../global.css'

  onMount(() => {
    document.body.style.backgroundColor = '#FFFFFF';
    document.body.style.margin = '0';
    document.body.style.padding = '0';
    document.body.style.fontFamily = 'Roboto, sans-serif';
  });
</script>

<style>
  .main-content {
    margin-left: 300px; /* Adjust based on the width of sidenav */
  }

  .content-wrapper {
    display: flex;
    flex-direction: column;
  }

  main {
    flex-grow: 1;
    box-sizing: border-box; 
  }
</style>

<Sidenav />
<div class="main-content">
  <FilterHeader />
  <div class="content-wrapper">
    <main>
      <slot />
    </main>
  </div>
</div>
<!--<FilterNav />-->

