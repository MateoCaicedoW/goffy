package converter

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	maxUploadSize = 50 << 20 // 50 MB
	uploadsPath   = "./uploads"
	outputsPath   = "./outputs"
)

func init() {
	// Create necessary directories
	dirs := []string{uploadsPath, outputsPath}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			panic(fmt.Sprintf("failed to create directory %s: %v", dir, err))
		}
	}

}

// DocxToPDF converts DOCX to PDF format
func docxToPDF(inputPath string) (string, error) {
	outputFilename := strings.TrimSuffix(filepath.Base(inputPath), filepath.Ext(inputPath)) + ".pdf"
	outputPath := filepath.Join(outputsPath, outputFilename)

	// Use LibreOffice for conversion
	cmd := exec.Command("libreoffice",
		"--headless",
		"--convert-to", "pdf",
		"--outdir", outputsPath,
		inputPath,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("conversion failed: %w, output: %s", err, string(output))
	}

	// Verify output file exists
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		return "", fmt.Errorf("output file not created")
	}

	return outputPath, nil
}

func pdfToDocx(inputPath string) (string, error) {
	outputFilename := strings.TrimSuffix(filepath.Base(inputPath), filepath.Ext(inputPath)) + ".docx"
	outputPath := filepath.Join(outputsPath, outputFilename)

	// Get absolute paths
	absInput, err := filepath.Abs(inputPath)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path for input: %w", err)
	}

	absOutput, err := filepath.Abs(outputPath)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path for output: %w", err)
	}

	// Run pdf2docx conversion
	cmd := exec.Command("pdf2docx", "convert", absInput, absOutput)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("conversion command failed: %w", err)
	}

	return outputPath, nil
}

// cleanupFiles removes temporary files after a delay
func cleanupFiles(paths ...string) {
	time.Sleep(5 * time.Second)
	for _, path := range paths {
		if err := os.Remove(path); err != nil {
			log.Printf("Failed to remove file %s: %v", path, err)
		}
	}
}
