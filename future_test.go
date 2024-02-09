package async

import (
	"reflect"
	"testing"
)

func TestAwaitAll(t *testing.T) {
	tests := []struct {
		name    string
		futures []Future[int]
		want    []int
	}{
		{
			name: "덧셈 task 목록 실행",
			futures: []Future[int]{
				RunTask(func(args ...any) int {
					lhs := args[0].(int)
					rhs := args[1].(int)

					return lhs + rhs
				}, 10, 20),
				RunTask(func(args ...any) int {
					lhs := args[0].(int)
					rhs := args[1].(int)

					return lhs + rhs
				}, 15, 2),
				RunTask(func(args ...any) int {
					lhs := args[0].(int)
					rhs := args[1].(int)

					return lhs + rhs
				}, 1, 20),
			},
			want: []int{
				30, 17, 21,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AwaitAll(tt.futures...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AwaitAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
