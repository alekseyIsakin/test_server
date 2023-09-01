package config

type Config struct {
	secret         string
	db             string
	colletionUsers string
	dbURI          string
	tokenDelimiter string
}

var (
	cfg Config
)

func Init() {
	cfg = Config{
		secret:         "S0|||e_\\/ery_5Ecret_K3y",
		db:             "test_server",
		colletionUsers: "users",
		dbURI:          "mongodb://127.0.0.1:27017/",
		tokenDelimiter: "%",
	}
}

func GetConfig() Config {
	return cfg
}

func (c *Config) GetSecret() string {
	return c.secret
}

func (c *Config) GetDBPath() string {
	return c.db
}

func (c *Config) GetUsersCollectionPath() string {
	return c.colletionUsers
}

func (c *Config) GetDBURI() string {
	return c.dbURI
}
func (c *Config) GetTokenDelimiter() string {
	return c.tokenDelimiter
}
