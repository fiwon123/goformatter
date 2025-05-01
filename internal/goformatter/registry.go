package goformatter

var registry = make(map[string]FormatterInterface)

func Register(ext string, f FormatterInterface) {
	registry[ext] = f
}

func GetFormatter(ext string) (FormatterInterface, bool) {
	f, ok := registry[ext]
	return f, ok
}
