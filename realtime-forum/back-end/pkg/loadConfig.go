package pkg

import "reflect"

func LoadDataToModelConfig(model interface{}) {
	reflect.ValueOf(model)
}
