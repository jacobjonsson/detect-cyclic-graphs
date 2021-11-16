test/ts:
	(cd ts && npm run test)

test/go:
	(cd go && go test)

test/all: test/ts test/go
