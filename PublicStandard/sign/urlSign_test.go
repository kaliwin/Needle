package sign

import (
	"fmt"
	"net/url"
	"testing"
)

func TestUrlSign(t *testing.T) {

	parse, err := url.Parse("https://132.100.175.191:8443/ba/common/assets/js/BigInt.js")
	if err != nil {
		panic(err)
	}

	fmt.Println(UrlSign(parse))

}
