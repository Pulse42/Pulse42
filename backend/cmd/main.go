package main

import (
	"backend/pkg/database"
	"backend/pkg/models"
	"backend/pkg/routes/auth"
	"backend/pkg/routes/example"
	"backend/pkg/routes/user"
	"backend/pkg/store"
	"backend/pkg/utils"
	"backend/pkg/validator"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// Config is the config structure for the application. yaml and environment matches must be defined. Special characters are stripped in environment variables
type Config struct {
	SessionStorage string `yaml:"session_storage" envConfig:"SESSION_STORAGE"`
	DatabaseDriver string `yaml:"database_driver" envConfig:"DATABASE_DRIVER"`
	DatabaseDsn    string `yaml:"database_dsn" envConfig:"DATABASE_DSN"`
}

// config is the config for the application, default values needs to be set explicitly
var config = Config{
	SessionStorage: "memory",
	DatabaseDriver: "sqlite",
	DatabaseDsn:    "database.sqlite",
}

// loadConfig returns the config for the application, tries to load from a yaml config file first, then from environment variables
func loadConfig() *Config {
	newConfig := Config{}
	utils.AutoPanic(envconfig.Process("", &newConfig))

	_, err := os.Stat("./config.yaml")
	if err == nil {
		yamlConfig := Config{}
		data, err := os.ReadFile("./config.yaml")
		utils.AutoPanic(err)
		utils.AutoPanic(yaml.Unmarshal(data, &yamlConfig))
		utils.FillAgainst(&newConfig, &yamlConfig)
	}

	utils.FillAgainst(&newConfig, &config)
	config = newConfig
	return &config
}

func autoMigrate() {
	utils.AutoPanic(database.Database.AutoMigrate(models.UserModel{}))
}

func main() {
	loadConfig()

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(logger.New())

	store.Store = store.NewStore(config.SessionStorage)
	validator.Validator = validator.NewValidator()
	database.Database = database.NewDatabase(config.DatabaseDriver, config.DatabaseDsn)

	autoMigrate()

	example.Register(app)
	auth.Register(app)
	user.Register(app)

	log.Fatal(app.Listen(":3000"))
}
