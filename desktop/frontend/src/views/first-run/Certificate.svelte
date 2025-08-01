<script lang="ts">
  import { _ } from 'svelte-i18n';
  import Base from "./Base.svelte";
  import { StoreCertificate } from '@go';
  import { AddCertificate, CertPassphrase } from '@src/components/index';

  let path : string = $state('');
  let passphrase : string = $state('');
  let valid : boolean = $state(false);
  let open : boolean = $state(false);
  let invalid : boolean = $state(false);

  let addCertificate : AddCertificate | undefined = $state();

  function onPrimary() {
    StoreCertificate(path, passphrase).catch((err) => {
      console.error(err);
    });
  }
</script>

<div class:hidden={open}>
  <Base
    title={$_("upload-certificate")}
    primary={$_("next")}
    hasForm
    primaryDisabled={!valid}
    {onPrimary}
  >
    <AddCertificate bind:this={addCertificate} bind:path bind:passphrase bind:valid bind:open bind:invalid />
  </Base>
</div>

<CertPassphrase bind:open bind:passphrase {invalid} onModalClose={addCertificate.reset} unlockCertificate={addCertificate.unlockCertificate} />

<style>
  .hidden {
    visibility: hidden;
  }
</style>