package factory

import (
	"testing"
)

func Test_Draw(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "Circle", fields: fields{Name: "Circle"}},
		{name: "Square", fields: fields{Name: "Square"}},
		{name: "Rectangle", fields: fields{Name: "Rectangle"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sp := ShapeFactory(tt.fields.Name)
			sp.Draw()
		})
	}
}
