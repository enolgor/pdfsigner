<script lang="ts">
  import { _ } from "svelte-i18n";
  import {
    Content,
    Grid,
    Row,
    Column,
    Accordion,
    AccordionItem,
  } from "carbon-components-svelte";
  import { onMount } from 'svelte';
  import { stamps } from "@models";
  import { NewDefaultStampConfig, SetUnsavedStamp } from "@go";
  import ContentSection from "./ContentSection.svelte";
  import ExtraContent from "./ExtraContent.svelte";
  import ImageSection from "./ImageSection.svelte";
  import BorderSection from "./BorderSection.svelte";
  import AlignmentSection from "./AlignmentSection.svelte";

  let stamp : stamps.StampConfig | undefined = $state();
  let sectionStates : boolean[] = $state([false, false, false, false, false]);

  onMount(async () => {
    stamp = await NewDefaultStampConfig();
  });


  $effect(() => {
    if (stamp) {
      SetUnsavedStamp(stamp)
      renderStamp();
    }
  });

  let image : HTMLImageElement | undefined = $state();
  async function renderStamp() {
    if (image) {
      const resp = await fetch("/unsaved-stamp", { method: "POST" });
      const blob = await resp.blob();
      image.src = URL.createObjectURL(blob);
    }
  }
</script>

{#if stamp !== undefined}
<Content>
  <Grid>
    <Row>
      <Column>
        <img alt="stamp" class="stamp" bind:this={image} />
      </Column>
      <Column>
        <Accordion size="sm">
            <AccordionItem title={$_("stamp-editor.content")} bind:open={sectionStates[0]}>
              <ContentSection bind:stamp />
            </AccordionItem>
            <AccordionItem title={$_("stamp-editor.extra-content")} bind:open={sectionStates[1]}>
              <ExtraContent bind:stamp />
            </AccordionItem>
            <AccordionItem title={$_("stamp-editor.image")} bind:open={sectionStates[2]}>
              <ImageSection bind:stamp />
            </AccordionItem>
            <AccordionItem title={$_("stamp-editor.border")} bind:open={sectionStates[3]}>
              <BorderSection bind:stamp />
            </AccordionItem>
            <AccordionItem title={$_("stamp-editor.alignment")} bind:open={sectionStates[4]}>
              <AlignmentSection bind:stamp />
            </AccordionItem>
        </Accordion>
      </Column>
    </Row>
  </Grid>
</Content>
{/if}
<style>
  .stamp {
    max-width: 300px;
    max-height: 300px;
  }
</style>