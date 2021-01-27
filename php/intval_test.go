package php

import (
	"encoding/json"
	"testing"
)

func TestIntval(t *testing.T) {
	var testValueNil interface{}
	var testValueInt int
	var testValueInt16 int16
	var testValueInt32 int32
	var testValueInt64 int64
	var testValueFloat32 float32
	var testValueFloat64 float64
	var testValueUint uint
	var testValueUint16 uint16
	var testValueUint32 uint32
	var testValueUint64 uint64
	var testValueJsonNumber json.Number
	var testValueString string
	var testBool bool
	//var testComplex64 complex64
	//testComplex64 = complex(1,1)
	var testByte byte
	var testUintptr uintptr

	testValueNil = nil
	testValueInt = 1
	testValueInt16 = 1
	testValueInt32 = 1
	testValueInt64 = 1
	testValueFloat32 = 1
	testValueFloat64 = 1
	testValueUint = 1
	testValueUint16 = 1
	testValueUint32 = 1
	testValueUint64 = 1
	testValueJsonNumber = "1"
	testValueString = "1"
	testBool = true
	testByte = 1
	testUintptr = 1
	if _val, _ := Int64val(testValueNil); _val != int64(0) {
		t.Fatal("Nil to int64 fail")
	}
	if _val, _ := Int64val(testValueInt); _val != int64(1) {
		t.Fatal("Int to int64 fail")
	}
	if _val, _ := Int64val(testValueInt16); _val != int64(1) {
		t.Fatal("Int16 to int64 fail")
	}
	if _val, _ := Int64val(testValueInt32); _val != int64(1) {
		t.Fatal("Int32 to int64 fail")
	}
	if _val, _ := Int64val(testValueInt64); _val != int64(1) {
		t.Fatal("Int64 to int64 fail")
	}
	if _val, _ := Int64val(testValueFloat32); _val != int64(1) {
		t.Fatal("Float32 to int64 fail")
	}
	if _val, _ := Int64val(testValueFloat64); _val != int64(1) {
		t.Fatal("int64 to int64 fail")
	}
	if _val, _ := Int64val(testValueUint); _val != int64(1) {
		t.Fatal("Uint to int64 fail")
	}
	if _val, _ := Int64val(testValueUint16); _val != int64(1) {
		t.Fatal("Uint16 to int64 fail")
	}
	if _val, _ := Int64val(testValueUint32); _val != int64(1) {
		t.Fatal("Uint32 to int64 fail")
	}
	if _val, _ := Int64val(testValueUint64); _val != int64(1) {
		t.Fatal("Uint64 to int64 fail")
	}
	if _val, _ := Int64val(testValueJsonNumber); _val != int64(1) {
		t.Fatal("JsonNumber to int64 fail")
	}
	if _val, _ := Int64val(testValueString); _val != int64(1) {
		t.Fatal("String to int64 fail")
	}
	if _val, _ := Int64val(testBool); _val != int64(1) {
		t.Fatal("Bool to int64 fail")
	}
	if _val, _ := Int64val(testByte); _val != int64(1) {
		t.Fatal("Byte to int64 fail")
	}
	if _val, _ := Int64val(testUintptr); _val != int64(1) {
		t.Fatal("Uintptr to int64 fail")
	}




	if _val, _ := Int64val(testValueNil); _val != int64(0) {
		t.Fatal("Nil to int64 fail")
	}
	if _val, _ := Int64val(testValueInt); _val != int64(1) {
		t.Fatal("Int to int64 fail")
	}
	if _val, _ := Int64val(testValueInt16); _val != int64(1) {
		t.Fatal("Int16 to int64 fail")
	}
	if _val, _ := Int64val(testValueInt32); _val != int64(1) {
		t.Fatal("Int32 to int64 fail")
	}
	if _val, _ := Int64val(testValueInt64); _val != int64(1) {
		t.Fatal("Int64 to int64 fail")
	}
	if _val, _ := Int64val(testValueFloat32); _val != int64(1) {
		t.Fatal("Float32 to int64 fail")
	}
	if _val, _ := Int64val(testValueFloat64); _val != int64(1) {
		t.Fatal("int64 to int64 fail")
	}
	if _val, _ := Int64val(testValueUint); _val != int64(1) {
		t.Fatal("Uint to int64 fail")
	}
	if _val, _ := Int64val(testValueUint16); _val != int64(1) {
		t.Fatal("Uint16 to int64 fail")
	}
	if _val, _ := Int64val(testValueUint32); _val != int64(1) {
		t.Fatal("Uint32 to int64 fail")
	}
	if _val, _ := Int64val(testValueUint64); _val != int64(1) {
		t.Fatal("Uint64 to int64 fail")
	}
	if _val, _ := Int64val(testValueJsonNumber); _val != int64(1) {
		t.Fatal("JsonNumber to int64 fail")
	}
	if _val, _ := Int64val(testValueString); _val != int64(1) {
		t.Fatal("String to int64 fail")
	}
	if _val, _ := Int64val(testBool); _val != int64(1) {
		t.Fatal("Bool to int64 fail")
	}
	if _val, _ := Int64val(testByte); _val != int64(1) {
		t.Fatal("Byte to int64 fail")
	}
	if _val, _ := Int64val(testUintptr); _val != int64(1) {
		t.Fatal("Uintptr to int64 fail")
	}


	if _val, _ := Intval(testValueNil); _val != int32(0) {
		t.Fatal("Nil to int32 fail")
	}
	if _val, _ := Intval(testValueInt); _val != int32(1) {
		t.Fatal("Int to int32 fail")
	}
	if _val, _ := Intval(testValueInt16); _val != int32(1) {
		t.Fatal("Int16 to int32 fail")
	}
	if _val, _ := Intval(testValueInt32); _val != int32(1) {
		t.Fatal("Int32 to int32 fail")
	}
	if _val, _ := Intval(testValueInt64); _val != int32(1) {
		t.Fatal("Int64 to int32 fail")
	}
	if _val, _ := Intval(testValueFloat32); _val != int32(1) {
		t.Fatal("Float32 to int32 fail")
	}
	if _val, _ := Intval(testValueFloat64); _val != int32(1) {
		t.Fatal("int64 to int32 fail")
	}
	if _val, _ := Intval(testValueUint); _val != int32(1) {
		t.Fatal("Uint to int32 fail")
	}
	if _val, _ := Intval(testValueUint16); _val != int32(1) {
		t.Fatal("Uint16 to int32 fail")
	}
	if _val, _ := Intval(testValueUint32); _val != int32(1) {
		t.Fatal("Uint32 to int32 fail")
	}
	if _val, _ := Intval(testValueUint64); _val != int32(1) {
		t.Fatal("Uint64 to int32 fail")
	}
	if _val, _ := Intval(testValueJsonNumber); _val != int32(1) {
		t.Fatal("JsonNumber to int32 fail")
	}
	if _val, _ := Intval(testValueString); _val != int32(1) {
		t.Fatal("String to int32 fail")
	}
	if _val, _ := Intval(testBool); _val != int32(1) {
		t.Fatal("Bool to int32 fail")
	}
	if _val, _ := Intval(testByte); _val != int32(1) {
		t.Fatal("Byte to int32 fail")
	}
	if _val, _ := Intval(testUintptr); _val != int32(1) {
		t.Fatal("Uintptr to int32 fail")
	}
}
