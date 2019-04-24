import Vue from 'vue'
import VueI18n from 'vue-i18n'
import messages from '@/locales/en'

Vue.use(VueI18n)

// static translations loading
// function loadLocaleMessages () {
//   const locales = require.context('./locales', true, /[A-Za-z0-9-_,\s]+\.json$/i)
//   const messages = {}
//   locales.keys().forEach(key => {
//     const matched = key.match(/([A-Za-z0-9-_]+)\./i)
//     if (matched && matched.length > 1) {
//       const locale = matched[1]
//       messages[locale] = locales(key)
//     }
//   })
//   return messages
// }

// export const i18n = new VueI18n({
//   locale: process.env.VUE_APP_I18N_LOCALE || 'en',
//   fallbackLocale: process.env.VUE_APP_I18N_FALLBACK_LOCALE || 'en',
//   messages: loadLocaleMessages()
// })

// lazy loading of translations
export const i18n = new VueI18n({
  locale: 'en', // set locale
  fallbackLocale: 'en',
  messages // set locale messages
})

const loadedLanguages = ['en'] // our default language that is preloaded

function setI18nLanguage (lang) {
  i18n.locale = lang
  return lang
}

export function loadLanguageAsync (lang) {
  if (i18n.locale !== lang) {
    if (!loadedLanguages.includes(lang)) {
      return import(/* webpackChunkName: "lang-[request]" */ `@/locales/${lang}`).then(msgs => {
        i18n.setLocaleMessage(lang, msgs.default)
        loadedLanguages.push(lang)
        return setI18nLanguage(lang)
      })
    }
    return Promise.resolve(setI18nLanguage(lang))
  }
  return Promise.resolve(lang)
}
