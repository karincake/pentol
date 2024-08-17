package test

import (
	"reflect"
	"testing"

	"github.com/karincake/pentol"
)

func TestBool(t *testing.T) {
	got := pentol.Bool()
	v := reflect.ValueOf(got)
	e := v.Elem().Kind()

	if v.Kind() != reflect.Ptr || e != reflect.Bool {
		t.Error("failed to get pointer of bool")
	}
}

func TestString(t *testing.T) {
	got := pentol.String()
	v := reflect.ValueOf(got)
	e := v.Elem().Kind()

	if v.Kind() != reflect.Ptr || e != reflect.String {
		t.Error("failed to get pointer of string")
	}
}

func TestInt(t *testing.T) {
	got := pentol.Int()
	v := reflect.ValueOf(got)
	e := v.Elem().Kind()

	if v.Kind() != reflect.Ptr || e != reflect.Int {
		t.Error("failed to get pointer of int")
	}
}

func TestInt8(t *testing.T) {
	got := pentol.Int8()
	v := reflect.ValueOf(got)
	e := v.Elem().Kind()

	if v.Kind() != reflect.Ptr || e != reflect.Int8 {
		t.Error("failed to get pointer of int8")
	}
}

func TestInt16(t *testing.T) {
	got := pentol.Int16()
	v := reflect.ValueOf(got)
	e := v.Elem().Kind()

	if v.Kind() != reflect.Ptr || e != reflect.Int16 {
		t.Error("failed to get pointer of Int16")
	}
}

func TestInt32(t *testing.T) {
	got := pentol.Int32()
	v := reflect.ValueOf(got)
	e := v.Elem().Kind()

	if v.Kind() != reflect.Ptr || e != reflect.Int32 {
		t.Error("failed to get pointer of Int32")
	}
}

func TestInt64(t *testing.T) {
	got := pentol.Int64()
	v := reflect.ValueOf(got)
	e := v.Elem().Kind()

	if v.Kind() != reflect.Ptr || e != reflect.Int64 {
		t.Error("failed to get pointer of int64")
	}
}

func TestUInt(t *testing.T) {
	got := pentol.Uint()
	v := reflect.ValueOf(got)
	e := v.Elem().Kind()

	if v.Kind() != reflect.Ptr || e != reflect.Uint {
		t.Error("failed to get pouinter of uint")
	}
}

func TestUInt8(t *testing.T) {
	got := pentol.Uint8()
	v := reflect.ValueOf(got)
	e := v.Elem().Kind()

	if v.Kind() != reflect.Ptr || e != reflect.Uint8 {
		t.Error("failed to get pouinter of uint8")
	}
}

func TestUInt16(t *testing.T) {
	got := pentol.Uint16()
	v := reflect.ValueOf(got)
	e := v.Elem().Kind()

	if v.Kind() != reflect.Ptr || e != reflect.Uint16 {
		t.Error("failed to get pouinter of uint16")
	}
}

func TestUInt32(t *testing.T) {
	got := pentol.Uint32()
	v := reflect.ValueOf(got)
	e := v.Elem().Kind()

	if v.Kind() != reflect.Ptr || e != reflect.Uint32 {
		t.Error("failed to get pouinter of uint32")
	}
}

func TestUInt64(t *testing.T) {
	got := pentol.Uint64()
	v := reflect.ValueOf(got)
	e := v.Elem().Kind()

	if v.Kind() != reflect.Ptr || e != reflect.Uint64 {
		t.Error("failed to get pouinter of uint64")
	}
}

func TestFloat32(t *testing.T) {
	got := pentol.Float32()
	v := reflect.ValueOf(got)
	e := v.Elem().Kind()

	if v.Kind() != reflect.Ptr || e != reflect.Float32 {
		t.Error("failed to get pouinter of float32")
	}
}

func TestFloat64(t *testing.T) {
	got := pentol.Float64()
	v := reflect.ValueOf(got)
	e := v.Elem().Kind()

	if v.Kind() != reflect.Ptr || e != reflect.Float64 {
		t.Error("failed to get pouinter of float64")
	}
}
