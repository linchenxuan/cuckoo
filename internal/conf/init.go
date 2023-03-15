package conf

import (
	"cuckoo/internal/adapter/repository"
	"cuckoo/pkg/logger"
	"fmt"
	"github.com/spf13/viper"
)

var (
	repositoryOpt repository.RepositoryOpt
	loggerConf    logger.LoggerConfig
	jwtConfig     JWTConfig
)

// TODO 环境区分
func Init() {
	viper.SetConfigName("cuckoo")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	objects := map[string]interface{}{
		"Repository": &repositoryOpt,
		"Logger":     &loggerConf,
		"JWT":        &jwtConfig,
	}

	for k, v := range objects {
		if viper.UnmarshalKey(k, v) != nil {
			panic(fmt.Errorf("unable to decode into struct, %v\n", err))
		}
	}
}
