package converter

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/wawandco/gomui"
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

func indexEl() Node {
	return Div(
		Data("controller", "converter"),
		Class("container mx-auto px-4 py-8 max-w-4xl"),

		// Main Card
		gomui.Card(
			gomui.CardHeader(
				H1(Text("Conversion Type")),
				P(Text("Select the type of conversion you want to perform")),
			),

			gomui.CardContent(
				// Conversion Direction Selector
				Div(
					Class("mb-8"),
					gomui.Tabs(
						Input(Type("hidden"), Name("conversionType"), Value("docx-to-pdf"), Data("converter-target", "conversionTypeInput")),
						gomui.TabsList(
							gomui.TabItem("pdf-to-docx",
								true,
								Text("PDF → DOCX"),
								Data("type", "pdf-to-docx"),
								Data("converter-target", "pdfToDocxBtn"),
								Data("action", "click->converter#selectConversion"),
							),
							gomui.TabItem("pdf-to-docx",
								false,
								Text("DOCX → PDF"),
								Data("type", "docx-to-pdf"),
								Data("converter-target", "docxToPdfBtn"),
								Data("action", "click->converter#selectConversion"),
							),
						),
					),
				),
				// Upload Area
				Div(
					Input(
						Type("file"),
						Name("file"),
						Data("converter-target", "fileInput"),
						Data("action", "change->converter#handleFileSelect"),
						Class("hidden"),
						Accept(".docx,application/vnd.openxmlformats-officedocument.wordprocessingml.document"),
					),
					Div(
						Data("converter-target", "dropzone"),
						Data("action", "click->converter#triggerFileInput"),
						Class("border-3 border-dashed rounded-xl p-12 text-center  transition-all cursor-pointer"),

						H3(
							Class("text-xl font-semibold mb-2"),
							Text("Drop your file here"),
						),
						P(
							Class("mb-4 text-sm"),
							Text("or click to browse"),
						),

						gomui.ButtonEl(gomui.ButtonSecondary, gomui.ButtonSm, false, Text("Select File")),
						P(
							Class("text-xs  mt-4"),
							Text("Supported formats:"),
							Span(
								Data("converter-target", "supportedFormat"),
								Class("font-medium"),
								Text("PDF"),
							),
						),
					),
				),

				// File Preview Area
				Div(
					Data("converter-target", "filePreview"),
					Class("hidden mt-6 p-4  rounded-lg border"),
					Div(
						Class("flex items-center justify-between"),
						Div(
							Class("flex items-center gap-3"),
							Div(
								Class("w-12 h-12 rounded-lg flex items-center justify-center card !p-0"),
								lucide.File(Class("size-5")),
							),
							Div(
								P(
									Data("converter-target", "fileName"),
									Class("font-semibold "),
									Text("document.pdf"),
								),
								P(
									Data("converter-target", "fileSize"),
									Class("text-sm "),
									Text("2.4 MB"),
								),
							),
						),
						gomui.ButtonEl(
							gomui.ButtonGhost,
							gomui.ButtonSm,
							false,
							Data("action", "click->converter#clearFile"),
							Class("text-red-500 hover:text-red-700 transition-colors"),
							lucide.X(Class("size-5")),
						),
					),
				),
			),

			gomui.CardFooter(
				gomui.ButtonEl(
					gomui.ButtonPrimary+" w-full",
					gomui.ButtonLg,
					false,
					Text("Convert File"),
					hx.Post("/convert"),
					hx.Include("[name='file'], [name='conversionType']"),
					hx.Encoding("multipart/form-data"),
					hx.Ext("htmx-download"),
					hx.DisabledElt("this"),
					hx.Indicator("#download-indicator"),
					lucide.LoaderCircle(ID("download-indicator"), Class("loading-indicator size-5 animate-spin")),
				),
			),
		),
	)
}
