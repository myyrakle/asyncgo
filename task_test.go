package async

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRunTask(t *testing.T) {
	type args struct {
		f    func(args ...any) int
		lhs  int
		rhs  int
		args []any
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "RunTask",
			args: args{
				f: func(args ...any) int {
					lhs := args[0].(int)
					rhs := args[1].(int)

					return lhs + rhs
				},
				lhs: 1,
				rhs: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := RunTask(tt.args.f, tt.args.lhs, tt.args.rhs)

			if got := task.Await(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunPanicableTask(t *testing.T) {
	type args struct {
		f    func(args ...any) Result[int]
		lhs  int
		rhs  int
		args []any
	}
	tests := []struct {
		name string
		args args
		want Result[int]
	}{
		{
			name: "정상 동작시",
			args: args{
				f: func(args ...any) Result[int] {
					lhs := args[0].(int)
					rhs := args[1].(int)

					return Ok(lhs + rhs)
				},
				lhs: 1,
				rhs: 2,
			},
			want: Ok(3),
		},
		{
			name: "패닉 발생시",
			args: args{
				f: func(args ...any) Result[int] {
					lhs := args[0].(int)
					rhs := args[1].(int)

					if lhs == 1 {
						panic("lhs is 1")
					}

					return Ok(lhs + rhs)
				},
				lhs: 1,
				rhs: 2,
			},
			want: Err[int](fmt.Errorf("panic recover: lhs is 1")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := RunPanicableTask(tt.args.f, tt.args.lhs, tt.args.rhs)

			if got := task.Await(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunPanicableTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
