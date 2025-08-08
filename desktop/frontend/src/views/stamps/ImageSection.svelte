<script lang="ts">
  import { _ } from "svelte-i18n";
  import { stamps } from "@models";
  import ColorPicker from '@src/components/ColorPicker.svelte';
  import { Form, FormGroup, Select, SelectItem, Slider } from "carbon-components-svelte";
  import { onMount } from "svelte";
  import { config } from "@models";

  interface Props {
   stamp : stamps.StampConfig
  }
  let { stamp = $bindable() } : Props = $props();

</script>

<Form>
  <FormGroup>
    <ColorPicker label={$_("stamp-editor.background-color")} bind:color={stamp.backgroundColor} />
  </FormGroup>
  <FormGroup>
    <Slider
      labelText={$_("stamp-editor.dpi")}
      min={0}
      max={1000}
      maxLabel="1000 dpi"
      bind:value={stamp.dpi}
    />
  </FormGroup>
  <FormGroup>
    <Select labelText={$_("stamp-editor.rotation")} bind:selected={stamp.rotate} inline>
      <SelectItem value={config.Rotation.ROTATE_0} text="0º" />
      <SelectItem value={config.Rotation.ROTATE_90} text="90º" />
      <SelectItem value={config.Rotation.ROTATE_180} text="180º" />
      <SelectItem value={config.Rotation.ROTATE_270} text="270º" />
    </Select>
  </FormGroup>
</Form>