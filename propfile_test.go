package main

import (
	"reflect"
	"testing"
)

func Test_parseProps(t *testing.T) {
	type args struct {
		conts []byte
	}
	tests := []struct {
		name string
		args args
		want []user
	}{
		{
			name: "Test 1",
			args: args{conts: []byte(`
			# test
			# test
			admin:pass,role1,role2
			`)},
			want: []user{user{Username: "admin", Roles: []string{"role1", "role2"}}},
		},
		{
			name: "Test 2",
			args: args{conts: []byte(`
			# test
			# test
			# admin:pass,role1,role2
			`)},
			want: []user{},
		},
		{
			name: "Test 3",
			args: args{conts: []byte(`
			admin:pass,role1,role2
			arjun:pass,user,admin
			`)},
			want: []user{
				user{Username: "admin", Roles: []string{"role1", "role2"}},
				user{Username: "arjun", Roles: []string{"user", "admin"}},
			},
		},
		{
			name: "Test 4",
			args: args{conts: []byte(`
			admin:pass,role1,role2

			arjun:pass,user,admin
			`)},
			want: []user{
				user{Username: "admin", Roles: []string{"role1", "role2"}},
				user{Username: "arjun", Roles: []string{"user", "admin"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseProps(tt.args.conts)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseProps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propsFile_GetUsers(t *testing.T) {
	type fields struct {
		path string
	}

	f := fields{path: "/etc/rundeck/realm.properties"}
	tests := []struct {
		name    string
		fields  fields
		want    []user
		wantErr bool
	}{
		{
			name:   "Test 1",
			fields: f,
			want: []user{
				user{Username: "admin", Roles: []string{"user", "admin", "architect", "deploy", "build"}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pf := &propsFile{
				path: tt.fields.path,
			}
			got, err := pf.GetUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("propsFile.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("propsFile.GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_propsFile_UpdateUser(t *testing.T) {
	type fields struct {
		path string
	}
	type args struct {
		oldUsername string
		newUsername string
		newRoles    []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Test 1",
			fields: fields{path: "test-realm"},
			args: args{
				oldUsername: "arjun",
				newRoles:    []string{},
				newUsername: "arjun",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pf := &propsFile{
				path: tt.fields.path,
			}
			if err := pf.UpdateUser(tt.args.oldUsername, tt.args.newUsername, tt.args.newRoles); (err != nil) != tt.wantErr {
				t.Errorf("propsFile.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
