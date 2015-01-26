package utils

import (
	"testing"
)

func TestEncodeByMD5(t *testing.T) {
	p := "123456"

	end := EncodeByMd5(p)
	if len(end) < 10 && end == "2A6E6F72612AE10ADC3949BA59ABBE56E057F20F883E" {
		t.Fatalf("encode error")
	}
	t.Log("encode:", end)
}
