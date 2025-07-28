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
  } from "carbon-components-svelte";
  import Save from "carbon-icons-svelte/lib/Save.svelte";
  import i18n from '@src/app/i18n.svelte';
  import settings from '@src/app/settings.svelte';
  import theme from '@src/app/theme.svelte';
    import themeSvelte from '@src/app/theme.svelte';

  interface Props {

  }
  let {} : Props = $props();

  $effect(() => {
    i18n.lang = settings.settings.lang;
  });

</script>

<Content>
  <Grid>
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
        <Button icon={Save} disabled={!settings.modified} onclick={settings.save}>{$_('save')}</Button>
      </Column>
    </Row>
  </Grid>
</Content>