package filesystem

import (
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestDirectoryExists(t *testing.T) {
	workdir := GetWorkdirOrDie()
	dirPath := workdir + "/test"

	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test without directory",
			args: args{path: dirPath},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DirectoryExists(tt.args.path); got != tt.want {
				t.Errorf("DirectoryExists() = %v, want %v", got, tt.want)
			}
		})
	}

	if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
		log.Error("mkdir : ",err)
	}
	tests = []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test with directory",
			args: args{path: dirPath},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DirectoryExists(tt.args.path); got != tt.want {
				t.Errorf("DirectoryExists() %s = %v, want %v", tt.args.path, got, tt.want)
			}
		})
	}
	os.Remove(dirPath)
}

func TestRemoveDirAndFiles(t *testing.T) {
	workdir := GetWorkdirOrDie()
	dirPath := workdir + "/test"
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test without directory",
			args: args{path: dirPath},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RemoveDirAndFiles(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("RemoveDirAndFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
		log.Error("mkdir : ",err)
	}
	if _, err := os.Create(dirPath + "/test.txt"); err != nil {
		log.Error("create file : ",err)
	}

	tests = []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test with directory",
			args: args{path: dirPath},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RemoveDirAndFiles(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("RemoveDirAndFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	os.Remove(dirPath)
}