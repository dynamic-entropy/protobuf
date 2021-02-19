package filerw

import (
	"io/ioutil"
	"log"

	simplepb "go_grpc/source/simple"

	"google.golang.org/protobuf/proto"
)

//ReadFromFile Reads data from a binary file
func ReadFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Problem occured while writing to file", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Problem occured while marshaling", err2)
		return err2
	}

	return nil

}

//WriteToFile Writes data to a binary file
func WriteToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Problem occured while marshaling", err)
		return err
	}
	err2 := ioutil.WriteFile(fname, out, 0644)
	if err2 != nil {
		log.Fatalln("Problem occured while writing to file", err2)
		return err2
	}

	return nil

}

//CreateSimple creates an instance of SimpleMessage
func CreateSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Name",
		SampleList: []int32{1, 3, 5, 7, 9},
	}
	return &sm
}
