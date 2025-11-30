package provider

import (
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
)

type npm struct{}
type yarn struct{}
type pnpm struct{}
type gradle struct{}
type maven struct{}
type cargo struct{}
type pip struct{}
type gomod struct{}

func homeDir() string {
	h, _ := os.UserHomeDir()
	return h
}

func cacheDir() string {
	// XDG first, then OS default
	if xdg := os.Getenv("XDG_CACHE_HOME"); xdg != "" {
		return xdg
	}
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(os.Getenv("LOCALAPPDATA"), "cache")
	case "darwin":
		return filepath.Join(homeDir(), "Library", "Caches")
	default: // linux, freebsd, …
		return filepath.Join(homeDir(), ".cache")
	}
}

func (npm) Name() string    { return "npm" }
func (yarn) Name() string   { return "yarn" }
func (pnpm) Name() string   { return "pnpm" }
func (gradle) Name() string { return "gradle" }
func (maven) Name() string  { return "maven" }
func (cargo) Name() string  { return "cargo" }
func (pip) Name() string    { return "pip" }
func (gomod) Name() string  { return "go" }

func (npm) Paths() []string {
	return []string{filepath.Join(homeDir(), ".npm")}
}
func (yarn) Paths() []string {
	return []string{
		filepath.Join(homeDir(), ".yarn", "cache"),
		filepath.Join(cacheDir(), "yarn"),
	}
}
func (pnpm) Paths() []string {
	return []string{filepath.Join(homeDir(), ".local", "share", "pnpm")}
}

func (gradle) Paths() []string {
	return []string{filepath.Join(homeDir(), ".gradle", "caches")}
}

func (maven) Paths() []string {
	return []string{filepath.Join(homeDir(), ".m2", "repository")}
}

func (cargo) Paths() []string {
	return []string{filepath.Join(homeDir(), ".cargo", "registry")}
}

func (pip) Paths() []string {
	// both legacy and XDG
	return []string{
		filepath.Join(homeDir(), ".cache", "pip"),
		filepath.Join(homeDir(), ".pip", "cache"),
	}
}

func (gomod) Paths() []string {
	return []string{
		filepath.Join(cacheDir(), "go-build"),
		filepath.Join(homeDir(), "go", "pkg", "mod"),
	}
}

func Scan(root string, label string) (Usage, error) {
	var size int64
	var children []Usage

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil { // permission denied, deleted race, etc.
			return nil // skip
		}
		// do not cross mount points (optional safety)
		//if d.IsDir() && path != root {
		//	if st, ok := d.Info(); ok {
		//		if st.Mode()&os.ModeDir != 0 && isMountRoot(path) {
		//			return filepath.SkipDir
		//		}
		//	}
		//}
		if !d.Type().IsRegular() {
			return nil
		}
		info, err := d.Info()
		if err != nil {
			return nil
		}
		size += info.Size()
		return nil
	})

	if err != nil {
		return Usage{}, err
	}
	return Usage{
		Path:     root,
		Size:     size,
		Label:    label,
		Children: children,
	}, nil
}

// GlobalProviders returns all built-in global-cache providers.
func GlobalProviders() []Provider {
	return []Provider{
		npm{}, yarn{}, pnpm{},
		gradle{}, maven{},
		cargo{},
		pip{},
		gomod{},
	}
}
