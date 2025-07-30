<script lang="ts">
  import { _ } from 'svelte-i18n';
  import {
    Select,
    SelectItem,
  } from "carbon-components-svelte";
  import i18n from '@src/app/i18n.svelte';
  import settings from '@src/app/settings.svelte';

  $effect(() => {
    i18n.lang = settings.settings.lang;
  });

</script>

<Select labelText={$_("language")} bind:selected={settings.settings.lang}>
  {#each i18n.langs as lang}
  <SelectItem value={lang} text={$_(`lang.${lang}`)} />
  {/each}
</Select>