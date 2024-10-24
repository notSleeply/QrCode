module main

go 1.22.4

replace qrcode => ./qrcode

require (
	github.com/google/uuid v1.3.0
	github.com/maruel/rs v1.1.0
)

require (
	qrcode v0.0.0-00010101000000-000000000000 // indirect
	rsc.io/qr v0.2.0 // indirect
)
