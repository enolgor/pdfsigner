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
    import db from '@src/app/db.svelte';
  interface Props {

  }
  let {} : Props = $props();

  let previous : string = $state('');
  let current : string = $state('');
  let modified : boolean = $derived(previous !== current);

  let password : string = $state('');

  onMount(async () => {
    previous = await db.readTest();
  });

  async function save() {
    await db.writeTest(current);
    previous = await db.readTest();
    current = '';
  }

  async function unlock() {
    await db.unlock(password);
    password = '';
  }

  async function changePassword() {
    await db.changePassword(password);
    password = '';
  }
</script>

<Content>
  <Grid>
    <Row>
      <Column>
        {#if db.locked}
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