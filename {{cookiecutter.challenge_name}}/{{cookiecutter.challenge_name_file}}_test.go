package {{cookiecutter.challenge_name_pkg}}_test

import (
	"{{cookiecutter.challenge_name_pkg}}"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test{{cookiecutter.challenge_name_func}}(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		in  int32
		out  int32
	}{
		{
			name: "Test",
			in:  0,
			out:  0,
		},
	} {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			want := tt.out

			got := {{cookiecutter.challenge_name_pkg}}.{{cookiecutter.challenge_name_func}}(tt.in)

			assert.Equal(t, want, got)
		})
	}
}
