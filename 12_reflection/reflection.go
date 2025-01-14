package reflection

import "reflect"

// func walk(x interface{}, fn func(input string)) {
// 	val := getValue(x)
//
// 	numberOfValues := 0
// 	var getField func(int) reflect.Value
//
// 	switch val.Kind() {
// 	case reflect.Struct:
// 		numberOfValues = val.NumField()
// 		getField = val.Field
// 	case reflect.Slice, reflect.Array:
// 		numberOfValues = val.Len()
// 		getField = val.Index
// 	case reflect.String:
// 		fn(val.String())
// 	}
//
// 	for i := 0; i < numberOfValues; i++ {
// 		walk(getField(i).Interface(), fn)
// 	}
// }

func walk(x any, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}
}

func getValue(x any) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
