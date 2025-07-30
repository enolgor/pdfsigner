<script lang="ts">
  import { _ } from 'svelte-i18n';
  import {
    Column,
    Form,
    FormGroup,
    Grid,
    PasswordInput,
    Row
  } from 'carbon-components-svelte';
  import store from '@src/app/store.svelte';

  type strength = "very-weak" | "weak" | "strong" | "very-strong";

  interface Props {
    password : string
    retype : string
  };

  let { password = $bindable(''), retype = $bindable('') } : Props = $props();

  let isEmpty : boolean = $derived(password.trim() === "");
  let match : boolean = $derived(password === retype);

  let passwordStrength : strength | undefined = $derived(isEmpty ?
    undefined : getPasswordStrength(password));
  let strengthText : string | undefined = $derived(
    passwordStrength ?
      $_("master-password-form.strength.label",
      { 
        values: {
          strength: $_("master-password-form.strength." + passwordStrength)
        }
      })
    : undefined
  );
  let warn : boolean = $derived(
    !isEmpty && (passwordStrength === "very-weak" || passwordStrength === "weak")
  );
  let warnText : string | undefined = $derived(warn ? strengthText : undefined);
  let helperText : string | undefined = $derived(!warn ? strengthText : undefined);


  function getPasswordStrength(password: string) : strength {
    let score = 0;

    if (!password) return "very-weak";
    if (password.length >= 8) score++;
    if (password.length >= 12) score++;

    if (/[a-z]/.test(password)) score++;
    if (/[A-Z]/.test(password)) score++;
    if (/\d/.test(password)) score++;
    if (/[\W_]/.test(password)) score++;

    if (score <= 2) return "very-weak";
    if (score <= 4) return "weak";
    if (score <= 5) return "strong";
    return "very-strong";
  }
</script>

<Grid padding>
  <Row>
    {#if store.firstRun || store.protected}
     <Column><p>{$_("master-password-form.info-1")}</p></Column>
    {:else}
    <Column><p>{$_("master-password-form.info-2")}</p></Column>
    {/if}
    
  </Row>
  <Row>
    <Column>
      <Form>
        <FormGroup>
          <PasswordInput
            bind:value={password}
            labelText={$_("master-password-form.label-1")}
            placeholder={$_("master-password-form.placeholder-1")}
            {helperText}
            {warn}
            {warnText}
          />
        </FormGroup>
        <FormGroup>
          <PasswordInput
            bind:value={retype}
            labelText={$_("master-password-form.label-2")}
            placeholder={$_("master-password-form.placeholder-2")}
            invalid={!match}
            invalidText={$_("master-password-form.nomatch")}
          />
        </FormGroup>
      </Form>
    </Column>
  </Row>
</Grid>

