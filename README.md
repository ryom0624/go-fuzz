
# Fuzz test tutorial
https://go.dev/doc/tutorial/fuzz


# Execute test with Fuzzing

```shell
$ go test -fuzz=Fuzz -fuzztime 30s

fuzz: elapsed: 0s, gathering baseline coverage: 0/45 completed
fuzz: elapsed: 0s, gathering baseline coverage: 45/45 completed, now fuzzing with 8 workers
fuzz: elapsed: 3s, execs: 1327231 (442404/sec), new interesting: 13 (total: 58)
fuzz: elapsed: 6s, execs: 2679769 (450829/sec), new interesting: 17 (total: 62)
fuzz: elapsed: 9s, execs: 3891523 (403865/sec), new interesting: 18 (total: 63)
fuzz: elapsed: 12s, execs: 5208004 (438818/sec), new interesting: 20 (total: 65)
fuzz: elapsed: 15s, execs: 6463336 (418385/sec), new interesting: 20 (total: 65)
fuzz: elapsed: 18s, execs: 7730156 (422386/sec), new interesting: 21 (total: 66)
fuzz: elapsed: 21s, execs: 8979770 (416426/sec), new interesting: 22 (total: 67)
fuzz: elapsed: 24s, execs: 10256404 (425631/sec), new interesting: 22 (total: 67)
fuzz: elapsed: 27s, execs: 11539229 (427620/sec), new interesting: 23 (total: 68)
fuzz: elapsed: 30s, execs: 12831808 (430881/sec), new interesting: 24 (total: 69)
fuzz: elapsed: 30s, execs: 12831808 (0/sec), new interesting: 24 (total: 69)
PASS
ok      _/path/to/go-fuzz 30.149s

```