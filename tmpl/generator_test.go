package tmpl

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerator(t *testing.T) {
	t1 := []byte(`{{define "t1"}}This is t1{{end}}`)
	t2 := []byte(`{{define "t2"}}This is t2{{end}}`)
	t3 := []byte(`{{define "t3"}}{{template "t1"}} and {{template "t2"}}{{end}}{{template "t3"}}`)

	tt, err := CreateFromBytes([][]byte{t1, t2, t3})
	require.NoError(t, err)
	require.NotNil(t, tt)
	require.Len(t, tt.Templates(), 4)

	output := bytes.Buffer{}
	err = tt.ExecuteTemplate(&output, "t1", nil)
	require.NoError(t, err)
	require.Equal(t, "This is t1", output.String())

	output.Reset()
	err = tt.ExecuteTemplate(&output, "t2", nil)
	require.NoError(t, err)
	require.Equal(t, "This is t2", output.String())

	output.Reset()
	err = tt.ExecuteTemplate(&output, "t3", nil)
	require.NoError(t, err)
	require.Equal(t, "This is t1 and This is t2", output.String())

	output.Reset()
	err = tt.ExecuteTemplate(&output, "", nil)
	require.NoError(t, err)
	require.Equal(t, "This is t1 and This is t2", output.String())

	output.Reset()
	err = tt.Execute(&output, nil)
	require.NoError(t, err)
	require.Equal(t, "This is t1 and This is t2", output.String())
}
