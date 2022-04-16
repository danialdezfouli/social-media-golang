package service

import (
	"testing"
)

func TestFindEnglishHashtags(t *testing.T) {
	expected := []string{"every", "first_tweet"}
	tags := FindHashtags("hello #every one this is my #first_tweet")

	for i, item := range expected {
		if tags[i] != item {
			t.Fatalf("expected to be '%s' but got '%s'", item, tags[1])
		}
	}
}

func TestFindPersianHashtags(t *testing.T) {
	expected := []string{"اولین", "توییترـفارسی"}
	tags := FindHashtags("این #اولین پست من است #توییترـفارسی")

	for i, item := range expected {
		if tags[i] != item {
			t.Fatalf("expected to be '%s' but got '%s'", item, tags[1])
		}
	}
}
