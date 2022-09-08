package importer

import (
	"testing"
)

func TestGetFiles_success(t *testing.T) {
	_, err := getFiles("../../data")
	if err != nil {
		t.Error(err)
	}
}

func TestGetFiles_failed(t *testing.T) {
	_, err := getFiles("../../data1")
	if err == nil {
		t.Error("should fail")
	}
}

func TestParseFile_success(t *testing.T) {
	_, err := parseFile("../../data/ffa.xml")
	if err != nil {
		t.Error(err)
	}
}

func TestParseFiles_invalidFilePath(t *testing.T) {
	_, err := parseFile("unknown.xml")
	if err == nil {
		t.Error("should fail")
	}
}

func TestParseFiles_invalidFile(t *testing.T) {
	_, err := parseFile("../../data/invalid.xml")
	if err == nil {
		t.Error("should fail")
	}
}
