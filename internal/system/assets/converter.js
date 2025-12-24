import { Controller } from "stimulus"

export default class extends Controller {
    static targets = ['fileInput', 'dropzone', 'filePreview', 'fileName', 'fileSize', 'pdfToDocxBtn', 'docxToPdfBtn'];
    static values = { type: { type: String, default: 'pdf-to-docx' } };

    selectConversion(event) {
        this.typeValue = event.currentTarget.dataset.type;
        this.updateAcceptedFormats();
        this.clearFile();
    }

    triggerFileInput() {
        this.fileInputTarget.click();
    }

    handleFileSelect(event) {
        const file = event.target.files[0];
        if (file) {
            this.displayFile(file);
        }
    }

    displayFile(file) {
        this.fileNameTarget.textContent = file.name;
        this.fileSizeTarget.textContent = this.formatFileSize(file.size);
        this.filePreviewTarget.classList.remove('hidden');
    }

    clearFile() {
        this.fileInputTarget.value = '';
        this.filePreviewTarget.classList.add('hidden');
    }

    formatFileSize(bytes) {
        if (bytes === 0) return '0 Bytes';
        const k = 1024;
        const sizes = ['Bytes', 'KB', 'MB', 'GB'];
        const i = Math.floor(Math.log(bytes) / Math.log(k));
        return Math.round(bytes / Math.pow(k, i) * 100) / 100 + ' ' + sizes[i];
    }
}