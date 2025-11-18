package config

import "github.com/joho/godotenv"

func Load(fileName string) error {
	err := godotenv.Load(fileName)
	if err != nil {
		return err
	}

	return nil
}
