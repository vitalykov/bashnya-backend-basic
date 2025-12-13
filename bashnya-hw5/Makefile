RUN = go run 
TEST = go test

task1: summator/squares_test.go summator/squares.go
	cd summator && $(TEST)
	cd summator && $(TEST) -race

task2: cmd/printer/main.go printer/printer.go
	$(RUN) $<

task3: safemap/map_test.go safemap/map.go
	cd safemap && $(TEST)
	cd safemap && $(TEST) -race

task4: cmd/pipeline/main.go pipeline/pipeline.go
	$(RUN) $<

BITS_TESTS = $(wildcard bits/*_test.go)
BITS_SOURCES = bits/bits_int32.go bits/bits_int64.go bits/bits32.go bits/bits64.go
task5: $(BITS_TESTS) $(BITS_SOURCES)
	cd bits && $(TEST)
	cd bits && $(TEST) -race
