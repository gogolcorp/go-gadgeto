package config

import (
	"reflect"
	"testing"

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
