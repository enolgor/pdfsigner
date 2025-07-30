<script lang="ts">
  import { _ } from 'svelte-i18n';
  import Base from "./Base.svelte";
  import FileDrop from '@src/components/FileDrop.svelte';
  
</script>

<Base
  title="Upload certificate"
  primary={$_("next")}
  hasForm
>
  <FileDrop />
</Base>