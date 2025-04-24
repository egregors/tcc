# tcc competition

## Problem

You have a flight network with flow capacities. Find max flow from a hub to a goals.

```mermaid
graph TD
  A22 -->|19| A13
  A13 -->|61| A16
  A6 -->|97| A16
  A16 -->|84| A12
  A18 -->|61| A23
  A28 -->|63| A19
  A16 -->|90| A10
  A12 -->|21| A19
  A12 -->|64| A5
  A2 -->|73| A28
  A18 -->|21| A2
  A22 -->|43| A5
  A22 -->|19| A18
  A9 -->|36| A11
  A19 -->|93| A15
  A4 -->|17| A2
  A9 -->|95| A0
  A12 -->|95| A4
  A7 -->|54| A8
  A16 -->|81| A22
  A14 -->|51| A24
  A0 -->|23| A2
  A21 -->|46| A27
  A22 -->|89| A2
  A21 -->|56| A28
  A4 -->|26| A11
  A14 -->|74| A23
  A26 -->|96| A25
  A9 -->|67| A5
  A11 -->|25| A24
  A0 -->|33| A14
  A3 -->|93| A20
  A11 -->|23| A2
  A21 -->|74| A22
  A16 -->|59| A21
  A19 -->|51| A8
  A24 -->|83| A6
  A15 -->|56| A7
  A10 -->|72| A28
  A8 -->|69| A14
  A28 -->|35| A29
  A11 -->|84| A3
  A19 -->|98| A29
  A13 -->|55| A29
  A14 -->|72| A9
  A26 -->|100| A3
  A28 -->|83| A26
  A27 -->|36| A12
  A3 -->|57| A29
  A10 -->|74| A15
  A11 -->|64| A5
  A9 -->|95| A28
  A15 -->|12| A22
  A17 -->|53| A24
  A20 -->|58| A28

  A24 -->|∞| SUPER
  A15 -->|∞| SUPER
  A21 -->|∞| SUPER
  A10 -->|∞| SUPER
  A28 -->|∞| SUPER
```

## Solution

The solution is a max flow algorithm. The algorithm is based on the [Edmonds-Karp algorithm](https://en.wikipedia.org/wiki/Edmonds%E2%80%93Karp_algorithm)

## Output

```shell
➜  tcc git:(main) ✗ go run main.go 
hub: A0, goals: [A10 A28 A24 A15 A21]
nodes: [A24 A27 A28 A7 A14 A6 A23 A21 A10 A4 A0 A8 A3 A2 A22 A19 A5 A17 A13 A26 A25 A20 A29 A16 A12 A18 A9 A11 A15]
 A16 |- 81 -> A22
 A16 |- 59 -> A21
 A16 |- 84 -> A12
 A16 |- 90 -> A10
  A4 |- 17 -> A2
  A4 |- 26 -> A11
 A21 |- 46 -> A27
 A21 |- 56 -> A28
 A21 |- 74 -> A22
 A10 |- 72 -> A28
 A10 |- 74 -> A15
  A8 |- 69 -> A14
 A27 |- 36 -> A12
  A0 |- 23 -> A2
  A0 |- 33 -> A14
 A22 |- 89 -> A2
 A22 |- 19 -> A13
 A22 |- 43 -> A5
 A22 |- 19 -> A18
  A6 |- 97 -> A16
 A18 |- 61 -> A23
 A18 |- 21 -> A2
 A28 |- 63 -> A19
 A28 |- 35 -> A29
 A28 |- 83 -> A26
  A2 |- 73 -> A28
 A19 |- 93 -> A15
 A19 |- 51 -> A8
 A19 |- 98 -> A29
 A24 |- 83 -> A6
 A12 |- 21 -> A19
 A12 |- 64 -> A5
 A12 |- 95 -> A4
 A14 |- 51 -> A24
 A14 |- 74 -> A23
 A14 |- 72 -> A9
  A3 |- 93 -> A20
  A3 |- 57 -> A29
 A15 |- 12 -> A22
 A15 |- 56 -> A7
 A17 |- 53 -> A24
  A7 |- 54 -> A8
 A13 |- 61 -> A16
 A13 |- 55 -> A29
  A9 |- 36 -> A11
  A9 |- 95 -> A0
  A9 |- 67 -> A5
  A9 |- 95 -> A28
 A26 |- 96 -> A25
 A26 |- 100 -> A3
 A11 |- 25 -> A24
 A11 |- 23 -> A2
 A11 |- 84 -> A3
 A11 |- 64 -> A5
 A20 |- 58 -> A28
Maximum number of passengers: 56%  
```