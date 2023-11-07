package errors

import (
	"testing"
)

var ErrTest = New("test error")

type (
	testErr struct{}
	TestErr struct{}
)

func (e TestErr) Error() string {
	return ""
}

func (t testErr) Error() string {
	return ""
}

func TestIsAny(t *testing.T) {
	type args struct {
		err  error
		errs []error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "return true because err is match any of errs",
			args: args{
				err:  ErrTest,
				errs: []error{New("fake"), ErrTest},
			},
			want: true,
		},
		{
			name: "return false because err is not match any of errs",
			args: args{
				err:  ErrTest,
				errs: []error{New("fake"), New("fake2")},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAny(tt.args.err, tt.args.errs...); got != tt.want {
				t.Errorf("IsAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAsAny(t *testing.T) {
	tests := []struct {
		name string
		errs []any
		want bool
	}{
		{
			name: "return true because err is match any of targets",
			errs: []any{&TestErr{}, &testErr{}},
			want: true,
		},
		{
			name: "return false because err is not match any of targets",
			errs: []any{&TestErr{}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testErr{}
			if got := AsAny(err, tt.errs...); got != tt.want {
				t.Errorf("AsAny() = %v, want %v", got, tt.want)
			}
		})
	}
}
