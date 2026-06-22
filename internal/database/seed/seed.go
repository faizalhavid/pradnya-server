package seed

import "gorm.io/gorm"

type Seeder interface {
	Run(db *gorm.DB) error
}

func Run(db *gorm.DB) error {
	seeders := []Seeder{
		UserSeeder{},
	}

	for _, seeder := range seeders {
		if err := seeder.Run(db); err != nil {
			return err
		}
	}

	return nil
}
