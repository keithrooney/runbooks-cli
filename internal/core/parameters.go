package core

import "fmt"

type Parameters []Parameter

func (parameters Parameters) Valid(arguments map[string]string) error {
	if len(parameters) > len(arguments) {
		return fmt.Errorf("argument errors: insufficient number of arguments provided\n")
	}
	for _, parameter := range parameters {
		if _, found := arguments[parameter.Name]; !found {
			return fmt.Errorf("argument errors: %s was not provided\n", parameter.Name)
		}
	}
	return nil
}
