package pablo

import "testing"

func TestLoadSection(t *testing.T) {
	_, err := LoadSection("fixture/book")
	if err != nil {
		t.Fatal(err)
	}
}
