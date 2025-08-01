<script lang="ts">
  import { _ } from "svelte-i18n";
  import { Form, FormGroup, Modal, PasswordInput } from "carbon-components-svelte";

  interface Props {
    open: boolean;
    onModalClose: () => void;
    unlockCertificate: () => void;
    passphrase: string;
    invalid: boolean;
  }

  let {
    open = $bindable(false),
    onModalClose,
    unlockCertificate,
    passphrase = $bindable(''),
    invalid,
  } : Props = $props();

  let invalidText : string = $derived(invalid ? $_("invalid-passphrase") : "");
  
</script>


<Modal hasForm
  size="xs"
  preventCloseOnClickOutside
  shouldSubmitOnEnter
  bind:open
  modalHeading={$_("certificate-locked")}
  primaryButtonText={$_("unlock")}
  selectorPrimaryFocus="#cert-passphrase"
  on:close={onModalClose}
  secondaryButtonText={$_("cancel")}
  on:click:button--secondary={onModalClose}
  on:submit={unlockCertificate}
>
  <Form style="overflow:hidden" on:submit={(e) => {e.preventDefault();}}>
    <FormGroup>
      <PasswordInput
        id="cert-passphrase"
        bind:value={passphrase}
        labelText={$_("certificate-passphrase")}
        placeholder={$_("passphrase")}
        {invalid}
        {invalidText}
      />
    </FormGroup>
  </Form>
</Modal>