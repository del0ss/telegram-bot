package redis

import (
	"context"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/redis/go-redis/v9"
	"strconv"
	"strings"
	"tgbot/internal/storage/excel"
)

var ctx = context.Background()

type Redis struct {
	Client *redis.Client
	Excel  excel.Excel
}

func New(address, password string, db int, excel excel.Excel) *Redis {
	return &Redis{
		Client: redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password,
			DB:       db,
		}),
		Excel: excel,
	}
}

func (r *Redis) Write(f *excelize.File, err error) error {
	if err != nil {
		return err
	}
	sheets := f.GetSheetMap()
	for _, val := range sheets {
		for i := 3; i <= len(f.GetRows(val)); i++ {
			if f.GetCellValue(val, "D"+strconv.FormatInt(int64(i), 10)) != "" && !strings.Contains(f.GetCellValue(val, "D"+strconv.FormatInt(int64(i), 10)), "-") {
				err := r.Client.Set(ctx, strings.ToLower(f.GetCellValue(val, "D"+strconv.FormatInt(int64(i), 10))), f.GetCellValue(val, "C"+strconv.FormatInt(int64(i), 10))+" - "+
					f.GetCellValue(val, "F"+strconv.FormatInt(int64(i), 10)), 0).Err()
				if err != nil {
					panic(err)
				}
			}

		}

	}

	return nil
}

func (r *Redis) Find(c chan string, userText string) {
	val, err := r.Client.Get(ctx, userText).Result()
	fmt.Println(val)
	if err != nil {
		c <- "Нет такого"
		return
	}
	if val == "" {
		c <- "пусто"
		return
	}
	c <- val
}

func (r *Redis) Delete() {
	r.Client.FlushAll(ctx)
}
