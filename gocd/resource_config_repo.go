package gocd

// SetVersion sets a version string for this config repo
func (p *ConfigRepo) SetVersion(version string) {
	p.Version = version
}

// GetVersion retrieves a version string for this config repo
func (p *ConfigRepo) GetVersion() (version string) {
	return p.Version
}
