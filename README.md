# experiment-goroutines

The experiment tests two different ways to map a slice.
The first is using gorountines and the second using a simple loop.

> :warning: This is just an example of how goroutines work. Don't focus on the quality of the algorithm


Results (1'000'000 items):

| function       | exec time (ms) |
| -------------- | -------------- |
| mapSlice       | 390.6552       |
| simpleMapSlice | 13.5139        |