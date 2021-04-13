package config

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/edwinvautier/go-cli/services/filesystem"
	"github.com/gobuffalo/packr/v2"
)

func TestInstallCmdConfig_GetBox(t *testing.T) {
	type fields struct {
		Box *packr.Box
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test with box",
			fields: fields{
				Box: packr.New("testing box", "./"),
			},
			want: packr.New("testing box ", "./").Name,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := InstallCmdConfig{
				Box: tt.fields.Box,
			}
			name := cmd.Box.Name
			if got := cmd.GetBox().Name; !reflect.DeepEqual(got, name) {
				t.Errorf("InstallCmdConfig.GetBox() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstallCmdConfig_GetProjectPath(t *testing.T) {
	type fields struct {
		ProjectPath string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test with path",
			fields: fields{
				ProjectPath: "/hello/world",
			},
			want: "/hello/world",
		},
		{
			name: "test with empty path",
			fields: fields{
				ProjectPath: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := MakeCmdConfig{
				ProjectPath: tt.fields.ProjectPath,
			}
			if got := cmd.GetProjectPath(); got != tt.want {
				t.Errorf("InstallCmdConfig.GetProjectPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateConfigAfterInstalling(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test without viper",
			args: args{name: "authenticator"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateConfigAfterInstalling(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("UpdateConfigAfterInstalling() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	// create config file
	workdir:= filesystem.GetWorkdirOrDie()
	os.Create(workdir + "/.go-cli-config.yml")

	tests = []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test without viper",
			args: args{name: "authenticator"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateConfigAfterInstalling(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("UpdateConfigAfterInstalling() error = %v, wantErr %v", err, tt.wantErr)
			}
			box := packr.New("test box", workdir)
			fileContent, _ := box.FindString(".go-cli-config.yml")
			if !strings.Contains(fileContent, "authenticator: true") {
				t.Errorf("file was not written")
			}
		})
	}

	os.Remove(workdir + "/.go-cli-config.yml")
}
