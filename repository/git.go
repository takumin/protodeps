package repository

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

type Repository struct {
	repository *git.Repository
}

func NewRepository(dir string) (*Repository, error) {
	var repo *git.Repository

	if _, err := os.Stat(dir); err == nil {
		repo, err = git.PlainOpen(dir)
		if err != nil {
			return nil, err
		}
	} else {
		repo, err = git.PlainInit(dir, true)
		if err != nil {
			return nil, err
		}
	}

	return &Repository{
		repository: repo,
	}, nil
}

func (r *Repository) SetOrigin(url string) error {
	cfg, err := r.repository.Config()
	if err != nil {
		return err
	}
	cfg.Remotes["origin"] = &config.RemoteConfig{
		Name: "origin",
		URLs: []string{url},
	}
	if err := cfg.Validate(); err != nil {
		return err
	}
	if err := r.repository.SetConfig(cfg); err != nil {
		return err
	}
	return nil
}
