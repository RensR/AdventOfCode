package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var tests = execute.TestCases{
	{
		puzzleTest,
		`24`,
		`93`,
	},

	{
		puzzle,
		`763`,
		`23921`,
	},
}

var puzzleTest = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`

var puzzle = `504,62 -> 504,58 -> 504,62 -> 506,62 -> 506,54 -> 506,62 -> 508,62 -> 508,55 -> 508,62
514,69 -> 519,69
534,120 -> 534,110 -> 534,120 -> 536,120 -> 536,115 -> 536,120 -> 538,120 -> 538,117 -> 538,120 -> 540,120 -> 540,119 -> 540,120 -> 542,120 -> 542,116 -> 542,120
510,67 -> 515,67
489,36 -> 489,29 -> 489,36 -> 491,36 -> 491,34 -> 491,36 -> 493,36 -> 493,26 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,29 -> 497,36 -> 499,36 -> 499,31 -> 499,36
534,120 -> 534,110 -> 534,120 -> 536,120 -> 536,115 -> 536,120 -> 538,120 -> 538,117 -> 538,120 -> 540,120 -> 540,119 -> 540,120 -> 542,120 -> 542,116 -> 542,120
515,104 -> 520,104
514,86 -> 514,87 -> 530,87 -> 530,86
489,36 -> 489,29 -> 489,36 -> 491,36 -> 491,34 -> 491,36 -> 493,36 -> 493,26 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,29 -> 497,36 -> 499,36 -> 499,31 -> 499,36
534,120 -> 534,110 -> 534,120 -> 536,120 -> 536,115 -> 536,120 -> 538,120 -> 538,117 -> 538,120 -> 540,120 -> 540,119 -> 540,120 -> 542,120 -> 542,116 -> 542,120
489,36 -> 489,29 -> 489,36 -> 491,36 -> 491,34 -> 491,36 -> 493,36 -> 493,26 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,29 -> 497,36 -> 499,36 -> 499,31 -> 499,36
510,90 -> 510,92 -> 502,92 -> 502,97 -> 516,97 -> 516,92 -> 515,92 -> 515,90
504,62 -> 504,58 -> 504,62 -> 506,62 -> 506,54 -> 506,62 -> 508,62 -> 508,55 -> 508,62
534,120 -> 534,110 -> 534,120 -> 536,120 -> 536,115 -> 536,120 -> 538,120 -> 538,117 -> 538,120 -> 540,120 -> 540,119 -> 540,120 -> 542,120 -> 542,116 -> 542,120
497,71 -> 502,71
521,84 -> 521,82 -> 521,84 -> 523,84 -> 523,82 -> 523,84 -> 525,84 -> 525,77 -> 525,84
534,120 -> 534,110 -> 534,120 -> 536,120 -> 536,115 -> 536,120 -> 538,120 -> 538,117 -> 538,120 -> 540,120 -> 540,119 -> 540,120 -> 542,120 -> 542,116 -> 542,120
568,158 -> 573,158
497,13 -> 497,16 -> 496,16 -> 496,23 -> 511,23 -> 511,16 -> 503,16 -> 503,13
556,152 -> 561,152
497,13 -> 497,16 -> 496,16 -> 496,23 -> 511,23 -> 511,16 -> 503,16 -> 503,13
548,127 -> 553,127
524,107 -> 538,107 -> 538,106
544,125 -> 549,125
497,39 -> 497,43 -> 493,43 -> 493,49 -> 506,49 -> 506,43 -> 501,43 -> 501,39
553,138 -> 553,141 -> 546,141 -> 546,149 -> 558,149 -> 558,141 -> 557,141 -> 557,138
489,36 -> 489,29 -> 489,36 -> 491,36 -> 491,34 -> 491,36 -> 493,36 -> 493,26 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,29 -> 497,36 -> 499,36 -> 499,31 -> 499,36
553,138 -> 553,141 -> 546,141 -> 546,149 -> 558,149 -> 558,141 -> 557,141 -> 557,138
534,127 -> 539,127
521,84 -> 521,82 -> 521,84 -> 523,84 -> 523,82 -> 523,84 -> 525,84 -> 525,77 -> 525,84
537,125 -> 542,125
521,84 -> 521,82 -> 521,84 -> 523,84 -> 523,82 -> 523,84 -> 525,84 -> 525,77 -> 525,84
521,84 -> 521,82 -> 521,84 -> 523,84 -> 523,82 -> 523,84 -> 525,84 -> 525,77 -> 525,84
553,138 -> 553,141 -> 546,141 -> 546,149 -> 558,149 -> 558,141 -> 557,141 -> 557,138
489,36 -> 489,29 -> 489,36 -> 491,36 -> 491,34 -> 491,36 -> 493,36 -> 493,26 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,29 -> 497,36 -> 499,36 -> 499,31 -> 499,36
497,39 -> 497,43 -> 493,43 -> 493,49 -> 506,49 -> 506,43 -> 501,43 -> 501,39
514,86 -> 514,87 -> 530,87 -> 530,86
534,120 -> 534,110 -> 534,120 -> 536,120 -> 536,115 -> 536,120 -> 538,120 -> 538,117 -> 538,120 -> 540,120 -> 540,119 -> 540,120 -> 542,120 -> 542,116 -> 542,120
518,71 -> 523,71
551,131 -> 556,131
489,36 -> 489,29 -> 489,36 -> 491,36 -> 491,34 -> 491,36 -> 493,36 -> 493,26 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,29 -> 497,36 -> 499,36 -> 499,31 -> 499,36
489,36 -> 489,29 -> 489,36 -> 491,36 -> 491,34 -> 491,36 -> 493,36 -> 493,26 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,29 -> 497,36 -> 499,36 -> 499,31 -> 499,36
534,120 -> 534,110 -> 534,120 -> 536,120 -> 536,115 -> 536,120 -> 538,120 -> 538,117 -> 538,120 -> 540,120 -> 540,119 -> 540,120 -> 542,120 -> 542,116 -> 542,120
489,36 -> 489,29 -> 489,36 -> 491,36 -> 491,34 -> 491,36 -> 493,36 -> 493,26 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,29 -> 497,36 -> 499,36 -> 499,31 -> 499,36
522,104 -> 527,104
556,135 -> 561,135 -> 561,134
560,154 -> 565,154
564,156 -> 569,156
510,90 -> 510,92 -> 502,92 -> 502,97 -> 516,97 -> 516,92 -> 515,92 -> 515,90
510,90 -> 510,92 -> 502,92 -> 502,97 -> 516,97 -> 516,92 -> 515,92 -> 515,90
508,104 -> 513,104
541,127 -> 546,127
507,69 -> 512,69
521,84 -> 521,82 -> 521,84 -> 523,84 -> 523,82 -> 523,84 -> 525,84 -> 525,77 -> 525,84
521,84 -> 521,82 -> 521,84 -> 523,84 -> 523,82 -> 523,84 -> 525,84 -> 525,77 -> 525,84
497,39 -> 497,43 -> 493,43 -> 493,49 -> 506,49 -> 506,43 -> 501,43 -> 501,39
534,120 -> 534,110 -> 534,120 -> 536,120 -> 536,115 -> 536,120 -> 538,120 -> 538,117 -> 538,120 -> 540,120 -> 540,119 -> 540,120 -> 542,120 -> 542,116 -> 542,120
511,102 -> 516,102
504,62 -> 504,58 -> 504,62 -> 506,62 -> 506,54 -> 506,62 -> 508,62 -> 508,55 -> 508,62
553,138 -> 553,141 -> 546,141 -> 546,149 -> 558,149 -> 558,141 -> 557,141 -> 557,138
534,120 -> 534,110 -> 534,120 -> 536,120 -> 536,115 -> 536,120 -> 538,120 -> 538,117 -> 538,120 -> 540,120 -> 540,119 -> 540,120 -> 542,120 -> 542,116 -> 542,120
497,13 -> 497,16 -> 496,16 -> 496,23 -> 511,23 -> 511,16 -> 503,16 -> 503,13
504,62 -> 504,58 -> 504,62 -> 506,62 -> 506,54 -> 506,62 -> 508,62 -> 508,55 -> 508,62
489,36 -> 489,29 -> 489,36 -> 491,36 -> 491,34 -> 491,36 -> 493,36 -> 493,26 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,29 -> 497,36 -> 499,36 -> 499,31 -> 499,36
497,39 -> 497,43 -> 493,43 -> 493,49 -> 506,49 -> 506,43 -> 501,43 -> 501,39
554,158 -> 559,158
489,36 -> 489,29 -> 489,36 -> 491,36 -> 491,34 -> 491,36 -> 493,36 -> 493,26 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,29 -> 497,36 -> 499,36 -> 499,31 -> 499,36
506,65 -> 511,65
497,13 -> 497,16 -> 496,16 -> 496,23 -> 511,23 -> 511,16 -> 503,16 -> 503,13
550,156 -> 555,156
521,84 -> 521,82 -> 521,84 -> 523,84 -> 523,82 -> 523,84 -> 525,84 -> 525,77 -> 525,84
504,62 -> 504,58 -> 504,62 -> 506,62 -> 506,54 -> 506,62 -> 508,62 -> 508,55 -> 508,62
489,36 -> 489,29 -> 489,36 -> 491,36 -> 491,34 -> 491,36 -> 493,36 -> 493,26 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,29 -> 497,36 -> 499,36 -> 499,31 -> 499,36
553,138 -> 553,141 -> 546,141 -> 546,149 -> 558,149 -> 558,141 -> 557,141 -> 557,138
497,13 -> 497,16 -> 496,16 -> 496,23 -> 511,23 -> 511,16 -> 503,16 -> 503,13
561,158 -> 566,158
497,39 -> 497,43 -> 493,43 -> 493,49 -> 506,49 -> 506,43 -> 501,43 -> 501,39
540,123 -> 545,123
489,36 -> 489,29 -> 489,36 -> 491,36 -> 491,34 -> 491,36 -> 493,36 -> 493,26 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,29 -> 497,36 -> 499,36 -> 499,31 -> 499,36
553,138 -> 553,141 -> 546,141 -> 546,149 -> 558,149 -> 558,141 -> 557,141 -> 557,138
510,90 -> 510,92 -> 502,92 -> 502,97 -> 516,97 -> 516,92 -> 515,92 -> 515,90
514,86 -> 514,87 -> 530,87 -> 530,86
489,36 -> 489,29 -> 489,36 -> 491,36 -> 491,34 -> 491,36 -> 493,36 -> 493,26 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,29 -> 497,36 -> 499,36 -> 499,31 -> 499,36
534,120 -> 534,110 -> 534,120 -> 536,120 -> 536,115 -> 536,120 -> 538,120 -> 538,117 -> 538,120 -> 540,120 -> 540,119 -> 540,120 -> 542,120 -> 542,116 -> 542,120
500,69 -> 505,69
514,100 -> 519,100
518,102 -> 523,102
504,62 -> 504,58 -> 504,62 -> 506,62 -> 506,54 -> 506,62 -> 508,62 -> 508,55 -> 508,62
534,120 -> 534,110 -> 534,120 -> 536,120 -> 536,115 -> 536,120 -> 538,120 -> 538,117 -> 538,120 -> 540,120 -> 540,119 -> 540,120 -> 542,120 -> 542,116 -> 542,120
489,36 -> 489,29 -> 489,36 -> 491,36 -> 491,34 -> 491,36 -> 493,36 -> 493,26 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,29 -> 497,36 -> 499,36 -> 499,31 -> 499,36
504,62 -> 504,58 -> 504,62 -> 506,62 -> 506,54 -> 506,62 -> 508,62 -> 508,55 -> 508,62
534,120 -> 534,110 -> 534,120 -> 536,120 -> 536,115 -> 536,120 -> 538,120 -> 538,117 -> 538,120 -> 540,120 -> 540,119 -> 540,120 -> 542,120 -> 542,116 -> 542,120
553,154 -> 558,154
503,67 -> 508,67
489,36 -> 489,29 -> 489,36 -> 491,36 -> 491,34 -> 491,36 -> 493,36 -> 493,26 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,29 -> 497,36 -> 499,36 -> 499,31 -> 499,36
510,90 -> 510,92 -> 502,92 -> 502,97 -> 516,97 -> 516,92 -> 515,92 -> 515,90
504,62 -> 504,58 -> 504,62 -> 506,62 -> 506,54 -> 506,62 -> 508,62 -> 508,55 -> 508,62
556,135 -> 561,135 -> 561,134
489,36 -> 489,29 -> 489,36 -> 491,36 -> 491,34 -> 491,36 -> 493,36 -> 493,26 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,29 -> 497,36 -> 499,36 -> 499,31 -> 499,36
497,13 -> 497,16 -> 496,16 -> 496,23 -> 511,23 -> 511,16 -> 503,16 -> 503,13
557,156 -> 562,156
553,138 -> 553,141 -> 546,141 -> 546,149 -> 558,149 -> 558,141 -> 557,141 -> 557,138
510,90 -> 510,92 -> 502,92 -> 502,97 -> 516,97 -> 516,92 -> 515,92 -> 515,90
521,84 -> 521,82 -> 521,84 -> 523,84 -> 523,82 -> 523,84 -> 525,84 -> 525,77 -> 525,84
511,71 -> 516,71
489,36 -> 489,29 -> 489,36 -> 491,36 -> 491,34 -> 491,36 -> 493,36 -> 493,26 -> 493,36 -> 495,36 -> 495,35 -> 495,36 -> 497,36 -> 497,29 -> 497,36 -> 499,36 -> 499,31 -> 499,36
534,120 -> 534,110 -> 534,120 -> 536,120 -> 536,115 -> 536,120 -> 538,120 -> 538,117 -> 538,120 -> 540,120 -> 540,119 -> 540,120 -> 542,120 -> 542,116 -> 542,120
534,120 -> 534,110 -> 534,120 -> 536,120 -> 536,115 -> 536,120 -> 538,120 -> 538,117 -> 538,120 -> 540,120 -> 540,119 -> 540,120 -> 542,120 -> 542,116 -> 542,120
497,39 -> 497,43 -> 493,43 -> 493,49 -> 506,49 -> 506,43 -> 501,43 -> 501,39
524,107 -> 538,107 -> 538,106
547,158 -> 552,158
497,13 -> 497,16 -> 496,16 -> 496,23 -> 511,23 -> 511,16 -> 503,16 -> 503,13
510,90 -> 510,92 -> 502,92 -> 502,97 -> 516,97 -> 516,92 -> 515,92 -> 515,90
497,39 -> 497,43 -> 493,43 -> 493,49 -> 506,49 -> 506,43 -> 501,43 -> 501,39
504,71 -> 509,71`