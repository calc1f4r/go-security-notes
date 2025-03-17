# Go Security and Best Practices Notes

## Runtime Behavior Checklist

### Panics on Division by Zero
- Go panics when attempting to divide by zero, which can crash the program if not handled
- This behavior is consistent across all numeric types
- Always check for zero divisors:
  ```go
  if y == 0 {
      return 0, errors.New("division by zero")
  }
  ```

### No Overflow/Underflow Checks
- Integer types don't have runtime checks for overflow or underflow
- Example: `int16(32767) + 1` results in `-32768` due to wrapping
- High risk in financial or scientific applications
- Manually check bounds before operations:
  ```go
  if value > math.MaxInt16 - addend {
      return errors.New("operation would overflow")
  }
  ```

### String Length Measurement
- `len()` returns bytes, not characters for strings
- Example: `len("hello😊")` returns `8` (not 6) because emoji uses 4 bytes
- For character count, use:
  ```go
  utf8.RuneCountInString(str)  // or
  count := 0
  for range str { count++ }
  ```

### Slice References to Underlying Arrays
- Slices are views into underlying arrays
- Modifying one slice affects others referencing the same array
- Create independent copies using:
  ```go
  independentCopy := make([]int, len(original))
  copy(independentCopy, original)
  // or
  independentCopy := append([]int{}, original...)
  ```

### Maps Are Unordered
- Maps do not maintain insertion order
- Iteration order is unpredictable
- For ordered operations, use a slice of keys and sort it:
  ```go
  var keys []string
  for k := range myMap {
      keys = append(keys, k)
  }
  sort.Strings(keys)
  for _, k := range keys {
      value := myMap[k]
      // Process in order
  }
  ```

### Concurrency Safety
- Multiple goroutines require synchronization
- Use mutex to protect shared resources:
  ```go
  var mu sync.Mutex
  mu.Lock()
  // Access shared resource
  mu.Unlock()
  ```
- Detect race conditions: `go test -race ./...`

### Explicit Error Handling
- Always check returned errors
- Don't ignore with `_` unless appropriate
- Standard pattern:
  ```go
  result, err := someFunction()
  if err != nil {
      return nil, fmt.Errorf("context: %w", err)
  }
  ```

### Resource Management with Defer
- Use `defer` for cleanup operations
- Remember LIFO execution order
- Example:
  ```go
  file, err := os.Open("file.txt")
  if err != nil {
      return err
  }
  defer file.Close()
  // Work with file...
  ```

### Panic and Recovery
- Use for truly exceptional conditions
- Recover in deferred functions:
  ```go
  defer func() {
      if r := recover(); r != nil {
          log.Printf("Recovered from panic: %v", r)
      }
  }()
  ```

### Type Safety
- Explicit conversions required between numeric types
- Example: `float64(intValue)` not automatic
- Prevents unexpected behavior from implicit conversions

### String Operations
- Strings are immutable
- Use builders for efficient concatenation:
  ```go
  var b strings.Builder
  b.WriteString("Hello")
  b.WriteString(", ")
  b.WriteString("World")
  result := b.String()
  ```

### Channel Usage
- Unbuffered: Synchronous, blocks until both sides ready
- Buffered: Asynchronous up to capacity
- Close to signal completion:
  ```go
  close(ch) // Signal no more values
  
  // Receiver can check:
  v, ok := <-ch
  if !ok {
      // Channel closed
  }
  ```

### Goroutine Management
- Goroutines must be explicitly terminated
- Use context for cancellation:
  ```go
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel() // Ensure cancellation on function exit
  
  go func(ctx context.Context) {
      select {
      case <-ctx.Done():
          return // Exit goroutine
      case <-time.After(time.Second):
          // Do work
      }
  }(ctx)
  ```

### Switch Statement Behavior
- Cases do not fall through by default
- Use `fallthrough` keyword when needed:
  ```go
  switch value {
  case 1:
      fmt.Println("One")
      // No fallthrough, stops here
  case 2:
      fmt.Println("Two")
      fallthrough // Explicitly continue to next case
  case 3:
      fmt.Println("Printed for both 2 and 3")
  }
  ```

### Safe Type Assertions
- Use comma-ok idiom to prevent panics:
  ```go
  value, ok := someInterface.(string)
  if !ok {
      // Handle type mismatch
  }
  ```


  ### While using the make keyword to create a slice you have to explicity tell the capacity
  