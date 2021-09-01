package main

import (
	"fmt"
)

func main() {
	var oHas, aHas, pHas, oUse, aUse, pUse float32
	fmt.Scanln(&oHas, &aHas, &pHas)
	fmt.Scanln(&oUse, &aUse, &pUse)

	var oLeft, aLeft, pLeft float32
	oLeft = oHas / oUse
	aLeft = aHas / aUse
	pLeft = pHas / pUse

	var smollNum float32
	if oLeft <= aLeft && oLeft <= pLeft {
		smollNum = oLeft
	} else if aLeft <= oLeft && aLeft <= pLeft {
		smollNum = aLeft
	} else {
		smollNum = pLeft
	}

	var oUsed, aUsed, pUsed float32
	oUsed = oUse * smollNum
	aUsed = aUse * smollNum
	pUsed = pUse * smollNum

	fmt.Print(oHas-oUsed, aHas-aUsed, pHas-pUsed)
}
