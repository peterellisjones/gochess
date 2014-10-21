gochess
=======

Chess Engine in Go. Currently missing evaluation and search (move generation seems correct though).

To run the tests:

```
$ ginkgo -r

[1413920538] Bitboard Suite - 2/2 specs •• SUCCESS! 90.346us PASS
[1413920538] Board Suite - 11/11 specs ••••••••••• SUCCESS! 335.337us PASS
[1413920538] Fen Suite - 7/7 specs ••••••• SUCCESS! 124.588us PASS
[1413920538] Make Suite - 30/30 specs •••••••••••••••••••••••••••••• SUCCESS! 847.587us PASS
[1413920538] Move Suite - 21/21 specs ••••••••••••••••••••• SUCCESS! 379.52us PASS
[1413920538] Movegeneration Suite - 283/283 specs ••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••••• SUCCESS! 7.846083ms PASS
[1413920538] Movelist Suite - 0/0 specs  SUCCESS! 34.16us PASS
[1413920538] Perft Suite - 9/9 specs ••••••••• SUCCESS! 585.861778ms PASS
[1413920538] Stack Suite - 2/2 specs •• SUCCESS! 288.157us PASS
[1413920538] Traverse Suite - 1/1 specs • SUCCESS! 3.332492ms PASS

Ginkgo ran 10 suites in 1.972268982s
Test Suite Passed

```
