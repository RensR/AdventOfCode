package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var tests = execute.TestCases{
	{
		puzzleTest,
		`1651`,
		`1707`,
	},
	//{
	//	puzzle,
	//	`1673`,
	//	``,
	//},
} // 2177 low

var puzzleTest = `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II`

var puzzle = `Valve DB has flow rate=0; tunnels lead to valves AC, UN
Valve LC has flow rate=6; tunnels lead to valves UV, CM, RD, IM, YQ
Valve SU has flow rate=0; tunnels lead to valves OH, BX
Valve JS has flow rate=0; tunnels lead to valves GR, RW
Valve BX has flow rate=18; tunnels lead to valves PA, SU
Valve WI has flow rate=0; tunnels lead to valves GR, JI
Valve YQ has flow rate=0; tunnels lead to valves LC, SB
Valve HX has flow rate=10; tunnels lead to valves VR, GZ, ID
Valve RI has flow rate=0; tunnels lead to valves HF, UV
Valve JQ has flow rate=0; tunnels lead to valves AA, IF
Valve RK has flow rate=0; tunnels lead to valves AA, CM
Valve AC has flow rate=0; tunnels lead to valves DB, HF
Valve JI has flow rate=12; tunnels lead to valves WI, YH, ND, ID
Valve DF has flow rate=0; tunnels lead to valves JW, AA
Valve PA has flow rate=0; tunnels lead to valves BX, RB
Valve OU has flow rate=0; tunnels lead to valves OH, XM
Valve YO has flow rate=0; tunnels lead to valves OK, HF
Valve YY has flow rate=0; tunnels lead to valves UN, MC
Valve OJ has flow rate=0; tunnels lead to valves SC, GR
Valve VR has flow rate=0; tunnels lead to valves IR, HX
Valve EY has flow rate=0; tunnels lead to valves HR, OK
Valve LE has flow rate=0; tunnels lead to valves GV, GZ
Valve HF has flow rate=14; tunnels lead to valves DS, YO, AC, RI, WP
Valve OM has flow rate=0; tunnels lead to valves DS, GV
Valve JW has flow rate=0; tunnels lead to valves UN, DF
Valve OK has flow rate=9; tunnels lead to valves IF, EY, OV, YO, WM
Valve RB has flow rate=0; tunnels lead to valves PA, XM
Valve HR has flow rate=0; tunnels lead to valves EY, CQ
Valve YM has flow rate=0; tunnels lead to valves GB, NB
Valve UN has flow rate=5; tunnels lead to valves RD, DB, JW, YY, WC
Valve SO has flow rate=0; tunnels lead to valves AA, RH
Valve RW has flow rate=0; tunnels lead to valves JS, GV
Valve IF has flow rate=0; tunnels lead to valves OK, JQ
Valve WP has flow rate=0; tunnels lead to valves HF, CQ
Valve YK has flow rate=0; tunnels lead to valves MO, GV
Valve MQ has flow rate=0; tunnels lead to valves AA, HI
Valve RH has flow rate=0; tunnels lead to valves SO, GB
Valve GB has flow rate=7; tunnels lead to valves YM, RH, PU
Valve XM has flow rate=16; tunnels lead to valves OU, ND, NB, RB
Valve RD has flow rate=0; tunnels lead to valves UN, LC
Valve HI has flow rate=0; tunnels lead to valves MQ, GR
Valve OH has flow rate=19; tunnels lead to valves OU, SU
Valve DS has flow rate=0; tunnels lead to valves OM, HF
Valve GV has flow rate=24; tunnels lead to valves RW, MC, YK, OM, LE
Valve AA has flow rate=0; tunnels lead to valves SO, DF, RK, MQ, JQ
Valve CQ has flow rate=17; tunnels lead to valves SB, MO, WP, SC, HR
Valve UV has flow rate=0; tunnels lead to valves LC, RI
Valve OV has flow rate=0; tunnels lead to valves OK, WC
Valve CM has flow rate=0; tunnels lead to valves RK, LC
Valve YH has flow rate=0; tunnels lead to valves NW, JI
Valve GZ has flow rate=0; tunnels lead to valves LE, HX
Valve WC has flow rate=0; tunnels lead to valves UN, OV
Valve ID has flow rate=0; tunnels lead to valves JI, HX
Valve SC has flow rate=0; tunnels lead to valves OJ, CQ
Valve GR has flow rate=11; tunnels lead to valves OJ, WI, HI, PU, JS
Valve IM has flow rate=0; tunnels lead to valves LC, WM
Valve NB has flow rate=0; tunnels lead to valves YM, XM
Valve TS has flow rate=20; tunnel leads to valve NW
Valve SB has flow rate=0; tunnels lead to valves CQ, YQ
Valve MC has flow rate=0; tunnels lead to valves GV, YY
Valve ND has flow rate=0; tunnels lead to valves JI, XM
Valve MO has flow rate=0; tunnels lead to valves CQ, YK
Valve PU has flow rate=0; tunnels lead to valves GB, GR
Valve IR has flow rate=13; tunnel leads to valve VR
Valve NW has flow rate=0; tunnels lead to valves YH, TS
Valve WM has flow rate=0; tunnels lead to valves IM, OK`
