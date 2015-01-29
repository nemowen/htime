package models

import (
	"testing"
)

func TestInsertTopic(t *testing.T) {
	topic := new(Topic)
	topic.Title = "123123"
	topic.Text = "hello..llasdfas"
	topic.Flags = "test"
	err := CreateTopic(topic)
	if err != nil {
		t.Error("insert topic error")
	}

}

func TestGetTopics(t *testing.T) {
	topics, err := GetTopics(0, 20)
	if err != nil {
		t.Error("GetTopics topics error")
	}
	if len(topics) < 1 {
		t.Fatal("GetTopics return error want 1")
	}
	t.Log("len:", len(topics))
}

func TestGetOne(t *testing.T) {
	to, err := GetTopicById(1)
	if err != nil {
		t.Fatal("get topic by id is error")
	}
	t.Log(to.Title)
}

func BenchmarkGetOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetTopicById(1)
	}
}
