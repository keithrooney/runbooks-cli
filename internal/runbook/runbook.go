package runbook

type Runbook struct {
	Title      string     `yaml:"title"`
	Details    string     `yaml:"details"`
	Mitigation Mitigation `yaml:"mitigation"`
}
