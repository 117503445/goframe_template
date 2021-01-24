package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GDgtLUIYkACIgycDMXlienpqUX6UFovqzg/LzSElYExSYU34U1EfLa0u0D9+eUR62oLHF4ssVi4YLkUV+bkC21LlvGvzbxW/vhy2+7bZ9/FtUh/ESi4bbKlUMFPkc83eZZI8Iq7F5petz3QX7w7eJ/rzqt3Q3fdDnW5Yx/wRD3z2/b9c+f+PNsTbu6n0DJBdY4ei8Iavl8RklM2aqtuvnJ9qhvf68TzzHmlE0pX3dxUpM15fcVVR/N/t142FVgkTrpgcDju4wup44w1r5ifPRS9kre/dNnWil9fJKfk5S+oETxq/9/CoiZHfbf1yXJ3/XK+y+3VHMaT5i+qUJ1ol9Ij3JLx8lH1Y432NiHrokOqWwUeObzVOy9/eymD4PfNX988kr397mvVWaFbU6aEPhRIWGR2Qujbtuk3Z/HxrhG6Lf/5rDpfC9OZfXNr6yXn7q1g1DWLrmp7fcBV7nrZ5tgFRS1HN7zpOply4+thmxXu6VvmBpsXaTnNcEg4VvPogc7Wb6GOzqvZojOvVKp+W1ZQdvD7yxwXJS6mM9c33+bv3/aovf/C7qzNbKccDaS4BG6t1jqxMuz0jg4Tp38P9jxbILSmi8NfNWSHueWa1Jz8G3prJk5IeXNI707uPGGTPVPqV7Iv/fE6O/zVNU/veG1NaeMkdyfRb2u8BB+4zDlRmhy23OXrHXf5tNthaXua/Nj5eF1D3kYId1wXflUfzKD1zfhjz0v7u7d882R53sscmX2fb/8fTbuKaXn5PH1lk2b//lxkEKn9NCDzRQ9X853nVae/TvuzS7r+41a/BJ505Ugmc3mb3VumsX44fWC9MLuesH5M2vQHt0w2ij4uLpXq0HgV+mqdv3PhDjFBvnVXHSe5f+jXtFmg98P3rWu/bnGs1nPeBXmVR2221L7SvBiy7NEe/5CJW+0m2lXI7d/xf/2XTlU3vZTKX/bbOz5ay+1me3b3YIppGfvXlRUpvZZGP69cer76mt0Tq2CvJavDD068Gq76eArj9CMTuLoTFaum7H7mIvrHqUio6uuKL6Y/z3vZrp9wPdX06LqF3AG5z84uDil6cfSvOfuJhe+mZId4/Xx6rfnhv7pfP54dclo9+c4u6fv/9jXemvvj5OaH+f7SLyavuBrVUra8Za5NpsbfS5lCXh9KLruKl4l7Fy9f8cvqV6zCzv/en87eO/Tg+9sf3/yKnmfu+em8Z7utxR+tf38SPjy6l2FZW6dfueWjV/xnhvNPshT2zLT1NXu/zP+fgd3NCxu/LLETlV1dsXZ/8quYxeo7HRNmTPwcINvx1Ca2K2V60aeZv64v5pI6e1kvodzm+s5i9V/1bAwM//8HeLNzHHyfunEGMwNDvQgDAyyLM2BkcXZEFofnapBuZDUB3oxMIsyIIgLZZFARAQPbGkEk3gIDYRR2p0CAAMN/x8fMDFgcxsoGkmdiYGLoZGBgUGQB8QABAAD//4nOtNHABAAA"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
