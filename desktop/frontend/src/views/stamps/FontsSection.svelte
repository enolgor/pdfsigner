<script lang="ts">
  import { _ } from "svelte-i18n";
  import { stamps } from "@models";
  import { Form, FormGroup, Select, SelectItem, SelectItemGroup } from "carbon-components-svelte";
  import { ListFonts } from "@go";
  import ColorPicker from "@src/components/ColorPicker.svelte";

  interface Props {
   stamp : stamps.StampConfig
  }

  type FontProp = "titleFont" | "keyFont" | "valueFont";

  let { stamp = $bindable() } : Props = $props();

</script>



{#await ListFonts()}
<!--TODO loading-->
{:then fonts}
{#snippet fontSelector(label: string, prop: FontProp)}
<Select labelText={label} bind:selected={stamp[prop]}>
  {#each Object.keys(fonts) as key}
    <SelectItemGroup label={key}>
      {#each fonts[key] as font}
        <SelectItem value={font} text={font} />
      {/each}
    </SelectItemGroup>
  {/each}
</Select>
{/snippet}
<Form>
  <FormGroup>
    {@render fontSelector($_("stamp-editor.title-font"), "titleFont")}
    <ColorPicker label={$_("stamp-editor.title-color")} bind:color={stamp.titleColor} />
  </FormGroup>
  <FormGroup>
    {@render fontSelector($_("stamp-editor.key-font"), "keyFont")}
     <ColorPicker label={$_("stamp-editor.key-color")} bind:color={stamp.keyColor} />
  </FormGroup>
  <FormGroup>
    {@render fontSelector($_("stamp-editor.value-font"), "valueFont")}
     <ColorPicker label={$_("stamp-editor.value-color")} bind:color={stamp.valueColor} />
  </FormGroup>
</Form>
{/await}

