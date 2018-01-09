package places

import (
	"testing"

	"github.com/HenkCord/GOServicePlaces/entities"
)

func Test_createValidation(t *testing.T) {
	type args struct {
		item *entities.Place
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "complete with error",
			args: args{
				item: &entities.Place{
					Name:   "name",
					City:   "city",
					Rating: 100,
					Menu: []entities.Menu{
						{
							Name: "Name",
							Cost: 1500,
						},
						{
							Cost: 1500,
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "complete",
			args: args{
				item: &entities.Place{
					Name:   "name",
					City:   "city",
					Rating: 5,
					Menu: []entities.Menu{
						{
							Name: "Name",
							Cost: 1500,
						},
						{
							Name: "Name",
							Cost: 1500,
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createValidation(tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("createValidation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
