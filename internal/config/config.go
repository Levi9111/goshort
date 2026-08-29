package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config represents our environment variables.
// Think of a Struct like a JavaScript Object blueprint, but with strict typing.
// The `mapstructure` tags tell Viper how to map the .env keys to these fields.

type config struct {
	Port  			string `mapstructure:"PORT"`
	MongoURI		string `mapstructure:"MONGO_URI"`
	RedisAddr		string `mapstructure:"REDIS_ADDR"`
	JWTSecret		string `mapstructure:"JWT_SECRET"`
}

// LoadConfig reads the .env file and populates our Config struct.
// It returns a POINTER to the Config struct (*Config) and an error.	

func LoadConfig() (*config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv() // ---> reads from OS environment vairables

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Warning: .env file not found, relying on OS env vars: %v",err)
	}

	var cfg config // Create empty instances of our COnfig struct


	// POINTERS EXPLAINED:
	// In Node.js, objects are always passed by reference.
	// In Go, everything is passed by VALUE (a complete clone is made) by default.
	// If we passed 'cfg' directly, Viper would populate a clone, and our original 'cfg' would stay empty!
	// By putting '&' in front of cfg, we pass the memory ADDRESS of cfg (a pointer).
	// Viper can now reach into that exact memory slot and populate our actual struct.
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	// We return '&cfg' (the address) so the rest of our app shares this single config object
	// rather than duplicating it everywhere. The function signature says we return '*Config',
	// which means "a pointer to a Config struct".
	return &cfg, nil
}