package utility

import "testing"

func TestZipSheetsData(t *testing.T) {
	keys := make([]interface{}, 2)
	keys[0] = "type"
	keys[1] = "name"

	data := make([][]interface{}, 2)
	firstRow := make([]interface{}, 2)
	secondRow := make([]interface{}, 2)

	firstRow[0] = "cat"
	firstRow[1] = "Aslan"
	secondRow[0] = "dog"
	secondRow[1] = "Snoopy"

	data[0] = firstRow
	data[1] = secondRow

	actual := ZipSheetsData(keys, data)

	if actual[0]["type"] != "cat" {
		t.Error("Expected first entry to have type of cat")
	}

	if actual[0]["name"] != "Aslan" {
		t.Error("Expected first entry to have name of Aslan")
	}

	if actual[1]["type"] != "dog" {
		t.Error("Expected second entry to have type of dog")
	}

	if actual[1]["name"] != "Snoopy" {
		t.Error("Expected second entry to have name of Snoopy")
	}
}
