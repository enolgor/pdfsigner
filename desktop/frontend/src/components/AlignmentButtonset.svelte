<script lang="ts">
  import { _ } from 'svelte-i18n';
  import { config } from "@models";
  import { Button, ButtonSet, Column, FormGroup, Grid, Row } from "carbon-components-svelte";
  import { TextAlignCenter, TextAlignLeft, TextAlignRight, TextBold, TextItalic, TextUnderline } from "carbon-icons-svelte";

  interface Props {
    alignment: config.Alignment;
    label: string;
    disabled?: boolean;
  }

  let { label, alignment = $bindable(), disabled = false } : Props = $props();

</script>

<Grid condensed>
  <Row>
    <Column>
      <div class="label">{label}</div>
    </Column>
    <Column>
      <Button
        size="small"
        isSelected={alignment === config.Alignment.LEFT}
        kind="ghost"
        iconDescription={$_("alignment.left")}
        icon={TextAlignLeft}
        onclick={() => (alignment = config.Alignment.LEFT)}
        {disabled}
      />
      <Button
        size="small"
        isSelected={alignment === config.Alignment.CENTER}
        kind="ghost"
        iconDescription={$_("alignment.center")}
        icon={TextAlignCenter}
        onclick={() => (alignment = config.Alignment.CENTER)}
        {disabled}
      />
      <Button
        size="small"
        isSelected={alignment === config.Alignment.RIGHT}
        kind="ghost"
        iconDescription={$_("alignment.right")}
        icon={TextAlignRight}
        onclick={() => (alignment = config.Alignment.RIGHT)}
        {disabled}
      />
    </Column>
  </Row>
</Grid>

<style>
  .label {
    display: flex;
    justify-content: left;
    align-items: center;
    width: 100%;
    height: 100%;
  }
</style>