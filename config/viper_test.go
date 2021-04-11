package config

import (
	"os"
	"testing"
	log "github.com/sirupsen/logrus"
)

func Test_initViper(t *testing.T) {
	type args struct {
		config interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test without config file",
			args: args{config: "hello"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := initViper(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("initViper() error = %v, wantErr %v", err, tt.wantErr)
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
			name: "Test with config file",
			args: args{config: "hello"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := initViper(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("initViper() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	if err := os.Remove(workdir + "/.go-cli-config.yml"); err != nil {
		log.Error("couldnt remove fake config file")
		return
	}
}
