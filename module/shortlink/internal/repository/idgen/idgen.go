package idgen

import (
	"encoding/base64"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"strconv"
	"time"
)

var node *snowflake.Node

func init() {
	st, err := time.Parse("2006-01-02", "2020-01-20")
	if err != nil {
		fmt.Println(err)
		return
	}
	snowflake.Epoch = st.UnixNano() / 1e6
	node, err = snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func NextId() string {
	return base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(node.Generate().Int64(), 10)))
}
