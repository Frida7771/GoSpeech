package convertutil

import (
	"errors"
	"reflect"
)

var (
	// ErrSrcDstCannotBeNil 当 src 或 dst 为 nil 时返回
	ErrSrcDstCannotBeNil = errors.New("src and dst cannot be nil")
	// ErrDstMustBePointerStruct 当 dst 不是指向结构体的指针时返回
	ErrDstMustBePointerStruct = errors.New("dst must be a pointer to struct")
)

// CopyProperties 复制结构体的属性，注意：dst必须是指针类型，且指向的结构体类型与src类型相同
//
// # Params:
//
//	src: 源对象
//	dst: 目标对象
//
// # Examples:
//
//	type src struct {
//		Name string
//		Map  map[string]int
//	}
//	type dst struct {
//		Name string
//		Map  map[string]int
//		Age  int
//	}
//	s1 := src{Name: "test", Map: map[string]int{"a": 1}}
//	s2 := new(dst)
//	CopyProperties(s1, s2)
func CopyProperties(src, dst any) error {
	if src == nil || dst == nil {
		return ErrSrcDstCannotBeNil
	}

	srcValue := reflect.ValueOf(src)
	if srcValue.Kind() == reflect.Ptr {
		srcValue = srcValue.Elem()
	}
	dstValue := reflect.ValueOf(dst)

	if dstValue.Kind() != reflect.Ptr || dstValue.Elem().Kind() != reflect.Struct {
		return ErrDstMustBePointerStruct
	}

	// 使用 Elem 解引用指针（获取指针指向的实际结构体值）
	dstValue = dstValue.Elem()

	srcType := srcValue.Type()
	for i := 0; i < srcType.NumField(); i++ {
		srcField := srcType.Field(i)
		srcFieldValue := srcValue.Field(i)

		// 跳过不可导出的字段（即首字母小写）
		if !srcFieldValue.CanInterface() {
			continue
		}

		dstField := dstValue.FieldByName(srcField.Name)
		if dstField.IsValid() && dstField.CanSet() && dstField.Type() == srcFieldValue.Type() {
			dstField.Set(srcFieldValue)
		}
	}

	return nil
}
