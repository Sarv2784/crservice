package library

import (
	"reflect"
	"testing"
)

func TestStructToStringMap(t *testing.T) {
	type TestStruct struct {
		Name  string
		Age   int
		Score float64
		Alive bool
	}

	tests := []struct {
		name    string
		input   any
		want    map[string]string
		wantErr bool
	}{
		{
			name: "Valid struct",
			input: TestStruct{
				Name:  "John",
				Age:   30,
				Score: 99.5,
				Alive: true,
			},
			want: map[string]string{
				"name":  "John",
				"age":   "30",
				"score": "99.5",
				"alive": "true",
			},
			wantErr: false,
		},
		{
			name: "Struct with empty fields",
			input: TestStruct{
				Name:  "",
				Age:   0,
				Score: 0.0,
				Alive: false,
			},
			want: map[string]string{
				"name":  "",
				"age":   "0",
				"score": "0",
				"alive": "false",
			},
			wantErr: false,
		},
		{
			name:    "Non-struct input",
			input:   "Not a struct",
			want:    nil,
			wantErr: true,
		},
		{
			name:  "Pointer to struct",
			input: &TestStruct{Name: "Alice", Age: 25, Score: 87.5, Alive: true},
			want: map[string]string{
				"name":  "Alice",
				"age":   "25",
				"score": "87.5",
				"alive": "true",
			},
			wantErr: false,
		},
		{
			name:    "Nil pointer",
			input:   (*TestStruct)(nil),
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StructToStringMap(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("StructToStringMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructToStringMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}
