package tests

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) { //nolint:paralleltest
	ts := newTester(t)
	defer ts.teardown()

	ts.initStore()

	out, err := ts.run("insert")
	assert.Error(t, err)
	assert.Equal(t, "\nError: Usage: "+filepath.Base(ts.Binary)+" insert name\n", out)

	_, err = ts.runCmd([]string{ts.Binary, "insert", "some/secret"}, []byte("moar"))
	assert.NoError(t, err)

	_, err = ts.runCmd([]string{ts.Binary, "insert", "some/newsecret"}, []byte("and\nmoar"))
	assert.NoError(t, err)

	t.Run("Regression test for #1573 without actual pipes", func(t *testing.T) { //nolint:paralleltest
		out, err = ts.run("show -f some/secret")
		assert.NoError(t, err)
		assert.Equal(t, "moar", out)

		out, err = ts.run("show -f some/newsecret")
		assert.NoError(t, err)
		assert.Equal(t, "and\nmoar", out)

		out, err = ts.run("show -f some/secret")
		assert.NoError(t, err)
		assert.Equal(t, "moar", out)

		out, err = ts.run("show -f some/newsecret")
		assert.NoError(t, err)
		assert.Equal(t, "and\nmoar", out)
	})

	t.Run("Regression test for #1595", func(t *testing.T) { //nolint:paralleltest
		t.Skip("TODO")

		_, err = ts.runCmd([]string{ts.Binary, "insert", "some/other"}, []byte("nope"))
		assert.NoError(t, err)

		out, err = ts.run("insert some/other")
		assert.Error(t, err)
		assert.Equal(t, "\nError: not overwriting your current secret\n", out)

		out, err = ts.run("show -o some/other")
		assert.NoError(t, err)
		assert.Equal(t, "nope", out)

		out, err = ts.run("--yes insert some/other")
		assert.NoError(t, err)
		assert.Equal(t, "Warning: Password is empty or all whitespace", out)

		out, err = ts.run("insert -f some/other")
		assert.NoError(t, err)
		assert.Equal(t, "Warning: Password is empty or all whitespace", out)

		out, err = ts.run("show -o some/other")
		assert.Error(t, err)
		assert.Equal(t, "\nError: empty secret\n", out)

		_, err = ts.runCmd([]string{ts.Binary, "insert", "-f", "some/other"}, []byte("final"))
		assert.NoError(t, err)

		out, err = ts.run("show -o some/other")
		assert.NoError(t, err)
		assert.Equal(t, "final", out)

		// This is arguably not a good behaviour: it should not overwrite the password when we are only working on a key:value.
		out, err = ts.run("insert -f some/other test:inline")
		assert.NoError(t, err)
		assert.Equal(t, "", out)

		out, err = ts.run("show some/other test")
		assert.NoError(t, err)
		assert.Equal(t, "inline", out)

		out, err = ts.run("insert some/other test:inline2")
		assert.Error(t, err)
		assert.Equal(t, "\nError: not overwriting your current secret\n", out)

		out, err = ts.run("show some/other Test")
		assert.NoError(t, err)
		assert.Equal(t, "inline", out)

		out, err = ts.run("--yes insert some/other test:inline2")
		assert.NoError(t, err)
		assert.Equal(t, "", out)

		out, err = ts.run("show some/other Test")
		assert.NoError(t, err)
		assert.Equal(t, "inline2", out)
	})

	t.Run("Regression test for #1650 with JSON", func(t *testing.T) { //nolint:paralleltest
		json := `Password: SECRET
--
glossary": {
    "title": "example glossary",
    "GlossDiv": {
        "title": "S",
        "GlossList": {
            "GlossEntry": {
                "ID": "SGML",
                "SortAs": "SGML",
                "GlossTerm": "Standard Generalized Markup Language",
                "Acronym": "SGML",
                "Abbrev": "ISO 8879:1986",
                "GlossDef": {
                    "para": "A meta-markup language, used to create markup languages such as DocBook.",
                    "GlossSeeAlso": ["GML", "XML"]
                },
                "GlossSee": "markup"
            }
        }
    }
}`
		_, err = ts.runCmd([]string{ts.Binary, "insert", "some/json"}, []byte(json))
		assert.NoError(t, err)

		// using show -n to disable parsing
		out, err = ts.run("show -f -n some/json")
		assert.NoError(t, err)
		assert.Equal(t, json, out)
	})

	t.Run("Regression test for #1600", func(t *testing.T) { //nolint:paralleltest
		input := `test1
test2
{
  "Creator": "the creator"
}`
		_, err = ts.runCmd([]string{ts.Binary, "insert", "some/multilinewithbraces"}, []byte(input))
		assert.NoError(t, err)

		// using show -n to disable parsing
		out, err = ts.run("show -f -n some/multilinewithbraces")
		assert.NoError(t, err)
		assert.Equal(t, input, out)
	})

	t.Run("Regression test for #1601", func(t *testing.T) { //nolint:paralleltest
		input := `thepassword
user: a user
web: test.com
user: second user`

		_, err = ts.runCmd([]string{ts.Binary, "insert", "some/multikey"}, []byte(input))
		assert.NoError(t, err)

		// using show -n to disable parsing
		out, err = ts.run("show -f -n some/multikey")
		assert.NoError(t, err)
		assert.Equal(t, input, out)
	})

	t.Run("Regression test full support for #1601", func(t *testing.T) { //nolint:paralleltest
		t.Skip("Skipping until we support actual key-valueS for KV")

		input := `thepassword
user: a user
web: test.com
user: second user`

		output := `thepassword
web: test.com
user: a user
user: second user`

		_, err = ts.runCmd([]string{ts.Binary, "insert", "some/multikeyvalues"}, []byte(input))
		assert.NoError(t, err)

		out, err = ts.run("show -f some/multikeyvalues")
		assert.NoError(t, err)
		assert.Equal(t, output, out)
	})

	t.Run("Regression test for #1614", func(t *testing.T) { //nolint:paralleltest
		input := `yamltest
---
user: 0123`

		output := `yamltest
---
user: 83`

		_, err = ts.runCmd([]string{ts.Binary, "insert", "some/yamloctal"}, []byte(input))
		assert.NoError(t, err)

		// with parsing we have 0123 interpreted as octal for 83
		out, err = ts.run("show -f some/yamloctal")
		assert.NoError(t, err)
		assert.Equal(t, output, out)

		// using show -n to disable parsing
		out, err = ts.run("show -f -n some/yamloctal")
		assert.NoError(t, err)
		assert.Equal(t, input, out)
	})

	t.Run("Regression test for #1594", func(t *testing.T) { //nolint:paralleltest
		input := `somepasswd
---
Test / test.com
user:myuser
url: test.com/`

		_, err = ts.runCmd([]string{ts.Binary, "insert", "some/kvwithspace"}, []byte(input))
		assert.NoError(t, err)

		out, err = ts.run("show -f some/kvwithspace")
		assert.NoError(t, err)
		assert.Equal(t, input, out)

		out, err = ts.run("show -f some/kvwithspace url")
		assert.NoError(t, err)
		assert.Equal(t, "test.com/", out)

		out, err = ts.run("show -f some/kvwithspace user")
		assert.Error(t, err)
	})
}
