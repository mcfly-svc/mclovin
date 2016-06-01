package main

import (
	"os"
	"os/user"
	"path"

	"github.com/BurntSushi/toml"
)

type UserCreds struct {
	Token string `json:"token" toml:"token"`
	//Environment string `json:"env" toml:"env"`
	//Name        string `json:"name" toml:"name"`
	//Email       string `json:"email" toml:"email"`
}

type CredentialStore interface {
	GetUserCreds() (*UserCreds, error)
	SaveUserCreds(user *UserCreds) error
	ClearUserCreds() error
}

type SimpleCredentialStore struct {
	configPath string
}

type mcflyConfigFile struct {
	UserCreds *UserCreds `toml:"user"`
}

func NewSimpleCredentialStore() (CredentialStore, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	dirPath := path.Join(usr.HomeDir, ".config")
	err = os.MkdirAll(dirPath, os.FileMode(0755))
	if err != nil {
		return nil, err
	}

	configPath := path.Join(dirPath, "mcfly.toml")

	s := &SimpleCredentialStore{
		configPath: configPath,
	}

	return s, nil
}

func (s *SimpleCredentialStore) GetUserCreds() (*UserCreds, error) {
	var f mcflyConfigFile

	if _, err := toml.DecodeFile(s.configPath, &f); err != nil {
		if os.IsNotExist(err) {
			// simply not signed in
			return nil, nil
		} else {
			return nil, err
		}
	}

	return f.UserCreds, nil
}

func (s *SimpleCredentialStore) SaveUserCreds(u *UserCreds) error {
	cf := mcflyConfigFile{
		UserCreds: u,
	}

	f, err := os.OpenFile(s.configPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(0600))
	if err != nil {
		return err
	}

	defer f.Close()

	e := toml.NewEncoder(f)
	err = e.Encode(&cf)
	if err != nil {
		return err
	}

	return nil
}

func (s *SimpleCredentialStore) ClearUserCreds() error {
	// for now just wipe the configuration file, since it only contains
	// user information.
	if err := os.Remove(s.configPath); err != nil && !os.IsNotExist(err) {
		return err
	}

	return nil
}
