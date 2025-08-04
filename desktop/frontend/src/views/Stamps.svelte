<script lang="ts">
    import {
    Content,
    Grid,
    Row,
    Column,
    TextInput,
    Button,
  } from "carbon-components-svelte";
  import { onMount } from 'svelte';
  import { stamps } from "@models";
  import { NewDefaultStampConfig, SetUnsavedStamp } from "@go";


  interface Props {

  }
  let {} : Props = $props();
  let stamp : stamps.StampConfig = $state(new stamps.StampConfig());
  onMount(async () => {
    stamp = await NewDefaultStampConfig();
    console.log($state.snapshot(stamp));
  });


  $effect(() => {
    SetUnsavedStamp(stamp)
  })
</script>

<Content>
  <Grid>
    <Row>
      <Column>
        <Button onclick={() => console.log($state.snapshot(stamp))}>Test</Button>
      </Column>
      <Column>
        <TextInput bind:value={stamp.title} />
      </Column>
    </Row>
  </Grid>
</Content>