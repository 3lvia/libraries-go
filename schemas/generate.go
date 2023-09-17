package schemas

//go:generate mkdir -p ./avro
//go:generate $GOPATH/bin/gogen-avro -containers ./avro ./avsc/100044_dp_examples-value.avsc
