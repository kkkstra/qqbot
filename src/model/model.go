package model

import (
	"reflect"
)

func GetModel(Model interface{}) interface{} {
	t := reflect.TypeOf(Model)
	switch t.String() {
	case "*model.StickerModel":
		return &StickerModel{db}
	}
	return nil
}
