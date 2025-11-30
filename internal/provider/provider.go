package provider

type Usage struct {
	Path     string
	Size     int64
	Label    string
	Children []Usage
}

type Provider interface {
	Name() string
	Paths() []string
}
