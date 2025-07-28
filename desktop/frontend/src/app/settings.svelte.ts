import { Settings, SaveSettings } from "@go";
import isEqual from 'lodash/isEqual';

let persisted : Record<string,string> = $state({});
let settings : Record<string,string> = $state({});
let modified : boolean = $derived(!isEqual(persisted, settings));

export async function initialize() {
  settings = await Settings();
  persisted = $state.snapshot(settings);
}

export async function save() : Promise<void> {
  persisted = await SaveSettings(settings);
}

export default {
  get modified() : boolean {
    return modified;
  },
  get settings(): Record<string,string> {
    return settings;
  },
  save,
};
