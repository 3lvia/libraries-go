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

func Test_descriptor_GenerationFolder(t *testing.T) {
	type fields struct {
		Subj      string
		V         int
		I         int
		S         string
		ErrorCode int
		Message   string
		T         string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{"private dp", fields{Subj: "private.dp.edna.examples-value"}, "dp/edna/examples/private/", false},
		{"public kunde", fields{Subj: "public.kunde.kunde-extensions.dialogue-value"}, "kunde/kunde-extensions/dialogue/public/", false},
		{"malformed", fields{Subj: "oot_tmp1_delete_me"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := descriptor{
				Subj:      tt.fields.Subj,
				V:         tt.fields.V,
				I:         tt.fields.I,
				S:         tt.fields.S,
				ErrorCode: tt.fields.ErrorCode,
				Message:   tt.fields.Message,
				T:         tt.fields.T,
			}
			got, err := d.GenerationFolder()
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerationFolder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerationFolder() got = %v, want %v", got, tt.want)
			}
		})
	}
}
