package cmd

import (
	"testing"

	_ "github.com/mattes/migrate/source/file"
	"github.com/spf13/viper"
)

func Test_connectToDbs(t *testing.T) {
	viper.AddConfigPath("../")
	viper.SetConfigName(".config")
	if err := viper.ReadInConfig(); err != nil {
		panic("Fatal error config file:" + err.Error())
	}

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Connect",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connectToDbs(); got == nil {
				t.Errorf("connectToDbs() = %v", got)
			}
		})
	}
}
