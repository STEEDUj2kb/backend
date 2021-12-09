package utils

import (
	"reflect"
)

// Reference:
//  - https://stackoverflow.com/questions/40091142/is-there-an-established-pattern-name-for-golang-code-which-seems-similar-to-a-mi
//  - https://gist.github.com/hasanozgan/2598e6abe80803e4573b
//  - https://stackoverflow.com/questions/60143465/how-to-change-value-of-empty-interface-that-pass-as-struct-reference-in-golang

// ApplyData func for get value from database data.
func ApplyData(toModel interface{}, fromModel interface{}) {
	toModelElements := reflect.ValueOf(toModel).Elem()
	fromModelElements := reflect.ValueOf(fromModel).Elem()
	for i := 0; i < fromModelElements.NumField(); i++ {
		fromModelType := fromModelElements.Type().Field(i)
		toModelElement := toModelElements.FieldByName(fromModelType.Name)
		if toModelElement.IsValid() {
			toModelElement.Set(fromModelElements.Field(i))
		}
	}
}
