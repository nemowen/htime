package models

import (
	"testing"
)

func init() {
	InitDatabase()
}

func TestInsertTopic(t *testing.T) {
	topic := new(Topic)
	topic.Title = "123123"
	topic.Text = "hello..llasdfas"
	topic.Flags = "test"
	err := CreateTopic(topic)
	if err != nil {
		t.Error("insert topic error")
	}
	topics, err := GetTopics(0, 20)
	if err != nil {
		t.Error("GetTopics topics error")
	}
	if len(topics) < 1 {
		t.Fatal("GetTopics return error want 1")
	}
}
