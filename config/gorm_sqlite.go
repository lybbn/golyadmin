package config

type Sqlite struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (m *Sqlite) Dsn() string {
	return m.Path
}

func (m *Sqlite) GetLogMode() string {
	return m.LogMode
}
