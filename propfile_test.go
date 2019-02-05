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
		name    string
		args    args
		want    []user
		wantErr bool
	}{
		{
			name: "Test 1",
			args: args{conts: []byte(`
			# test
			# test
			admin:pass,role1,role2
			`)},
			want:    []user{user{username: "admin", roles: []string{"role1", "role2"}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseProps(tt.args.conts)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseProps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseProps() = %v, want %v", got, tt.want)
			}
		})
	}
}
