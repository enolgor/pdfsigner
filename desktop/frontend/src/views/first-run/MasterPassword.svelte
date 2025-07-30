<script lang="ts">
  import { _ } from 'svelte-i18n';
  import Base from "./Base.svelte";
  import { MasterPassword } from '@src/components';

  interface Props {
    password : string;
  }

  let { password = $bindable('') } : Props = $props();
  let retype : string = $state('');

  let isEmpty : boolean = $derived(password.trim() === "");
  let match : boolean = $derived(password === retype);
</script>

<Base
  title={$_("first-run.master-password.title")}
  primary={$_("next")}
  primaryDisabled={isEmpty || !match}
  hasForm
  secondary={$_("first-run.master-password.risk")}
  secondaryDanger
  secondaryDisabled={!isEmpty}
>
  <MasterPassword bind:password bind:retype />
</Base>