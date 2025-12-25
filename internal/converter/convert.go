package converter

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"go.leapkit.dev/core/server"
)

type conversionType string

const (
	DocxToPDF conversionType = "docx-to-pdf"
	PDFToDocx conversionType = "pdf-to-docx"
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

	conversionType := r.FormValue("conversionType")
	fmt.Println("conversionType", conversionType)
	switch conversionType {
	case string(DocxToPDF):
		convertToPdf(file, header)(w, r)
	case string(PDFToDocx):
		convertToDocx(file, header)(w, r)
	default:
		server.Errorf(w, http.StatusBadRequest, "invalid conversion type: %s", conversionType)
		return
	}

}

func convertToPdf(file multipart.File, header *multipart.FileHeader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate file extension
		ext := strings.ToLower(filepath.Ext(header.Filename))
		if ext != ".docx" && ext != ".doc" {
			server.Errorf(w, http.StatusBadRequest, "invalid file type: %s", ext)
			return
		}

		contentType := header.Header.Get("Content-Type")
		if contentType != "application/vnd.openxmlformats-officedocument.wordprocessingml.document" && contentType != "application/msword" {
			server.Errorf(w, http.StatusBadRequest, "invalid content type: %s", contentType)
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
			server.Errorf(w, http.StatusInternalServerError, "conversion error: %w", err)
			return
		}

		// Serve the converted file
		filenameWithoutExt := strings.TrimSuffix(header.Filename, filepath.Ext(header.Filename))
		pdfFilename := filenameWithoutExt + ".pdf"
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", pdfFilename))
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
}

func convertToDocx(file multipart.File, header *multipart.FileHeader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate file extension
		ext := strings.ToLower(filepath.Ext(header.Filename))
		// Validate file extension
		if !strings.HasSuffix(strings.ToLower(header.Filename), ".pdf") {
			http.Error(w, "Only PDF files are allowed", http.StatusBadRequest)
			return
		}

		contentType := header.Header.Get("Content-Type")
		if contentType != "application/pdf" {
			server.Errorf(w, http.StatusBadRequest, "invalid content type: %s", contentType)
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

		// Convert PDF to DOCX
		outputPath, err := pdfToDocx(uploadPath)
		if err != nil {
			server.Errorf(w, http.StatusInternalServerError, "conversion error: %w", err)
			return
		}

		// Serve the converted file
		filenameWithoutExt := strings.TrimSuffix(header.Filename, filepath.Ext(header.Filename))
		docxFilename := filenameWithoutExt + ".docx"
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", docxFilename))
		w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
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
}
