package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GAIrzYLYkACIgycDMXlienpqUX6UFovqzg/LzSElYFx+zuuhJ5JB/0OOQi0vbc5uiOdiVVR7c6el2dnszunBLeVTS4r32y9ouXLM+WKSRlPkva8mCusr7wgP3HXWqdc1W9Xfj6fXV9Wlh7FzJBwjd/Q7XRVwUMW9TaXWTs3xJ/fmC3HkhHBH7H74cYY0aIlEfzLbp9SjLrP4et0coosS1ZKRc90wSJl9gtZsqV7D/st5579O/rE5A1VtXXxNqbX9T9/PfypeN+NM/casieWF+5T0Th0ZWMCx/H05bunrFcpnH9JWdX6k6jUbPkVVz9IxZlF+Rmc1dxyVHNW6wVPxiDPM6d6pfr+xptHKHDkr1iprKeizfi36ubEig3H37Is9vh7/kTf4xmzPa2+H/+UZ1rdsmYHk2nG1kliuU8n8GVW35vvuaKi+jr3uamdbeyiuvWn7rB92nutb3FnmXTCSQ7JKfkPfntrGaduXZGz+I3+StV4RgaG//8DvNk5qi2a7zkzMjAoMTMwwGKDASM22BGxAY8AkG5kNQHejEwizIjYRDYZFJswsK0RROKNW4RR2J0CAQIM/x37GBmwOIyVDSTPxMDE0MnAwHCGEcQDBAAA//+eNO4hawIAAA=="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
