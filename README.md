# Goffy

A modern web application built with Go for document conversion and processing, featuring:

- **Backend**: Built with Go using the Leapkit framework for robust server-side functionality
- **Frontend**: Interactive UI powered by HTMX for dynamic, responsive interactions
- **Styling**: Tailwind CSS for utility-first styling and modern design
- **Document Processing**: LibreOffice integration for document conversion and manipulation
- **Component System**: Gomponents and Gomui for composable, type-safe UI components
- **Asset Management**: Integrated asset pipeline with CSS and JavaScript bundling

## Features

- Document conversion utilities leveraging LibreOffice
- Dynamic server-side rendering with HTMX integration
- RESTful HTTP API
- Toast notifications and UI components from Basecoat
- Stimulus.js for interactive components
- File upload and download capabilities

## Getting Started

### Requirements
- Go 1.24.0 or later
- LibreOffice (for document processing)
- Tailwind CSS (binary)

## Tailwind CSS Installation
```bash
go tool tailo download -v v4.0.6 
```

## LibreOffice Installation
# Ubuntu/Debian
```bash
sudo apt-get install libreoffice
```

# macOS
```bash
brew install --cask libreoffice
```

# Windows
Download from https://www.libreoffice.org/download/

### Installation

```bash
go mod download
```

### Running the Application

```bash
go tool dev --watch.extensions=.go,.css,.js  
```


The server will start on `http://localhost:3000` by default.

### Environment Variables

- `HOST`: Server host (default: `0.0.0.0`)
- `PORT`: Server port (default: `3000`)
- `SESSION_SECRET`: Session secret for security (default: auto-generated)
