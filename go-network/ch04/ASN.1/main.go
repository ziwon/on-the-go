package main

/**

ASN.1, http://www.obj-sys.com/asn1tutorial/node1.html
Universal Tags

	0	reserved for BER
	1	BOOLEAN
	2	INTEGER
	3	BIT STRING
	4	OCTET STRING
	5	NULL
	6	OBJECT IDENTIFIER
	7	ObjectDescriptor
	8	INSTANCE OF, EXTERNAL
	9	REAL
	10	ENUMERATED
	11	EMBEDDED PDV
	12	UTF8String
	13	RELATIVE-OID
	16	SEQUENCE, SEQUENCE OF
	17	SET, SET OF
	18	NumericString
	19	PrintableString
	20	TeletexString, T61String
	21	VideotexString
	22	IA5String
	23	UTCTime
	24	GeneralizedTime
	25	GraphicString
	26	VisibleString, ISO646String
	27	GeneralString
	28	UniversalString
	29	CHARACTER STRING
	30	BMPString
*/

import (
	"encoding/asn1"
	"fmt"
	"os"
)

type T1 struct {
	int1 int
	str1 string
}

func main() {
	mdata, err := asn1.Marshal(12)
	checkError(err)
	fmt.Println(mdata) // [2 1 12]

	var n int
	_, err1 := asn1.Unmarshal(mdata, &n)
	checkError(err1)

	fmt.Println("After marshal/unmarshal: ", n)

	t1 := T1{16, "hello"}
	mdata, err = asn1.Marshal(t1) // [48 10 2 1 16 19 5 104 101 108 108 111]
	fmt.Println(mdata)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
