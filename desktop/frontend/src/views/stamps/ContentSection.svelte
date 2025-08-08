<script lang="ts">
  import { _ } from "svelte-i18n";
  import { stamps } from "@models";
  import { Checkbox, Form, FormGroup, Select, SelectItem, TextInput } from "carbon-components-svelte";
  import moment from "moment";

  const dateFormats = [
    'DD/MM/YYYY',
    'D/M/YYYY',
    'MM/DD/YYYY',
    'M/D/YYYY',
    'YYYY-MM-DD',
    'DD-MM-YYYY',
    'D-M-YYYY',
    'DD.MM.YYYY',
    'D.M.YYYY',
    'YYYY/MM/DD',
    'YYYY.MM.DD'
  ];

  const timeFormats = [
    '',
    'HH:mm:ss',
    'HH:mm',
    'hh:mm:ss A',
    'hh:mm A'
  ];

  interface Props {
   stamp : stamps.StampConfig
  }

  let { stamp = $bindable() } : Props = $props();

  let dateFormat : string = $state(dateFormats[0]);
  let timeFormat : string = $state('');
  let includeOffset: boolean = $state(false);
  const golangDate = moment("2006-01-02T15:04:05", "YYYY-MM-DDTHH:mm:ss");

  let dateTimeFormat : string = $derived(timeFormat === '' ? dateFormat : `${dateFormat} ${timeFormat}`);
  $effect(() => {
    stamp.dateFormat = golangDate.format(dateTimeFormat) + (includeOffset && timeFormat !== '' ? ' -07:00' : '');
  });
</script>

<Form>
  <FormGroup>
    <Checkbox labelText={$_("stamp-editor.include-title")} bind:checked={stamp.includeTitle} />
    <TextInput labelText={$_("stamp-editor.title")} bind:value={stamp.title} disabled={!stamp.includeTitle} inline/>
  </FormGroup>
  <FormGroup>
    <Checkbox labelText={$_("stamp-editor.include-subject")} bind:checked={stamp.includeSubject} />
    <TextInput labelText={$_("stamp-editor.subject-key")} bind:value={stamp.subjectKey} disabled={!stamp.includeSubject} inline/>
  </FormGroup>
  <FormGroup>
    <Checkbox labelText={$_("stamp-editor.include-issuer")} bind:checked={stamp.includeIssuer} />
    <TextInput labelText={$_("stamp-editor.issuer-key")} bind:value={stamp.issuerKey} disabled={!stamp.includeIssuer} inline/>
  </FormGroup>
  <FormGroup>
    <Checkbox labelText={$_("stamp-editor.include-date")} bind:checked={stamp.includeDate} />
    <TextInput labelText={$_("stamp-editor.date-key")} bind:value={stamp.dateKey} disabled={!stamp.includeDate} inline/>
    
  </FormGroup>
  <FormGroup>
    <Select labelText={$_("stamp-editor.date-format")} bind:selected={dateFormat} disabled={!stamp.includeDate} inline>
      {#each dateFormats as df}
        <SelectItem value={df} text={moment().format(df)} />
      {/each}
    </Select>
     <Select labelText={$_("stamp-editor.time-format")} bind:selected={timeFormat} disabled={!stamp.includeDate} inline>
      {#each timeFormats as tf}
        {#if tf == ''}
        <SelectItem value={tf} text='' />
        {:else}
        <SelectItem value={tf} text={moment().format(tf)} />
        {/if}
      {/each}
    </Select>
    <Checkbox labelText={$_("stamp-editor.include-offset")} bind:checked={includeOffset} disabled={!stamp.includeDate || timeFormat === ''} />
  </FormGroup>
</Form>