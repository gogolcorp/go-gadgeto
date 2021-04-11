package helpers

import (
	"reflect"
	"testing"
)

func TestJoinString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test hello world",
			args: args{str: "hello world"},
			want: "hello-world",
		},
		{
			name: "test with space at end",
			args: args{str: "hello world "},
			want: "hello-world",
		},
		{
			name: "test long sentence",
			args: args{str: "hey guys i am a test string"},
			want: "hey-guys-i-am-a-test-string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinString(tt.args.str); got != tt.want {
				t.Errorf("JoinString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFilePartsFromName(t *testing.T) {
	type args struct {
		name       string
		outputName string
	}
	tests := []struct {
		name string
		args args
		want FileParts
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFilePartsFromName(tt.args.name, tt.args.outputName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFilePartsFromName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperCaseFirstChar(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpperCaseFirstChar(tt.args.word); got != tt.want {
				t.Errorf("UpperCaseFirstChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerCase(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LowerCase(tt.args.name); got != tt.want {
				t.Errorf("LowerCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
