<script lang="ts">
  import { _ } from "svelte-i18n";
  import { stamps, config } from "@models";
  import { Button, Column, Form, FormGroup, Grid, Row, TextInput, Tooltip } from "carbon-components-svelte";
  import { Add, TrashCan } from "carbon-icons-svelte";

  interface Props {
   stamp : stamps.StampConfig
  }

  let { stamp = $bindable() } : Props = $props();

  let keys : string[] = $state([]);
  let values : string[] = $state([]);

  function addExtraLine() {
    keys = [...keys, ''];
    values = [...values, ''];
  }

  function deleteExtraLine(idx : number) {
    keys = keys.filter((_, i) => i !== idx);
    values = values.filter((_, i) => i !== idx);
  }

  $effect(() => {
    stamp.extraLines = keys.map((_, idx) => {
      return new config.TextLine({Key: keys[idx], Value: values[idx]});
    });
  });
</script>

<Form>
  <FormGroup>
    <Grid>
      <Row>
        <Column>
          <Button size="small" icon={Add} onclick={addExtraLine}>{$_("stamp-editor.add-extra-line")}</Button>
        </Column>
        <Column sm={1} md={1} lg={1}>
          <Tooltip triggerText={$_("stamp-editor.variables-info-label")}>
            <p>{$_("stamp-editor.variables-info", { values: { vars: "{{.Issuer}} {{.Subject}} {{.Date}}" } })}</p>
          </Tooltip>
        </Column>
      </Row>
    </Grid>
  </FormGroup>
  {#each keys as key, idx}
  <FormGroup>
    <TextInput labelText={$_("stamp-editor.key")} bind:value={keys[idx]} inline/>
    <TextInput labelText={$_("stamp-editor.value")} bind:value={values[idx]} inline/>
    <Button size="small" iconDescription={$_("stamp-editor.delete-extra-line")} icon={TrashCan} kind="danger" onclick={() => deleteExtraLine(idx)}/>
  </FormGroup>
  {/each}
</Form>