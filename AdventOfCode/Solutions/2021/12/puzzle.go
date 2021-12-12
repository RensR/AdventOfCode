package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var tests = execute.TestCases{
	{
		puzzle,
		`5958`,
		`150426`,
	},
	{
		micro,
		`10`,
		`36`,
	},
	{
		puzzleTest,
		`226`,
		`3509`,
	},
}

var micro = `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

var puzzleTest = `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`

var puzzle = `pg-CH
pg-yd
yd-start
fe-hv
bi-CH
CH-yd
end-bi
fe-RY
ng-CH
fe-CH
ng-pg
hv-FL
FL-fe
hv-pg
bi-hv
CH-end
hv-ng
yd-ng
pg-fe
start-ng
end-FL
fe-bi
FL-ks
pg-start`
