# Astro and Wails using TypeScript Template ♥

![Preview Template](https://raw.githubusercontent.com/DiegPS/diegps-page/main/public/images/projects/app.webp)

## Installations in your machine

you must have installed the following tools:

- [Node.js](https://nodejs.org/)
- [Go](https://golang.org/)

### Install Wails

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

## Using this template

To use this template, you must run the following command:
```bash
wails init -n test -t https://github.com/DiegPS/wails-astro-ts
```
Note: We could optionally add `-ide vscode` or `-ide goland` to the end of this command if you wanted to add IDE support.

## Development

To start the development server, you must run the following command:
```bash
wails dev
```

## Build

To build the application, you must run the following command:
```bash
wails build
```
