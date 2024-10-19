.IGNORE = run
TRAP = trap 'exit 0' INT;

run:
	@$(TRAP)go run cmd/go-gpt-api/main.go -config config/local.yaml