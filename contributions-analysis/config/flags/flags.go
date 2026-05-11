package flags

type Configuration struct {
	LogLevel string `yaml:"loglevel" default:"info"`

	LiveURL string `yaml:"liveurl"`

	Address string `yaml:"address" default:":8080"`

	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`

	Postgres struct {
		Database string `yaml:"database"  required:"true"`
		Port     string `yaml:"port"      required:"true"`
		Host     string `yaml:"host"      required:"true"`
		Driver   string `yaml:"driver"    required:"true"`
		MaxConn  int    `yaml:"maxconn"   default:"70"`
		Username string `yaml:"username"  required:"true"`
		Password string `yaml:"password"  required:"true"`
	} `yaml:"postgres"`

	PostgresCredential struct {
		URI string `yaml:"url"`
	} `yaml:"postgresCredentials"`

	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`

	SMTP struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"smtp"`
}
