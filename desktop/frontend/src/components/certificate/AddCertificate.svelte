<script lang="ts">
  import { _ } from 'svelte-i18n';
  import { FileDrop } from '@src/components/index';
  import { CodeSnippet, Form, FormGroup, PasswordInput } from 'carbon-components-svelte';
  import { GetCertificateID } from '@go';

  interface Props {
    path: string;
    passphrase: string;
    valid: boolean;
    open: boolean;
    invalid: boolean;
  }

  let {
    path = $bindable(''),
    passphrase = $bindable(''),
    valid = $bindable(false),
    open = $bindable(false),
    invalid = $bindable(false),
  } : Props = $props();

  let issuer : string = $state('');
  let subject : string = $state('');

  let filedrop : FileDrop | undefined = $state(undefined);

  function onFileChosen(_path: string) {
    path = _path;
    open = true;
    valid = false;

  }

  export function reset() {
    if (!valid) {
      open = false;
      invalid = false;
      passphrase = "";
      path = "";
      filedrop?.reset();
    }
  }

 export async function unlockCertificate() {
    try {
      const id = await GetCertificateID(path, passphrase);
      invalid = false;
      open = false;
      valid = true;
      issuer = id.Issuer;
      subject = id.Subject;
    } catch (err) {
      passphrase = '';
      invalid = true;
    }
  }
</script>

<Form>
  {#if !valid}
  <FileDrop bind:this={filedrop} label={$_("select-p12")} extensions={[".p12"]} {onFileChosen} />
  {:else}
    <FormGroup legendText={$_("certificate-subject")}>
      <CodeSnippet code={subject} hideCopyButton/>
    </FormGroup>
    <FormGroup legendText={$_("certificate-issuer")}>
      <CodeSnippet code={issuer} hideCopyButton/>
    </FormGroup>
  {/if}
</Form>