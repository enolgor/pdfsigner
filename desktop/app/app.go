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
	"net/http"
	"os"
	"path"
	"time"

	"github.com/enolgor/pdfsigner/desktop/app/certs"
	"github.com/enolgor/pdfsigner/desktop/app/settings"
	"github.com/enolgor/pdfsigner/desktop/app/stamps"
	"github.com/enolgor/pdfsigner/desktop/app/store"
	"github.com/enolgor/pdfsigner/desktop/app/translations"
	"github.com/enolgor/pdfsigner/signer"
	"github.com/enolgor/pdfsigner/signer/config"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var t = translations.Translate

// App struct
type App struct {
	ctx          context.Context
	appKey       string
	settings     *settings.Settings
	dataDir      string
	db           *store.DB
	unsavedStamp *stamps.StampConfig
	Mux          *http.ServeMux
}

// NewApp creates a new App application struct
func NewApp(appKey string) *App {
	app := &App{appKey: appKey}
	app.Mux = http.NewServeMux()
	app.Mux.HandleFunc("POST /unsaved-stamp", app.serveUnsavedStamp)
	return app
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

func (a *App) OpenFileDialog(extensions string) (string, error) {
	filters := []runtime.FileFilter{}
	if extensions != "" {
		filters = append(filters, runtime.FileFilter{
			DisplayName: extensions,
			Pattern:     extensions,
		})
	}
	return runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Filters: filters,
	})
}

func (a *App) GetCertificateID(path, passphrase string) (id certs.StoredCertificateID, err error) {
	fmt.Printf("trying to open %s \n", path)
	data, err := os.ReadFile(path)
	if err != nil {
		a.handleErr(err)
		err = nil
		return
	}
	var sc *certs.StoredCertificate
	if sc, err = certs.NewStoredCertificate(data, passphrase); err != nil {
		return
	}
	id = sc.StoredCertificateID
	return
}

func (a *App) StoreCertificate(path, passphrase string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	sc, err := certs.NewStoredCertificate(data, passphrase)
	if err != nil {
		return err
	}
	return a.db.Certs().Set(fmt.Sprintf("%s / %s", sc.Subject, sc.Issuer), *sc)
}

func (a *App) ListCertificates() []string {
	return a.db.Certs().Keys()
}

func (a *App) DeleteCertificate(key string) error {
	return a.db.Certs().Delete(key)
}

func (a *App) SetDefaultCertificate(key string) {
	a.db.Certs().Move(key, 0)
}

func (a *App) GetDefaultCertificate() (cert certs.StoredCertificate, err error) {
	keys := a.db.Certs().Keys()
	if len(keys) == 0 {
		err = errors.New("no certificates") //TODO
		return
	}
	return a.db.Certs().Get(keys[0])
}

func (a *App) GetStoredCertificateID(key string) (id certs.StoredCertificateID, err error) {
	var sc certs.StoredCertificate
	if sc, err = a.db.Certs().Get(key); err != nil {
		return
	}
	id = sc.StoredCertificateID
	return
}

func (a *App) NewDefaultStampConfig() stamps.StampConfig {
	sc := stamps.StampConfig{}
	sc.FromConfig(config.New())
	return sc
}

func (a *App) SetUnsavedStamp(sc *stamps.StampConfig) {
	a.unsavedStamp = sc
}

func (a *App) serveUnsavedStamp(w http.ResponseWriter, req *http.Request) {
	if a.unsavedStamp == nil {
		http.Error(w, "unsaved stamp not found", http.StatusNotFound)
		return
	}
	cfg, err := a.unsavedStamp.ToConfig("")
	if err != nil {
		http.Error(w, "internal", http.StatusInternalServerError)
		return
	}
	cert, err := a.GetDefaultCertificate()
	if err != nil {
		http.Error(w, "default certificate not found", http.StatusNotFound)
		return
	}
	unlocked, err := cert.Unlock()
	if err != nil {
		http.Error(w, "internal", http.StatusInternalServerError)
		return
	}
	date := time.Now()
	w.Header().Add("Content-Type", "image/png")
	if err := signer.DrawPngImage(w, date, unlocked, cfg); err != nil {
		http.Error(w, "internal", http.StatusInternalServerError)
		return
	}
}

func (a *App) handleErr(err error) {
	fmt.Printf("error: %s\n", err.Error()) //TODO
}
