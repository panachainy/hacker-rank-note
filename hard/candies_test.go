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
		// {
		// 	name: "when_1",
		// 	args: args{
		// 		input: []int32{4, 6, 4, 5, 6, 2},
		// 	},
		// 	want: 10,
		// },
		// {
		// 	name: "when_2",
		// 	args: args{
		// 		input: []int32{1, 2, 2},
		// 	},
		// 	want: 4,
		// },
		{
			name: "when_3_now_error",
			args: args{
				input: []int32{2, 4, 2, 6, 1, 7, 8, 9, 2, 1},
			},
			// 1+2+3+1+1+1+2+1+2+1

			want: 19,
		},
		// {
		// 	name: "when_4",
		// 	args: args{
		// 		input: []int32{2, 4, 3, 5, 2, 6, 4, 5},
		// 	},
		// 	want: 12,
		// },
		// {
		// 	name: "when_5",
		// 	args: args{
		// 		input: []int32{1, 2, 2},
		// 	},
		// 	want: 4,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputs := tt.args.input
			lenth := int32(len(inputs))
			res := candies(lenth, inputs)

			assert.Equal(t, tt.want, res, tt.name)
		})
	}
}
