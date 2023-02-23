package main

import (
	"encoding/json"
	"os"

	"github.com/kfelter/protobuf-example/publish"
	"google.golang.org/protobuf/proto"
)

var event = publish.EventList{
	Events: []*publish.Event{
		{
			Content: []byte(`some event content!`),
			Tags:    []string{"tag1", "tag2"},
		},
		{
			Content: []byte(`some event content2!`),
			Tags:    []string{"tag3", "tag4"},
		},
	},
}

func main() {
	buf, _ := proto.Marshal(&event)
	os.WriteFile("events.protobuf", buf, os.ModePerm)

	buf, _ = json.Marshal(&event)
	os.WriteFile("events.json", buf, os.ModePerm)
}

func MarshalJson() ([]byte, error) {
	return json.Marshal(&event)
}

func MarshalProto() ([]byte, error) {
	return proto.Marshal(&event)
}

func UnmarshalJson(buf []byte) error {
	events := publish.EventList{}

	return json.Unmarshal(buf, &events)
}

func UnmarshalProto(buf []byte) error {
	events := publish.EventList{}

	return proto.Unmarshal(buf, &events)
}
