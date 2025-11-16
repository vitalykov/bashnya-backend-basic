package safemap

import (
	"sync"
	"testing"
)

// tests basic operations in sequential mode
func TestSafeMap_Sequential(t *testing.T) {
	sm := NewSafeMap[string, int]()

	sm.Store("key1", 100)
	if val, ok := sm.Get("key1"); !ok || val != 100 {
		t.Errorf("Get() = (%v, %v), want (100, true)", val, ok)
	}

	if val, ok := sm.Get("nonexistent"); ok || val != 0 {
		t.Errorf("Get(nonexistent) = (%v, %v), want (0, false)", val, ok)
	}

	sm.Delete("key1")
	if val, ok := sm.Get("key1"); ok {
		t.Errorf("Get() after Delete = (%v, %v), want (0, false)", val, ok)
	}

	sm.Store("key2", 200)
	sm.Store("key2", 300)
	if val, ok := sm.Get("key2"); !ok || val != 300 {
		t.Errorf("Get() after overwrite = (%v, %v), want (300, true)", val, ok)
	}
}

// tests concurrent reads and writes
func TestSafeMap_ConcurrentReadWrite(t *testing.T) {
	sm := NewSafeMap[int, string]()
	const numGoroutines = 100
	const numOperations = 1000

	var wg sync.WaitGroup

	// Writer goroutines
	for i := range numGoroutines {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := range numOperations {
				key := id*numOperations + j
				sm.Store(key, "value")
				sm.Get(key)
				sm.Delete(key)
			}
		}(i)
	}

	// Reader goroutines that read while others write
	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range numOperations {
				sm.Get(j)
			}
		}()
	}

	wg.Wait()
}

// specifically tests for data races
func TestSafeMap_DataRaceDetection(t *testing.T) {
	sm := NewSafeMap[string, int]()
	const numGoroutines = 50

	var wg sync.WaitGroup

	// Start multiple goroutines that heavily contend on the same keys
	for i := range numGoroutines {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()
			for j := range 1000 {
				key := "contended_key"
				sm.Store(key, goroutineID)
				sm.Get(key)
				sm.Store(key, goroutineID+j)
				sm.Get(key)
				sm.Delete(key)
				sm.Get(key)
			}
		}(i)
	}

	wg.Wait()
}

// tests concurrent stores to same key
func TestSafeMap_ConcurrentStore(t *testing.T) {
	sm := NewSafeMap[string, int]()
	const numGoroutines = 100

	var wg sync.WaitGroup

	// All goroutines store to the same key
	for i := range numGoroutines {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			sm.Store("same_key", val)
		}(i)
	}

	wg.Wait()

	// At least one value should be stored successfully
	if val, ok := sm.Get("same_key"); !ok {
		t.Error("Expected some value to be stored for 'same_key'")
	} else {
		t.Logf("Final value for 'same_key': %d", val)
	}
}

// tests race between Get and Delete
func TestSafeMap_ConcurrentGetDelete(t *testing.T) {
	sm := NewSafeMap[int, string]()
	for i := range 100 {
		sm.Store(i, "value")
	}

	var wg sync.WaitGroup

	// Reader goroutines
	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range 1000 {
				sm.Get(j % 100)
			}
		}()
	}

	// Deleter goroutines
	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range 1000 {
				sm.Delete(j % 100)
				// Immediately re-store to keep contention high
				sm.Store(j%100, "new_value")
			}
		}()
	}

	wg.Wait()
}

// tests with different types
func TestSafeMap_TypeSafety(t *testing.T) {
	sm1 := NewSafeMap[string, []int]()
	sm1.Store("slice", []int{1, 2, 3})
	if val, ok := sm1.Get("slice"); !ok || len(val) != 3 {
		t.Errorf("Get(slice) failed: %v, %v", val, ok)
	}

	type Person struct {
		Name string
		Age  int
	}
	sm2 := NewSafeMap[int, Person]()
	sm2.Store(1, Person{"Alice", 30})
	if val, ok := sm2.Get(1); !ok || val.Name != "Alice" {
		t.Errorf("Get(struct) failed: %v, %v", val, ok)
	}

	sm3 := NewSafeMap[any, string]()
	sm3.Store(123, "int_key")
	sm3.Store("abc", "string_key")
	if val, ok := sm3.Get(123); !ok || val != "int_key" {
		t.Errorf("Get(interface key) failed: %v, %v", val, ok)
	}
}

// tests under extreme load
func TestSafeMap_Stress(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	sm := NewSafeMap[int, int]()
	const numGoroutines = 1000
	const operationsPerGoroutine = 10000

	var wg sync.WaitGroup
	start := make(chan struct{})

	for i := range numGoroutines {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()
			<-start // Wait for all goroutines to be ready

			for j := range operationsPerGoroutine {
				key := (goroutineID * operationsPerGoroutine) + j

				// Mix of operations
				sm.Store(key, key*2)
				sm.Get(key)
				if j%3 == 0 {
					sm.Delete(key)
				}
			}
		}(i)
	}

	// Start all goroutines at once to maximize contention
	close(start)
	wg.Wait()

	// Verify no panics occurred and basic integrity
	t.Log("Stress test completed without panics")
}

// TestSafeMap_EmptyMap tests operations on empty map
func TestSafeMap_EmptyMap(t *testing.T) {
	sm := NewSafeMap[string, int]()

	// Get from empty map
	if val, ok := sm.Get("any"); ok {
		t.Errorf("Get from empty map returned (%v, %v), expected (0, false)", val, ok)
	}

	// Delete from empty map (should not panic)
	sm.Delete("any")

	// Concurrent operations on empty map
	var wg sync.WaitGroup
	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sm.Get("key")
			sm.Delete("key")
		}()
	}
	wg.Wait()
}

func BenchmarkSafeMap_ConcurrentReadWrite(b *testing.B) {
	sm := NewSafeMap[int, int]()
	b.RunParallel(func(pb *testing.PB) {
		counter := 0
		for pb.Next() {
			sm.Store(counter, counter)
			sm.Get(counter)
			counter++
		}
	})
}

func BenchmarkSafeMap_ConcurrentReads(b *testing.B) {
	sm := NewSafeMap[int, int]()
	// Pre-populate
	for i := range 1000 {
		sm.Store(i, i)
	}

	b.RunParallel(func(pb *testing.PB) {
		counter := 0
		for pb.Next() {
			sm.Get(counter % 1000)
			counter++
		}
	})
}
