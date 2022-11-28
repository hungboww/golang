package postgres

import "testing"

func TestConnection(t *testing.T) {
	dataConnection := Connection()
	if dataConnection == nil {
		t.Errorf("Connection() = %v, want %v", dataConnection, nil)
	}

}
