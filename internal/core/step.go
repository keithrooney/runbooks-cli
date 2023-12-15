package core

type Step struct {
	Name  string `yaml:"name"`
	Shell Shell  `yaml:"shell"`
}

func (step Step) Execute() error {
	return step.Shell.Execute()
}
