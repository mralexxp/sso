package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// Для пакета Cleanenv

type Config struct {
	Env         string        `yaml:"env" env-default:"local"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-required:"true"`
	GRPC        GRPCConfig    `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

// Must в названии используется в случае, когда функция не возвращает ошибки и обрабатывает ее внутри приложения
func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("no config path provided")
	}

	// Проверяем наличие файла по полученному пути. Если ошибка IsNotExist, то паника
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to load config: " + err.Error())
	}

	// TODO: Проверить обязательные поля конфига

	return &cfg
}

// Загружает путь к файлу конфигурации через фалг или env.
// Флаг является приоритетным в загрузке
func fetchConfigPath() string {
	var res string

	// --config="path_to_config.yaml"
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res

}
