<script lang="ts">
  import { _ } from "svelte-i18n";
  import { config, stamps } from "@models";
  import AlignmentButtonset from '@src/components/AlignmentButtonset.svelte';
  import { Button, ButtonSet, Form, FormGroup } from "carbon-components-svelte";

  interface Props {
   stamp : stamps.StampConfig
  }
  let { stamp = $bindable() } : Props = $props();


</script>

<Form>
  <FormGroup>
    <AlignmentButtonset label={$_("stamp-editor.title-alignment")} bind:alignment={stamp.titleAlignment} />
  </FormGroup>
  <FormGroup>
    <AlignmentButtonset label={$_("stamp-editor.line-alignment")} bind:alignment={stamp.lineAlignment} />
  </FormGroup>
  <FormGroup>
    <AlignmentButtonset label={$_("stamp-editor.key-alignment")} bind:alignment={stamp.keyAlignment} disabled={stamp.lineAlignment !== config.Alignment.CENTER} />
  </FormGroup>
  <FormGroup>
    <AlignmentButtonset label={$_("stamp-editor.value-alignment")} bind:alignment={stamp.valueAlignment} disabled={stamp.lineAlignment !== config.Alignment.CENTER} />
  </FormGroup>
  <FormGroup>
    <AlignmentButtonset label={$_("stamp-editor.logo-alignment")} bind:alignment={stamp.logoAlignment} />
  </FormGroup>
</Form>