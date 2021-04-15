package updateCommand

import (
	"testing"

	"github.com/edwinvautier/go-cli/prompt/entity"
)

func TestParseEntity(t *testing.T) {
	type args struct {
		model       *entity.NewEntity
		fileContent string
	}
	tests := []struct {
		name       string
		args       args
		wantFields bool
	}{
		{
			name: "Test with bad file",
			args: args{
				model:       &entity.NewEntity{Name: "customer"},
				fileContent: "",
			},
			wantFields: false,
		},
		{
			name: "Test with goo file file",
			args: args{
				model:       &entity.NewEntity{Name: "customer"},
				fileContent: "type Customer struct {\nID int\n Name string\n}",
			},
			wantFields: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ParseEntity(tt.args.model, tt.args.fileContent)
			if tt.wantFields && len(tt.args.model.Fields) < 1 {
				t.Error("No fields : ", tt.args.model)
			}
		})
	}
}

func Test_parseField(t *testing.T) {
	type args struct {
		model *entity.NewEntity
		field *entity.EntityField
		line  string
	}
	tests := []struct {
		name string
		args args
		wantName bool
		wantType bool
		wantHasDate bool
		wantIsSlice bool
	}{
		{
			name: "test with bad line",
			args: args{
				model: &entity.NewEntity{Name: "Customer"},
				field: &entity.EntityField{},
				line: "blop",
			},
			wantName: false,
			wantType: false,
			wantHasDate: false,
			wantIsSlice: false,
		},
		{
			name: "test with standard type",
			args: args{
				model: &entity.NewEntity{Name: "Customer"},
				field: &entity.EntityField{},
				line: "ID int",
			},
			wantName: true,
			wantType: true,
			wantHasDate: false,
			wantIsSlice: false,
		},
		{
			name: "test with date",
			args: args{
				model: &entity.NewEntity{Name: "Customer"},
				field: &entity.EntityField{},
				line: "\tCreatedAt time.Time",
			},
			wantName: true,
			wantType: true,
			wantHasDate: true,
			wantIsSlice: false,
		},
		{
			name: "test with slice",
			args: args{
				model: &entity.NewEntity{Name: "Customer"},
				field: &entity.EntityField{},
				line: "TodoList []string",
			},
			wantName: false,
			wantType: false,
			wantHasDate: false,
			wantIsSlice: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseField(tt.args.model, tt.args.field, tt.args.line)
		
			if tt.wantName && len(tt.args.field.Name) < 2 {
				t.Error("No field name : ", tt.args.field)
			}

			if tt.wantType && len(tt.args.field.Type) < 2 {
				t.Error("No field type : ", tt.args.field)
			}

			if tt.wantIsSlice && tt.args.field.IsSlice == false {
				t.Error("expected field is slice : ", tt.args.field)
			}

			if tt.wantHasDate && tt.args.model.HasDate == false {
				t.Error("expected model to have date : ", tt.args.field)
			}
		})
	}
}
