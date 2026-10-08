## goroutine leaks ..
## solution of the go routine leak 
Explicit Exit
Goroutine exits when its function returns.
Basic way a goroutine finishes.

work → return → exit

2. Channel Closing
Tells workers no more work is coming.
for range over the channel ends after remaining jobs are processed.

close(channel) → no more jobs → worker exits

3. Context Cancellation
Used when the caller wants to stop an operation.
Useful for timeouts, HTTP requests, DB operations, etc.

cancel() → stop operation

4. WaitGroup
Used to wait for goroutines to finish.
It does not stop goroutines.

Wait() → wait until all finish

5. errgroup
Coordinates multiple goroutines and handles errors.
Can also coordinate cancellation.

goroutines + errors + cancellation