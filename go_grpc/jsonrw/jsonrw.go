package jsonrw

import (
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

//ToJSON converts protobuf Messages to JSON strings
func ToJSON(pb proto.Message) []byte {
	jsonString, err := protojson.Marshal(pb)
	if err != nil {
		log.Fatalln("Failed to Marshal to JSON")
		return nil
	}
	return jsonString

}

//FromJSON converts JSON strings to protobuf Messages
func FromJSON(jsonString []byte, pb proto.Message) {
	err := protojson.Unmarshal(jsonString, pb)
	if err != nil {
		log.Fatalln("Failed to UnMarshal fron JSON")
	}
}
