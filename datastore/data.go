package datastore

import "github.com/SubrotoRoy/wrestlers/model"

var wrestlers []model.Wrestler

func GetAllWrestlers() []model.Wrestler {
	return wrestlers
}

func AddWrestlers(wrestler model.Wrestler) {
	wrestlers = append(wrestlers, wrestler)
}
