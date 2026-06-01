package business

import (
	"project02/model"
)

func LoadPreyRecords(filePath string) ([]model.PreyRecord, error) {
	return nil, nil
}

func SavePreyRecords(filePath string, records []model.PreyRecord) error {
	return nil
}

func DisplayPreyRecords(index int, records []model.PreyRecord) {
}

func CreatePreyRecord() model.PreyRecord {
	return model.PreyRecord{}
}

func EditPreyRecord(record *model.PreyRecord) {
}

func DeletePreyRecord(records *[]model.PreyRecord, index int) {
}
