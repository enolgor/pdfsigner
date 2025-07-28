<script lang="ts">
  import { _ } from 'svelte-i18n';
  import {
    Header,
    HeaderUtilities,
    HeaderGlobalAction,
  } from "carbon-components-svelte";
  import SettingsAdjust from "carbon-icons-svelte/lib/SettingsAdjust.svelte";
  import Certificate from "carbon-icons-svelte/lib/Certificate.svelte";
  import DocumentPdf from "carbon-icons-svelte/lib/DocumentPdf.svelte";
  import Stamp from "carbon-icons-svelte/lib/Stamp.svelte";
  import Help from "carbon-icons-svelte/lib/Help.svelte";
  import type { Snippet } from "svelte";
  import { views, controller } from "@src/app/views.svelte";

  interface Props {
    emptyShell?: boolean,
    content: () => ReturnType<Snippet>,
  }

  let {
    emptyShell = false,
    content
  } : Props = $props();

  let isSideNavOpen : boolean = $state(false);
</script>

<Header company="@enolgor" platformName="PDFSigner" bind:isSideNavOpen>
  {#if !emptyShell}
  <HeaderUtilities>
    <HeaderGlobalAction
      iconDescription={$_("signpdf")}
      tooltipAlignment="start"
      icon={DocumentPdf}
      onclick={() => controller.view = views.sign}
    />
    <HeaderGlobalAction
      iconDescription={$_("certificates")}
      icon={Certificate}
      onclick={() => controller.view = views.certificates}
    />
    <HeaderGlobalAction
      iconDescription={$_("stamps")}
      icon={Stamp}
      onclick={() => controller.view = views.stamps}
    />
    <HeaderGlobalAction
      iconDescription={$_("settings")}
      icon={SettingsAdjust}
      onclick={() => controller.view = views.settings}
    />
    <HeaderGlobalAction
      iconDescription={$_("help")}
      tooltipAlignment="end"
      icon={Help}
      onclick={() => controller.view = views.help}
    />
  </HeaderUtilities>
  {/if}
</Header>

{@render content()}
