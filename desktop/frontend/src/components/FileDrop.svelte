<script lang="ts">
  import { _ } from 'svelte-i18n';
  import { OnFileDrop, LogPrint, OnFileDropOff } from "@runtime";
  import { onMount, onDestroy } from 'svelte';
    import { ClickableTile, FileUploaderDropContainer } from 'carbon-components-svelte';
    import { DocumentImport } from 'carbon-icons-svelte';

  onMount(() => {
    OnFileDrop((x, y, paths) => {
      LogPrint(`file dropped: x ${x}, y ${y}, paths ${paths}`)
      console.log("file dropped");
      console.log("x", x, "y", y, "paths", paths);
    }, true);
  });

  onDestroy(() => {
    OnFileDropOff();
  });
  
</script>

<div class="dropcontainer" style="--wails-drop-target: drop;">
<ClickableTile
  href="#"
  on:click={(e) => {
    e.preventDefault();
    // custom behavior
  }}
>
  <div style="width: 100%; height: 100%; display: flex; align-items: center; justify-content: center;">
    <DocumentImport />
    Drop a file or click to choose
  </div>
</ClickableTile>
</div>