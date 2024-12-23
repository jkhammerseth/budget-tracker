<script>
  import { onMount } from 'svelte';
  import Sidenav from "../components/Sidenav.svelte";
  import Header from '../components/Header.svelte';
  import FilterHeader from "../components/FilterHeader.svelte";
  import Fab from '../components/FAB.svelte';
  import '../global.css'

  onMount(() => {
    document.body.style.backgroundColor = '#ecf0f1';
    document.body.style.margin = '0';
    document.body.style.padding = '0';
    document.body.style.fontFamily = 'Roboto, sans-serif';
  });

</script>

<Header />
<Sidenav />
<div class="fab"><Fab isOpen=True /></div>

<div class="main-content">
  <div class="content-wrapper">
   
    <main>
      <slot />
    </main>
  </div>
</div>


<style>
  .main-content {
    
    padding-top: var(--header-height);
    background-color: var(--background-color);
    height: 100%;
  }

  .content-wrapper {
    display: flex;
    flex-direction: column;
    
  }

  main {
    flex-grow: 1;
    box-sizing: border-box; 
  }

  .fab {
    margin-right: 22rem; 
    z-index: 10;
  }
</style>