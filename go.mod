module goffy

go 1.24.0

require (
	github.com/eduardolat/gomponents-lucide v1.4.0
	github.com/google/uuid v1.6.0
	github.com/mattn/go-sqlite3 v1.14.32
	github.com/wawandco/gomui v0.0.0-20251224202848-27c4acf35fe6
	go.leapkit.dev/core v0.1.13
	maragu.dev/gomponents v1.2.0
	maragu.dev/gomponents-htmx v0.6.1
)

require (
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/gobuffalo/flect v1.0.3 // indirect
	github.com/gobuffalo/plush/v5 v5.0.11 // indirect
	github.com/gorilla/securecookie v1.1.2 // indirect
	github.com/gorilla/sessions v1.3.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	go.antoniopagano.com/tailo v0.0.11 // indirect
	go.leapkit.dev/tools v0.1.9 // indirect
	golang.org/x/sys v0.38.0 // indirect
)

tool (
	go.antoniopagano.com/tailo
	go.leapkit.dev/tools/db
	go.leapkit.dev/tools/dev
)

// replace github.com/wawandco/gomui => ../gomui
