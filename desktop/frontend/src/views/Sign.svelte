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
</script>

<Content>
  <Grid>
    <Row>
      <Column>
        <Form>
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
      </Column>
    </Row>
  </Grid>
</Content>