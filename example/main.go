package main

import (
	"fmt"
	"github.com/olongfen/gen-id/generator"
	"github.com/olongfen/gen-id/utils"
	"time"
)

func main() {
	fmt.Println(generator.NewGeneratorData(nil))
	fmt.Println(randDate().String())

}

func randDate() time.Time {
	begin, _ := time.Parse("2006-01-02 15:04:05", "1970-01-01 00:00:00")
	fmt.Println(begin)
	end, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(end)
	return time.Unix(utils.RandInt64(begin.Unix(), end.Unix()), 0).UTC()
}
