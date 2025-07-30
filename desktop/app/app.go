// MIT License
//
// Copyright (c) 2025 @enolgor
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package app

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/enolgor/pdfsigner/desktop/app/settings"
	"github.com/enolgor/pdfsigner/desktop/app/store"
	"github.com/enolgor/pdfsigner/desktop/app/translations"
	"github.com/goforj/godump"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var t = translations.Translate

// App struct
type App struct {
	ctx      context.Context
	appKey   string
	settings *settings.Settings
	dataDir  string
	db       *store.DB
}

// NewApp creates a new App application struct
func NewApp(appKey string) *App {
	godump.Dump(translations.Translations)
	return &App{appKey: appKey}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	var err error
	var configDir string
	if configDir, err = os.UserConfigDir(); err != nil {
		a.handleErr(err)
	}
	a.dataDir = path.Join(configDir, a.appKey)
	if err = os.MkdirAll(a.dataDir, os.ModePerm); err != nil {
		a.handleErr(err)
	}
	if a.settings, err = settings.New(path.Join(a.dataDir, "settings.json")); err != nil {
		a.handleErr(err)
	}
	if a.db, err = store.New(path.Join(a.dataDir, "data")); err != nil {
		a.handleErr(err)
	}
	if _, err := a.db.Flags().Get("first-run"); err != nil {
		if store.IsNotExist(err) {
			a.db.Flags().Set("first-run", true)
		}
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return t("hello", translations.Vars{"name": name})
}

func (a *App) Translations() map[string]map[string]string {
	return translations.Translations
}

func (a *App) GetFallbackLang() string {
	return translations.DefaultLang
}

func (a *App) GetLang() string {
	return translations.CurrentLang
}

func (a *App) SetLang(lang string) {
	translations.SetLang(lang)
}

func (a *App) Settings() map[string]string {
	return a.settings.All()
}

func (a *App) SaveSettings(values map[string]string) map[string]string {
	if err := a.settings.Save(values); err != nil {
		a.handleErr(err)
	}
	return values
}

func (a *App) IsStoreLocked() bool {
	return a.db.IsLocked()
}

func (a *App) UnlockStore(password string) error {
	if err := a.db.Unlock(password); err != nil {
		if store.IsInvalidPassword(err) {
			return errors.New(t("master-password.invalid"))
		}
		a.handleErr(err)
		return nil
	}
	return nil
}

func (a *App) ChangePassword(new string) {
	if err := a.db.Reencrypt(new); err != nil {
		a.handleErr(err)
		return
	}
}

func (a *App) ReadTest() string {
	val, err := a.db.TestBucket().Get("test")
	if err != nil {
		a.handleErr(err)
	}
	return val
}

func (a *App) WriteTest(value string) {
	if err := a.db.TestBucket().Set("test", value); err != nil {
		a.handleErr(err)
	}
}

func (a *App) IsFirstRun() bool {
	value, err := a.db.Flags().Get("first-run")
	if err != nil {
		a.handleErr(err)
	}
	return value
}

func (a *App) FirstRunCompleted() {
	if err := a.db.Flags().Set("first-run", false); err != nil {
		a.handleErr(err)
	}
}

func (a *App) IsStoreProtected() bool {
	return a.db.IsProtected()
}

func (a *App) OpenFileDialog(info, extensions string) {
	runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})
}

func (a *App) handleErr(err error) {
	fmt.Printf("error: %s\n", err.Error()) //TODO
}
