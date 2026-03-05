package fl

import "os"

type Flag struct {
	Value       *bool
	Description string
}

var flags = make(map[string]*Flag)

func Bool(name string, defaultValue bool, description string) *bool {
	v := defaultValue
	f := &Flag{
		Value:       &v,
		Description: description,
	}
	flags[name] = f
	return f.Value
}

func Parse() {
	args := os.Args[1:]

	for _, arg := range args {
		if f, exists := flags[arg]; exists {
			*f.Value = true
		}
	}
}
