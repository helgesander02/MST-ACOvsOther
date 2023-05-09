# MST-ACOvsOther
Simulate TSN and AVB data streams to find the best path in the topology.

## Installation
Clone this repo by `git clone https://github.com/helgesander02/MST-ACOvsOther`<br />

## Running
Quickstart: `go run main.go`<br />
More options:
| Option | Description |
| -------- | ---- | 
| --test_case | Conducting n experiments |
| --tsn | Number of TSN flows |
| --avb | Number of AVB flows |
| --HyperPeriod | Greatest Common Divisor of Simulated Time LCM |
| --show_topology | Display all topology information |
| --show_flows | Display all Flows information |


## Reference
[geeksforgeeks, "Steiner Tree Problem"](https://www.geeksforgeeks.org/steiner-tree/)<br />
[知乎, "viterbi"](https://www.zhihu.com/question/20136144)<br />
[Cisco, "OSPF"](https://www.cisco.com/c/zh_tw/support/docs/ip/open-shortest-path-first-ospf/7039-1.html)<br />
[geeksforgeeks, "Dijkstra"](https://www.geeksforgeeks.org/dijkstras-shortest-path-algorithm-greedy-algo-7/)<br />
[Sean Chou, "基礎演算法系列 — Graph 資料結構與Dijkstra’s Algorithm"](https://medium.com/%E6%8A%80%E8%A1%93%E7%AD%86%E8%A8%98/%E5%9F%BA%E7%A4%8E%E6%BC%94%E7%AE%97%E6%B3%95%E7%B3%BB%E5%88%97-graph-%E8%B3%87%E6%96%99%E7%B5%90%E6%A7%8B%E8%88%87dijkstras-algorithm-6134f62c1fc2)<br />

