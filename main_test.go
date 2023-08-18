package main

import "testing"

func TestSayHello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test01",
			want: "Hello World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SayHello(); got != tt.want {
				t.Errorf("SayHello() = %v, want %v", got, tt.want)
			}
		})
	}
}
