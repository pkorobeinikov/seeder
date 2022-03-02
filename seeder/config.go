package seeder

type (
	Spec struct {
		Seeder Seeder `yaml:"seeder"`
	}

	Seeder struct {
		State []State
	}

	State struct {
		Name   string `yaml:"name"`
		Type   string `yaml:"type"`
		Config []Config
	}

	Config struct {
		File       string `yaml:"file"`
		Key        string `yaml:"key"`
		Bucket     string `yaml:"bucket"`
		ObjectName string `yaml:"object-name"`
		Option     Option `yaml:"option"`
	}

	Option struct {
		ContentType     string `yaml:"content-type"`
		ContentEncoding string `yaml:"content-encoding"`
	}
)
