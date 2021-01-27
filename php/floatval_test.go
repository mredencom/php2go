package php

import (
	"encoding/json"
	"testing"
)

func TestFloatval(t *testing.T) {

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
	if _val, _ := Floatval(testValueNil); _val != float64(0) {
		t.Fatal("Nil to float64 fail")
	}
	if _val, _ := Floatval(testValueInt); _val != float64(1) {
		t.Fatal("Int to float64 fail")
	}
	if _val, _ := Floatval(testValueInt16); _val != float64(1) {
		t.Fatal("Int16 to float64 fail")
	}
	if _val, _ := Floatval(testValueInt32); _val != float64(1) {
		t.Fatal("Int32 to float64 fail")
	}
	if _val, _ := Floatval(testValueInt64); _val != float64(1) {
		t.Fatal("Int64 to float64 fail")
	}
	if _val, _ := Floatval(testValueFloat32); _val != float64(1) {
		t.Fatal("Float32 to float64 fail")
	}
	if _val, _ := Floatval(testValueFloat64); _val != float64(1) {
		t.Fatal("Float64 to float64 fail")
	}
	if _val, _ := Floatval(testValueUint); _val != float64(1) {
		t.Fatal("Uint to float64 fail")
	}
	if _val, _ := Floatval(testValueUint16); _val != float64(1) {
		t.Fatal("Uint16 to float64 fail")
	}
	if _val, _ := Floatval(testValueUint32); _val != float64(1) {
		t.Fatal("Uint32 to float64 fail")
	}
	if _val, _ := Floatval(testValueUint64); _val != float64(1) {
		t.Fatal("Uint64 to float64 fail")
	}
	if _val, _ := Floatval(testValueJsonNumber); _val != float64(1) {
		t.Fatal("JsonNumber to float64 fail")
	}
	if _val, _ := Floatval(testValueString); _val != float64(1) {
		t.Fatal("String to float64 fail")
	}
	if _val, _ := Floatval(testBool); _val != float64(1) {
		t.Fatal("Bool to float64 fail")
	}
	if _val, _ := Floatval(testByte); _val != float64(1) {
		t.Fatal("Byte to float64 fail")
	}
	if _val, _ := Floatval(testUintptr); _val != float64(1) {
		t.Fatal("Uintptr to float64 fail")
	}
}
