package main

import (
"fmt"
"os"

```
"github.com/YOUR_USERNAME/devvault/internal/storage"
```

)

func main() {
if len(os.Args) < 2 {
printUsage()
return
}

```
command := os.Args[1]

switch command {
case "init":
	if len(os.Args) < 3 {
		fmt.Println("Usage: devvault init <project>")
		return
	}

	project := os.Args[2]

	if err := storage.Init(project); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Project %q initialized.\n", project)

case "set":
	if len(os.Args) < 4 {
		fmt.Println("Usage: devvault set <key> <value>")
		return
	}

	key := os.Args[2]
	value := os.Args[3]

	if err := storage.Set(key, value); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("%s saved.\n", key)

case "get":
	if len(os.Args) < 3 {
		fmt.Println("Usage: devvault get <key>")
		return
	}

	value, err := storage.Get(os.Args[2])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(value)

case "list":
	vars, err := storage.List()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, variable := range vars {
		fmt.Println(variable)
	}

default:
	printUsage()
}
```

}

func printUsage() {
fmt.Println("DevVault - Local environment variable manager")
fmt.Println()
fmt.Println("Usage:")
fmt.Println("  devvault init <project>")
fmt.Println("  devvault set <key> <value>")
fmt.Println("  devvault get <key>")
fmt.Println("  devvault list")
}
