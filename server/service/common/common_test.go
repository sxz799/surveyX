package common

import (
	"github.com/sxz799/surveyX/config"
	"github.com/sxz799/surveyX/model/entity"
	"github.com/sxz799/surveyX/utils"
	"gorm.io/gorm"
	"testing"
)

func TestService_Login(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		user entity.LoginUser
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantU   entity.User
		wantErr bool
	}{
		{
			name: "test login",
			fields: fields{
				db: utils.InitDB(config.NewConfig()),
			},
			args: args{
				entity.LoginUser{
					Username: "33",
					Password: "33",
				},
			},
			wantU: entity.User{
				Username: "33",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				db: tt.fields.db,
			}
			gotU, err := s.Login(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotU.Username != tt.wantU.Username {
				t.Errorf("Login() gotU = %v, want %v", gotU, tt.wantU)
			}
		})
	}
}
