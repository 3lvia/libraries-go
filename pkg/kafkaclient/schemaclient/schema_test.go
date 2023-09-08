package schemaclient

import (
	"github.com/linkedin/goavro/v2"
	"testing"
)

func TestAvro(t *testing.T) {
	codec, err := goavro.NewCodec(avroSchemaEdnaExamples)
	if err != nil {
		t.Fatal(err)
	}
	_ = codec
}

const (
	avroSchemaEdnaExamples = `{"type":"record","name":"Person","namespace":"dp.demoapp","fields":[{"name":"Cars","type":["null",{"type":"array","items":["null",{"type":"record","name":"Car","fields":[{"name":"Color","type":["null","string"],"default":null},{"name":"Gearbox","type":["null",{"type":"enum","name":"Gearbox","symbols":["Automatic","Manual"]}],"default":null},{"name":"model","type":["null","string"],"default":null},{"name":"Person","type":["null","Person"],"default":null}]}]}],"default":null},{"name":"Houses","type":["null",{"type":"array","items":["null",{"type":"record","name":"House","fields":[{"name":"Address","type":["null","string"],"default":null},{"name":"Buildingtype","type":{"type":"enum","name":"Buildingtype","symbols":["House","Apartment","Cabin","Other"]}},{"name":"Color","type":["null","string"],"default":null}]}]}],"default":null},{"name":"Id","type":["null","string"],"default":null},{"name":"Name","type":["null","string"],"default":null},{"name":"Updated","type":["null","string"],"default":null}]}`
)