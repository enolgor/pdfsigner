import {initialize as initi18n} from './i18n.svelte';

export default async function initialize() {
  await initi18n();
}