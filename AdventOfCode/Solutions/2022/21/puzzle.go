package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var tests = execute.TestCases{
	{
		puzzleTest,
		`152`,
		`301`,
	},
	{
		puzzle,
		`155708040358220`,
		`3342154812537`,
	},
}

var puzzleTest = `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32`

var puzzle = `qmqn: 5
dnln: szrq * gzsr
whls: lwzw * lswb
tsgp: tgnc + jvcs
jqcj: zbhg + zlsf
hcbv: 3
nzpc: fhhz * tjvm
qdbg: 2
pbbm: 3
mhbb: 18
dllz: 4
tcnz: jjwd * pzsl
rfbj: hwdb * sqmj
wtjv: 7
rljg: 5
bbsl: 2
rppv: 4
rdlf: 1
jvmb: 7
shdn: fdng * mvlq
ddtp: 2
dzpq: bpnf * hvgt
dcbh: 2
qfzc: jgvc * tcfs
pzbs: 5
mmpl: cqwg * swqw
vdlh: cmrw + fbbj
qtcf: fnnt * vtlb
qtrv: 3
njrm: 4
gbjw: crqf * rdtc
dfcf: sgzl * nghz
ntcf: bjrh + mjvz
rcrd: 2
mslh: ndrj * pmpv
bcnj: 2
nswf: gtpb + gvln
fzfq: 3
pzfg: gfwg * rvjn
ldtn: mrmt / jwgc
vmwn: fqqr * twzq
nphv: 5
nnts: 6
zqjq: vmsd * mzmh
plrt: dlpj + vhtz
vlzz: 2
jztw: 2
lrsg: 9
qprs: fctb * vwph
vbtd: 4
rmng: 4
rgtf: 8
dhvr: 4
ngqn: 2
sjtp: 4
vtwc: 2
jqjh: 9
tmjh: 5
hcqw: 3
fqmg: 18
bljg: scpw * gjld
lhmm: lsmb * hdgz
wwlj: 2
qwmg: fjrb + fppq
whnl: dvwt + lgcs
dpfd: jzjn + mpgg
crcv: jldb * qlsb
zbsv: 11
bnrp: dljj - fbpz
dnzf: bcvl * cthh
hdws: 9
vvcm: mhbl + qvzn
ctgb: bfpc + bvgr
vtlm: tnfn * lswp
scgv: jbln + lttt
hvfs: 5
zvfr: 2
dtjm: 11
cffr: 4
mrlb: gpgp * jcdn
qght: 4
qhzg: bdrl / zbgg
hjmr: 3
mvlq: bzjv + fljh
bstr: 2
vpwm: 11
wssd: 4
qwdt: 3
dnwc: svzn * dfcs
hlsr: 5
vtzc: mvzn * jccl
nnps: cmwc * hgdp
qvzn: 4
gszz: vvwr * jvjf
dqzl: bvps * wdwj
fsff: 2
ccrc: 3
mwpb: 19
grpz: 4
rrgh: 5
rhnn: nnps + dfjg
mnhq: dllz + jbrr
tjgm: qnjv + rhmj
bbtr: sscc / mctz
vtjv: 5
ggds: qmtb + twbw
jwqg: jmzf + hwtw
qpgc: rlqh * dsqz
zwsn: 14
gcgz: njlv * bczs
nhnp: 5
zmnt: 6
qrqb: 2
cmll: 2
dftb: rdgj - hpjj
rdpq: hnbj / hcqq
wccj: fpnl * fhfc
gwfp: 3
bsvc: 2
hdsc: zflt + wgvl
nscc: 4
lmjp: 5
ljmd: 2
fnrs: zvsg + gvjh
hwgv: 5
hdpg: jpmj * zhpj
jjhd: 7
nzdw: 2
mclp: 3
jptw: bjrt - rrhp
wcpd: qmzd * fhsp
djmd: qcgh * qdvj
lcjs: 3
dfzh: gfpq + fnpn
jcmj: 2
clgd: 3
tplh: sbzr + ddtr
rmdz: 3
qsnb: nbbw * bmdp
rdjm: sllw * dwjw
dqpl: cjqq * sdgd
cglw: dmvh - rjbt
cchb: 6
flzd: 3
fswh: qdbg * bpmn
gfpq: 4
whqj: 2
ndsv: 5
ccbh: jnvh * fhld
zmth: jtwc * rdzt
ntdn: 4
fpwr: 5
ddlw: 16
bwhs: qgwt * jnhr
tvhs: pmgl + zdsc
nbtc: 11
lfmn: 9
gtpb: tmsd * bmwd
lsdl: 9
fjqg: 5
wrsn: 3
wvmp: 2
zzrh: 3
flfg: dwqf * vgjv
rvgf: 5
blsg: 3
dtcf: 2
qhqc: 2
tfhg: zlch + znfl
bcmc: fjlt * dbrj
qmtb: zhtb + fccd
dqwl: mvdm + ncdr
nfvz: 2
pczp: zwnf + tvfb
qwqn: 3
frtp: 4
prpb: wjlp + brgz
brgz: 6
ljpj: 16
vljf: 2
bggf: 6
nsdv: nqhd + gpfw
ccbt: wssd * mjdv
vnld: vtnh * jjwf
trsl: 7
ndnf: pccn - fzht
jbgn: jtqq * pqsc
hfdr: slgz * lzgf
lbpb: 1
vjvn: lzmj * tvhs
zspf: 2
rjpr: mfdq + gczl
zvdr: tlsb * rcbh
llfs: 2
ppfr: rbdh * mzbl
qzln: mhbb * wfjt
tzsl: 2
dqlf: qnch * qsjs
gpsj: 2
zwsl: lszt * wbbd
mrmt: pnfn * cfzd
zbjt: cwth / hvgs
hrnm: 18
vthb: 3
mqtp: phbs + dwfm
nnnj: ndnf * cndd
rsdg: frtp * dsmj
znsg: 2
mlrq: 2
pmtt: mrml * gwzv
tsjl: hfdr / jrmq
llfp: vzhl * pvmt
vqpb: wlhz / frfw
fffd: dfdq + rcth
dfqz: 5
jfzh: ssmd * lbqf
wfjt: 6
pnwl: 3
sdgd: 3
sbwg: 3
fbwc: 2
lmrz: 19
rcbh: 3
dwjw: 2
pbvr: 4
rwwd: qqzh + wwcz
ttrp: 4
jtwc: 2
vtsq: ddsm + ssnj
lwrf: mwpd + qbmc
sbfm: cfgm * cwnl
dfcs: dcbh + blwr
mgtl: 3
jfct: 10
ptdj: 5
lswb: qpcp + njmd
humn: 2037
lljf: pgvf / mjnz
nflw: 2
vtng: 5
qmvs: jnsp + zrcq
bpwf: 6
gwhz: zpmg * vbsw
pnjz: 11
slgz: 7
jhnj: vtht + pqsv
ttzw: gpsh * cmzg
rlqh: 5
tdjq: 5
gddg: 2
bgsp: 14
jpwq: hlht * hblh
qtwn: mbtb + mhhn
bqpl: 4
fqhh: hrnm * znwj
cnwt: rdlj * bjqh
wzcd: cqvv * mdfq
wzql: djqs * srbf
fqsn: wrjw + mcqw
lttt: zczb * qvzz
mzdl: jhnj + ttzw
gpgp: 2
nbgj: 12
djzq: 9
rgzs: 1
rszv: 13
fcwc: 3
nvcs: 10
hwsj: grcl * dbmc
hdpz: vlwq * vpns
zgmn: gpsg + gdjm
hdtf: dcnr + dfmv
jtsl: wqjj * wtds
tprt: pspc * qsqf
rbwb: jgqj * zfcb
gdls: 18
ltrd: hlvv + jzng
hzpb: sbhg + qshs
dslq: 17
qrnp: 4
qnjv: 4
wpbl: tdjq + zstb
clpc: 7
pzlt: fdsd * lwrf
cfgg: 20
ghtl: 2
vhmd: gtdc * brph
vcpv: 3
sswc: gsmq - ltdc
lfgg: 3
ndvr: pppl * ngng
vzws: 2
vpld: ncvr / vhwq
jbsr: 6
cjtf: flfg * qqfs
ntgv: 2
qqvg: jqvj * mgmw
gpng: nwjg * fzfq
sdfb: 4
lvcq: 5
dfbc: 1
mwhc: 3
fzsp: nqgl * lfql
ftnh: 4
sdjq: 2
hgnj: gqbd + sltr
gtbv: zpfd + zlrc
bjbt: mhwm * wlnz
pblc: 8
gnbw: 5
fmwl: trrw + qpgc
qvzz: 2
cshb: nsdv * thdr
gdnh: 2
cmnw: vnvb * ptgl
hcjh: 3
nprp: tmns * wdsp
hdfq: 11
zznh: dhrb * jjdc
qzjr: lqwm + dwbq
tnnb: mzwm - bvqv
trfb: 2
hpjq: gqmb + bzlb
pmgl: 11
wjlp: 11
sllw: 13
dngq: rljg * jtqt
rngg: spjf * cshb
hfjj: lljf - zrwl
lszt: 10
gfms: wzcq * lpmb
vfhz: mntl * gpsj
lgcz: nmjh * cswg
dncm: mvpz + bsjv
gldn: stzn * bcdz
cwsp: cvnz + sbfm
lrlc: flbz * dtcn
gvjq: gsvd / vfrn
rvrr: dsjr + gfpg
mqhc: 3
rgpg: 2
vbql: rpvs * cqfn
gcql: 12
mzcf: wdnr + sscs
jpvw: fddj + wjtz
hphc: cpsp + srvn
ncvr: tdrz * fqwl
jtgt: jtff - frbq
gsvv: cgbw - rqbf
gmlz: ngnw + cjvd
zqst: fdpr * hglv
tqhz: zplh * cjnc
pvvs: 2
qdpb: pscg + sttp
rsnv: 5
cpvd: 3
sqzj: 2
wrhb: 5
bzqj: vcjl + mmsg
qhcf: dgnb + vnqz
frll: wccj + qgqg
qrqh: zzrh * lbdf
wnfh: 2
vbqc: 2
crqf: dbmh / nlnt
bchc: 5
wpjs: jdpf - fszn
nvgt: hrvs + tjbt
dfdq: tcmm * wsbp
gsnq: 18
hgnm: lbbp * pfzg
zmcd: 14
svtl: hdsc + qsjt
hvfp: pwdn * qgqv
bzzz: 5
jtdt: qpmm + tfqp
tsgv: hzzz + fmng
chjp: frcz + mmpj
qpgd: 5
rlfj: czsc + cmmf
dtgg: gszz * hlft
jccc: 1
rplp: wfsd * gjln
nmtf: 3
jzjn: 16
clbm: tjbc / hczl
hmtv: 5
dsqb: 2
lvgz: 5
hhrw: zvdr + nfrt
sgtq: 7
wppf: hcjh * wnfh
jpmj: vpmn + qtcf
vsbc: rplp * llfs
fbbc: vqml - zplm
bqff: 2
pspc: 3
qpsj: 4
fsws: vdln + ncpf
fqwp: cjht + wdgd
ncpd: 5
slpm: qrlz + crzh
wfbh: 2
bslg: ljfw + dtgs
wqjj: 4
wvqp: vzhb + wscp
wgvl: 5
twsz: qbqm + jjld
qqfs: rlqr + zmwg
mfhb: wpjs + wzjb
hglv: 6
njmn: fnvn - qmbc
vvrq: 2
rhhb: 3
dhpc: mjlq + qmss
mmnb: dzjr + zrhm
wvlw: vhjz + tjng
dvmv: 2
dgwr: hgdl + mgvm
mfcz: hhbf + fsrq
qsqf: hfcg * swcj
fstw: 9
zhrm: 9
brnz: jpwq / jfqr
qmwb: 4
ldds: prvn - gqlz
wlgf: jcmj + qmqn
tjng: hnjh * rsmv
rqvg: dwnm + tngn
pltn: 17
wmgq: 5
zvsg: 6
jgvc: 4
rnbt: qtcw / nmpf
zrcq: 11
sfdd: 2
mtzd: mphh + njfb
bbgf: nflw * ccjz
ftzz: qjzz + pbvr
bvtw: tpnz + ddzb
jlqd: fjcq - vfph
grzn: 5
wwng: lznj + cbrj
trnp: jgfv + bqvn
lnpr: 2
dtcj: 2
vpmn: hlrl * ttns
bjlm: hhbp * phdw
vwsj: 2
gjln: 2
jtzl: cwph * dbrh
bljh: 4
vgwj: dvfn + zwqv
lqdt: 2
jtsj: qwmg * fdvf
zqmf: vlgv * wnrz
wfvj: hzcl * lshs
rzmm: nndr / sqrh
dpss: 5
bhfq: 3
wpfg: 3
lqvp: 5
fjcq: lljv * chjp
srpt: 6
nfth: ltrd * nhml
wzzw: 10
dcnr: 6
mgcl: 14
qbmc: 3
dqpm: jjhd * hbnq
vdss: hgvw + bsqq
tzfz: 2
rnjl: 4
zbgg: 2
cfln: nvcs + plbl
nvtj: qrqb * rhsm
cbrj: 20
btgq: mjvl * tlhr
jsss: 5
ppns: 2
qbqp: gldn + pjjr
wfsr: 4
cjld: thfg * pbbm
bpvq: 3
httg: 2
jqmw: zhlp + tdbf
btqv: 2
sjlz: zzwd * jjjw
wlpr: 2
fbrt: 8
nlnt: 7
hhbj: 3
jrlb: lmcj + zdsg
ndrj: zlfz + fdpq
nltb: 2
jvsr: 4
vtsh: wzzw + bgsp
ndnr: vgwj * dfhr
ssnb: 3
cnwq: gldj * tnnb
fcts: 3
vcnb: 13
fsqw: cbjh * jnbv
hjnj: sngp + jssj
zqwm: zmth - tnnp
znzf: fsqw * mmml
fdth: 3
gpgt: 3
mlwh: 2
svlh: zgmz - zbhr
ssmd: cbpl * pgrz
lvsp: 17
rhlf: 2
bwtj: 5
btmb: 3
vbpl: 2
pmlb: btdh * bchc
hnbz: 3
htdq: 9
pdlw: llms + vttr
zbdr: lqdt * hpjq
gjnq: 2
sjrj: htfl + drtr
vfpj: pvvp * cwsp
grsn: rtrd - dvlf
rfzj: 4
ddlb: 2
rdlj: 4
pdvq: 4
ltvh: gnhs / gwfp
fsrq: 19
jnrg: 3
nwfz: vbql + rfqt
jjbv: 10
hbmc: 4
lgbg: 3
wjgf: drnn + pjwb
wmwq: ghtl + dhvr
fvld: 2
wtws: 3
wqrm: ncvb * ctbd
rjvj: 3
dgnb: 17
bdqh: 2
mpfh: 13
dmvh: jwgs / hwgv
mrml: jmgd + wmzd
pmhz: 2
mpzl: bfnd + qrgt
pnns: 3
nfsp: lbnn * rbmv
cwmz: 3
njlv: bljh * jqcj
cjfn: vsbc + mclp
fbzr: 18
hcvn: 5
zbhr: 2
fvtb: 2
grnj: tgzt + cjvz
hpwz: jztw + jqbg
ftmd: 4
cpbf: wlpr + qrqh
wpfw: pcgf + mnwc
dvqc: bvgq + nvtj
jvbt: fbbc * fpgp
dpvq: 2
plzp: nmsp / lsdl
qmzd: vvnr + lrdp
sqmj: 3
cvcz: 5
tvwv: jjht * wjtj
fmrr: wslp * wmgq
mqsw: wvmp * tcbr
ntqv: cwvz * cmll
bwgc: 5
mtlc: 2
pfzp: gmng * dpss
nsgd: 3
bcns: nzpc / tmnf
hhht: 5
hzhp: 11
jgww: 1
vtmd: 3
fhsp: tzvp * grsn
mrcc: 3
gmjr: 2
ncgv: 2
qwsz: bvgj + clgd
qsjr: 5
jmgd: rwhn * zfwq
hcfn: 2
blms: gwvp + ltvh
jvlv: 2
fjlt: 3
czwr: 2
ldhs: vmbg * pnqf
brwb: 2
sdjv: 14
bwbb: 4
lpmb: fzfz + bvsn
tbtp: 3
wzlf: 4
vvgn: dpfd + vpwm
dqpn: jtwf + ztgp
hlmg: wcnd / vwsj
vtrw: cjpm / gbtl
pbhv: qmnp * ggsq
vgql: 7
vglr: rlqc * rngg
zqhr: gjnq * nscc
lwlv: fhdr * wpbl
hwtw: 18
tdbr: 2
nrfw: 3
gdjm: rsnz * bmjl
rnfm: fcmn + hrzr
bdvq: nhnp * qhhz
rnvp: nfsp + mlgd
jszv: 14
bgjp: wjwh * qnpv
drtr: rpwp + djzq
phgw: 2
bvgr: vldp * hvsc
hzzz: 4
jjbz: nrts * crhl
nqhv: hpfd * btgq
mhbf: 8
qbcj: vcsh / lpll
mjcn: mmlz * scll
rlwh: mgfp + sqtq
qpnt: hhht * zchq
jvjf: fzsp + tsjt
hrcj: vbtd * wrfm
tnwv: 11
gswg: sdjq * zfcs
czff: cjld * dwhg
vbbw: mdbg * lvgz
bnzm: rqws + wqll
htmz: 4
ssnj: lqbq + ntdn
cqnm: 1
htmh: 9
dlhr: 4
wnrz: wsnb + nbgj
qdsm: 17
nclz: 4
cmzp: 2
cfgb: sdpv * lzsl
wslp: 5
qpcp: cvsb * mnhq
nwdl: 3
bvqv: stnd + jhht
wfqv: 7
vdsb: 3
zhvn: 4
jphn: jtjm + hcvn
ljfw: 4
zddg: 3
lzsl: npqz * gdnv
nrgv: npwm * mqsc
qmdw: cffr + jvmb
jcrq: zzsg + bnrp
tmsd: 7
wwhs: 6
mqsc: 3
csng: 1
wfjm: qrnp * hdlr
cnhz: dgsz * rnjs
jzwt: 2
spbv: 6
drvn: 2
gwrw: 17
sbgq: 2
lwmg: pmlb - hltm
ngch: 5
rszs: 4
bfng: jmnr + clwv
gwst: 4
dfjg: 4
dvlf: jfzf + lscd
lvjd: jwft * znnr
gtdc: 2
nrzm: mpfv * llft
tjbt: wswr - cdhs
wnfs: 14
msdt: znzf + rnvp
nzbv: 17
wwhg: sdld * wtjv
ztqn: 4
clqd: 8
wctn: rdlf + jphn
vdlr: 2
hrpt: 2
hlwg: hffg * qhmn
gsmq: bltd * zhwp
njmd: 11
zjrw: 2
vzmp: nnnj + rdpq
ljrg: 3
zmqc: 15
tfgg: fqzp / fsws
lzgf: 7
ngzr: 3
vfrn: 2
nbln: fnrs * qpbl
gmmp: 3
wjqd: 6
mtsq: ndnt * bhdg
hbvn: 4
ngng: 7
bmpw: hgnm - bcns
pfzg: vljf * dqpn
hlht: 4
hrvs: 2
gjpt: gwgt * bqpc
pgvf: cjtf - gdjv
vtzj: 2
qrgt: jdsw * bnjl
fqql: cqrd + fgds
tvmv: zqlr * rhlf
pcgf: scfh / vbnq
thql: zzfz * qtwn
qvdz: mccq * vggs
lwwl: 5
zrwl: hlgc * swqg
sjqt: dvsj * qqvh
zsdd: mqhc * qtsr
zrqs: 5
mhmd: 4
gqmb: 5
vldp: 3
cqvg: lwlv - hdfb
mjtg: vcbb * vptr
dzpj: 16
tqhq: 11
rrdl: 4
ltwv: 1
swqw: ztws * twpg
tmnf: 2
sjdm: wwlj * zczf
wjjv: cnsn * sgtq
jzqr: 2
wdsq: 5
qzhz: jcqp * mlwh
gczl: tbtp + njrm
fqwl: ptcj * zzcd
fhdr: 2
zfwq: 2
pztp: dnmd * jvwm
jgbd: sljh + fnwr
jssd: ctnr * mgms
tsbj: 13
vggs: 8
nzlr: sblb + vvzg
fbbj: jfqd * tdst
mgpb: zfdn + wdcj
cbqr: zrdg + cddq
bjzc: 16
jslm: zsdd + msmv
pzds: rwcf + twsz
mmsg: npjp + wwgb
gzbn: qgdn - fhsl
trfp: dtgg + fdnj
bgpp: nvlt + jqrl
nghz: 2
rjbt: cfnh * jmvv
zhlp: 5
pqvq: 2
fcrf: wzlf * zchf
bbnw: 13
vmvl: 6
vwgc: 3
cgqv: 2
cjht: 11
jsdj: 5
qthf: pcpt + fvsm
cbdz: qmwm * vcnb
gzsr: 8
fbcs: jgww + wdps
nzts: 3
zplm: 2
cqfn: zrqn + vwlh
mbtb: rgtf + nwdw
pgzb: bbsd / bqff
nndr: bdnv * nqjg
njfb: srrq * dvmv
trrw: fhsb * bphr
cggt: 2
hjvj: 5
hlft: 4
cjqq: 3
fnnt: lwrb * tjsd
glfg: lcsp + clrf
dtcg: lshv + wctn
bgnl: 5
cthh: 2
cnfz: nhvd * zzzz
tngn: ttmz + lzrs
lzbz: 17
tcsl: 3
nrts: spbv + nclz
zgwt: frgn + trsl
zfdn: bzqj * tsgv
llqf: 2
qrsr: fdcf * svdt
nfrt: bpzr * mjff
bvgq: twgb + mlmt
tzpc: qjtr + ttfq
njfw: 4
rczt: cppq + vbqc
ddsm: vcgg * bsvc
blrz: 2
mjlq: jtdr + jgtz
jfqm: 8
zwnf: rqds * nrmr
nlvc: 2
nmzb: 2
qzcd: njfl + vhbt
qrhh: bvtw * jtsj
ptgl: 3
lzmj: 6
vbpw: pdjw - qzrf
gqbq: 2
dwhm: 8
nthb: lchw + nzlr
thqr: drcd + tzdj
bbsd: gzbn + bfws
wdps: cfgg + qnlh
jrpr: zzbp + rnfm
djsl: 2
bfjz: wcvr + vshv
hvgg: 2
cmzg: bzcb + mgcl
lfhl: 5
cjvd: 11
mjdv: 4
zfjm: 3
cpsp: sspm / ddpm
rpwp: ncpr * vtsb
mgjc: mfnw * lcvn
drcd: bscw * fgfl
qfjq: hdpg + hvzf
vffb: 8
blwr: 4
qbgz: 6
gcsg: tsjl - lpdn
jgtn: mctd * ddww
mznj: nbsl - cjss
vhmp: 2
qwzv: mhmd + hcfn
qmsf: rszv * tthb
sptp: 8
bldq: zvrw + qwsv
hfcg: djmd / vwch
ttmf: 3
hhbf: 3
jthf: 6
vwvg: rfhl / vvrq
vcbb: 3
fnhl: 11
zjtd: 3
czbt: 2
jqvj: 2
gmng: sbwg * fshz
scll: 3
vhsb: tlzj * zqpc
sgmv: bvzc * qgcp
rhjh: 4
tjrq: 20
mrqw: rntn + mswf
wcvr: 4
dzps: lvnt * mzcf
lqwm: 5
hgwn: snrs - brqq
zqpc: 5
qnch: 3
dvvz: 17
gfbl: qprs * ncml
bmjl: qzjr * rgwj
bdhw: vshl * nmhq
bndh: 6
stzn: qspl + wzcd
dvhc: 2
dbvw: nhrg + qlcw
jldb: fhsm / sfcm
ghsq: bdhr * hcbv
tfjd: 2
sfzg: wtmg * ngzr
hhmq: mbqf / mrjj
lgcs: 12
jnsp: sptp * mlht
njbb: 3
gjmj: 8
ncwl: 2
ldpp: nzts * bnhg
fqqr: 2
qljt: rlwh + trnp
njmv: 2
scpw: 2
wcvs: mwsp + gphs
fvsm: 3
gtss: zlwf + cbqr
tvcv: 2
zqtl: 8
jhwt: htjj + ddsn
jzrd: 16
vzhg: fmwl + gbjw
dgvd: mdsq * fszv
wsnb: 11
zzsg: bfng * pbtj
mhpj: slbh + czps
tqqm: fmlv * frll
tnbz: vhrh + qfzc
fjff: ldws * ngnp
brqq: 16
cfhg: 1
tfnd: bctw + mpfl
mdbg: 5
wlnz: brvt + fstw
rlcv: 2
pjrb: 3
gpfp: 3
hpjj: 6
cmwc: 6
sbbh: wfls + flhg
bhwr: fdth * wfjm
lshv: 3
pdcg: 6
ztlg: jjsc + ptlc
vgdn: 2
twgb: 5
bzhd: 2
ddpm: 2
dwhg: 5
jlmr: 7
ttns: wlgf * rfmw
gvsw: wwhs + fqvd
rpht: 2
zmjv: mtzd + jbrg
pchb: 2
vhjz: cbtq * tlql
fbds: 2
wjhp: bfqh + gfhj
frgn: 1
fmjc: 5
gblv: tprt * qbqp
tjvm: 2
slbh: htmv * rzft
dbrh: 2
hzmb: 6
hzgn: 2
zgwz: 7
qlsn: vcpv * mbdb
zphs: phgw * dnqz
gphs: wmsh * ttrp
tnhr: 3
jjsc: 11
nwqb: 4
tttq: 5
slzt: 15
wwlt: 5
frsn: mndm + nqdr
hbfr: 3
mhdd: nwqb * gvlf
tzqh: 9
fgmp: thfn * hmtv
vjps: fbds * dbvw
pnqf: 3
zczb: qvpv - lphl
sblb: 10
zsdr: 20
fgds: mlhv + pjgp
cmbq: fsdv * wbmw
jcqp: 5
svfq: 2
hlfg: 2
ngsn: hhrw * lpdd
lslq: 5
dpqw: 3
vbwl: 2
mbdb: 3
qsjd: jjbv + dzps
jfqr: 3
gqbd: 2
rcgz: qlsn * wtgl
jzhc: 3
qgqg: qdsm + vdmn
wjtz: ccbh * ncwl
mgms: 14
zqlr: 13
bcfr: 11
qjtr: 8
tnfn: wppf * jqjh
lsmb: 2
pbds: dwhm * zvhg
vcsh: vdrt * ntqv
grrh: pmhz * vbgq
wmsh: 4
hchb: bgnl * dftm
fcnz: ngqn * sfcd
rrhp: sbql * cgln
brph: 3
pdjw: mnsm / tzfz
gntb: 4
fmng: 3
jjhp: 10
qjdv: 3
qsjt: dzvp + jnlt
znzl: gjmj + tbcb
rdzt: shdn + fvnn
tmtg: lslb + nqhv
gcsd: wpbm + rvrt
dwqf: 4
bdtr: wrtw * prcp
mbpp: jvvd + drjm
pvmt: clpc - nrhj
cqwg: 2
qmbc: 1
rjnm: 1
nrhj: 1
wttj: tpvw + vvhw
qbqm: 1
wflt: nght * fgsb
vrdl: pjhd + tftf
tdwc: nrgv * jwst
bhwn: 2
zdsg: 5
tngp: rnhb - snmw
grbw: 5
mnsm: nnrz * tcrd
whqf: 4
pvvp: 2
jjwd: mjcn + qtvj
rcls: wsrn * jnwh
gdnv: 5
rnhb: rwwd * rdjm
zvnw: sjdm - blms
ltsz: cvcz + bwhs
qvps: 5
nhfm: 3
pvsb: 3
gtmw: 2
zsvn: 2
pbcd: 2
szjd: zvnw / qvfl
cptg: jzrd + rppv
zplh: 2
znwj: jdtm * ppps
jrrm: gscs * dbmb
vwph: njbb * vndn
scgm: 2
nrls: 5
bchh: tmjh * fcts
gtnb: wdnl / zcdf
hchj: 8
gzvp: 3
gbdw: 10
hwrf: trfp + wzhw
vwnt: 4
mjnz: 2
drqh: 12
rrpd: 3
gtjw: 19
vptr: 13
htrw: 2
bhdg: 2
dngw: 5
zgrp: bjlm + dznq
hvzf: sjqt * nwfz
pjcp: ljmd * cqlt
mswf: pwlg + mmpl
vjzc: 3
brvt: 2
grpw: pbwb * nmtf
mdsq: qsjd * hzhp
hvgs: 2
zmrp: 2
hqfq: 5
qcgh: qqqr + jmbh
cjss: trtg * svqh
zmps: 2
ttzj: 2
jwvh: hmqm + ddjv
qrhr: 2
fzfz: nmds / zfjm
zzhq: 2
gwzv: 2
zqhz: wqlz + fswh
fdng: dzsr + cqnm
dbjn: bfmb * gjpt
wptn: zdms - dfbc
cjpm: gvlp + sjlz
jcdg: 2
sqrh: 2
rvbw: rjjd / zmps
mvzz: lwmg + svlh
mlvg: qvdz - zlpd
qqqr: gtzd + nbfz
vshl: 12
fdqt: dlzs - wfpf
mhhn: sssf + jnvd
dtdc: hgrg * nhgf
gzbh: rfhz + qrhh
zzfz: mffv / hzmb
wcnd: tfgg * vlqp
njfl: zphs + rwtn
sscs: 6
vmzj: ndhd + hjfj
vdjt: 17
wwmr: 7
vvzg: bglg + zmnt
szhg: 4
fnqf: 5
hfvn: cqgb + dznl
rwlv: nfth / gflj
cdff: 2
drtl: rrpd * dtdc
lwdp: 4
ssbs: gmcf + qwmj
bmsz: pnjz * pqvq
mrsr: stsl * ccmp
dlml: gvjq * vbds
vhwq: 7
pqsv: 2
gvbf: 2
vhrh: vsfb * btmb
lchn: hgwn - ttds
gvlp: wwsb * rdwt
bzlb: 2
bcjl: 18
lcqp: hvzl + mjbm
drcv: grgq * wsqd
tccd: 2
wldm: 5
qvcc: rjfr + bzrc
jmzf: vmwn / tzsl
gpsh: 2
mlmt: hfwr * lnpr
hdfb: 5
wdnr: 1
sssf: 13
nbsl: fnln * zbjt
lrzc: qhpg + jmcd
msmv: 10
mlgd: gdnh * tplh
dvsj: 3
dzsr: nhcc * gpfp
pbwb: jbjf - dqsf
hpfd: wldm + sbtj
qzmz: qlnq + zzmr
hvgt: 5
smcr: jldc * rnjl
hffh: fgtr + mbpp
sfcd: 5
qlnq: znbd * wtrg
mwhw: pglw * dbsg
rwtn: 5
vmsd: 5
lcgn: 5
mtbj: wsgl + lgfc
cmmf: 4
qhbp: wpfw * jmws
hqmq: 2
gscs: lmjp + vzws
gpsg: dsqb + ftrf
bpmn: dnzf + dhpc
qsgq: mhht - bcfr
jmdc: drcv + bpcp
mmml: qrhr + fbcs
qhhz: vsmd * lcqp
wrtm: rhjh + ndvr
nvqd: thcv * wjbp
rpsh: zgwt + tmpz
nqgl: ppfr + bjbt
ltdc: gntb * rcrd
gddp: 17
lznj: vrdl + bncd
bzls: wvlw / rgnq
lfvq: 3
qphj: 13
llms: vjvn * qtcl
tqzp: 8
pcjz: hgnj + tjrq
cgjv: qvsn + tnzp
wgzp: vvcm * tccd
lwzw: 5
dbdz: 3
lsdt: 6
sqtq: clqd * vjps
lcng: zbsv * psnz
jrrd: rmmh * fjff
fhhz: pdcz * nrls
rfmw: vzvm * qrcj
mbvz: nzln + znzr
hdrn: mpqn / bdcp
rzft: 5
jbrg: pnwl * bffz
dqlt: 2
pscf: dzpq + zmpm
drjm: 1
fnvn: 12
nrbw: dcnf * vphh
dlns: 3
gqlz: 2
fzht: tznp * jlqd
sljh: lsrd * dshp
trfz: rjlc - lpgz
pflp: 9
qpfc: 3
bctw: qwdt * cffj
vdcs: fvqm - dngw
tpbh: dfcf / tltm
dtgt: 2
jmvv: 7
fbwp: bfjz + gfwv
jbbc: 2
gldj: 2
vnbd: 5
zzzz: bggc * fwjt
wnpb: 3
lpmc: 5
fjnj: 3
qhff: wdcw * lqvp
bvzc: 9
sscc: tnbz * szzc
mmjs: dvhc * cpbf
shfn: 6
mvzn: 17
jfqd: 2
qcwh: 1
jqbg: stcc * rszs
rwhn: znfp + cntt
wrzp: psfg + trnj
gvcw: jcnt * dtjm
tjjv: rlhf + gfms
fghz: 14
zpww: 2
mfnw: 7
sgzl: jtgt - dqlf
cmvw: 5
phwf: 3
wqhd: 7
stnd: fmmv * lcgn
cbjh: 7
rtnp: rqvg + wqrm
hdgz: 17
fmgj: 14
lljv: 5
vmbg: 2
rgnq: 3
rvrt: 5
rjlj: zjfj + wbnm
jgfv: qpsj * mslh
lsrd: nfhb * gtjw
tftf: zrqc / tdbr
tjsd: 3
qnlh: 8
rrgg: sfgp + fcrf
wtjw: mcqf * tgjs
wfsd: 4
gmnp: hffh * qhqc
vtvh: 2
sjbw: 18
nnsq: 2
gsjg: 5
jgwp: 5
qrcj: 5
jtff: nvnq + fchs
dgsz: 3
pzsl: 3
hznq: gmgj * zvth
vhmn: slcc * mznj
pszp: 3
vqwz: dbdz * jslm
qnpv: mmzv * wzql
rfhz: ndnr + dpmh
wqll: 6
twzc: 7
rhvz: 5
qfnj: jjdl / dtcj
lngz: hvfp * wfvj
wvfq: fhbj * gmjr
qbbv: 2
fcqp: 2
bfws: jhwt / frtd
bhdm: 3
tlvv: rhbt * gddg
vttr: jpdm * vdwj
mtlr: 4
smjq: 4
htfl: 20
vzhl: 3
mvwh: 1
fhpb: llqf * mqtp
jmzw: 4
vbhw: qmsg / bpgn
sfsf: 9
fnpn: 2
wldg: fvrw * jvch
hshw: vdlh + msdt
qhfq: gblv + bgjp
bjrh: mdqh + mbvz
jgtz: 2
vbds: mwhw + mrlb
jjtb: zwsl + chcm
fdpr: jlbf + tqzp
gbtl: gtwz * tnlq
jgcr: 6
qmjh: 2
rlqc: whtv + hwrf
rfqt: 4
cnrt: ncjl * rcqv
sfgp: tvwv / rlgg
qtsr: 9
ccjz: jtdt - hfvn
dndj: tlzl * wcqb
jfzf: 15
jjdc: rjnm + pzbs
mlhv: 6
nqdr: nrbw - vjfw
lwpz: bwtr + sndn
fvsv: zbzg * blmp
dvfn: pscv * ptsl
psld: bslg * crzl
hmjv: qght * gqrp
phbs: qmsf + fsnr
svrd: vvgn + wvfq
jdsw: 2
qvtv: 3
rgcp: 14
hssl: 2
vjfw: brnz - rfgf
vlqp: 2
qnmd: hrjp + qpmp
nnch: 9
ghvj: pqtf - nbtc
svqh: 3
pgzg: 19
jtqt: fcpw + jjjs
ftvr: dqpm * tfzf
wscp: pfzp * jwdq
mctz: 2
pftw: jrrd + mwqw
glhb: gjrf * pdvq
fscc: nhzj - mrsr
pscv: 7
wvjt: fpvn * lchn
zmjc: qhcf - szhg
cqlt: ddjn + zwcq
sfqp: 3
lgfc: 5
srvn: rvwz * qzmz
fvnn: gnvj * vbpl
jvwm: pmtt / dbpz
bncd: 9
bmwd: 2
jjwf: 4
jnlt: 4
tqnj: gtss + cffn
qgcp: 3
sbvp: 2
jbjf: nwms * crcv
tlql: 17
phdw: tqnj + qzhz
jzng: lmql * fpqm
sdsz: srpt * gtmw
brpv: lhmm * ldds
pfvt: pcjg * hsvg
lgqz: wfqv * fbrg
djqs: 2
qqgj: 20
bhdb: 10
tbcb: rczt * gjdd
zlpd: 1
lmls: 9
bhmc: qdpb * czbt
swqg: tpbh - bljg
wdnl: dwtm + pshj
bscw: 5
rbgq: 7
zzbp: qmvs * zbwl
wwmd: 3
pmsj: 5
rsnz: 2
fzmb: qpnt + tlvv
dshp: 2
zmwg: pbcd * jjtb
pqtf: dvvz * gwzt
mgvm: vjzc + jzdr
jdpz: 4
ddsn: vzrj + jjzg
pfbf: cnwq - vzhg
tnnp: ssbs * ptdj
gqdm: 2
hlwq: 2
jvvd: 6
lclq: 2
cswg: 3
fgtp: 2
vdln: 4
ddgv: bhmc - mjtg
lbnn: rwlv * rwqz
jlwf: vtrw - rjnj
cqrd: 7
pcpt: 4
jlvf: nrmv + dhnv
cnsn: hffw + jcdl
wsbp: dshh * csdl
gqcm: chhf - tsjz
qmtj: 2
wzhw: mnwj * bsdt
vcjl: bzls + hznq
wzcq: 7
dztg: tcrs * hgqn
zwqv: rrqw + sdjv
jwft: 10
tgzt: sjtp * swbc
zlrc: 10
tjbr: rlcv * ffzr
znbd: 4
hprz: 4
ppbl: bmpw + rvsc
vtlb: 2
wlhz: ftrh * ppzq
dpvb: vfbp + vbmz
qhmn: 5
vpsw: dfzh + lbvf
fhsm: phmz * hqmq
jbln: mdhl + ssqb
ftrf: zvdf / flzd
ccmp: jvbt + rvgf
vdmn: gcgz / jdpz
jqnf: 3
fmmv: 5
qpmp: vvfd + glhb
lqbq: 3
jmcd: 5
rcth: hdpz * pdcg
jwvn: 2
pjgp: sgmv + rgpj
hnbj: njdj * grrh
dbrj: zcdc + gcql
jhhq: 3
bbqq: 3
qvsn: mlrq * wqpq
wtrg: 2
dzjr: jhhq * sfdd
dfhr: 18
wfpf: 1
hqfl: brwb + hpth
rrqw: fghz + pflp
rqws: 5
cjnc: mmjs + wtlw
wcts: jdbp + dncm
bmqf: fnpj * jgln
rhmw: 1
srrq: dcrq + gtbv
nbpn: 3
qfss: 3
qssl: cjvt * gwst
qhcv: zgdn + hrpt
mctd: 2
cwmg: hchj * rhhb
fnpj: 2
wdcw: 2
vsbt: ggtj * tsbj
zwbd: wnfs + wjgf
jdtm: rhmw + wjlc
swcj: 3
psnz: 2
pwdn: 7
grpm: lgqz + whnl
bdcp: 2
cnrj: cgqv * cbcf
ldws: whqj * mwhc
jcdn: njfw * dvqc
clwv: 12
rpvs: 3
vvwp: 4
wmzd: mpqs * fzwv
vphh: 4
gmcf: 6
rjlc: njdc * jbsr
nhzj: tqqm - vjfj
fcmn: dmcq * vdss
bvsn: 8
jzqm: 2
qwcj: dsdm * zspf
sngp: zmmf - bfrg
qmnp: 4
lcvn: 2
qsjw: 11
vtfl: 2
dvzq: mhps - vsfw
jjjw: ngch + wfvv
vzhb: qlzg * rhnn
vndn: 2
jbtc: 2
cghv: 5
bjrt: jwqg + hccn
blmp: 3
pglw: tnwv + tjbr
mpfl: zqwm * qphp
wtgl: 3
fpqm: 4
lzrs: lrvl * lngg
lqtn: 4
htbp: lgbg * lpdp
qmss: sdrf * gsvv
bzjv: 3
dnqz: 9
qjlg: 3
mdhl: mtbj * rphf
lgsq: vtjv * nzbv
dbmh: zgwz * gfzh
cpvv: whqf * drvn
nndq: 4
fzbg: 5
dzvp: nwdl * vnzp
svvb: vtzc * rpht
gwcf: 2
tlzj: 3
jwgs: rjwl + jzdj
jtwf: svrd * rgpg
chcm: 17
brrv: 17
qpwn: 3
zlwf: 1
pnfn: 3
sznv: 2
zjfj: zznh + lvpg
thcv: 3
spjf: wcts / lnbr
vfjr: 6
hggg: qwcj * gqbq
cmrw: wgzp * nrns
jzdr: 3
pbtj: gvsw + gbdz
hzcl: 5
jdpf: 8
cthr: 2
trnj: fzzg * fmjl
svdt: 4
jhht: zqpz * mzdl
vbgq: zsdr + hqfl
jnvd: ttmf * bmsz
mdqh: smcr * bldq
fcvw: wlmd * qbjf
fptv: 5
lrdp: gpst * lzbz
nqjg: jmdc / vhfr
bpgn: 2
ctbd: 10
qzfg: hchb + vfpj
zrhm: 1
qmsh: tcnz * vthb
cbsh: ctgb + njmn
fbpz: fpwr * qlzh
vtst: 6
ggtj: 7
rnvj: qjdv * rjlj
bwtr: 11
jpqn: hdws + wttj
jldc: 2
vwch: 2
dqds: 5
fccd: rjvj * mlvg
thfg: 5
rppl: 17
zqsv: nndq * fbwc
tdrz: wflt * httg
sndn: 2
gmll: pzds * brvf
ppps: 2
pjwb: qhrn + qwrj
hjcf: 2
dfmv: 1
jgln: qtmv * rpwm
bdnv: 2
bsdt: zvls * bpvq
qjzz: 7
flhg: jvph * pfvt
jvcs: nthb * tcsl
rdvj: 1
nvzb: pqfz * qbcj
hvsc: 8
wccs: 6
psfg: rpsh * sjbw
swjm: mhbf * wqhd
bhpw: sbbh / wfbh
jqnp: fbwp + qrsr
vvwr: 2
fctb: 4
hzzq: 7
rlqr: mgpb / nprp
hdlr: 2
flhz: 7
hgdp: 7
rbdh: znsg * pgzb
jjdl: qqvg - wlhj
thln: pcjz * vbwl
rvwz: 19
jvph: hjvj + wfsr
tmns: 3
tpvw: 19
zvth: 3
qwrj: 2
hczl: 5
lpdn: 1
jznn: cmpr * grzn
ndnt: 5
hmjq: 4
czps: pzlt * rgcp
zrvq: hzzq + vwbc
dwnm: dfqz * hssl
zhtb: cwmg + vtng
fqvd: 5
dvpm: pvcn * cpsv
rpqq: mvzz * ldmf
glmw: tlcb * vtfl
rhmj: 3
sdpl: 3
nght: 2
jnwh: 4
llft: 4
zdms: bngm * dfdm
ctnr: 3
fdsd: 2
vdlw: 3
nzln: bzhd * fscc
bdrl: bdhw - gzfp
lvnt: 7
pszl: hzpb + qmnn
pscg: nnwr / vfnm
cwgv: qmtj + fzbg
qtdq: 2
ngvj: 5
npqz: 2
nwdw: 2
wbsd: 2
chct: 19
hzcz: 1
mgmw: wjvj + vtlm
jjjs: wrhb * flhz
gmpq: vgdn + vqpb
tcmm: 3
fnwr: hfjj / hlwq
lczl: 3
fchs: 5
zchq: zbdr / jwvn
vbnq: 4
cggn: 7
jnfj: 3
mnwd: 3
jbrr: 9
frbq: jflv * fvtb
qdvj: 2
ptcj: 2
lpdp: 3
lslb: nhvz + ctzh
gfvd: 6
qlcv: 4
zmwt: bdqh * jvsr
bsjv: 2
hsvg: 3
vspg: mnwd * vdsb
znzr: zqjq * bsjg
nmsp: humn - dnhg
zszf: 4
fvrw: nbqz + tdzr
fljh: 4
sbql: 2
ttmz: pftw + vtsh
bdhr: 7
bglg: vffb - rdvj
qhzj: 7
ppzq: 4
zlch: jzvd * cthr
wdmg: vwgc * hwvm
djvp: 2
ddjn: zhmc / fcqp
cqgb: tffz + qsjr
cssj: 4
dlzs: trfb * jthf
fqlz: bsjj * brjs
rtrd: mqsw * qmjm
cjvz: qhzj * bqlr
bmdp: trfz + zwsn
wlhj: hwlp * vtsq
thdr: 5
vnzp: 2
gwvp: bsql + thcc
mvpz: nbpn * hnbz
rwqz: 3
wjtj: 2
gpst: 2
ncdr: mgnf + cdff
rfhl: bhwr + mfcz
zzmr: 3
rswp: 4
zfmj: 5
dfdm: 2
rlrm: 5
dsjr: bmqf - bgpp
pdth: 10
tznp: bbsl * dslq
cvbl: 8
wgdz: 17
ztws: vchv + bchf
zwcq: fqwp + ldtn
bfcj: 13
ftrh: hcqw * djsl
zljt: 7
qtmv: 2
phmz: jssd + wgdz
vcgg: 3
hwdb: 5
gbdz: nfdf * lclq
ptsl: sznv * lfvq
zlfz: 2
zlsf: nlvc * jlmr
fpvn: 5
dhrb: 12
cjrc: wjdr * rlrm
fhld: gwcf * nvgt
lqzp: mrqw / hjcf
zcnl: 13
ndhd: 13
ngvr: 2
nhml: 2
zvrw: fglv + ltsz
znpf: hjmr * trbr
dljj: rfzj * jlwf
lfmf: vzmp / pjdt
bvgj: hrcj * djvp
nbbw: 5
ncpf: 3
zzwd: gtnb - ftvr
glvz: qmsh - cbdz
qshs: ncgv + htmh
vlcf: cggt * lwwl
bwmh: 2
ngnp: 2
hmwg: 13
mwsp: jlmh * ncpd
rbmv: 2
tpnz: vhmd * bcmc
bnjl: 5
mzwm: bdlj / grpz
gggq: 3
wlvg: 17
bcjc: gdsf + pjrb
ncml: zqsv * gtmh
jnbv: mvwh + vzdg
trbr: 4
lchw: jzqr * dlns
smww: cnhz + wwtc
jflv: 3
nmjh: hlwg / tttq
znnr: 5
brjs: grnj / qpwn
fwjt: wjtf + lngz
sspm: hzmp - zqjj
mwqw: lbpb + cmnw
wmcm: qzln + wtjw
ndlr: 4
lqqd: 4
slcc: 2
zqpz: 2
pshj: zmjc + ggds
fzzg: gwrw + wccs
gnvj: pvvs * rnbt
nnrz: 6
bzrc: cfln * rqrj
hgdl: 7
wjbp: 12
bphr: 5
tcrs: 4
jdqd: jcdg + cwfw
ncvb: 6
hgvw: 5
dbpz: 2
llsq: ztlg * shfb
wctf: 4
svzn: scgm * gzvp
bptf: jwzl * rbwb
zrqc: dgwr * vlzz
rbrw: jrrm + dfvw
rtvn: bggf + pbds
gcgm: gpzm + ldzw
fszv: dtgt * bnzh
wltv: sbvp * mhpj
ctzh: vfjr * hthb
dgwm: fbrt * bwbb
clhv: 2
gsvd: rcls * lqqd
psnj: 2
csbh: 4
gqvj: htdq * mllr
qtvj: 20
szzc: rhvz * lsrv
pcdw: rrgg + gcgm
zhpj: 2
zbzg: rfbj + bjzc
fhbj: 12
qgdr: 4
pmpv: 2
fhsl: 8
sfcm: 2
qhrn: 9
wcqb: 2
tjns: 9
rvjn: 2
gvln: 10
svfc: 5
rqds: ljmp * nltb
wjwh: ljlj * gzbh
sdld: svtl + cnff
zsfw: 12
pbtp: jjhp * jbbc
hnjh: qgdr * rtpw
dfvw: qwqn + hbmc
nghc: 19
jlzp: 4
csdl: 3
wpbm: 2
wsqd: 2
cfzd: dqzl - hdfq
vrzs: cnrg * dhqh
vjmv: fhcn * wrzp
cvjh: 5
nhfs: 11
fwfp: jzqm * fqsn
mhwm: tvrt + jjbz
bfmb: 3
btdh: 3
mpgg: zjrw * svfc
vzvm: 5
cwvz: 3
hhbp: 4
bpnf: 2
cfgm: 4
hhhc: 2
jmth: qhzg + pczp
dsmj: mfhb + zljt
hblh: clfm * wtws
srbf: ppbl + hfmg
rntr: 2
tzvp: 2
hfwr: jswv * dtcf
qszd: cglw / nltj
mjvz: dlrg * gbcg
nmds: jzhc * vwvg
qhsj: 10
cwfw: vbzt * wcwt
dwzh: bptf + vcsj
zvhg: 2
gjwp: 2
bfqh: fbzr + jsdj
rjfr: tngp + ntcf
rqbf: 3
czsc: qwzv + qhcv
gpfw: fmbf - dlml
stcc: 2
bqlr: 5
mcqf: 2
njvj: gggq * qlcv
cqvv: qvcw * qfjq
mzsv: bhpw + thql
jjdr: 2
tzdj: 12
mbqf: fmrr * gtcc
znpj: 2
trtg: rgzs + vmvl
crhl: znpf + znzl
mrjj: ddtp * rswp
vhfr: 2
lphl: 12
ddtr: qjfb + jlvf
npwm: 3
tsjz: 3
dznq: bwgt * gqcm
tlzl: qnmd + hhmq
bbcb: ngvj * bwjp
ssqb: fdqt * gjwp
cnff: dlhr * zqtl
frtd: 4
sbhg: 4
rwqv: lgcz + dgwm
vsfw: hshw * jmzw
qjfb: lmls * bcjc
lnbr: 2
sdrf: 2
zdpc: qcwh + fnqf
vwbc: 4
tltm: 2
nbfz: 13
dcnf: 20
hrfs: 2
cphf: 2
gtjl: dqpl * bhwn
rlzz: 1
wftw: tmmg * wnpb
whhp: 5
nhcc: 2
dhqh: smww + rsdg
lgtm: rrdl + nsgd
ttds: 9
vvnr: 3
zcfp: 2
ljlj: vdjt + qbgz
vzqt: sthq + zddg
wsgl: 16
bztq: 17
wtmg: rmdz * cwgv
wdsp: 2
ttqt: 3
smpt: fwfp / ttzj
dnrt: 5
twbw: swjm + fffd
cwnl: cggn * nrfw
zbzd: 6
dsmb: 7
mlht: 3
cndd: 2
fmjl: zdpc + bztq
hltm: 5
pslg: lffz * rcgl
ddww: fzmb + lvjd
vbmz: mpwj * pszl
fzwv: 7
gfzh: ftmj * ntgv
gtmh: 4
tdbf: 4
zczf: fhpb + rnvj
mllr: 4
fmlv: 3
gfwv: gdls * cvjh
bnmh: 2
mzmh: 9
snmw: cbsh + grpm
lpgz: 13
fnln: 2
hlrl: 5
vjfj: fptv * twzc
zllb: fzdf / vgql
lshs: pdth + hmwg
gqjs: gmlz + jccc
fpgp: 2
njdc: 10
gfwg: 13
cjvt: wbsd * zvgw
htmv: gpgt + mtsq
vlgv: dpvb + nfvs
lpll: 2
rlhf: tvmv * gsjg
jdbp: 3
rntn: szjd - vrzs
zbhg: 3
qwsv: 16
wjvj: jmth * hbds
ncjl: 2
zrqn: 3
wmwl: mrcc + qqgj
jjzg: ljrg * zqst
fdcf: tfjd * lfgg
gmgj: phwf * nhfm
qlcw: 5
jmws: 2
rmfd: zqhr + fgmp
vzdg: 6
wszs: 2
tgjs: ftpt * brrv
dlrg: 3
lwrb: dwzh + jwvh
lvpg: fcvw + qhsj
znfp: 2
fszn: 1
jccl: 17
dwfm: jcrq / hlfg
nqhd: dgvd + spmr
fmbf: sdpl * glmw
lbbp: 2
plbl: bpwf + qvvl
wfls: grbw * qzfg
vwqq: tfjp * rcgz
mdfq: ljwl * ttqt
jmbh: bbgf - gbdw
jbpd: 2
lzsh: hwsj + rlzz
jggl: hbjf * vtzj
lbdf: 3
vdwj: bwgc * whhp
fcpw: wctf * mpzl
dcrs: mzfg * hdrn
zgdn: bdvq / pwdz
jlnq: 7
btjh: rwwm + nvqd
qmjm: 2
nvlt: jqmw * psnj
zjwh: 2
hwvm: jtzl + msnl
mmpj: 3
dhnv: tdwc * dqds
jcdl: 2
dwtm: sdfb * jgbd
chhf: rntr * vzqt
jdfw: qhfq * rbrw
jtqq: 3
vvfd: 3
rphf: 5
msnl: lrsg * zdnc
hqlb: 2
ptlc: gfvd * sfqp
pcvd: 2
rfgf: 10
dbsg: 11
tthb: 20
gjdd: 3
vghc: 3
dqsf: nnch + cgjv
dsdm: qsjw * hvgg
gpzm: ptlg * smpt
brwc: jbtc * qsnb
bnzh: vbhw / vtlc
vfbp: vdlw + pgzg
wwcz: wjjv - sdsz
cddq: 2
lbqf: wpfg + cpvv
qvpv: clbm / btqv
vnvb: 2
nmpf: 2
wwsb: lczl * frjm
zstb: 3
nbsd: jdqd + wvjt
mmlz: 9
zcdc: 11
zmpm: 9
mmzv: qsgq * bqpl
mnwc: wmwq * tqhz
svjc: zqmf + wcpd
tdzr: zdql * vjmv
sdhh: 2
hcqq: 3
zflt: fcwc * ndlr
jvch: 2
rcqv: qvtv * nghc
cbpl: 5
bnhg: 2
ngnw: rppl * zfmj
cnsr: bbtr / mnbt
pccn: pdlw / qmdw
tmmg: 2
bngm: 19
mhbl: 5
bdlw: 6
bzjb: 4
ldmf: 2
vfph: htmz * bnmh
rjwl: tmtg * mtlr
jnvh: 2
sbtj: 9
scfh: jdfw - pbhv
sbzr: zgmn * lfhl
qvcw: 5
npjp: 17
jmgv: 2
gtzd: ngth + fqmg
wcwt: dnrt * qvps
vgqz: jpvw * czwr
mbgf: 10
tgnc: 1
pvcn: 7
tvrt: nbln + dndj
gbcg: wptn * cmzp
qvfl: 3
tfjp: mgjc / zpww
nfvs: 10
jcnt: ngvr * dqwl
qcfp: bstr * zrvq
mgfp: gvcw + lrlc
lswp: 13
tlsb: 7
dtcn: 5
qlzg: tmql / pchb
cbtq: 5
gfpg: 10
hfmg: wmcm + glfg
lffz: 2
gnhs: mstw + nvzb
cwth: hlmg * zvfr
qgwt: 3
pqsc: 7
gzfp: cnwt * lbjh
vdrt: 9
pcjg: 9
zchf: rpqq - vghc
vcsj: mmhm + tfhg
vvhw: zmwt + vnwh
cpsv: 3
mccq: 4
jswv: 4
nnwr: pcdw + slpm
cwph: nphv + llfp
bfrg: 12
gbhl: hvfs * djmw
mhht: znpj * jbgn
mnps: 3
rlgg: 2
vchv: ssnb + mhdd
bltd: 11
djmw: 2
tcfs: 2
cbcf: mdcz + dngq
hbds: 2
ggsq: lqzp + jfqm
vtnh: lvcq + pzfg
fbrg: 5
vsfb: 3
nfhb: 17
rhsm: zcnl + ldpp
mjvl: 3
qtph: 1
lpdd: zqhz * dvzq
ptlg: 2
nwjg: 7
ncpr: hprz + bhdm
bqpc: 4
zdql: 13
jrmq: wdsq + ljnb
fhdn: 5
bvrh: 2
jwgc: 3
lrvd: 3
fdpq: 4
zmmf: bfbh - fcnz
fshz: 3
mcqw: hggg + jptw
lrvl: vnbd + rrgh
vhmh: 11
qmsg: bwmh * gmnp
bcvl: jpqn + vvwp
vtlc: 2
fppq: jbpd * jnfj
gfhj: jrlb * bblv
cfnh: gzvd - rsnv
rdtc: 3
bwjp: 3
vlwq: 2
sdpv: 5
qpbl: 3
zvdf: gqsv + qssl
fsnr: wwhg + hmsc
fhsb: lpmc * vdlr
wjlc: fmjc + ndsv
hrzr: jgtn * qhff
lcsp: ghvj * zjtd
pwlg: 17
rhbt: vtmd * lcjs
thcc: hlsr * jqnp
dvwt: nrzm / qtdq
wtds: ldhs + ltwv
wdwj: 15
qhpg: lwpz * jsss
wsrn: 2
zlwv: dftb * zrqs
tnzp: zcfp * mpfh
ffzr: pvsb * bfcj
pjdt: 2
ngth: jvlv * bczp
bchf: 12
qqvh: dqlt + qpgd
lngg: ddlw + vwjq
lscd: nbdp * cssj
bjqh: 2
root: qhbp + vglr
bczs: 8
wjtf: wdmg * dnln
stsl: 3
njdj: nswf - lrvd
clfm: 8
fsrh: zbzd * ggcl
jlmh: csbh + hdtf
qmsm: 2
gjld: 3
dnmd: 3
mpqn: sqzj * sjrj
mnbt: 5
vnwh: 3
cffn: 20
wbbd: rmlf * smjq
jwst: 3
zvgw: 4
gjrf: 4
fpnl: pscf + gbhl
fhcn: 2
bdlj: qfnj + grpw
jmnr: bndh + vccw
zhmc: rdqj * mbgf
qlzh: qzcd + rwqv
clrf: svfq * gqjs
qqzh: 14
dpmh: svjc - rvbw
grcl: 7
vccw: hqfq + bcnj
rjnj: 6
tffz: jgwp * hzgn
tfqp: wlvg * mgtl
hgqn: 11
wjdr: 3
fdnj: jrpr * lmrz
mphz: pcvd * bwtj
qjqj: nhfs + jlnq
zgmz: bzzz * tnhr
wbmw: 4
nvnq: blrz * gpng
fgsb: qcfp + mtwm
fjrb: cqvg * jzwt
lbjh: vspg + lgsq
jzdj: llsq + zmjv
frjm: 3
bsjj: 4
rhpm: qzwj + slgg
vqml: gmmp * tjgm
jzvd: 4
cvnz: tnnj * tjns
vnqz: bhdb * wrsn
vpns: 4
rsmv: fmgj / hhhc
cnrg: 3
pqfz: 5
qdlm: vpld / zmrp
sthq: qnsd + ftnh
rtpw: 2
wlmd: gwhz + qmwb
zqjj: tqhq * sswc
jqwm: 5
vbzt: 2
qbjf: 4
bzcb: 11
tcjb: 2
sdnw: thln / jjdr
tmql: dzpj + gqvj
hbvw: prpb + nnts
vbsw: 5
rcgl: 3
jlbf: 1
hvzl: 1
spmr: sfzg * pztp
bfnd: 1
wzjb: 17
hffg: hmjv - hbfr
lfql: 2
tlhr: wwmd * bhfq
wqlz: bdtr / gcsd
vshv: 3
fsdv: jlbr + fhdn
qrlz: qmsm * pfbf
mphh: nfvz * vqwz
jtjm: 1
jqrl: 5
gvlf: 2
hccn: jnrg + pbtp
tcrd: 6
cgbw: lrzc / pmsj
dtgs: 3
zdnc: 13
nrmr: 5
mjff: 5
wswr: vlcf * qtrv
rnjs: jqwm + ccbt
qmwm: 2
jhlm: pblc - zsvn
tvfb: dnwc * bbnw
jjht: ppml + qdlm
cffj: fnhl + njvj
fhfc: 4
wqpq: 4
zzcd: 7
hlvv: vtwc * wjhp
pgrz: vsbt / qthf
bffz: vvjs + mzsv
qmnn: cmvw * dpvq
tsjt: qphj * qszd
fdvf: qjqj - hzcz
vvjs: jfzh + mrfl
drnn: zhvn * tvcv
vsmd: 3
hthb: zjwh * nbsd
tfzf: 2
nhvz: vhmh * thqr
pjqn: 9
bpzr: 2
gvjh: wjqd + cnsr
mstw: 9
hmqm: fsrh + sfsf
pjhd: 1
lfth: gnbw * qwsz
bwgt: vnld + vwqq
prvn: 19
fglv: 19
ggcl: wftw + qpfc
hbjf: 3
zcdf: 2
dwbq: fvld * cwmz
rgpj: 3
gflj: 4
ttfq: 3
bpcp: 10
jlbr: 8
mpfv: tcjb * mwpb
dshh: 5
gzvd: fqhh / sbgq
gtwz: 2
dsqz: 9
jjld: psld + cjrc
thfn: 7
mwpd: 20
rvsc: brwc + lfth
whtv: qvcc * tjjv
nfdf: 13
hjfj: bdlw * rmng
nltj: 2
zllw: 12
ssbv: gddp * ftmd
qwmj: vtst + lgtm
bsjg: 19
dnhg: vhmn * jgcr
ftpt: 13
fgfl: 5
bggc: rhpm + lzsh
lmql: pszp + fgpv
nmhq: lfmf - nbbn
zhwp: 3
bblv: 2
wrtw: 9
rwcf: jlzp + chct
pwdz: 5
mmhm: qbbv * ftrw
fzdf: dsmb * bnzm
shfb: drqh * nnsq
nhvd: jfct + rzmm
cmpr: 5
frcz: 4
gwzt: 2
zfcs: 9
ftmj: jdzl + jggl
qnsd: 3
ldzw: zlwv * blsg
vwlh: 6
cgln: 5
tmpz: 1
pmzs: 14
rmlf: 2
wfvv: 2
dbmb: 3
nhrg: 3
pdcz: 11
jdzl: 1
rztt: 4
ljmp: qnrp * ftzz
rwwm: cvbl * wcvs
wtlw: hjnj + wszs
vwjq: 3
fvqm: jszv + clhv
ddzb: lsdt * bchh
wjdt: htrw + zhrm
bqvn: hbvn + czff
gdsf: fsff * lqtn
jtdr: vhsb + gqdm
dftm: 11
gdjv: gvbf * wrtm
qtcw: cphf * sdnw
cvsb: 2
gqsv: 1
drzq: lwdp * bzjb
vfnm: hhbj * pnns
cdhs: 3
zvls: bfhc + dtcg
gqrp: 11
mdcz: njmv * rvrr
mhps: tzqh * wvqp
nrmv: fqlz - vbpw
swsj: 3
mzfg: 5
pjjr: ngsn + cnfz
hmsc: gswg * gcsg
mrfl: dbjn + rjpr
rgwj: 3
bsqq: swsj * dvpm
tljh: plrt * sdhh
cppq: 11
vtsb: 2
zpfd: 1
rdwt: bcjl + lslq
tdst: 13
vhbt: 1
mgnf: 5
hrjp: 4
wdcj: qljt + ddgv
nbbn: svvb + rtnp
rmmh: 3
qpmm: 3
qzwj: 2
bcdz: whls + wltv
prcp: zllw + vmzj
twzq: mmnb * fjqg
dznl: 2
bvps: 3
tnnj: 13
tnlq: 3
slgg: mnps * rbgq
bczp: vwnt + drtl
tcbr: mphz + htbp
zbwl: wjdt * dpqw
bfhc: 1
swbc: 4
crzh: tsgp + slzt
jwdq: wwng * jqnf
lmcj: ppns * pslg
snrs: drzq * pgfw
zpmg: 3
vhtz: 4
szrq: 3
mntl: hmjq * zllb
pppl: jtsl / zszf
dlpj: 13
znfl: 3
hlgc: 3
pgfw: 3
nhgf: 3
jbtq: qmjh * scgv
cntt: 11
nbqz: ccrc * glvz
gtcc: cptg - ztqn
qvvl: rztt + hrfs
qtcl: gfbl + plzp
rpwm: wmwl + jhlm
fgpv: 8
qphp: 2
wwgb: rtvn + pmzs
hwlp: gtjl + ljpj
mtwm: 15
dcrq: fvsv - bbcb
ztgp: rmfd * qfss
zftr: 5
tjbc: gmll * zzhq
lbvf: 1
wdgd: 2
wbnm: zftr * ghsq
zldd: 7
ddjv: fqql + dztg
mjbm: cghv + qtph
hzmp: zgrp / bbqq
ljwl: 7
htjj: pjcp / hqlb
hgrg: 3
hffw: 5
mvdm: 4
rdqj: 4
mzbl: hbvw + cfhg
wrjw: fjnj * cjfn
vzrj: lvsp * vpsw
crzl: 2
dmcq: rlfj * jmgv
tlcb: wldg / vhmp
mpqs: 3
rdgj: 19
jlsj: 7
fddj: brpv / bvrh
twpg: 4
rqrj: vfhz + btjh
gwgt: 2
qsjs: 4
mnwj: tfnd * fgtp
qgdn: cpvd * lfmn
zrdg: 4
sttp: jznn * ddlb
hbnq: zwbd + tljh
sltr: 9
ppml: lcng + pjqn
nbdp: 2
mndm: 11
mpwj: 3
zfcb: 2
bfpc: nzdw * vbbw
nwms: 2
ljnb: 2
qlsb: 2
qzrf: 4
flbz: cfgb / nmzb
vtht: ssbv - pltn
jnhr: 2
jpdm: jlsj + cchb
qnrp: 2
mffv: shfn * vdcs
fqzp: cnrt + dcrs
jgqj: hpwz + csng
bfbh: frsn - gsnq
hpth: wwmr * wwlt
ftrw: 3
qspl: hphc * vgqz
brvf: 5
rjjd: jbtq + cnrj
vgjv: 2
mfdq: tzpc * mtlc
qgqv: 5
dbmc: 5
bsql: zmqc * gmpq
wwtc: 15
jwzl: 2
fgtr: cmbq + zsfw
frfw: 2
nrns: 11
jssj: 6
grgq: zldd * vtvh
lsrv: 2
wrfm: 2
zdsc: zmcd * qjlg`
