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
		return windows_fontDirectories()
	case "linux":
		return linux_fontDirectories()
	case "darwin":
		return darwin_fontDirectories()
	default:
		return []string{}
	}
}

func windows_fontDirectories() []string {
	return []string{
		filepath.Join(os.Getenv("windir"), "Fonts"),
		filepath.Join(os.Getenv("localappdata"), "Microsoft", "Windows", "Fonts"),
	}
}

func linux_fontDirectories() []string {
	directories := linux_userFontDirs()
	directories = append(directories, linux_systemFontDirs()...)
	return directories
}

func darwin_fontDirectories() []string {
	return []string{
		expandUser("~/Library/Fonts/"),
		"/Library/Fonts/",
		"/System/Library/Fonts/",
	}
}

func linux_userFontDirs() (paths []string) {
	if dataPath := os.Getenv("XDG_DATA_HOME"); dataPath != "" {
		return []string{expandUser("~/.fonts/"), filepath.Join(expandUser(dataPath), "fonts")}
	}
	return []string{expandUser("~/.fonts/"), expandUser("~/.local/share/fonts/")}
}

func linux_systemFontDirs() (paths []string) {
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
