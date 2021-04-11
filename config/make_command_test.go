package config

import (
	"os"
	"reflect"
	"testing"

	"github.com/edwinvautier/go-cli/prompt/entity"
	"github.com/gobuffalo/packr/v2"
	log "github.com/sirupsen/logrus"
)

func TestAddModelToConfig(t *testing.T) {
	type args struct {
		newEntity entity.NewEntity
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				newEntity: entity.NewEntity{
					Name: "",
					NamePascalCase: "",
					NameLowerCase: "",
					HasDate: false,
					HasCustomTypes: false,
					Fields: []entity.EntityField{
						{
							Type: "string",
							Name: "Name",
							IsSlice: false,
							SliceType: "",
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddModelToConfig(tt.args.newEntity); (err != nil) != tt.wantErr {
				t.Errorf("AddModelToConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	// Create config file
	workdir, err := os.Getwd()
	if err != nil {
		return
	}
	if _, err := os.Create(workdir + "/.go-cli-config.yml"); err != nil {
		log.Error(err)
		return
	}

	tests = []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				newEntity: entity.NewEntity{
					Name: "",
					NamePascalCase: "",
					NameLowerCase: "",
					HasDate: false,
					HasCustomTypes: false,
					Fields: []entity.EntityField{
						{
							Type: "string",
							Name: "Name",
							IsSlice: false,
							SliceType: "",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddModelToConfig(tt.args.newEntity); (err != nil) != tt.wantErr {
				t.Errorf("AddModelToConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	if err := os.Remove(workdir + "/.go-cli-config.yml"); err != nil {
		log.Error(err)
		return
	}
}

func TestMakeCmdConfig_GetBox(t *testing.T) {
	type fields struct {
		GoPackageFullPath string
		Box               *packr.Box
		Entity            entity.NewEntity
		ProjectPath       string
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
			cmd := MakeCmdConfig{
				GoPackageFullPath: tt.fields.GoPackageFullPath,
				Box:               tt.fields.Box,
				Entity:            tt.fields.Entity,
				ProjectPath:       tt.fields.ProjectPath,
			}
			name := cmd.Box.Name
			if got := cmd.GetBox().Name; !reflect.DeepEqual(got, name) {
				t.Errorf("MakeCmdConfig.GetBox() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeCmdConfig_GetProjectPath(t *testing.T) {
	type fields struct {
		GoPackageFullPath string
		Box               *packr.Box
		Entity            entity.NewEntity
		ProjectPath       string
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
				GoPackageFullPath: tt.fields.GoPackageFullPath,
				Box:               tt.fields.Box,
				Entity:            tt.fields.Entity,
				ProjectPath:       tt.fields.ProjectPath,
			}
			if got := cmd.GetProjectPath(); got != tt.want {
				t.Errorf("MakeCmdConfig.GetProjectPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
