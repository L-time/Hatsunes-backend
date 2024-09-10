package configs

type Log struct {
	Level string `yaml:"level"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Config struct {
	Log Log `yaml:"log"`
	DB  DB  `yaml:"db"`
}
