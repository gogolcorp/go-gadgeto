package authenticator

import "testing"

func Test_getPublicPemFile(t *testing.T) {
	type args struct {
		fileData *[]byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "no file",
			args:    args{fileData: &[]byte{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := getPublicPemFile(tt.args.fileData); (err != nil) != tt.wantErr {
				t.Errorf("getPublicPemFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getPrivatePemFile(t *testing.T) {
	type args struct {
		fileData *[]byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "no file",
			args:    args{fileData: &[]byte{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := getPrivatePemFile(tt.args.fileData); (err != nil) != tt.wantErr {
				t.Errorf("getPrivatePemFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
