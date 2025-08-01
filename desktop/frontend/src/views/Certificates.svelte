<script lang="ts">
    import { _ } from "svelte-i18n";
    import {
    Content,
    Grid,
    Row,
    Column,
    ExpandableTile,
    ButtonSet,
    Button,
    Modal,
  } from "carbon-components-svelte";
  import { onMount } from "svelte";
  import { ListCertificates, StoreCertificate, DeleteCertificate, SetDefaultCertificate, GetStoredCertificateID } from "@go";
  import Add from "carbon-icons-svelte/lib/Add.svelte";
  import { AddCertificate, CertPassphrase } from "@src/components";
  import Loading from "./Loading.svelte";
  
  let certificates : string[] = $state([]);

  let path : string = $state('');
  let passphrase : string = $state('');
  let valid : boolean = $state(false);
  let openPassword : boolean = $state(false);
  let invalid : boolean = $state(false);
  let open : boolean = $state(false);

  let addCertificate : AddCertificate | undefined = $state();

  async function init() {
    certificates = await ListCertificates();
  }

  function onModalClose() {
    path = "";
    passphrase = "";
    valid = false;
    openPassword = false;
    invalid = false;
    open = false;
    addCertificate?.reset();
  }

  function onPrimary() {
    StoreCertificate(path, passphrase).then(() => {
      onModalClose();
      return init();
    }).catch((err) => {
      console.error(err); //TODO
    });
  }

  async function deleteCertificate(key : string) {
    try {
      await DeleteCertificate(key);
      await init();
    } catch(err) {
      console.error(err); //TODO
    }
  }

  async function setDefaultCertificate(key: string) {
    await SetDefaultCertificate(key);
    await init();
  }

  function 

</script>

{#await init()}
<Loading />
{:then}
<div class:hidden={open}>
<Content>
  <Grid padding>
    <Row>
      <Column>
        <h1>TODO // Certificates</h1>
      </Column>
    </Row>
    <Row>
      <Column>
        <Button size="small" icon={Add} onclick={() => (open = true)}>{$_("add-cert")}</Button>
      </Column>
    </Row>
    {#each certificates as cert, idx}
    <Row>
      <Column>
        <ExpandableTile onclick={}>
          <div slot="above">
            <div>{cert}</div>
            {#if idx === 0}
            <br/><div>{$_("default-cert")}</div>
            {/if}
          </div>
          <div slot="below">
            <br/>
            <ButtonSet>
              {#if idx !== 0}
              <Button size="small" kind="tertiary" onclick={async () => await setDefaultCertificate(cert)}>{$_("set-default")}</Button>
              <Button size="small" kind="danger" onclick={async () => await deleteCertificate(cert)}>{$_("delete")}</Button>
              {/if}
            </ButtonSet>
          </div>
        </ExpandableTile>
      </Column>
    </Row>
    {/each}
  </Grid>
</Content>
</div>

<div class:hidden={openPassword}>
<Modal hasForm
  bind:open
  modalHeading={$_("upload-certificate")}
  primaryButtonText={$_("save")}
  on:close={onModalClose}
  secondaryButtonText={$_("cancel")}
  on:click:button--secondary={onModalClose}
  on:submit={onPrimary}
  primaryButtonDisabled={!valid}
>
  <AddCertificate bind:this={addCertificate} bind:path bind:passphrase bind:valid bind:open={openPassword} bind:invalid />
</Modal>
</div>

<CertPassphrase bind:open={openPassword} bind:passphrase {invalid} onModalClose={addCertificate.reset} unlockCertificate={addCertificate.unlockCertificate} />

{/await}
<style>
  .hidden {
    visibility: hidden;
  }
</style>