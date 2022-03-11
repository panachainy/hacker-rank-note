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
		want int32
	}{
		{
			name: "when_1",
			args: args{
				input: []int32{3, -7, 0},
			},
			want: 3,
		},
		{
			name: "when_2",
			args: args{
				input: []int32{-59, -36, -13, 1, -53, -92, -2, -96, -54, 75},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputs := tt.args.input
			res := minimumAbsoluteDifference(inputs)

			assert.Equal(t, tt.want, res, tt.name)
		})
	}
}
