<script lang="ts">
  import initialize from '@src/app/initialize';
  import { _ } from 'svelte-i18n';
  import { Shell } from '@src/components';
  import { views, controller } from '@src/app/views.svelte';
  import store from '@src/app/store.svelte';
  import {
    Certificates,
    FirstRun,
    Locked,
    Settings,
    Sign,
    Stamps,
    Loading,
    Help,
  } from '@src/views';

  let emptyShell : boolean = $derived(store.locked || store.firstRun);
</script>

{#await initialize()}
  <Shell emptyShell><Loading /></Shell>
{:then}
{@const view = controller.view}
<Shell {emptyShell}>
  {#if store.firstRun}
    <FirstRun />
  {:else if store.locked}
    <Locked />
  {:else if view === views.settings}
    <Settings />
  {:else if view === views.certificates}
    <Certificates />
  {:else if view === views.stamps}
    <Stamps />
  {:else if view === views.sign}
    <Sign />
  {:else if view === views.help}
    <Help />
  {:else}
    <Loading />
  {/if}
</Shell>
{/await}