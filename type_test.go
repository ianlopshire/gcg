package gcg

import (
	"testing"
)

func TestNamedType(t *testing.T) {
	getPtr := func(nt NamedType) *NamedType {
		return &nt
	}

	tests := []struct {
		name      string
		namedType *NamedType
		want      []byte
		wantErr   bool
	}{
		{
			name:      "General Case",
			namedType: getPtr(NewNamedType("Empty", "string")),
			want:      []byte("type Empty string"),
			wantErr:   false,
		},
		{
			name:      "Empty Underlying Type",
			namedType: getPtr(NewNamedType("EmptyUnderlyingType", "")),
			wantErr:   true,
		},
		{
			name:      "Empty Name",
			namedType: getPtr(NewNamedType("", "string")),
			wantErr:   true,
		},
		{
			name:      "Nil Pointer",
			namedType: nil,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.namedType.MarshalGo()

			if (err != nil) != tt.wantErr {
				t.Errorf("NamedType.MarshalGo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(tt.want) != string(tt.want) {
				t.Errorf("NamedType.MarshalGo() = %s, want %s", got, tt.want)
			}

		})
	}
}
