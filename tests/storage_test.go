package tests

import (
"os"
"path/filepath"
"testing"

```
"github.com/YOUR_USERNAME/devvault/internal/storage"
```

)

func TestSetAndGet(t *testing.T) {
tempDir := t.TempDir()

```
oldDir, err := os.Getwd()
if err != nil {
	t.Fatal(err)
}

if err := os.Chdir(tempDir); err != nil {
	t.Fatal(err)
}

defer os.Chdir(oldDir)

err = storage.Set("API_KEY", "test-value")
if err != nil {
	t.Fatalf("Set() failed: %v", err)
}

value, err := storage.Get("API_KEY")
if err != nil {
	t.Fatalf("Get() failed: %v", err)
}

if value != "test-value" {
	t.Fatalf("expected test-value, got %s", value)
}
```

}

func TestGetMissingVariable(t *testing.T) {
tempDir := t.TempDir()

```
oldDir, err := os.Getwd()
if err != nil {
	t.Fatal(err)
}

if err := os.Chdir(tempDir); err != nil {
	t.Fatal(err)
}

defer os.Chdir(oldDir)

_, err = storage.Get("MISSING")

if err == nil {
	t.Fatal("expected error for missing variable")
}
```

}

func TestInit(t *testing.T) {
tempDir := t.TempDir()

```
oldDir, err := os.Getwd()
if err != nil {
	t.Fatal(err)
}

if err := os.Chdir(tempDir); err != nil {
	t.Fatal(err)
}

defer os.Chdir(oldDir)

err = storage.Init("my-project")
if err != nil {
	t.Fatalf("Init() failed: %v", err)
}

expected := filepath.Join(".devvault", "my-project", ".env")

if _, err := os.Stat(expected); err != nil {
	t.Fatalf("expected %s to exist: %v", expected, err)
}
```

}
