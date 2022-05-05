package types

import (
	"reflect"
	"strconv"
)

func Int64ToString(num int64) string  {
	return strconv.FormatInt(num, 10)
}

func UInt64ToString(num uint64) string  {
	return strconv.FormatUint(num, 10)
}

func gStringToInt(str string) (int, error)  {
	return strconv.Atoi(str)
}

// Fill 通过反射将对象2的值填充给对象1
func Fill(obj1, obj2 interface{})  {
	v1 := reflect.ValueOf(obj1).Elem()
	v2 := reflect.ValueOf(obj2).Elem()

	for i := 0; i < v1.NumField(); i++  {
		fieldInfo1 := v1.Type().Field(i)
		field1Name := fieldInfo1.Name
		field1Type := fieldInfo1.Type

		for j := 0; j < v2.NumField(); j++  {
			fieldInfo2 := v2.Type().Field(j)
			if field1Name == fieldInfo2.Name && field1Type == fieldInfo2.Type {
				if v2.FieldByName(fieldInfo2.Name).IsValid() {
					switch v2.FieldByName(fieldInfo2.Name).Type().String() {
					case "int":
						if v2.FieldByName(fieldInfo2.Name).Int() == 0 {
							continue
						}
					case "string":
						if v2.FieldByName(fieldInfo2.Name).String() == "" {
							continue
						}
					}
				}

				newValue := v2.FieldByName(fieldInfo2.Name)
				if newValue.IsValid() {
					v1.FieldByName(fieldInfo1.Name).Set(newValue)
				}
			}
		}
	}
}

