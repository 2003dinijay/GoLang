# GoLang Learning Workspace

This repository contains small Go programs organized by topic.

## Project Structure

- `01_first-program` - First Go program
- `02_variables_and_types` - Variables and basic types
- `03_packages_imports` - Packages and imports
- `04_var_vs_short_declare` - `var` vs short declaration (`:=`)
- `05_basic_types_string` - String basics
- `06_basic_types_int` - Integer basics

## Initial Setup

1. Install Go from the official site: https://go.dev/dl/
2. Verify installation:

```bash
go version
```

3. From the project root, verify the module file exists:

```bash
ls go.mod
```

## Run a Program

From the project root, run any lesson folder with:

```bash
go run ./01_first-program
```

You can replace the folder name with any other lesson:

```bash
go run ./02_variables_and_types
go run ./03_packages_imports
go run ./04_var_vs_short_declare
go run ./05_basic_types_string
go run ./06_basic_types_int
```

## Optional: Build a Program

To compile a lesson instead of running directly:

```bash
go build ./01_first-program
```

## Learning Tip

Edit a lesson's `main.go`, save, and run it again with `go run` to quickly test changes.

## Troubleshooting

If you get `zsh: command not found: go` on macOS, Go is usually installed but not in your PATH.

1. Add Go to your shell PATH:

```bash
echo 'export PATH="/usr/local/go/bin:$PATH"' >> ~/.zshrc
```

2. Reload terminal config:

```bash
source ~/.zshrc
```

3. Verify again:

```bash
go version
```
