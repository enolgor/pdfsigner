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

package fonts

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

func getSystemFontDirectories() []string {
	switch runtime.GOOS {
	case "windows":
		return windowsFontDirectories()
	case "linux":
		return linuxFontDirectories()
	case "darwin":
		return darwinFontDirectories()
	default:
		return []string{}
	}
}

func windowsFontDirectories() []string {
	return []string{
		filepath.Join(os.Getenv("windir"), "Fonts"),
		filepath.Join(os.Getenv("localappdata"), "Microsoft", "Windows", "Fonts"),
	}
}

func linuxFontDirectories() []string {
	directories := linuxUserFontDirs()
	directories = append(directories, linuxSystemFontDirs()...)
	return directories
}

func darwinFontDirectories() []string {
	return []string{
		expandUser("~/Library/Fonts/"),
		"/Library/Fonts/",
		"/System/Library/Fonts/",
	}
}

func linuxUserFontDirs() (paths []string) {
	if dataPath := os.Getenv("XDG_DATA_HOME"); dataPath != "" {
		return []string{expandUser("~/.fonts/"), filepath.Join(expandUser(dataPath), "fonts")}
	}
	return []string{expandUser("~/.fonts/"), expandUser("~/.local/share/fonts/")}
}

func linuxSystemFontDirs() (paths []string) {
	if dataPaths := os.Getenv("XDG_DATA_DIRS"); dataPaths != "" {
		for _, dataPath := range filepath.SplitList(dataPaths) {
			paths = append(paths, filepath.Join(expandUser(dataPath), "fonts"))
		}
		return paths
	}
	return []string{"/usr/local/share/fonts/", "/usr/share/fonts/"}
}

func expandUser(path string) (expandedPath string) {
	if strings.HasPrefix(path, "~") {
		if u, err := user.Current(); err == nil {
			return strings.Replace(path, "~", u.HomeDir, -1)
		}
	}
	return path
}
