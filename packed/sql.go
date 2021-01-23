package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GDY1GgexIAExBg4GZKLUhNLUvWKC3P09BHs0BBWBkYDb56EN2He+Y88BPYpf5R/JcA5wSnF+uYGzyMHN3JnJN4K3H2Td+cN3rxN5z5NyjXyrPu74qw2/7FED6WzHo+FQ8prQ7+G1upVn2icJfslYR7ff7nv/npdrUmOkdJ6zpeW1ElX77vqtPqto+yWNzuiWrPOTY59sMh+0oQlH2smnEvmPJzJ87FEZbLsndxncYzm+rIslxj2qMl9Vzixp7k76WBSRlFMXMQsxlNHmf584Fxk0FBXIu0b9CKo30TzlkMs35bXFgllO18LZxm+UXjavtOhOaIrpG2fBzOHq9blDJ3Tx6eyTujl3fxaJ+Bz4rPNH4wTnCwnzWcwFkyWPJBhtdP3qTNX2OmUZUcnNkq1PE1oeGrV96j25/3r50zOik+077jXX3DU43iKbUYy/6lVkqIvqlrMHk4yPhIT7pS7J4y35sPmsN/ns/Ts1R+wpmwMYFzi6yT4ykFt8sran6vOBrD/DV361Xzrte2hp+/U/207bbXs0wZ59SQZXrsl3ZVKj9fskFXYlp3c+yFMfvbichOmmzGrblcmnXwbyrYwqZbn0dMta4Ks7ji8t7PdMHPBw///ZnWzKj61ZdSdXjVnZsun3BWrzZ8b/fXlaONN8dCN2X8mLKpA88YLBfeEF4VuOut87mrfrkm7Za/79XVplHZY/LKla+xnXjNb+tdq0636uzorbmgtt3/9LNmg67B+vOVBH5WT2zNl6nJFCvWDBNkL9B8drFfSez1rWeg37i3WzfM3+20O/Hf2uubk5kwZNQPZdzrfs0Q5vk4oVJiyO9v1f7JV7JcFa+SdTl7g2bvqZehjCZO3dwzObd2s/PjLkinVU0p5vZRk3IVLPx8QYN9VWPDW5e3xqOVh0l/+WdleYp03kT16W2zWtl/+v3rNDksVHvfPsdj9eG3ZlZ0vVct9n8avelLf/Vr86KHvDAwM//8HeLNzfM071TSXiYFhPzcDAyylMzBItaKmdC6UlA5O3VMCeRJABiArC/BmZBJhRmQWZMNBmQUGtjWCSAJZB2EYdvdAgADDf8dXTAzYXcfKBlLCxMDE0MfAwKDNDOIBAgAA//92NyJ6zwMAAA=="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
