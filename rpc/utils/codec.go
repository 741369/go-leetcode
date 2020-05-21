/**********************************************
** @Des: encode decode utils
** @Author: lzw
** @Last Modified time: 2020/5/21 下午5:25
***********************************************/

package utils

import (
	"bytes"
	"encoding/gob"
)

type Data struct {
	Name string        // service name
	Args []interface{} // request's or response's body except error
	Err  string        // remote server error
}

func encode(data Data) (b []byte, err error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err = encoder.Encode(data); err != nil {
		return nil, err
	}
	b = buf.Bytes()
	return
}

func decode(b []byte) (data Data, err error) {
	buf := bytes.NewBuffer(b)
	decoder := gob.NewDecoder(buf)
	err = decoder.Decode(&data)
	return
}
