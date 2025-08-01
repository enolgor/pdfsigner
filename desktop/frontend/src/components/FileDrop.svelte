<script lang="ts">
  import { _ } from 'svelte-i18n';
  import { OnFileDrop, OnFileDropOff } from "@runtime";
  import { OpenFileDialog } from "@go";
  import { onMount, onDestroy } from 'svelte';
  import { ClickableTile, FormGroup } from 'carbon-components-svelte';
  import { DocumentImport } from 'carbon-icons-svelte';

  interface Props {
    extensions?: string[],
    label: string,
    onFileChosen?: (path: string) => void,
  }

  let { label, extensions, onFileChosen = () => {} } : Props = $props();

  let errMsg : string = $state("");
  let filename : string = $state("");
  let messageText : string = $derived(errMsg !== "" ? errMsg : filename !== "" ? $_("selected-file", {values: { filename } }) : "");

  let nativeExtensions : string = $derived(extensions && extensions.length > 0 ? extensions.map(ext => `*${ext}`).join(";"): "");

  onMount(() => {
    OnFileDrop((x, y, paths) => {
      if (paths.length != 1) {
        errMsg = $_("cant-drop-multiple");
        return;
      }
      const path = paths[0];
      if (extensions && extensions.length > 0) {
        if (extensions.find((ext) => path.endsWith(ext)) === undefined) {
          errMsg = $_("wrong-file-drop");
          return;
        }
      }
      errMsg = "";
      filename = path.split(/[/\\]/).pop() || "";
      onFileChosen(path);
    }, true);
  });

  onDestroy(() => {
    OnFileDropOff();
  });

  async function onFileChoose(e: MouseEvent) {
    e.preventDefault();
    try {
      const path = await OpenFileDialog(nativeExtensions);
      if (path === "") {
        return;
      }
      errMsg = "";
      filename = path.split(/[/\\]/).pop() || "";
      onFileChosen(path);
    }catch(err) {
      errMsg = err as string;
    }
  }

  export function reset() {
    errMsg = "";
    filename = "";
  }
  
</script>
<FormGroup
  message
  {messageText}
  legendText={label}
>


  <div class="dropcontainer" style="--wails-drop-target: drop;">
  <ClickableTile
    href="#"
    onclick={onFileChoose}
  >
    <div style="width: 100%; height: 100%; display: flex; align-items: center; justify-content: center;">
      <DocumentImport />
      {$_("drop-a-file")}
    </div>
  </ClickableTile>
  </div>
</FormGroup>