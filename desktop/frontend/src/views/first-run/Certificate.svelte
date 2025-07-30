<script lang="ts">
  import { _ } from 'svelte-i18n';
  import Base from "./Base.svelte";
</script>

<Base
  title="Upload certificate"
  primary={$_("next")}
  hasForm
>
  <div style="width: 100%; height: 500px;">Drop your certificate here</div>
</Base>