package function

import (
	"fatsharkbot/src/util"
	"fatsharkbot/src/util/cqhttp"
	"fmt"
	"strconv"
)

func CodeRandomFace() string {
	faceId := util.RandInt32(222)
	return fmt.Sprintf(cqhttp.CodeFace, strconv.Itoa(int(faceId)))
}

func CodePoke(qq int64) string {
	return fmt.Sprintf(cqhttp.CodePoke, strconv.Itoa(int(qq)))
}

func CodeAt(qq int64) string {
	return fmt.Sprintf(cqhttp.CodeAt, strconv.Itoa(int(qq)))
}
