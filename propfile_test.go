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
			want: []user{user{username: "admin", roles: []string{"role1", "role2"}}},
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
				user{username: "admin", roles: []string{"role1", "role2"}},
				user{username: "arjun", roles: []string{"user", "admin"}},
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
