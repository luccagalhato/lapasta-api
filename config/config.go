package config

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// Config representa a estrutura do arquivo de configuração.
type Config struct {
	API struct {
		Host string `yaml:"host,omitempty"`
		Port string `yaml:"port,omitempty"`
	} `yaml:"api,omitempty"`
	SQL struct {
		Host     string `yaml:"host,omitempty"`
		Port     string `yaml:"port,omitempty"`
		User     string `yaml:"username,omitempty"`
		Password string `yaml:"password,omitempty"`
	} `yaml:"sql,omitempty"`
}

// Yml é a variável global que armazena a configuração carregada.
var Yml Config

// LoadConfig carrega a configuração do arquivo config.yaml.
func LoadConfig() error {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "config.yaml"
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("erro ao abrir config: %w", err)
	}

	return yaml.Unmarshal(data, &Yml)
}

// CreateConfigFile cria o arquivo config.yaml se não existir.
func CreateConfigFile() {
	if _, err := os.Stat("config.yaml"); err == nil {
		fmt.Println("O arquivo 'config.yaml' já existe. Deseja sobrescrever? (y/N)")
		var rsp string
		fmt.Scan(&rsp)
		if strings.ToLower(rsp) == "y" {
			writeFile()
		}
		return
	}
	writeFile()
}

// writeFile escreve a configuração no arquivo config.yaml.
func writeFile() {
	data, err := yaml.Marshal(Yml)
	if err != nil {
		fmt.Printf("Erro ao gerar o YAML: %v\n", err)
		return
	}
	if err := os.WriteFile("config.yaml", data, 0644); err != nil {
		fmt.Printf("Erro ao escrever no arquivo config.yaml: %v\n", err)
	}
}
