package converter

import (
	"goffy/internal/system/assets"
	"goffy/internal/system/helpers"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/wawandco/gomui"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

type layout struct {
	Title       string
	Description string
	Yield       Node
}

func (props layout) New() Node {
	return HTML5(
		HTML5Props{
			Title:       props.Title,
			Description: props.Description,
			Language:    "en",
			Head: []Node{

				Link(Rel("stylesheet"), Href(assets.Manager.Path("/public/application.css"))),
				Link(Rel("stylesheet"), Href(assets.Manager.Path("/public/basecoatui.css"))),

				Link(Rel("proconnect"), Href("https://fonts.googleapis.com")),
				Link(Rel("preconnect"), Href("https://fonts.gstatic.com"), Attr("crossorigin", "")),
				Link(Rel("stylesheet"), Href("https://fonts.googleapis.com/css2?family=Dancing+Script:wght@500&display=swap")),
				Raw(helpers.Importmap()),
				Script(Src(assets.Manager.Path("/public/basecoatui.js")), Defer()),
				Script(Src(assets.Manager.Path("/public/basecoattoast.js")), Defer()),
				Script(Src(assets.Manager.Path("/public/htmx.js")), Defer()),
				Script(Src(assets.Manager.Path("/public/htmx-download.js")), Defer()),
			},
			Body: []Node{

				Div(
					Class("min-h-screen flex flex-col bg-base-100 text-base-content"),
					props.Yield,
				),

				gomui.ThemeToggle("theme-toggle", lucide.Sun(Class("size-5")), Class("btn-icon-outline size-8 absolute top-5 right-5")),
				gomui.DarkModeScript(),
			},
		},
	)
}
