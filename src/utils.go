package tonpaygo

import (
	"encoding/base64"
	"fmt"

	"github.com/xssnick/tonutils-go/tvm/cell"
)

func BuildURLForWallet(reciver string, cell_ *cell.Cell, amount uint64, comission uint64) string {
	bin := base64.StdEncoding.EncodeToString(cell_.ToBOCWithFlags(false))
	total := amount + comission

	return fmt.Sprintf("ton://transfer/%s?bin=%s&amount=%d", reciver, bin, total)
}
