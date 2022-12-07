package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var tests = execute.TestCases{
	{
		puzzle,
		`1297683`,
		`5756764`,
	},
	{
		puzzleTest,
		`95437`,
		`24933642`,
	},
}
var puzzle = `$ cd /
$ ls
dir dscbfp
283653 fsdfddfv
dir mjzqq
241330 rcm.psp
dir sjbpgc
dir zfsbvs
$ cd dscbfp
$ ls
dir fgtvzpl
dir hgfrgbv
dir hmwqgjnl
dir jvr
dir lcvdgm
dir mmhtpz
dir wqc
dir znl
dir zph
dir zwlpm
$ cd fgtvzpl
$ ls
dir fvrghzfg
28513 lbbg.rhq
$ cd fvrghzfg
$ ls
212295 cjb.nwg
dir ftqs
$ cd ftqs
$ ls
250415 mmhtpz
$ cd ..
$ cd ..
$ cd ..
$ cd hgfrgbv
$ ls
86365 cdrgnzrz.hwf
175318 wqtmwb
$ cd ..
$ cd hmwqgjnl
$ ls
dir dscbfp
dir dwtfrgj
130223 fnl.whg
66339 mtcv
dir rgvvz
$ cd dscbfp
$ ls
146043 wpzr
$ cd ..
$ cd dwtfrgj
$ ls
dir jwmw
dir pntqg
$ cd jwmw
$ ls
dir dscbfp
243410 lbbg.phv
dir mmhtpz
$ cd dscbfp
$ ls
dir mmhtpz
$ cd mmhtpz
$ ls
dir mdrddz
$ cd mdrddz
$ ls
167704 fgjsq.bpb
$ cd ..
$ cd ..
$ cd ..
$ cd mmhtpz
$ ls
233626 dtcmsq.pdl
163642 jczs.rgg
111667 msfmjd.vlr
23137 ndhvh.jbq
$ cd ..
$ cd ..
$ cd pntqg
$ ls
68578 lnjvpcgq.zqs
62492 rcm.psp
$ cd ..
$ cd ..
$ cd rgvvz
$ ls
dir hmwqgjnl
dir jrjgnch
110656 lnjvpcgq.zqs
206537 mmhtpz.wgd
198736 msfmjd.vlr
110172 rrl.wqz
dir wrb
$ cd hmwqgjnl
$ ls
141962 bsgljmww.whq
dir hsmr
123313 msfmjd.vlr
223573 rcm.psp
$ cd hsmr
$ ls
229737 fzptbrzb.lqv
$ cd ..
$ cd ..
$ cd jrjgnch
$ ls
4426 lbbg.fzh
dir zbnp
$ cd zbnp
$ ls
65658 cjqbfv.fjf
270282 pwv.bcz
$ cd ..
$ cd ..
$ cd wrb
$ ls
dir fgwnpp
249291 lbbg
dir mbd
dir rdnf
dir rgvvz
dir rnrr
dir vzjh
$ cd fgwnpp
$ ls
dir gdfrtgl
230060 lbbg.qfm
275250 rbd.fqj
dir zlgpdb
$ cd gdfrtgl
$ ls
dir gvw
dir rgvvz
$ cd gvw
$ ls
123285 mmhtpz.vbl
167405 qhfqz
$ cd ..
$ cd rgvvz
$ ls
254625 rcm.psp
$ cd ..
$ cd ..
$ cd zlgpdb
$ ls
72966 cjjl.wzp
217828 ndhvh.jbq
141199 wbnl.qdc
$ cd ..
$ cd ..
$ cd mbd
$ ls
dir hmwqgjnl
$ cd hmwqgjnl
$ ls
13420 mmhtpz.snz
$ cd ..
$ cd ..
$ cd rdnf
$ ls
dir frqb
dir jjjzcjdh
dir mwpd
dir nvz
152893 rcm.psp
$ cd frqb
$ ls
dir dscbfp
231903 gwlj.lgn
dir mmhtpz
298674 nhc.hpl
193867 rfmv.dzd
dir sqt
$ cd dscbfp
$ ls
280820 rgvvz.glj
$ cd ..
$ cd mmhtpz
$ ls
137192 jczs.rgg
273426 rcm.psp
$ cd ..
$ cd sqt
$ ls
dir fwvgbgbs
245416 rgvvz.cgh
12489 wnjgq
$ cd fwvgbgbs
$ ls
dir mfpd
170052 msfmjd.vlr
278273 mwr.stv
177661 vgwstqlt.zml
$ cd mfpd
$ ls
264130 gmwhsj.jvp
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd jjjzcjdh
$ ls
dir cnjnnzw
dir flprjhlm
266398 fqclzm.mfw
dir gftmlvmh
113237 jgnmbml.rnr
148723 ljfg.vmc
237722 rrzvq.cqr
$ cd cnjnnzw
$ ls
231497 ldqdwn.bvf
$ cd ..
$ cd flprjhlm
$ ls
185486 hmwqgjnl.jdn
$ cd ..
$ cd gftmlvmh
$ ls
dir blpt
41840 qbr.tjw
$ cd blpt
$ ls
dir sbtg
$ cd sbtg
$ ls
7518 wsh.vrp
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd mwpd
$ ls
35589 flgh
$ cd ..
$ cd nvz
$ ls
147018 mpmcpwnl.hqg
$ cd ..
$ cd ..
$ cd rgvvz
$ ls
280057 dhwzlmgn
156280 gmzvl.gtg
dir hmwqgjnl
dir pfcjgrgp
dir prhzdjqv
132543 qlwqmd.wmj
223249 rcm.psp
124899 zvzjh.sbg
$ cd hmwqgjnl
$ ls
99756 mvdpfcvs.wft
$ cd ..
$ cd pfcjgrgp
$ ls
dir lbbg
dir nlgf
130799 wtrdvzhs
$ cd lbbg
$ ls
193125 lnjvpcgq.zqs
dir vnlmr
$ cd vnlmr
$ ls
39935 lgd.hzz
$ cd ..
$ cd ..
$ cd nlgf
$ ls
106563 jfr
$ cd ..
$ cd ..
$ cd prhzdjqv
$ ls
dir jhhq
$ cd jhhq
$ ls
135155 mzjdptfm.fln
$ cd ..
$ cd ..
$ cd ..
$ cd rnrr
$ ls
249039 jczs.rgg
154853 whlwl
$ cd ..
$ cd vzjh
$ ls
dir hjrs
$ cd hjrs
$ ls
146991 mmhtpz
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd jvr
$ ls
13268 gwsldhjc.vbf
dir hbpwd
34369 mjtwr.rpv
dir tfvztnb
$ cd hbpwd
$ ls
253823 bzqc
dir dscbfp
237051 ggvvcpg.gmj
23097 mhl
211524 ndhvh.jbq
133936 qzslmmz.fzp
$ cd dscbfp
$ ls
184190 msfmjd.vlr
238872 ndhvh.jbq
dir pfsq
dir pljsm
dir tlgtcb
$ cd pfsq
$ ls
dir hmwqgjnl
249351 mmhtpz
$ cd hmwqgjnl
$ ls
301132 rcm.psp
$ cd ..
$ cd ..
$ cd pljsm
$ ls
263279 bqh.vhl
$ cd ..
$ cd tlgtcb
$ ls
50558 bgzrlnzz.rfm
83169 gtrjzl.hhh
212764 mgf.zlg
210599 mmhtpz.dqm
dir tqsvhc
$ cd tqsvhc
$ ls
264513 lnjvpcgq.zqs
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd tfvztnb
$ ls
dir mdz
$ cd mdz
$ ls
209155 dqjh
$ cd ..
$ cd ..
$ cd ..
$ cd lcvdgm
$ ls
135165 lbbg.gsz
282211 vlbwsps.plg
$ cd ..
$ cd mmhtpz
$ ls
dir dscbfp
dir mmhtpz
70564 msfmjd.vlr
dir nzdvqb
59232 rgvvz
dir ztbht
$ cd dscbfp
$ ls
dir cddzd
dir dscbfp
dir hmwqgjnl
142689 qnnc
dir whbjm
$ cd cddzd
$ ls
255928 dtnn.hzf
$ cd ..
$ cd dscbfp
$ ls
dir brvf
182587 fhmzbc.nlb
129514 lbbg
177841 ngw.tlj
dir qttp
dir zdmb
269964 zqn.htj
$ cd brvf
$ ls
24071 tzsvg.pwc
174118 vtdntn
$ cd ..
$ cd qttp
$ ls
54613 ccdnjnwz
131277 zzppc
$ cd ..
$ cd zdmb
$ ls
117682 bwjhr.gvw
dir hmwqgjnl
138816 jczs.rgg
273129 mmhtpz
89710 rcm.psp
dir rgvvz
$ cd hmwqgjnl
$ ls
102907 mmhtpz.dnp
235885 ndhvh.jbq
239856 rcm.psp
$ cd ..
$ cd rgvvz
$ ls
dir fpzzfc
292842 lbbg
70603 msfmjd.vlr
157797 qjs.zlm
dir zgn
$ cd fpzzfc
$ ls
dir rgvvz
$ cd rgvvz
$ ls
222407 crrwlp.zcd
$ cd ..
$ cd ..
$ cd zgn
$ ls
70915 hmwqgjnl.zdg
276918 lbbg.wlg
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd hmwqgjnl
$ ls
207726 cqn.rds
$ cd ..
$ cd whbjm
$ ls
44954 ndhvh.jbq
136649 pphznj
dir ptq
215264 rmpbqgrl.mqt
276795 tbtv
292696 vddnpnp
dir wnqgpj
$ cd ptq
$ ls
191976 jczs.rgg
59747 ncbs.mjc
77874 njqmf
dir zss
$ cd zss
$ ls
209276 vnv.nvm
$ cd ..
$ cd ..
$ cd wnqgpj
$ ls
211738 jczs.rgg
$ cd ..
$ cd ..
$ cd ..
$ cd mmhtpz
$ ls
dir hmwqgjnl
$ cd hmwqgjnl
$ ls
dir jzb
dir mmhtpz
$ cd jzb
$ ls
138393 drqp.ttd
250098 jczs.rgg
dir mfq
$ cd mfq
$ ls
dir nfn
$ cd nfn
$ ls
239922 rcm.psp
$ cd ..
$ cd ..
$ cd ..
$ cd mmhtpz
$ ls
158815 jvwqttqn
72929 jvzfccm
18353 lwz.clt
144822 rnvfllt.fwn
$ cd ..
$ cd ..
$ cd ..
$ cd nzdvqb
$ ls
226139 rcm.psp
$ cd ..
$ cd ztbht
$ ls
279766 jczs.rgg
222685 msfmjd.vlr
$ cd ..
$ cd ..
$ cd wqc
$ ls
dir cndf
dir hmwqgjnl
dir lbbg
$ cd cndf
$ ls
dir hmwqgjnl
284014 qsr.pjg
$ cd hmwqgjnl
$ ls
91921 zrsp.qwd
$ cd ..
$ cd ..
$ cd hmwqgjnl
$ ls
dir fccbrtcn
23864 msfmjd.vlr
286035 nrbmbpm
dir pwvgqth
91650 rgvvz
dir wllwhm
$ cd fccbrtcn
$ ls
87931 bmjdngzq.zbv
dir cwr
dir djngvdp
266568 lnjvpcgq.zqs
dir zdwhqb
$ cd cwr
$ ls
dir cmnv
$ cd cmnv
$ ls
dir dscbfp
dir dwdnwdz
291794 rcm.psp
$ cd dscbfp
$ ls
dir gtcg
$ cd gtcg
$ ls
73017 pgzcm.qbz
$ cd ..
$ cd ..
$ cd dwdnwdz
$ ls
dir hvgwfj
$ cd hvgwfj
$ ls
235064 rnjjh.qnp
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd djngvdp
$ ls
264940 jczs.rgg
162343 rgvvz
$ cd ..
$ cd zdwhqb
$ ls
233875 ggd
230766 ggqrt.pqn
$ cd ..
$ cd ..
$ cd pwvgqth
$ ls
dir nrm
dir rgvvz
289164 sgdsg.fbs
$ cd nrm
$ ls
135745 hmwqgjnl.fwb
$ cd ..
$ cd rgvvz
$ ls
269675 rcm.psp
$ cd ..
$ cd ..
$ cd wllwhm
$ ls
14033 rptrszg.lfh
$ cd ..
$ cd ..
$ cd lbbg
$ ls
dir fsz
dir hmwqgjnl
2856 jpmw.tsp
156026 ndhvh.jbq
dir rgmmpm
dir vzg
$ cd fsz
$ ls
29744 cdsr
dir frjzv
dir hmwqgjnl
75723 hmwqgjnl.jcj
222555 jczs.rgg
dir lbbg
dir pplp
184370 vzb
$ cd frjzv
$ ls
87553 lnjvpcgq.zqs
$ cd ..
$ cd hmwqgjnl
$ ls
94991 hsjmzpq
$ cd ..
$ cd lbbg
$ ls
177441 dscbfp
9637 lnjvpcgq.zqs
$ cd ..
$ cd pplp
$ ls
230787 gcjfcbg.sds
54478 twnslqqv.gtw
54723 wzwcw.pfj
$ cd ..
$ cd ..
$ cd hmwqgjnl
$ ls
dir drpt
dir dscbfp
18560 jczs.rgg
dir tpt
$ cd drpt
$ ls
dir gmcpd
10052 jczs.rgg
87927 mmhtpz.jdt
dir rgvvz
$ cd gmcpd
$ ls
dir rgvvz
$ cd rgvvz
$ ls
dir tjqcj
$ cd tjqcj
$ ls
273175 bgsjwb
$ cd ..
$ cd ..
$ cd ..
$ cd rgvvz
$ ls
104452 fbv
$ cd ..
$ cd ..
$ cd dscbfp
$ ls
dir bjn
dir rvm
$ cd bjn
$ ls
119619 mmhtpz.slf
$ cd ..
$ cd rvm
$ ls
dir hgl
$ cd hgl
$ ls
9146 jczs.rgg
294696 rcm.psp
$ cd ..
$ cd ..
$ cd ..
$ cd tpt
$ ls
102194 tmlbnm
$ cd ..
$ cd ..
$ cd rgmmpm
$ ls
25462 dcp.zfg
99756 frmlbqp
71786 msfmjd.vlr
275455 mvrszdhp.jbr
$ cd ..
$ cd vzg
$ ls
dir hmwqgjnl
160612 msfmjd.vlr
103726 pprnm.rmw
$ cd hmwqgjnl
$ ls
130188 ndhvh.jbq
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd znl
$ ls
256149 qcd
dir rgrhdq
dir wfhcc
$ cd rgrhdq
$ ls
dir nqpz
17930 rglm.wrj
$ cd nqpz
$ ls
253269 rcm.psp
$ cd ..
$ cd ..
$ cd wfhcc
$ ls
dir dpds
dir hmwqgjnl
82595 rcm.psp
dir rgvvz
$ cd dpds
$ ls
143467 bljj.ddw
283261 rcm.psp
$ cd ..
$ cd hmwqgjnl
$ ls
dir fnmqpt
$ cd fnmqpt
$ ls
288739 dscbfp
$ cd ..
$ cd ..
$ cd rgvvz
$ ls
41475 srcnvmj.tqb
$ cd ..
$ cd ..
$ cd ..
$ cd zph
$ ls
dir hmwqgjnl
dir tbwhzrtt
$ cd hmwqgjnl
$ ls
67069 rcm.psp
219165 rgvvz
$ cd ..
$ cd tbwhzrtt
$ ls
198892 gwlw.hbb
94996 mmhtpz
dir wwwj
$ cd wwwj
$ ls
dir lcttc
14000 lnjvpcgq.zqs
$ cd lcttc
$ ls
141874 jhlrtbjw.thr
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd zwlpm
$ ls
206595 ctvmprqd.nwz
292861 dscbfp.bpv
171719 ldmrl.zhz
23134 lnjvpcgq.zqs
dir mmhtpz
283001 msfmjd.vlr
dir pwrfgjdg
dir qdztzhl
$ cd mmhtpz
$ ls
90712 lzv.smr
102181 mcntll.fgt
260630 nhn.fll
219684 tlws
dir tpsvqmgb
$ cd tpsvqmgb
$ ls
dir bpmqlq
$ cd bpmqlq
$ ls
67400 lnjvpcgq.zqs
$ cd ..
$ cd ..
$ cd ..
$ cd pwrfgjdg
$ ls
30612 fflm
174894 jczs.rgg
$ cd ..
$ cd qdztzhl
$ ls
dir dscbfp
dir gcpjnrb
dir nlffhf
dir vnrgqg
dir zrpgb
$ cd dscbfp
$ ls
dir gdjvclf
dir mbhdvsq
dir tqb
dir vlqqslp
$ cd gdjvclf
$ ls
104663 hwwlf.mhv
$ cd ..
$ cd mbhdvsq
$ ls
98749 gqsjtd.rbv
200237 lbbg.nwb
$ cd ..
$ cd tqb
$ ls
101550 cqjmfvd.grg
dir hmwqgjnl
$ cd hmwqgjnl
$ ls
dir mlmmr
$ cd mlmmr
$ ls
dir rgvvz
$ cd rgvvz
$ ls
188297 gprgzd.rjf
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd vlqqslp
$ ls
191754 lbbg
$ cd ..
$ cd ..
$ cd gcpjnrb
$ ls
254902 nzzsg.bds
$ cd ..
$ cd nlffhf
$ ls
295353 hhjqvqnc
dir mmhb
58446 rgvvz.hsj
120084 rvgctqpr.qpp
302397 wtmlrm
$ cd mmhb
$ ls
57621 ndhvh.jbq
$ cd ..
$ cd ..
$ cd vnrgqg
$ ls
dir fpmpzvj
dir gtqwbhc
213709 hmwqgjnl
283893 hmwqgjnl.bbc
dir lnw
dir mmhtpz
177958 rcm.psp
68642 vzr.gwd
297210 zqcvpgfm.sgv
$ cd fpmpzvj
$ ls
62127 mmhtpz.mwv
$ cd ..
$ cd gtqwbhc
$ ls
174559 rgvvz
$ cd ..
$ cd lnw
$ ls
dir lbbg
162886 mlsb
dir rgvvz
dir szgmtgb
253283 zpvqj.crr
$ cd lbbg
$ ls
dir jgwj
dir ndljhbv
dir smhzdbn
68451 sznhpr
$ cd jgwj
$ ls
290222 fnpmb
148122 lnjvpcgq.zqs
dir ptwrtdcf
4669 rcm.psp
76708 thgtnwhq
$ cd ptwrtdcf
$ ls
dir gcjpt
$ cd gcjpt
$ ls
dir hmwqgjnl
$ cd hmwqgjnl
$ ls
259660 lnjvpcgq.zqs
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd ndljhbv
$ ls
226511 qmtttllc.wgt
$ cd ..
$ cd smhzdbn
$ ls
161029 mjwqn.bfs
49691 msfmjd.vlr
77941 ncwrv.fjc
12615 ndhvh.jbq
19222 rlrnrcqw.cgh
$ cd ..
$ cd ..
$ cd rgvvz
$ ls
270852 msfmjd.vlr
$ cd ..
$ cd szgmtgb
$ ls
167819 ndhvh.jbq
$ cd ..
$ cd ..
$ cd mmhtpz
$ ls
6153 msfmjd.vlr
$ cd ..
$ cd ..
$ cd zrpgb
$ ls
dir dgcrgg
dir dscbfp
dir gcdtbrnj
dir hmwqgjnl
284848 hnfqrh.nll
dir jsjrg
16735 lbbg.lpg
274897 vrp.fvm
$ cd dgcrgg
$ ls
dir gpvhhh
210466 lbbg
161257 tqc.qgg
dir zggnvwcb
$ cd gpvhhh
$ ls
100124 msfmjd.vlr
$ cd ..
$ cd zggnvwcb
$ ls
76550 dsmvrsp.hht
106682 swj.pnq
$ cd ..
$ cd ..
$ cd dscbfp
$ ls
297842 msfmjd.vlr
$ cd ..
$ cd gcdtbrnj
$ ls
78674 jmc.sst
$ cd ..
$ cd hmwqgjnl
$ ls
280262 hmwqgjnl
dir hrl
dir mmhtpz
159812 rcm.psp
$ cd hrl
$ ls
dir cqwbgn
dir nlqd
dir wmgqpt
$ cd cqwbgn
$ ls
91489 msfmjd.vlr
$ cd ..
$ cd nlqd
$ ls
1982 jczs.rgg
172223 rcm.psp
$ cd ..
$ cd wmgqpt
$ ls
dir hmwqgjnl
91172 jczs.rgg
125726 lbbg
147883 rgvvz
99929 rhzwrpw.jml
145844 tzzqwr
$ cd hmwqgjnl
$ ls
198790 mlccths
284058 mmhtpz.dvb
$ cd ..
$ cd ..
$ cd ..
$ cd mmhtpz
$ ls
60918 hmwqgjnl.jhw
220709 lnjvpcgq.zqs
$ cd ..
$ cd ..
$ cd jsjrg
$ ls
230757 czg
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd mjzqq
$ ls
33792 fsrws.tdp
dir lbbg
50734 lnz.fwp
268057 wbgjwnb
$ cd lbbg
$ ls
207413 cqf
$ cd ..
$ cd ..
$ cd sjbpgc
$ ls
73862 lbbg.ddv
$ cd ..
$ cd zfsbvs
$ ls
dir lbbg
10390 lqgh.fpl
147426 mqw.fmr
273227 rcm.psp
$ cd lbbg
$ ls
dir mmhtpz
$ cd mmhtpz
$ ls
25218 dtdt`

var puzzleTest = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`
