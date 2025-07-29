<script lang="ts">
    import {
    Content,
    Grid,
    Row,
    Column,
    Form,
    FormGroup,
    TextInput,
    Button,
  } from "carbon-components-svelte";
    import { onMount } from "svelte";
    import store from '@src/app/store.svelte';
  interface Props {

  }
  let {} : Props = $props();

  let previous : string = $state('');
  let current : string = $state('');
  let modified : boolean = $derived(previous !== current);

  let password : string = $state('');

  onMount(async () => {
    if (!store.locked) {
      previous = await store.readTest();
    }
  });

  async function save() {
    await store.writeTest(current);
    previous = await store.readTest();
    current = '';
  }

  async function unlock() {
    await store.unlock(password);
    password = '';
    previous = await store.readTest();
  }

  async function changePassword() {
    await store.changePassword(password);
    password = '';
  }
</script>

<Content>
  <Grid>
    <Row>
      <Column>
        {#if store.locked}
        <Form>
          <FormGroup>
            <TextInput
              bind:value={password}
              labelText="Unlock db"
              placeholder="Enter password or leave empty if no password"
            />
          </FormGroup>
        </Form>
        <Button onclick={unlock}>Unlock</Button>
        {:else}
        <Form>
          <FormGroup>
            <TextInput
              bind:value={password}
              labelText="Set master password"
              placeholder="Enter password or leave empty if no password"
            />
          </FormGroup>
          <Button onclick={changePassword}>Change password</Button>
          <FormGroup>
            <TextInput
              bind:value={current}
              labelText="DB TEST"
              helperText={previous}
              placeholder="Enter new value"
            />
          </FormGroup>
        </Form>
        <Button disabled={!modified} onclick={save}>Save</Button>
        {/if}
      </Column>
    </Row>
  </Grid>
</Content>