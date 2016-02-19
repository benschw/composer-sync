package config

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strings"

	"github.com/benschw/opin-go/config"
	"golang.org/x/crypto/ssh/terminal"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Stash StashConfig `yaml:"stash"`
	Satis SatisConfig `yaml:"satis"`
}
type StashConfig struct {
	RepoTpl     string `yaml:"repo_tpl"`
	ApiUrl      string `yaml:"api_url"`
	AuthEncoded string `yaml:"auth"`
	Login       string `-`
	Password    string `-`
	ProjKey     string `yaml:"proj_key"`
}
type SatisConfig struct {
	ApiUrl string `yaml:"api_url"`
}

func normalizePath(cfgPath string) (string, error) {
	if cfgPath == "~/.composer-sync.yaml" {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}
		cfgPath = (usr.HomeDir + "/.composer-sync.yaml")
	}
	return cfgPath, nil
}

func Load(cfgPath string) (*Config, error) {
	cfgPath, err := normalizePath(cfgPath)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := config.Bind(cfgPath, cfg); err != nil {
		return nil, err
	}
	b, err := base64.StdEncoding.DecodeString(cfg.Stash.AuthEncoded)
	if err != nil {
		return nil, err
	}
	parts := strings.Split(string(b[:]), ":")
	if len(parts) != 2 {
		return cfg, nil
	}
	cfg.Stash.Login = parts[0]
	cfg.Stash.Password = parts[1]

	return cfg, err
}

func Login(cfg *Config, cfgPath string, loginCmd bool) error {
	fmt.Println("Authenticate With Stash:")

	if err := loadCreds(cfg); err != nil {
		return err
	}

	if loginCmd || promptToSave(cfgPath) {
		if err := save(cfg, cfgPath); err != nil {
			return err
		}
		fmt.Printf("Auth Updated in '%s'\n", cfgPath)
	}
	fmt.Println()
	return nil
}

func loadCreds(cfg *Config) error {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print(" Login: ")
	login, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	cfg.Stash.Login = strings.TrimSpace(login)

	fmt.Print(" Password: ")
	pass, err := terminal.ReadPassword(0)
	fmt.Println("")
	if err != nil {
		return err
	}
	cfg.Stash.Password = string(pass[:])
	return nil
}
func promptToSave(cfgPath string) bool {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Save to %s?\n y/n: ", cfgPath)
	str, err := reader.ReadString('\n')
	if err != nil {
		return false
	}
	resp := strings.TrimSpace(str)
	return resp == "y"
}

func save(cfg *Config, cfgPath string) error {
	cfgPath, err := normalizePath(cfgPath)
	if err != nil {
		return err
	}

	str := cfg.Stash.Login + ":" + cfg.Stash.Password
	strEnc := base64.StdEncoding.EncodeToString([]byte(str))
	cfg.Stash.AuthEncoded = strEnc

	b, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(cfgPath, b, 0400)
}
