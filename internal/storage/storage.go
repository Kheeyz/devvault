package storage

import (
"bufio"
"errors"
"fmt"
"os"
"path/filepath"
"strings"

```
"github.com/YOUR_USERNAME/devvault/internal/config"
```

)

func Init(project string) error {
if project == "" {
return errors.New("project name cannot be empty")
}

```
dir := filepath.Join(config.DataDirectory, project)

if err := os.MkdirAll(dir, 0700); err != nil {
	return err
}

file := filepath.Join(dir, config.EnvironmentFile)

_, err := os.OpenFile(file, os.O_CREATE, 0600)

return err
```

}

func Set(key, value string) error {
if key == "" {
return errors.New("key cannot be empty")
}

```
dir := filepath.Join(config.DataDirectory, "default")

if err := os.MkdirAll(dir, 0700); err != nil {
	return err
}

file := filepath.Join(dir, config.EnvironmentFile)

f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
if err != nil {
	return err
}
defer f.Close()

_, err = fmt.Fprintf(f, "%s=%s\n", key, value)

return err
```

}

func Get(key string) (string, error) {
vars, err := readVariables()
if err != nil {
return "", err
}

```
value, exists := vars[key]
if !exists {
	return "", fmt.Errorf("variable %q not found", key)
}

return value, nil
```

}

func List() ([]string, error) {
vars, err := readVariables()
if err != nil {
return nil, err
}

```
result := make([]string, 0, len(vars))

for key := range vars {
	result = append(result, key)
}

return result, nil
```

}

func readVariables() (map[string]string, error) {
file := filepath.Join(
config.DataDirectory,
"default",
config.EnvironmentFile,
)

```
f, err := os.Open(file)
if err != nil {
	if os.IsNotExist(err) {
		return map[string]string{}, nil
	}

	return nil, err
}
defer f.Close()

vars := make(map[string]string)

scanner := bufio.NewScanner(f)

for scanner.Scan() {
	line := strings.TrimSpace(scanner.Text())

	if line == "" || strings.HasPrefix(line, "#") {
		continue
	}

	parts := strings.SplitN(line, "=", 2)

	if len(parts) != 2 {
		continue
	}

	vars[parts[0]] = parts[1]
}

return vars, scanner.Err()
```

}
