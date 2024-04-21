module github.com/janmmiranda/go_adventures

go 1.22.2

require github.com/peterh/liner v1.2.2

require (
	github.com/mattn/go-runewidth v0.0.3 // indirect
	golang.org/x/sys v0.0.0-20211117180635-dee7805ff2e1 // indirect
)

require github.com/janmmiranda/go_adventures/internal/common v0.0.0

replace github.com/janmmiranda/go_adventures/internal/common => ./internal/common

require github.com/janmmiranda/go_adventures/internal/dices v0.0.0

replace github.com/janmmiranda/go_adventures/internal/dices => ./internal/dices

require github.com/janmmiranda/go_adventures/internal/entities v0.0.0

replace github.com/janmmiranda/go_adventures/internal/entities => ./internal/entities
