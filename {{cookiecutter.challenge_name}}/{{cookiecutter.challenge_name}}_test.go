package {{cookiecutter.challenge_name}}_test

import (
	"{{cookiecutter.challenge_name}}"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test{{cookiecutter.challenge_name.replace('-',' ').title().replace(' ', '')}}(t *testing.T) {
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

			got := {{cookiecutter.challenge_name}}.{{cookiecutter.challenge_name}}(tt.in)

			assert.Equal(t, want, got)
		})
	}

}
