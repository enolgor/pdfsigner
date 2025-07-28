
import { addMessages, init, locale } from 'svelte-i18n';
import { Translations, SetLang, GetFallbackLang, GetLang } from '@go'

let _langs : string[] = $state([]);
let _lang : string = $state('');


export async function initialize() {
  const translations = await Translations();
  _langs = Object.keys(translations);
  for (let lang of Object.keys(translations)) {
    addMessages(lang, translations[lang]);
  }
  const fallbackLocale = await GetFallbackLang();
  const initialLocale = await GetLang();
  _lang = initialLocale;
  init({fallbackLocale, initialLocale});
}

const i18n = {
  get langs() : string[] {
    return _langs;
  },
  get lang() : string {
    return _lang;
  },
  set lang(value: string) {
    SetLang(value).then(() => {
      _lang = value;
      locale.set(value);
    });
  }
};

export const { langs, lang } = i18n;
export default i18n;