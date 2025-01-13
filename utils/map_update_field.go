package utils

import "reflect"

func MapUpdateField(existingProduct interface{}, updateData interface{}) {
	existingVal := reflect.ValueOf(existingProduct).Elem()
	updateVal := reflect.ValueOf(updateData).Elem()

	for i := 0; i < updateVal.NumField(); i++ {
		field := updateVal.Field(i)
		fieldName := updateVal.Type().Field(i).Name

		if !field.IsNil() {
			existingField := existingVal.FieldByName(fieldName)
			if existingVal.IsValid() && existingVal.CanSet() {
				existingField.Set(field.Elem())
			}
		}
	}
}
