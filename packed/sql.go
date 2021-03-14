package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GBw0c0LYkACYgycDMlFqYklqXrFhTl6+gh2aAgrA6PBIt+EJ2G+/YscBPZ+1n7VcdLnUMr3DdKPGReWWCUlz1Kznjcr7FakEYdwTnr6/79sgvaSvZvlzPTUV6wL/7rrVWgrw05GBoe3imwbrmzO/3zy+rbEK9/mnZ/wZaXEgwiL/hc8TC8arpr1/K7eYf9xyvMYvyBW7lid5VWRGjOesopUrjqXFPusv+F2LH/jF4b5HgeYO7wyrpjxtC88W1gbF3kwZmLvJIsmz09zUqScpCbYuYqvON/YNo+97ESt44rDBR1T7od2ZC5cErZxglVTRNbSyTuN9m/Raum2tf6i9eNW/U+p+Cfyq0NiH80/ubNL4ezm9LX3TC36zTbL9vEwPSmtPvDRdr3OwV/REfd1vU+4+EyeG6Ds1mYmfZjheMK0qZmOEgK32vXSO6byL9Szm5paLZ4qe5Qvunj6L7krGYyX7BZyyCpETtt9t/mSE7sc34zyvh3ls24Y/29/8TMnYzrH93K15ye/fI6K1gtR/OdUIP1n46tt7lu/hE0L/hVtXhH6/HPd77urm+/nfpQvL/6mZPZFq++Z5pTHlxuMZwfcY1krfmDWCTO+w3JGS/jYZkuucXkWXHDRue+AwYMeiz0mrcrX7xnO9YzdJ7tsXfaG+bs+7VEpba75L7h23tE6nctprfcuxOmsL957V2LvqaOPwuOOhHOp318yzZUlf9OMf1NafyZEhr0uFpTWq3+oFs97Kje3SFdcP33j2cuBTv3Lk1P3zn15odNvvalU4byVHGqV87Rm/mWfZPNz+e8L/4QZGP7/D/Bm53hXXL9KjYmBQZGbgQGW5BgYotCSHBdKkgMns5hFvgkgA5CVBXgzMokwI1ItsuGgVAsD2xpBJIE0jDAMu3sgQIDhv2MxEwN217GygZQwMTAx9DEwMGxhAvEAAQAA//+dW6qIWAMAAA=="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
