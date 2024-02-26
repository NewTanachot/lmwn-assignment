package irepository

import "lmwn-assignment/domain/entities"

type ICovidRepository interface {
	GetCovidCaseRecord() (*[]entities.Covid, error)
}
