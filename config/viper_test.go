package config

import (
	"os"
	"testing"

	"github.com/edwinvautier/go-gadgeto/services/filesystem"
	log "github.com/sirupsen/logrus"
)

func Test_initViper(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Test without config file",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitViper(); (err != nil) != tt.wantErr {
				t.Errorf("initViper() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	// Create config file
	workdir := filesystem.GetWorkdirOrDie()
	if _, err := os.Create(workdir + "/.go-gadgeto-config.yml"); err != nil {
		log.Error(err)
		return
	}

	tests = []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Test with config file",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitViper(); (err != nil) != tt.wantErr {
				t.Errorf("initViper() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	if err := os.Remove(workdir + "/.go-gadgeto-config.yml"); err != nil {
		log.Error("couldnt remove fake config file")
		return
	}
}
