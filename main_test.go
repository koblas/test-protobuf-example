package main

import (
	"bytes"
	_ "embed"
	"testing"
)

var result error

//go:embed events.json
var jsonData []byte

//go:embed events.protobuf
var protoData []byte

func TestMarshalJson(t *testing.T) {
	buf, err := MarshalJson()
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(buf, jsonData) {
		t.Errorf("arrays differ")
	}
}

func TestMarshalProto(t *testing.T) {
	buf, err := MarshalProto()
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(buf, protoData) {
		t.Errorf("arrays differ")
	}
}

func BenchmarkMarshalJson(b *testing.B) {
	var r error
	for i := 0; i < b.N; i++ {
		_, r = MarshalJson()
	}
	result = r
}

func BenchmarkMarshalProtobuf(b *testing.B) {
	var r error
	for i := 0; i < b.N; i++ {
		_, r = MarshalProto()
	}
	result = r
}

func BenchmarkUnmarshalJson(b *testing.B) {
	var r error
	for i := 0; i < b.N; i++ {
		r = UnmarshalJson(jsonData)
	}
	result = r
}

func BenchmarkUnmarshalProtobuf(b *testing.B) {
	var r error
	for i := 0; i < b.N; i++ {
		r = UnmarshalProto(protoData)
	}
	result = r
}
