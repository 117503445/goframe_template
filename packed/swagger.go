package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GBYKmcexIAERBg4GYrLE9PTU4v0obReVnF+XmgIKwPjsrncCT2TDvodchBoe29zdEc6E6ui2p09L8/OZndOCW4rm1xWvtl6RcuXZ8oVkzKeJO15MVdYX3lBfuKutU65qt+u/Hw+u76sLD2KmSHhGr+h2+mqgocs6m0us3ZuiD+/MVuOJSOCP2L3w40xokVLIviX3T6lGHWfw9fp5BRZlqyUip7pgkXK7BeyZEv3HvZbzj37d/SJyRuqauvibUyv63/+evhT8b4bZ+41ZE8sL9ynonHoysYEjuPpy3dPWa9SOP+Ssqr1J1Gp2fIrrn6QepuzZU7CrOCM7mAT7oWTmZZMmjFV2lTm9/68DQ4C9wKC3WM8wph+fd2k/DGh/zan28Rf8yfKHjc0n1zyrv/p3axvHJGFrFkXslR1b51SlL30bY/95ICP39aJz9EW5uXXivo3ZSfvk+q1su7Cb80OTBI00rh/4Pv0kLRTWYE33W7GB3nvY2Jg+P8/wJudQ3he31RnRgYGJWYGBlhsMGDEBjsiNuARANKNrCbAm5FJhBkRm8gmg2ITBrY1gki8cYswCrtTIECA4b9jHyMDFoexsoHkmRiYGDoZGBjOMIJ4gAAAAP//bUKt82sCAAA="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
