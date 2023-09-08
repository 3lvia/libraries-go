package schemaclient

//type SchemaType string
//
//const (
//	Protobuf SchemaType = "PROTOBUF"
//	Avro     SchemaType = "AVRO"
//	Json     SchemaType = "JSON"
//)
//
//type SchemaRegistry interface {
//	Get(topic string) (*Schema, error)
//}
//
//type Schema struct {
//	ID      int    `json:"id"`
//	Subject string `json:"subject"`
//	Version int    `json:"version"`
//	Schema  string `json:"schema"`
//}
//
////
////// Schema is a data structure that holds all
////// the relevant information about schemas.
////type Schema struct {
////	id         int
////	schema     string
////	schemaType *SchemaType
////	version    int
////	references []Reference
////	codec      *goavro.Codec
////	jsonSchema *jsonschema.Schema
////}
////
////// Reference references use the import statement of Protobuf and
////// the $ref field of JSON Schema. They are defined by the name
////// of the import or $ref and the associated subject in the registry.
////type Reference struct {
////	Name    string `json:"name"`
////	Subject string `json:"subject"`
////	Version int    `json:"version"`
////}
