package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GC4UGkWxIAERBg4GYrLE9PTU4v0obReVnF+XmgIKwPjjddcCT2TDvodchBoe29zdEc6E6ui2p09L8/OZndOCW4rm1xWvtl6RcuXZ8oVkzKeJO15MVdYX3lBfuKutU65qt+u/Hw+u76srCyKmSHhGr+h2+mqgocs6m0us3ZuiD+/MVuOJSOCP2L3w40xokVLIviX3T6lGHWfw9fp5BRZlqyUip7pgkXK7BeyZEv3HvZbzj37d/SJyRuqauvibUyv63/+evhT8b4bZ+41ZE8sL9ynonHoysYEjuPpy3dPWa9SOP+Ssqr1J1Gp2fIrrn7YdPCTzeU5TtFzFkZyzBVRmHpkzZ0VD2/8L61nbfrQNsvqkdMmoQvrp8t2slff8TPouvK/541F386ZS4Lrdd6HbXY/0aUd0BymnXkxjO/lhG3P/85s6Tq7O65gnY+JzYqMn8HP7snfz35rMStje72324kFNcz344Nmaq52TZVOrffU/cnFwPD/f4A3O4fWhKVRjowMDArMDAywuGDAiAt2RFzAgx+kG1lNgDcjkwgzIi6RTQbFJQxsawSReGMWYRR2p0CAAMN/xx5GBiwOY2UDyTMxMDF0MjAwnGIE8QABAAD//zXlxphpAgAA"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
