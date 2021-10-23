- Garbage collection happens automatically.
- Out of scope for all.
- The GOGC variable sets the initial garbage collection target percentage. A collection is triggered when the ratio of freshly allocated data to live data remaining after the previous collection reaches this percentage. The default is GOFC=100. Setting GOGC=off disables the garbage collector entirely. The runtime/debug package's SetGCPercent function allows changing this percentage at run time.

## There are two ways to manage memory:
- new() -->
    - Allocate memory but no INIT (no initialized)
    - you will get a memory address
    - zeroed storage --> you cannot put any data initially
- make() -->
    - Allocate memory and INIT (initialized)
    - you will get memory address
    - non-zeroed storage --> can put data initially

