# Markdown Preview Tool

A command-line tool that converts Markdown files into HTML documents.

## Usage

```bash
go run main.go -in <input.md> [-out <output_name>]
```

### Flags

| Flag  | Required | Description |
|-------|----------|-------------|
| `-in` | Yes      | Path to the input Markdown file |
| `-out` | No      | Name for the output HTML file (without extension). If omitted, the output file is named after the input file (e.g. `README.md` → `README.md.html`) |

### Examples

Convert `README.md` and specify the output name:

```bash
go run main.go -in README.md -out readme
# Creates: readme.html
```

Convert `README.md` using the default output name:

```bash
go run main.go -in README.md
# Creates: README.md.html
```
