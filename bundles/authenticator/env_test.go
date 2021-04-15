package authenticator

import (
	"os"
	"os/exec"
	"testing"

	"github.com/edwinvautier/go-cli/services/filesystem"
	log "github.com/sirupsen/logrus"
)

func Test_goDotEnvVariable(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no env is set",
			args: args{key: "TOKEN_DURATION"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goDotEnvVariable(tt.args.key); got != tt.want {
				t.Errorf("goDotEnvVariable() = %v, want %v", got, tt.want)
			}
		})
	}
	workdir := filesystem.GetWorkdirOrDie()
	createEnv := exec.Command("touch", workdir+"/.env")
	createEnv.Run()

	file, _ := os.OpenFile(workdir+"/.env", os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	if _, err := file.WriteString("TOKEN_DURATION=10"); err != nil {
		log.Fatal(err)
	}

	tests = []struct {
		name string
		args args
		want string
	}{
		{
			name: "no env is set",
			args: args{key: "TOKEN_DURATION"},
			want: "10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goDotEnvVariable(tt.args.key); got != tt.want {
				t.Errorf("goDotEnvVariable() = %v, want %v", got, tt.want)
			}
		})
	}
	deleteEnv := exec.Command("rm", workdir+"/.env")
	deleteEnv.Run()
}
