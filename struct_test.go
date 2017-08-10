package gcg

import (
	"testing"
)

func TestStruct_MarshalGo(t *testing.T) {
	getPtr := func(s Struct) *Struct {
		return &s
	}
	tests := []struct {
		name    string
		Struct  *Struct
		want    []byte
		wantErr bool
	}{
		{
			name:    "Empty Struct",
			Struct:  getPtr(NewStruct("EmptyStruct")),
			want:    []byte("type EmptyStruct struct {}"),
			wantErr: false,
		},
		{
			name:    "Single Field",
			Struct:  getPtr(NewStruct("EmptyStruct", NewStructField("GeneralCase", "string", NewStructTags("json", "jsonValue")))),
			want:    []byte("type EmptyStruct struct {GeneralCase string `json:\"jsonValue\"`;}"),
			wantErr: false,
		},
		{
			name: "Multiple Field",
			Struct: getPtr(
				NewStruct("EmptyStruct",
					NewStructField("GeneralCase", "string", NewStructTags("json", "jsonValue")),
					NewStructField("GeneralCase2", "time.Time", NewStructTags("json", "jsonValue")),
				),
			),
			want:    []byte("type EmptyStruct struct {GeneralCase string `json:\"jsonValue\"`;GeneralCase2 time.Time `json:\"jsonValue\"`;}"),
			wantErr: false,
		},
		{
			name:    "Nil Struct",
			Struct:  nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.Struct.MarshalGo()
			if (err != nil) != tt.wantErr {
				t.Errorf("Struct.MarshalGo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(got) != string(tt.want) {
				t.Errorf("Struct.MarshalGo() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestStructField_MarshalGo(t *testing.T) {
	getPtr := func(sf StructField) *StructField {
		return &sf
	}
	tests := []struct {
		name        string
		structField *StructField
		want        []byte
		wantErr     bool
	}{
		{
			name:        "General Case No Tags",
			structField: getPtr(NewStructField("GeneralCase", "string", nil)),
			want:        []byte("GeneralCase string "),
			wantErr:     false,
		},
		{
			name:        "General Case with Tags",
			structField: getPtr(NewStructField("GeneralCase", "string", NewStructTags("json", "jsonValue"))),
			want:        []byte("GeneralCase string `json:\"jsonValue\"`"),
			wantErr:     false,
		},
		{
			name:        "Empty Name (Embedded Type)",
			structField: getPtr(NewStructField("", "EmbeddedType", nil)),
			want:        []byte(" EmbeddedType "),
			wantErr:     false,
		},
		{
			name:        "Nil Struct Field",
			structField: nil,
			wantErr:     true,
		},
		{
			name:        "Empty Type",
			structField: getPtr(NewStructField("GeneralCase", "", nil)),
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.structField.MarshalGo()
			if (err != nil) != tt.wantErr {
				t.Errorf("StructField.MarshalGo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(got) != string(tt.want) {
				t.Errorf("StructField.MarshalGo() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestStructTags(t *testing.T) {
	getPtr := func(st StructTags) *StructTags {
		return &st
	}

	tests := []struct {
		name    string
		t       *StructTags
		want    []byte
		wantErr bool
	}{
		// TODO: Fix test for multiple types. The Map does not always return the tags in a specific order.
		// {
		// 	name:    "Multiple Tags",
		// 	t:       getPtr(NewStructTags("json", "jsonField", "xml", "xmlField")),
		// 	want:    []byte("`json:\"jsonField\" xml:\"xmlField\"`"),
		// 	wantErr: false,
		// },
		{
			name:    "Single Tags",
			t:       getPtr(NewStructTags("json", "jsonField")),
			want:    []byte("`json:\"jsonField\"`"),
			wantErr: false,
		},
		{
			name:    "Nil Tags",
			t:       nil,
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.t.MarshalGo()
			if (err != nil) != tt.wantErr {
				t.Errorf("StructTags.MarshalGo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(got) != string(tt.want) {
				t.Errorf("StructTags.MarshalGo() = %s, want %s", got, tt.want)
			}
		})
	}
}
