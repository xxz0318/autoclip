// Package conf 请修改此处包名注释
// @author: xiexinzhong
// @create: 2024-01-26 18:21
// @description:
package conf

import "testing"

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LoadConfig()
		})
	}
}
