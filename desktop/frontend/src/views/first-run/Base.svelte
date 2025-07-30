<script lang="ts">
  import { _ } from 'svelte-i18n';
  import {
    ComposedModal,
    ModalHeader,
    ModalBody,
    ModalFooter,
    Button,
  } from "carbon-components-svelte";
  import type { Snippet } from 'svelte';
  import firstrun from './firstrun.svelte';

  interface Props {
    children: Snippet
    title: string,
    primary: string,
    onPrimary?: () => void,
    primaryDisabled?: boolean,
    secondary?: string,
    onSecondary?: () => void,
    secondaryDanger?: boolean,
    secondaryDisabled?: boolean,
    hasForm?: boolean,
  }
  const {
    children,
    title,
    primary,
    onPrimary = () => {},
    primaryDisabled = false,
    secondary = "",
    onSecondary = () => {},
    secondaryDanger = false,
    secondaryDisabled = false,
    hasForm = false,
  } : Props = $props();

  function wrap(fn : () => void) : () => void {
    return () => {
      onPrimary();
      firstrun.advance();
    }
  }
</script>

<ComposedModal open preventCloseOnClickOutside>
  <ModalHeader label={$_("initial-setup")} {title} closeClass="first-run-hidden" />
  <ModalBody style="overflow: hidden;">
    {@render children()}
  </ModalBody>
  <ModalFooter>
    {#if secondary !== ""}
      <Button kind={secondaryDanger ? "danger" : "secondary"} disabled={secondaryDisabled} onclick={wrap(onSecondary)}>{secondary}</Button>  
    {/if}
    <Button kind="primary" disabled={primaryDisabled} onclick={wrap(onPrimary)}>{primary}</Button>  
  </ModalFooter>
</ComposedModal>

<style>
  :global {
    .first-run-hidden {
      display: none;
    }
  }
</style>