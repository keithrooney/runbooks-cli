package core

type Mitigation struct {
	Steps  []Step `yaml:"steps"`
	Clause Clause `yaml:"when"`
}

func (mitigation Mitigation) Evaluate() (bool, error) {
	return mitigation.Clause.Evaluate()
}

func (mitigation Mitigation) Execute(printf func(format string, i ...interface{})) error {
	for _, step := range mitigation.Steps {
		printf("STEP [%s] => [%s]\n\n", step.Name, step.Shell)
		if err := step.Execute(); err != nil {
			return err
		}
	}
	return nil
}
