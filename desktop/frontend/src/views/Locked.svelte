<script lang="ts">
  import { _ } from 'svelte-i18n';
  import {
    Content,
    Grid,
    Row,
    Column,
    Form,
    FormGroup,
    Button,
    PasswordInput,
    ComposedModal,
    ModalHeader,
    ModalBody,
    ModalFooter,
  } from "carbon-components-svelte";
  import store from '@src/app/store.svelte';
  import { views, controller } from "@src/app/views.svelte";
  import Unlocked from "carbon-icons-svelte/lib/Unlocked.svelte";

  let password : string = $state('');
  let passwordInvalid : boolean = $state(false);
  let passwordInvalidText : string = $state('');

  async function unlock() {
    try {
      await store.unlock(password, () => { controller.view = views.sign });
    } catch(err) {
      passwordInvalid = true;
      passwordInvalidText = err as string;
    }
  }
</script>

<ComposedModal open preventCloseOnClickOutside size="xs">
  <ModalHeader title={$_("master-password.title")} closeClass="first-run-hidden" />
  <ModalBody style="overflow: hidden;">
    <PasswordInput
              bind:value={password}
              labelText={$_("master-password.label")}
              placeholder={$_("master-password.placeholder")}
              invalid={passwordInvalid}
              invalidText={passwordInvalidText}
            />
  </ModalBody>
  <ModalFooter>
    <Button onclick={unlock} icon={Unlocked}>{$_("unlock")}</Button>
  </ModalFooter>
</ComposedModal>