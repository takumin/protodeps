package ensure

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/urfave/cli/v2"

	"github.com/takumin/protodeps/config"
	"github.com/takumin/protodeps/repository"
)

func NewCommands(c *config.Config, f []cli.Flag) []*cli.Command {
	flags := []cli.Flag{}
	return []*cli.Command{
		{
			Name:   "ensure",
			Usage:  "ensure proto files",
			Flags:  append(flags, f...),
			Action: action(c),
		},
	}
}

func action(cfg *config.Config) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		if err := createCacheDir(cfg.CacheDir); err != nil {
			return err
		}

		for k, v := range cfg.Dependencies {
			log.Println("Cloning:", k, "(", v.Source, ")")

			repo, err := repository.NewRepository(path.Join(cfg.CacheDir, k))
			if err != nil {
				return err
			}

			if err := repo.SetOrigin(v.Source); err != nil {
				return err
			}
		}

		return nil
	}
}

func createCacheDir(d string) error {
	fi, err := os.Stat(d)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(d, 0700); err != nil {
			return err
		}
	} else {
		if !fi.IsDir() {
			return errors.New(fmt.Sprintf("Failed to exists path: %s\n", fi.Name()))
		}
	}
	return nil
}
