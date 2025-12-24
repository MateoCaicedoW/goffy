package converter

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"go.leapkit.dev/core/server"
)

func Convert(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		server.Errorf(w, http.StatusBadRequest, "file too big: %w", err)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		server.Errorf(w, http.StatusBadRequest, "error retrieving file: %w", err)
		return
	}
	defer file.Close()

	// Validate file extension
	ext := strings.ToLower(filepath.Ext(header.Filename))
	if ext != ".docx" && ext != ".doc" {
		server.Errorf(w, http.StatusBadRequest, "invalid file type: %s", ext)
		return
	}

	// Save uploaded file
	uploadPath := filepath.Join(uploadsPath, uuid.New().String()+ext)
	dst, err := os.Create(uploadPath)
	if err != nil {
		server.Errorf(w, http.StatusInternalServerError, "error saving file: %w", err)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		server.Errorf(w, http.StatusInternalServerError, "error saving file: %w", err)
		return
	}
	dst.Close()

	// Convert DOCX to PDF
	outputPath, err := docxToPDF(uploadPath)
	if err != nil {
		log.Printf("Conversion error: %v", err)
		server.Errorf(w, http.StatusInternalServerError, "conversion error: %w", err)
		return
	}

	// Serve the converted file
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filepath.Base(outputPath)))
	w.Header().Set("Content-Type", "application/pdf")

	outputFile, err := os.Open(outputPath)
	if err != nil {
		server.Errorf(w, http.StatusInternalServerError, "error opening converted file: %w", err)
		return
	}

	defer outputFile.Close()

	if _, err := io.Copy(w, outputFile); err != nil {
		server.Errorf(w, http.StatusInternalServerError, "error sending file: %w", err)
		return
	}

	// Cleanup temporary files
	go cleanupFiles(uploadPath, outputPath)
}
