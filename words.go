package main

type Words struct {
	CaseSensitive bool     `yaml:"case_sensitive"`
	Targets       []string `yaml:"targets,flow"`
}
