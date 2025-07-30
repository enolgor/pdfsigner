<script lang="ts">
  import { _ } from 'svelte-i18n';
  import firstrun from './first-run/firstrun.svelte';
  import Base from "./first-run/Base.svelte";
  import store from '@src/app/store.svelte';
  import settings from "@src/app/settings.svelte";
  import { LanguageSelector } from '@src/components/index';
  import MasterPassword from './first-run/MasterPassword.svelte';
    import Certificate from './first-run/Certificate.svelte';

  let password : string = $state('');

  firstrun.steps = 5;
  $effect(() => {
    if (firstrun.done) {
      settings.save()
        .then(store.firstRunCompleted)
        .then(() => store.changePassword(password)).catch((err) => {
        console.error(err);
      });
    }
  });

</script>

{#if firstrun.step === 0}
  <Base title={$_("first-run.intro.title")} primary={$_("next")} >
    <p>{$_("first-run.intro.setup")}</p>
  </Base>
  {:else if firstrun.step === 1}
  <Base title={$_("first-run.language.title")} primary={$_("next")} >
    <LanguageSelector />
  </Base>
  {:else if firstrun.step === 2}
  <Certificate />
  {:else if firstrun.step === 3}
  <MasterPassword bind:password />
  {:else if firstrun.step === 4}
  <Base title={$_("first-run.done.title")} primary={$_("proceed")} >
    <p>{$_("first-run.done.allset")}</p>
  </Base>
{/if}