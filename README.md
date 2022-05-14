# gqlgen-memory-leak


```bash
cd ./test && go run .


Goroutines leakage:

count: 1	goroutines num: 2	Alloc = 0 MiB	TotalAlloc = 0 MiB	Sys = 8 MiB	NumGC = 0
count: 19	goroutines num: 44	Alloc = 1 MiB	TotalAlloc = 1 MiB	Sys = 8 MiB	NumGC = 0
count: 37	goroutines num: 80	Alloc = 2 MiB	TotalAlloc = 2 MiB	Sys = 12 MiB	NumGC = 0
count: 56	goroutines num: 118	Alloc = 1 MiB	TotalAlloc = 3 MiB	Sys = 13 MiB	NumGC = 1
count: 75	goroutines num: 156	Alloc = 2 MiB	TotalAlloc = 4 MiB	Sys = 13 MiB	NumGC = 1
count: 94	goroutines num: 194	Alloc = 2 MiB	TotalAlloc = 6 MiB	Sys = 13 MiB	NumGC = 2
count: 113	goroutines num: 232	Alloc = 3 MiB	TotalAlloc = 7 MiB	Sys = 13 MiB	NumGC = 2
count: 132	goroutines num: 270	Alloc = 3 MiB	TotalAlloc = 8 MiB	Sys = 13 MiB	NumGC = 3
count: 151	goroutines num: 308	Alloc = 3 MiB	TotalAlloc = 10 MiB	Sys = 13 MiB	NumGC = 4
count: 170	goroutines num: 346	Alloc = 3 MiB	TotalAlloc = 11 MiB	Sys = 17 MiB	NumGC = 5
```