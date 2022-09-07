package importer

import (
	"testing"
)

func TestGetFiles_success(t *testing.T) {
	_, err := getFiles("../../stats")
	if err != nil {
		t.Error(err)
	}
}

func TestGetFiles_failes(t *testing.T) {
	_, err := getFiles("../../stats1")
	if err == nil {
		t.Error("should fail")
	}
}
