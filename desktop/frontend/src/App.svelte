<script lang="ts">
  import initialize from '@src/app/initialize';
  import { _ } from 'svelte-i18n';
  import { Shell } from '@src/components';
  import { views, controller } from '@src/app/views.svelte';
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

  let emptyShell : boolean = $derived(controller.view === views.locked || controller.view === views.firstRun);
</script>

{#await initialize()}
<Loading />
{:then}
<Shell {emptyShell} {content} />
{/await}

{#snippet content()}
{@const view = controller.view}
{#if view === views.settings}
  <Settings />
{:else if view === views.certificates}
  <Certificates />
{:else if view === views.stamps}
  <Stamps />
{:else if view === views.sign}
  <Sign />
{:else if view === views.firstRun}
  <FirstRun />
{:else if view === views.locked}
  <Locked />
{:else if view === views.help}
  <Help />
{:else}
  <Loading />
{/if}
{/snippet}