package OracleTranscoder

// 中文作为oracle的查询条件、更新参数时，需要转为该数据库可识别的16进制
// 对应的列名需要用rawtohex函数包裹
import (
	"encoding/hex"
	"fmt"

	"github.com/axgle/mahonia"
)

// Oracle select时列名与函数rawtohex的组合
// select rawtohex(?) as ?
func OracleRawtohex(column string) string {
	return fmt.Sprintf("rawtohex(%v) as %v", column, column)
}

// oracle where条件与hextoraw的组合
// where column=utl_raw.cast_to_varchar2(hextoraw(?))
func OracleWhereHex() string {
	return "utl_raw.cast_to_varchar2(hextoraw(?))"
}

// 中文->16进制
func StrToOracelHex(str string) string {
	encoder := mahonia.NewEncoder("GBK")
	strGbk := encoder.ConvertString(str)
	return hex.EncodeToString([]byte(strGbk))
}

// 16进制->中文
func OracelHexToStr(hexStr string) (str string, err error) {
	decoder := mahonia.NewDecoder("GBK")
	bt, err := hex.DecodeString(hexStr)
	if err != nil {
		return
	}
	str = decoder.ConvertString(string(bt))
	return
}
