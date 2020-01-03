package redis

import (
	"testing"
)

func TestSignGenerator_GetSID(t *testing.T) {
	tests := []struct {
		name    string
		want    int64
		wantErr bool
	}{
		{
			"",
			10001,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			si := InitSignGenerator()
			got, err := si.GetSID()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got < tt.want {
				t.Errorf("GetSID() got = %v, want %v", got, tt.want)
			}
		})
	}
}