// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     100423_3.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type MessageMetadata struct {
	AsyncTimeoutAfter *UnionNullLong `json:"AsyncTimeoutAfter"`

	CurrentStatus *UnionNullMessageStatus `json:"CurrentStatus"`

	CurrentStatusTime *UnionNullLong `json:"CurrentStatusTime"`

	Destination *UnionNullString `json:"Destination"`

	DetailedErrorMessage *UnionNullString `json:"DetailedErrorMessage"`

	Id *UnionNullInt `json:"Id"`

	RequestHeader *UnionNullString `json:"RequestHeader"`

	RequestId *UnionNullString `json:"RequestId"`

	RequestPayload *UnionNullString `json:"RequestPayload"`

	SentTime *UnionNullLong `json:"SentTime"`
}

const MessageMetadataAvroCRC64Fingerprint = "Q.$cm\x96\x19`"

func NewMessageMetadata() MessageMetadata {
	r := MessageMetadata{}
	r.AsyncTimeoutAfter = nil
	r.CurrentStatus = nil
	r.CurrentStatusTime = nil
	r.Destination = nil
	r.DetailedErrorMessage = nil
	r.Id = nil
	r.RequestHeader = nil
	r.RequestId = nil
	r.RequestPayload = nil
	r.SentTime = nil
	return r
}

func DeserializeMessageMetadata(r io.Reader) (MessageMetadata, error) {
	t := NewMessageMetadata()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeMessageMetadataFromSchema(r io.Reader, schema string) (MessageMetadata, error) {
	t := NewMessageMetadata()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeMessageMetadata(r MessageMetadata, w io.Writer) error {
	var err error
	err = writeUnionNullLong(r.AsyncTimeoutAfter, w)
	if err != nil {
		return err
	}
	err = writeUnionNullMessageStatus(r.CurrentStatus, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.CurrentStatusTime, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Destination, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DetailedErrorMessage, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.Id, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.RequestHeader, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.RequestId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.RequestPayload, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.SentTime, w)
	if err != nil {
		return err
	}
	return err
}

func (r MessageMetadata) Serialize(w io.Writer) error {
	return writeMessageMetadata(r, w)
}

func (r MessageMetadata) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"AsyncTimeoutAfter\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"CurrentStatus\",\"type\":[\"null\",{\"name\":\"MessageStatus\",\"namespace\":\"AdmsExtensionsLibMessagetracker.Enums\",\"symbols\":[\"CouldNotSendToSaf\",\"SentToSaf\",\"ReceivedOkFromSaf\",\"ReceivedErrorFromSaf\"],\"type\":\"enum\"}]},{\"default\":null,\"name\":\"CurrentStatusTime\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"Destination\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DetailedErrorMessage\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"RequestHeader\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"RequestId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"RequestPayload\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SentTime\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"AdmsExtensionsLibMessagetracker.Models.MessageMetadata\",\"type\":\"record\"}"
}

func (r MessageMetadata) SchemaName() string {
	return "AdmsExtensionsLibMessagetracker.Models.MessageMetadata"
}

func (_ MessageMetadata) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ MessageMetadata) SetInt(v int32)       { panic("Unsupported operation") }
func (_ MessageMetadata) SetLong(v int64)      { panic("Unsupported operation") }
func (_ MessageMetadata) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ MessageMetadata) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ MessageMetadata) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ MessageMetadata) SetString(v string)   { panic("Unsupported operation") }
func (_ MessageMetadata) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *MessageMetadata) Get(i int) types.Field {
	switch i {
	case 0:
		r.AsyncTimeoutAfter = NewUnionNullLong()

		return r.AsyncTimeoutAfter
	case 1:
		r.CurrentStatus = NewUnionNullMessageStatus()

		return r.CurrentStatus
	case 2:
		r.CurrentStatusTime = NewUnionNullLong()

		return r.CurrentStatusTime
	case 3:
		r.Destination = NewUnionNullString()

		return r.Destination
	case 4:
		r.DetailedErrorMessage = NewUnionNullString()

		return r.DetailedErrorMessage
	case 5:
		r.Id = NewUnionNullInt()

		return r.Id
	case 6:
		r.RequestHeader = NewUnionNullString()

		return r.RequestHeader
	case 7:
		r.RequestId = NewUnionNullString()

		return r.RequestId
	case 8:
		r.RequestPayload = NewUnionNullString()

		return r.RequestPayload
	case 9:
		r.SentTime = NewUnionNullLong()

		return r.SentTime
	}
	panic("Unknown field index")
}

func (r *MessageMetadata) SetDefault(i int) {
	switch i {
	case 0:
		r.AsyncTimeoutAfter = nil
		return
	case 1:
		r.CurrentStatus = nil
		return
	case 2:
		r.CurrentStatusTime = nil
		return
	case 3:
		r.Destination = nil
		return
	case 4:
		r.DetailedErrorMessage = nil
		return
	case 5:
		r.Id = nil
		return
	case 6:
		r.RequestHeader = nil
		return
	case 7:
		r.RequestId = nil
		return
	case 8:
		r.RequestPayload = nil
		return
	case 9:
		r.SentTime = nil
		return
	}
	panic("Unknown field index")
}

func (r *MessageMetadata) NullField(i int) {
	switch i {
	case 0:
		r.AsyncTimeoutAfter = nil
		return
	case 1:
		r.CurrentStatus = nil
		return
	case 2:
		r.CurrentStatusTime = nil
		return
	case 3:
		r.Destination = nil
		return
	case 4:
		r.DetailedErrorMessage = nil
		return
	case 5:
		r.Id = nil
		return
	case 6:
		r.RequestHeader = nil
		return
	case 7:
		r.RequestId = nil
		return
	case 8:
		r.RequestPayload = nil
		return
	case 9:
		r.SentTime = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ MessageMetadata) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ MessageMetadata) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ MessageMetadata) HintSize(int)                     { panic("Unsupported operation") }
func (_ MessageMetadata) Finalize()                        {}

func (_ MessageMetadata) AvroCRC64Fingerprint() []byte {
	return []byte(MessageMetadataAvroCRC64Fingerprint)
}

func (r MessageMetadata) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["AsyncTimeoutAfter"], err = json.Marshal(r.AsyncTimeoutAfter)
	if err != nil {
		return nil, err
	}
	output["CurrentStatus"], err = json.Marshal(r.CurrentStatus)
	if err != nil {
		return nil, err
	}
	output["CurrentStatusTime"], err = json.Marshal(r.CurrentStatusTime)
	if err != nil {
		return nil, err
	}
	output["Destination"], err = json.Marshal(r.Destination)
	if err != nil {
		return nil, err
	}
	output["DetailedErrorMessage"], err = json.Marshal(r.DetailedErrorMessage)
	if err != nil {
		return nil, err
	}
	output["Id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["RequestHeader"], err = json.Marshal(r.RequestHeader)
	if err != nil {
		return nil, err
	}
	output["RequestId"], err = json.Marshal(r.RequestId)
	if err != nil {
		return nil, err
	}
	output["RequestPayload"], err = json.Marshal(r.RequestPayload)
	if err != nil {
		return nil, err
	}
	output["SentTime"], err = json.Marshal(r.SentTime)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *MessageMetadata) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["AsyncTimeoutAfter"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AsyncTimeoutAfter); err != nil {
			return err
		}
	} else {
		r.AsyncTimeoutAfter = NewUnionNullLong()

		r.AsyncTimeoutAfter = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["CurrentStatus"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CurrentStatus); err != nil {
			return err
		}
	} else {
		r.CurrentStatus = NewUnionNullMessageStatus()

		r.CurrentStatus = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["CurrentStatusTime"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CurrentStatusTime); err != nil {
			return err
		}
	} else {
		r.CurrentStatusTime = NewUnionNullLong()

		r.CurrentStatusTime = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Destination"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Destination); err != nil {
			return err
		}
	} else {
		r.Destination = NewUnionNullString()

		r.Destination = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["DetailedErrorMessage"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DetailedErrorMessage); err != nil {
			return err
		}
	} else {
		r.DetailedErrorMessage = NewUnionNullString()

		r.DetailedErrorMessage = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		r.Id = NewUnionNullInt()

		r.Id = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["RequestHeader"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RequestHeader); err != nil {
			return err
		}
	} else {
		r.RequestHeader = NewUnionNullString()

		r.RequestHeader = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["RequestId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RequestId); err != nil {
			return err
		}
	} else {
		r.RequestId = NewUnionNullString()

		r.RequestId = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["RequestPayload"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RequestPayload); err != nil {
			return err
		}
	} else {
		r.RequestPayload = NewUnionNullString()

		r.RequestPayload = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["SentTime"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SentTime); err != nil {
			return err
		}
	} else {
		r.SentTime = NewUnionNullLong()

		r.SentTime = nil
	}
	return nil
}
