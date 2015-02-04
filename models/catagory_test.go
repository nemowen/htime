package models

import "testing"

func TestCreateCategory(t *testing.T) {
	ct := Category{
		Name: "日志",
	}

	ct.Save()
}
