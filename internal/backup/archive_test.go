package backup

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"os"
	"path/filepath"
	"testing"
)

func TestTarGzRoundTrip(t *testing.T) {
	srcDir := t.TempDir()

	mustWriteFile(t, filepath.Join(srcDir, "top.txt"), "hello world", 0o644)
	mustWriteFile(t, filepath.Join(srcDir, "empty.txt"), "", 0o644)
	if err := os.MkdirAll(filepath.Join(srcDir, "nested", "deeper"), 0o755); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}
	mustWriteFile(t, filepath.Join(srcDir, "nested", "deeper", "leaf.txt"), "leaf contents", 0o600)

	var buf bytes.Buffer
	if err := tarGzDir(srcDir, &buf); err != nil {
		t.Fatalf("tarGzDir: %v", err)
	}

	destDir := t.TempDir()
	if err := untarGz(&buf, destDir); err != nil {
		t.Fatalf("untarGz: %v", err)
	}

	cases := []struct {
		rel     string
		content string
		mode    os.FileMode
	}{
		{"top.txt", "hello world", 0o644},
		{"empty.txt", "", 0o644},
		{"nested/deeper/leaf.txt", "leaf contents", 0o600},
	}

	for _, c := range cases {
		path := filepath.Join(destDir, c.rel)
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("ReadFile(%s): %v", c.rel, err)
		}
		if string(data) != c.content {
			t.Fatalf("content of %s = %q, want %q", c.rel, data, c.content)
		}

		info, err := os.Stat(path)
		if err != nil {
			t.Fatalf("Stat(%s): %v", c.rel, err)
		}
		if info.Mode().Perm() != c.mode {
			t.Fatalf("mode of %s = %o, want %o", c.rel, info.Mode().Perm(), c.mode)
		}
	}
}

func TestUntarGzRejectsPathTraversal(t *testing.T) {
	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gw)

	if err := tw.WriteHeader(&tar.Header{
		Name:     "../escape.txt",
		Typeflag: tar.TypeReg,
		Mode:     0o644,
		Size:     int64(len("gotcha")),
	}); err != nil {
		t.Fatalf("WriteHeader: %v", err)
	}
	if _, err := tw.Write([]byte("gotcha")); err != nil {
		t.Fatalf("Write: %v", err)
	}
	if err := tw.Close(); err != nil {
		t.Fatalf("tw.Close: %v", err)
	}
	if err := gw.Close(); err != nil {
		t.Fatalf("gw.Close: %v", err)
	}

	destDir := t.TempDir()
	if err := untarGz(&buf, destDir); err == nil {
		t.Fatalf("untarGz should reject a path-traversal entry")
	}

	escaped := filepath.Join(filepath.Dir(destDir), "escape.txt")
	if _, err := os.Stat(escaped); !os.IsNotExist(err) {
		t.Fatalf("path-traversal entry should not be written outside destDir, stat err = %v", err)
	}

	entries, err := os.ReadDir(destDir)
	if err != nil {
		t.Fatalf("ReadDir(destDir): %v", err)
	}
	if len(entries) != 0 {
		t.Fatalf("destDir should remain empty, got %v", entries)
	}
}

func mustWriteFile(t *testing.T, path, content string, mode os.FileMode) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), mode); err != nil {
		t.Fatalf("WriteFile(%s): %v", path, err)
	}
}
