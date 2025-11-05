UNIQ_MAIN = cmd/uniq/main.go
UNIQ_EXE = uniq
ACCEPTANCE_TESTS_SH = acceptance-tests/tests.sh
TEST_FILES = acceptance-tests/user/*.txt

.PHONY: all
all: build

.PHONY: build
build: $(UNIQ_EXE)

$(UNIQ_EXE): $(UNIQ_MAIN)
	go build -o $@ $^

.PHONY: run
run: $(UNIQ_MAIN)
	go run $^

.PHONY: test
test: $(UNIQ_EXE)
	go test ./... -v
	bash $(ACCEPTANCE_TESTS_SH)

.PHONY: clean
clean:
	rm -f $(UNIQ_EXE)
	rm -rf $(TEST_FILES)
