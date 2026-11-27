package db

import (
	"path/filepath"
	"testing"
)

func TestCheckpointProducesConsistentCopy(t *testing.T) {
	baseDir := t.TempDir()

	src, err := NewRocksDB(filepath.Join(baseDir, "src"))
	if err != nil {
		t.Fatalf("NewRocksDB: %v", err)
	}
	defer src.Close()

	if err := src.Put("a", []byte("1")); err != nil {
		t.Fatalf("Put a: %v", err)
	}
	if err := src.Put("b", []byte("2")); err != nil {
		t.Fatalf("Put b: %v", err)
	}

	checkpointDir := filepath.Join(baseDir, "checkpoint")
	if err := src.Checkpoint(checkpointDir); err != nil {
		t.Fatalf("Checkpoint: %v", err)
	}

	// The original database must remain open and writable after taking the
	// checkpoint, and writes afterward must not appear in the checkpoint.
	if err := src.Put("c", []byte("3")); err != nil {
		t.Fatalf("Put c after checkpoint: %v", err)
	}

	cp, err := NewRocksDB(checkpointDir)
	if err != nil {
		t.Fatalf("NewRocksDB(checkpoint): %v", err)
	}
	defer cp.Close()

	for key, want := range map[string]string{"a": "1", "b": "2"} {
		got, exists, err := cp.Get(key)
		if err != nil {
			t.Fatalf("Get(%s): %v", key, err)
		}
		if !exists {
			t.Fatalf("Get(%s): key not found in checkpoint", key)
		}
		if string(got) != want {
			t.Fatalf("Get(%s) = %q, want %q", key, got, want)
		}
	}

	if _, exists, err := cp.Get("c"); err != nil {
		t.Fatalf("Get(c): %v", err)
	} else if exists {
		t.Fatalf("Get(c): key written after checkpoint leaked into checkpoint")
	}

	// Original must still see its own post-checkpoint write.
	got, exists, err := src.Get("c")
	if err != nil {
		t.Fatalf("Get(c) on source: %v", err)
	}
	if !exists || string(got) != "3" {
		t.Fatalf("Get(c) on source = %q, %v, want 3, true", got, exists)
	}
}

func TestCheckpointRefusesExistingDir(t *testing.T) {
	baseDir := t.TempDir()

	src, err := NewRocksDB(filepath.Join(baseDir, "src"))
	if err != nil {
		t.Fatalf("NewRocksDB: %v", err)
	}
	defer src.Close()

	existing := filepath.Join(baseDir, "already-here")
	if err := src.Checkpoint(existing); err != nil {
		t.Fatalf("Checkpoint: %v", err)
	}

	if err := src.Checkpoint(existing); err == nil {
		t.Fatalf("Checkpoint into an existing directory should fail")
	}
}

func TestGetTellsAnEmptyValueFromAMissingKey(t *testing.T) {
	r, err := NewRocksDB(filepath.Join(t.TempDir(), "db"))
	if err != nil {
		t.Fatalf("NewRocksDB: %v", err)
	}
	defer r.Close()

	if err := r.Put("empty", []byte{}); err != nil {
		t.Fatalf("Put empty: %v", err)
	}
	if v, ok, err := r.Get("empty"); err != nil || !ok || len(v) != 0 {
		t.Errorf("Get(empty) = %q, %v, %v; want an empty value that exists", v, ok, err)
	}
	if _, ok, err := r.Get("missing"); err != nil || ok {
		t.Errorf("Get(missing) found = %v, err = %v; want not found", ok, err)
	}
}
