# DevVault 

DevVault is a simple command-line tool for managing environment variables for local development projects.

## Features

* Initialize a project
* Store environment variables
* List stored variables
* Retrieve individual variables
* Export variables to a `.env` file
* Keep sensitive values out of Git

## Requirements

* Go 1.22+

## Installation

Clone the repository:

```bash
git clone https://github.com/YOUR_USERNAME/devvault.git
cd devvault
```

Run the application:

```bash
go run .
```

## Usage

Initialize a project:

```bash
devvault init my-project
```

Add a variable:

```bash
devvault set API_KEY "your-secret-key"
```

List variables:

```bash
devvault list
```

Get a variable:

```bash
devvault get API_KEY
```

Export variables:

```bash
devvault export
```

## Project Structure

```text
devvault/
├── main.go
├── internal/
│   ├── config/
│   └── storage/
└── tests/
```

## Security

DevVault is intended for local development. Never commit real secrets to GitHub.

Use `.env.example` to document required environment variables without exposing their values.

## License

MIT
