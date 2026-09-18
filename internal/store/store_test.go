package store

import "testing"

func TestSet(t *testing.T) {
	// Setting a new key stores its value.
	t.Run("stores a new key", func(t *testing.T) {
		s := New()
		s.Set("name", "alice")
		got, ok := s.Get("name")
		if !ok || got != "alice" {
			t.Fatalf("Get after Set = (%q, %v), want (%q, true)", got, ok, "alice")
		}
	})

	// Setting an existing key overwrites the old value.
	t.Run("overwrites an existing key", func(t *testing.T) {
		s := New()
		s.Set("name", "alice")
		s.Set("name", "bob")
		got, _ := s.Get("name")
		if got != "bob" {
			t.Fatalf("Get after overwrite = %q, want %q", got, "bob")
		}
	})
}

func TestGet(t *testing.T) {
	// Getting an existing key returns its value and true.
	t.Run("returns value for existing key", func(t *testing.T) {
		s := New()
		s.Set("k", "v")
		got, ok := s.Get("k")
		if got != "v" || !ok {
			t.Fatalf("Get = (%q, %v), want (%q, true)", got, ok, "v")
		}
	})

	// Getting a missing key returns empty string and false.
	t.Run("returns false for missing key", func(t *testing.T) {
		s := New()
		got, ok := s.Get("missing")
		if got != "" || ok {
			t.Fatalf("Get = (%q, %v), want (\"\", false)", got, ok)
		}
	})

	// A deleted key can no longer be retrieved.
	t.Run("returns false after delete", func(t *testing.T) {
		s := New()
		s.Set("k", "v")
		s.Del("k")
		if _, ok := s.Get("k"); ok {
			t.Fatal("Get returned ok=true for a deleted key")
		}
	})
}

func TestExist(t *testing.T) {
	// A key that was set exists.
	t.Run("true for existing key", func(t *testing.T) {
		s := New()
		s.Set("k", "v")
		if !s.Exist("k") {
			t.Fatal("Exist = false, want true for a set key")
		}
	})

	// A key that was never set does not exist.
	t.Run("false for missing key", func(t *testing.T) {
		s := New()
		if s.Exist("nope") {
			t.Fatal("Exist = true, want false for an unset key")
		}
	})

	// A key stops existing once deleted.
	t.Run("false after delete", func(t *testing.T) {
		s := New()
		s.Set("k", "v")
		s.Del("k")
		if s.Exist("k") {
			t.Fatal("Exist = true, want false after delete")
		}
	})
}

func TestDel(t *testing.T) {
	// Deleting an existing key removes it and counts 1.
	t.Run("deletes a single key", func(t *testing.T) {
		s := New()
		s.Set("k", "v")
		if n := s.Del("k"); n != 1 {
			t.Fatalf("Del = %d, want 1", n)
		}
		if s.Exist("k") {
			t.Fatal("key still exists after Del")
		}
	})

	// Deleting a missing key counts 0.
	t.Run("missing key counts zero", func(t *testing.T) {
		s := New()
		if n := s.Del("ghost"); n != 0 {
			t.Fatalf("Del = %d, want 0", n)
		}
	})

	// Deleting several keys counts only the ones that existed.
	t.Run("counts only existing keys", func(t *testing.T) {
		s := New()
		s.Set("a", "1")
		s.Set("b", "2")
		if n := s.Del("a", "b", "missing"); n != 2 {
			t.Fatalf("Del = %d, want 2", n)
		}
	})
}
