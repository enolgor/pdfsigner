import { initialize as initi18n, default as i18n } from './i18n.svelte';
import { initialize as initSettings, default as settings } from './settings.svelte';
import { initialize as initStore } from './store.svelte';
import theme from './theme.svelte';

export default async function initialize() {
  await initSettings();
  await initi18n();
  await initStore();
  i18n.lang = settings.settings.lang;
  theme.theme = settings.settings.theme;
}