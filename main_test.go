package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductDTO_ToEntity(t *testing.T) {
	type args struct {
		input []int32
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "when_",
			args: args{
				input: []int32{4, 6, 4, 5, 6, 2},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputs := tt.args.input
			lenth := int32(len(inputs))
			res := candies(lenth, inputs)

			assert.Equal(t, res, tt.want, tt.name)
		})
	}
}
