package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GBwFDMPYkACIgycDMXlienpqUX6UFovqzg/LzSElYHxVh93Qs+kg36HHATa3tsc3ZHOpOvks6Py8mxz/rYTy3hSJ5eVb7Ze0fLlmXLFpIwnSXtezBXWV16Qn7hrrVOu6rcrP5/Pri8rS49iZkgwk0jsmVb68BCnn4xg5rNDcbbH11VKLEqMTVjbdrwo9MTCpNgZedrsGczKJpyqwdvkVy18aBjndqL3TfN7+dI9h8t2M8/+HX1i8oaq2rp4G9Pr+p+/Hv5UvO/GmXsN2RPLC/epaBy6sjGB43j68t0hC1Q+NlxSVrX+JCo1W/5EfgWvvbSu+oxnJpuOas5qveCpGOR55lSvVE9zvHmEAkf+ipXKeirajHurbk6s2HH8Lctij7/nT/Q9njHb0+r78U95ptUta3YwmWZsnSSW+3QCX2blvfmeKyqqr3Ofm9rJxi6qW3/qDtunvdf6FneWSSec5JCclv/gt7eWcerWFTmL3+ivVI1nZGD4/z/Am53D/zXrb2dGBgYlZgYGWGwwYMQGOyI24BEA0o2sJsCbkUmEGRGbyCaDYhMGtjWCSLxxizAKu1MgQIDhv2MfIwMWh7GygeSZGJgYOhkYGM4wgniAAAAA///Yr4RuawIAAA=="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
