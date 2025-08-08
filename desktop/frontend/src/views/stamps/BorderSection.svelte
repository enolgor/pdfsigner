<script lang="ts">
  import { _ } from "svelte-i18n";
  import { stamps } from "@models";
  import ColorPicker from '@src/components/ColorPicker.svelte';
  import { Form, FormGroup, Slider } from "carbon-components-svelte";

  interface Props {
   stamp : stamps.StampConfig
  }
  let { stamp = $bindable() } : Props = $props();

  $effect(() => {
    stamp.borderSizePt = Math.floor(stamp.borderSizePt * 10) / 10;
  })
</script>

<Form>
  <FormGroup>
    <Slider
      labelText={$_("stamp-editor.border-size")}
      min={0}
      max={10}
      maxLabel="10 pts"
      step={0.1}
      bind:value={stamp.borderSizePt}
    />
  </FormGroup>
  <FormGroup>
    <ColorPicker label={$_("stamp-editor.border-color")} bind:color={stamp.borderColor} />
  </FormGroup>
</Form>