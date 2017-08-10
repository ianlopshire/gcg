package gcg

import "testing"

func TestComment(t *testing.T) {
	tests := []struct {
		name    string
		c       Comment
		want    []byte
		wantErr bool
	}{
		{
			name:    "Single Line Comment",
			c:       NewComment("This is a single line."),
			want:    []byte("// This is a single line."),
			wantErr: false,
		},
		{
			name:    "Multi-line Comment",
			c:       NewComment("This is a line 1." + "\n" + "This is line 2."),
			want:    []byte("// This is a line 1." + "\n" + "// This is line 2."),
			wantErr: false,
		},
		{
			name:    "Empty Comment",
			c:       NewComment(""),
			want:    []byte("// "),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.MarshalGo()
			if (err != nil) != tt.wantErr {
				t.Errorf("Comment.MarshalGo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(got) != string(tt.want) {
				t.Errorf("Comment.MarshalGo() = %v, want %v", got, tt.want)
			}
		})
	}
}
