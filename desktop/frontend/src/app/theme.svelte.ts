import type { CarbonTheme } from 'carbon-components-svelte/src/Theme/Theme.svelte';
import settings from './settings.svelte';

let carbontheme : CarbonTheme | undefined = $state(undefined);

export default {
  get carbontheme() : CarbonTheme | undefined {
    return carbontheme;
  },
  get carbonthemes() : [labelA: CarbonTheme, labelB: CarbonTheme] {
    return ["g10", "g80"];
  },
  set carbontheme(value: CarbonTheme | undefined) {
    carbontheme = value;
    switch (value) {
      case "g80":
        settings.settings.theme = "dark";
        break;
      default:
        settings.settings.theme = "light";
    }
  },
  set theme(value: string) {
    settings.settings.theme = value;
    switch (value) {
      case "dark":
        carbontheme = "g80";
        break;
      default:
        carbontheme = "g10";
    }
    document.documentElement.setAttribute("theme", carbontheme);
  }
}