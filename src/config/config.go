package config

type Config struct {
	secret_access  string
	db             string
	dbUsers        string
	dbtokens       string
	domain         string
	dbURI          string
	tokenDelimiter string
	maxAgeRefresh  int
	maxAgeAccess   int
}

var (
	cfg Config
)

func Init() {
	cfg = Config{
		secret_access:  "S0|||e_/ery_5Ecret_K3y_F0rG4iN4c3S5",
		db:             "test_server",
		dbUsers:        "users",
		dbtokens:       "tokens",
		dbURI:          "mongodb://127.0.0.1:27017/",
		domain:         "localhost",
		tokenDelimiter: "%",
		maxAgeAccess:   60 * 15,
		maxAgeRefresh:  60,
	}
}

func GetConfig() Config {
	return cfg
}

func (c *Config) GetAccessSecret() string {
	return c.secret_access
}

func (c *Config) GetMaxAgeAccess() int {
	return c.maxAgeAccess
}

func (c *Config) GetMaxAgeRefresh() int {
	return c.maxAgeRefresh
}

func (c *Config) GetDBPath() string {
	return c.db
}

func (c *Config) GetDBTokens() string {
	return c.dbtokens

}
func (c *Config) GetDBUsers() string {
	return c.dbUsers
}
func (c *Config) GetDomain() string {
	return c.domain
}

func (c *Config) GetDBURI() string {
	return c.dbURI
}
func (c *Config) GetTokenDelimiter() string {
	return c.tokenDelimiter
}
