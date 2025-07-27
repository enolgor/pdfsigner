package translations

import (
	"embed"
	"fmt"
	"io"
	"maps"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/kaptinlin/go-i18n"
)

//go:embed yaml
var fs embed.FS

type Vars = i18n.Vars

const DefaultLang = "en"

var Translations map[string]map[string]string = make(map[string]map[string]string)

var bundle *i18n.I18n

var localizer *i18n.Localizer

var CurrentLang string

func init() {
	entries, err := fs.ReadDir("yaml")
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || (!strings.HasSuffix(name, ".yaml") && !strings.HasSuffix(name, ".yml")) {
			continue
		}
		fmt.Println("open", name)
		f, err := fs.Open("yaml/" + name)
		if err != nil {
			panic(err)
		}
		lang := strings.TrimSuffix(strings.TrimSuffix(name, ".yaml"), ".yml")
		parseTranslationYaml(lang, f)
		f.Close()
	}
	langs := []string{}
	for lang := range maps.Keys(Translations) {
		langs = append(langs, lang)
	}
	bundle = i18n.NewBundle(
		i18n.WithDefaultLocale(DefaultLang),
		i18n.WithLocales(langs...),
	)
	if err := bundle.LoadMessages(Translations); err != nil {
		panic(err)
	}
	SetLang(DefaultLang)
}

func parseTranslationYaml(lang string, reader io.Reader) {
	data := map[string]any{}
	if err := yaml.NewDecoder(reader).Decode(&data); err != nil {
		panic(err)
	}
	flat, err := flatten(data)
	if err != nil {
		panic(err)
	}
	Translations[lang] = flat
}

func flatten(data map[string]any) (map[string]string, error) {
	var flatten map[string]string = make(map[string]string)
	doflat(flatten, "", data)
	return flatten, nil
}

func doflat(agg map[string]string, subkey string, data map[string]any) {
	subk := ""
	if subkey != "" {
		subk = subkey + "."
	}
	for k, v := range data {
		switch typedv := v.(type) {
		case string:
			agg[subk+k] = typedv
		case map[string]any:
			doflat(agg, subk+k, typedv)
		case map[string]string:
			for k2, v2 := range typedv {
				agg[subk+k+"."+k2] = v2
			}
		}
	}
}

func SetLang(lang string) {
	CurrentLang = lang
	localizer = bundle.NewLocalizer(lang)
}

func Localizer(lang string) func(key string, vars ...Vars) string {
	localizer := bundle.NewLocalizer(lang)
	return func(key string, vars ...Vars) string {
		return localizer.Get(key, vars...)
	}
}

func Translate(key string, vars ...Vars) string {
	return localizer.Get(key, vars...)
}
