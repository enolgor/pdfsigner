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
  import { NewUnsavedStamp, SetUnsavedStamp } from "@go";
  import ContentSection from "./ContentSection.svelte";
  import ExtraContent from "./ExtraContent.svelte";
  import ImageSection from "./ImageSection.svelte";
  import BorderSection from "./BorderSection.svelte";
  import AlignmentSection from "./AlignmentSection.svelte";
  import LogoSection from "./LogoSection.svelte";
  import FontsSection from "./FontsSection.svelte";

  let stamp : stamps.StampConfig | undefined = $state();
  let sectionStates : boolean[] = $state([false, false, false, false, false, false, false]);

  onMount(async () => {
    stamp = await NewUnsavedStamp();
  });


  $effect(() => {
    if (stamp) {
      SetUnsavedStamp(stamp).then(() => renderStamp()).catch((err) => {
        console.error(err); //TODO
      });
    }
  });

  let image : HTMLImageElement | undefined = $state();
  function renderStamp() {
    if (image) {
     image.src = `/unsaved-stamp.png?${new Date().getTime()}`;
    }
  }
</script>

{#if stamp !== undefined}
<Content>
  <Grid>
    <Row>
      <Column>
        <img alt="stamp" class="stamp" bind:this={image} src="/unsaved-stamp.png" />
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
             <AccordionItem title={$_("stamp-editor.logo")} bind:open={sectionStates[4]}>
              <LogoSection bind:stamp />
            </AccordionItem>
            <AccordionItem title={$_("stamp-editor.alignment")} bind:open={sectionStates[5]}>
              <AlignmentSection bind:stamp />
            </AccordionItem>
             <AccordionItem title={$_("stamp-editor.fonts")} bind:open={sectionStates[6]}>
              <FontsSection bind:stamp />
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
    /*max-height: 300px;*/
  }
</style>