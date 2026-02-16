package fl

import "os"

// Flag define la estructura para nuestras banderas booleanas
type Flag struct {
	value       *bool
	description string
}

var flags = make(map[string]*Flag)

func Bool(cmd string, defaultValue bool, description string) *bool {
	v := defaultValue

	flags[cmd] = &Flag{
		value:       &v,
		description: description,
	}
	return &v
}

// Parse recorre os.Args y actualiza los valores de las banderas encontradas
func Parse() {
	args := os.Args[1:]

	for _, arg := range args {
		if f, exists := flags[arg]; exists {
			*f.value = true
		}
	}
}
