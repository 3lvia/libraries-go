package mschema

import "testing"

func TestTypeName(t *testing.T) {
	type args struct {
		t Type
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"json", args{JSON}, "JSON"},
		{"protobuf", args{PROTOBUF}, "PROTOBUF"},
		{"avro", args{AVRO}, "AVRO"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TypeName(tt.args.t); got != tt.want {
				t.Errorf("TypeName() = %v, want %v", got, tt.want)
			}
		})
	}
}
