<script lang="ts">
  import { _ } from "svelte-i18n";
  import { stamps } from "@models";
  import { Checkbox, Form, FormGroup, Slider } from "carbon-components-svelte";
  import { StoreLogo } from "@go";
  import FileDrop from "@src/components/FileDrop.svelte";

  interface Props {
   stamp : stamps.StampConfig
  }
  let { stamp = $bindable() } : Props = $props();

  function onFileChosen(path: string) {
    StoreLogo(path).then((name) => {
      stamp.logo = name;
    }).catch((err) => {
      //TODO: do somethign
    })
  }

  $effect(() => {
    stamp.logoOpacity = Math.floor(stamp.logoOpacity * 10) / 10;
  })

</script>

<Form>
  <FileDrop label={$_("stamp-editor.logo-file")} {onFileChosen} extensions={[".png"]} />
  <FormGroup>
    <Slider
      labelText={$_("stamp-editor.logo-opacity")}
      min={0}
      max={1}
      maxLabel="1"
      step={0.1}
      bind:value={stamp.logoOpacity}
    />
  </FormGroup>
  <FormGroup>
    <Checkbox labelText={$_("stamp-editor.logo-grayscale")} bind:checked={stamp.logoGrayScale} />
  </FormGroup>
</Form>