package settings

import (
	"encoding/json"
	"os"

	"github.com/enolgor/pdfsigner/desktop/translations"
)

var defaultSettings map[string]string = map[string]string{
	"lang":  translations.DefaultLang,
	"theme": "light",
}

type Settings struct {
	path   string
	values map[string]string
}

func New(path string) (settings *Settings, err error) {
	var f *os.File
	settings = &Settings{
		path:   path,
		values: make(map[string]string),
	}
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		err = settings.Save(defaultSettings)
	} else {
		if f, err = os.Open(path); err != nil {
			return
		}
		defer f.Close()
		var values map[string]string
		if err = json.NewDecoder(f).Decode(&values); err != nil {
			return
		}
		err = settings.Save(values)
	}
	return
}

func (s *Settings) Save(values map[string]string) (err error) {
	var f *os.File
	if f, err = os.OpenFile(s.path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm); err != nil {
		return
	}
	defer f.Close()
	s.values = merge(values, defaultSettings)
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	err = enc.Encode(s.values)
	return
}

func (s *Settings) Get() map[string]string {
	return s.values
}

func merge(m1 map[string]string, m2 map[string]string) map[string]string {
	for k, v := range m2 {
		if _, ok := m1[k]; !ok {
			m1[k] = v
		}
	}
	return m1
}
