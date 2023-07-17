package dictionary_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/patrickhuber/go-collections/dictionary"
)

func TestDictionary(t *testing.T) {
	t.Run("set", func(t *testing.T) {

		d := dictionary.New()
		d.Set("test", "1")
		require.Equal(t, 1, d.Length())
	})
	t.Run("get", func(t *testing.T) {

		d := dictionary.New()
		d.Set("test", "1")
		value := d.Get("test")
		require.Equal(t, "1", value)
	})
	t.Run("list_keys", func(t *testing.T) {
		d := dictionary.New()
		d.Set("test", 1)
		d.Set("other", 2)
		keys := d.Keys()
		require.Equal(t, 2, len(keys))
		if keys[0] == "test" {
			require.Equal(t, "other", keys[1])
		} else if keys[0] == "other" {
			require.Equal(t, "test", keys[1])
		} else {
			t.Fatalf("unknown key")
		}
	})

	t.Run("list_values", func(t *testing.T) {
		d := dictionary.New()
		d.Set("test", "1")
		d.Set("other", "2")
		values := d.Values()
		require.Equal(t, 2, len(values))

		if values[0] == "1" {
			require.Equal(t, "2", values[1])
		} else if values[0] == "2" {
			require.Equal(t, "1", values[1])
		} else {
			t.Fatal("unknown value")
		}
	})

	t.Run("clear", func(t *testing.T) {
		d := dictionary.New()
		d.Set("test", "1")
		d.Set("other", "2")
		d.Clear()
		require.Equal(t, 0, d.Length())
	})

	t.Run("lookup_missing", func(t *testing.T) {

		d := dictionary.New()
		d.Set("test", "1")
		d.Set("other", "2")
		_, ok := d.Lookup("hello")
		require.False(t, ok)
	})

	t.Run("lookup_present", func(t *testing.T) {

		d := dictionary.New()
		d.Set("test", "1")
		d.Set("other", "2")
		value, ok := d.Lookup("test")
		require.True(t, ok)
		require.Equal(t, "1", value)
	})

	t.Run("remove", func(t *testing.T) {
		d := dictionary.New()
		d.Set("test", "1")
		d.Set("other", "2")
		d.Remove("other")
		require.Equal(t, 1, d.Length())
	})
}
