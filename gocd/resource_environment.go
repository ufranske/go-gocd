package gocd

func (er *EnvironmentsResponse) RemoveLinks() {
	er.Links = nil
	for _, env := range er.Embedded.Environments {
		env.RemoveLinks()
	}
}

func (env *Environment) RemoveLinks() {
	env.Links = nil
	for _, p := range env.Pipelines {
		p.RemoveLinks()
	}
	for _, a := range env.Agents {
		a.RemoveLinks()
	}
}

