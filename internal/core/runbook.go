package core

type Runbook struct {
	Title      string     `yaml:"title"`
	Details    string     `yaml:"details"`
	Impact     string     `yaml:"impact"`
	Mitigation Mitigation `yaml:"mitigation"`
}

func (runbook *Runbook) Apply(printf func(format string, i ...interface{})) error {
	printf("RUNBOOK [%s]\n\n", runbook.Title)
	return runbook.Mitigation.Execute(printf)
}
