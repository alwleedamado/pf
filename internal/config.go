package internal

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type config struct {
	Dirs []string
}

var dirs = []string{
	"~/.cache/ccache",
	"~/.cache/clangd",
	"~/.cargo/registry",
	"~/.cargo/git",
	"~/go/pkg/mod",
	"~/.npm",
	"~/.yarn",
	"~/.yarn/berry/cache",
	"~/.cache/pnpm",
	"~/.cache/pip",
	"~/.cache/poetry",
	"~/.local/share/virtualenvs",
	"~/.gradle/caches",
	"~/.gradle/wrapper",
	"~/.gradle/daemon",
	"~/.m2/repository",
	"~/.cache/JetBrains",
	"~/.config/JetBrains",
	"~/.config/Code/Cache",
	"~/.config/Code/CachedData",
	"~/.config/Code/CachedExtensions",
	"~/.config/Code/User/workspaceStorage",
	"~/Android/Sdk/.download",
}

func CreateCinfg() {
	var conf config
	var d []string
	for _, dir := range dirs {
		if _, err := os.Stat(ExpandTilde(dir)); err != nil {
			d = append(d, dir)
		}
	}
	conf.Dirs = d
	content, err := yaml.Marshal(&conf)
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(GetConfigPath(), content, 0644); err != nil {
		panic(err)
	}
}

func GetConfigPath() string {
	c, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	configPath := filepath.Join(c, "pf.cnf")
	if _, err := os.Stat(configPath); err != nil {
		os.Create(configPath)
	}
	return configPath
}

func ListConfig() config {
	path := GetConfigPath()
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var conf config
	if err := yaml.Unmarshal(content, &conf); err != nil {
		panic(err)
	}
	return conf
}

func AddDirectory(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}

	conf := ListConfig()
	conf.Dirs = append(conf.Dirs, path)
	conf.Dirs = RemoveDup(conf.Dirs)
	out, err := yaml.Marshal(&conf)
	if err != nil {
		return false
	}
	if err := os.WriteFile(GetConfigPath(), out, 0644); err != nil {
		return false
	}
	return true
}

func GetDirectories() []string {
	conf := ListConfig()
	return conf.Dirs
}

func RemoveDirectory(path string) bool {
	conf := ListConfig()
	conf.Dirs = RemoveElement(conf.Dirs, path)
	out, err := yaml.Marshal(conf)
	if err != nil {
		return false
	}
	if err := os.WriteFile(GetConfigPath(), out, 0644); err != nil {
		return false
	}
	return true
}
