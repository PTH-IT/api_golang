package usecase

import "gorm.io/gorm"

func NewInteractor(
	gormDb *gorm.DB,

) interactor {

	return interactor{
		gormDb,
	}
}

type interactor struct {
	gormDb *gorm.DB
}

func (i *interactor) Interactor() error {

	return nil
}
