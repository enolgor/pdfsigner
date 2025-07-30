<script lang="ts">
  import { _ } from 'svelte-i18n';
  import {
    Content,
    Grid,
    Row,
    Column,
    Form,
    FormGroup,
    Checkbox,
    RadioButtonGroup,
    RadioButton,
    Select,
    SelectItem,
    Button,
    Theme,
    Accordion,
    AccordionItem,
    ButtonSet,
  } from "carbon-components-svelte";
  import i18n from '@src/app/i18n.svelte';
  import settings from '@src/app/settings.svelte';
  import theme from '@src/app/theme.svelte';
  import themeSvelte from '@src/app/theme.svelte';
  import MasterPassword from '@src/components/MasterPassword.svelte';
  import Save from "carbon-icons-svelte/lib/Save.svelte";
  import Erase from "carbon-icons-svelte/lib/Erase.svelte";
  import store from '@src/app/store.svelte';

  interface Props {

  }
  let {} : Props = $props();

  let password : string = $state('');
  let retype : string = $state('');
  let isEmpty : boolean = $derived(password.trim() === "");
  let match : boolean = $derived(password === retype);

  $effect(() => {
    i18n.lang = settings.settings.lang;
  });

  async function changePassword() {
    try {
      await store.changePassword(password);
      password = '';
      retype = '';
    } catch (err) {
      console.error(err); //TODO
    }
  }

</script>

<Content>
  <Grid padding>
    <Row>
      <Column>
        <Form>
          <FormGroup>
            <Select labelText={$_("language")} bind:selected={settings.settings.lang}>
              {#each i18n.langs as lang}
              <SelectItem value={lang} text={$_(`lang.${lang}`)} />
              {/each}
            </Select>
          </FormGroup>
          <FormGroup>
          <Theme
            bind:theme={theme.carbontheme}
            render="toggle"
            toggle={{
              themes: themeSvelte.carbonthemes,
              labelA: $_("enable-darkmode"),
              labelB: $_("enable-darkmode"),
              hideLabel: true,
              size: "sm",
            }}
          />
          </FormGroup>
        </Form>
        <Button disabled={!settings.modified} onclick={settings.save}>
          {$_('save')}
        </Button>
      </Column>
    </Row>
    <Row>
      <Column>
        <Accordion>
          <AccordionItem title={store.protected ? $_("change-master-password") : $_("set-master-password")}>
            <MasterPassword bind:password bind:retype />
            <ButtonSet>
              {#if store.protected}
              <Button size="field" kind="danger" disabled={!isEmpty} icon={Erase} onclick={changePassword}>
                {$_("unset-master-password")}
              </Button>
              {/if}
              <Button size="field" disabled={isEmpty || !match} icon={Save} onclick={changePassword}>
                {$_('set-new-password')}
              </Button>
            </ButtonSet>
          </AccordionItem>
        </Accordion>
      </Column>
    </Row>
  </Grid>
</Content>