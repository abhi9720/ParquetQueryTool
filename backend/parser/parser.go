package parser

import (
	"log"

	"apica-search/models"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
)

func LoadParquetFile(filePath string) ([]models.Record, error) {
	fr, err := local.NewLocalFileReader(filePath)
	if err != nil {
		return nil, err
	}
	defer fr.Close()

	pr, err := reader.NewParquetReader(fr, new(models.Record), 4)
	if err != nil {
		return nil, err
	}
	defer pr.ReadStop()

	num := int(pr.GetNumRows())
	records := make([]models.Record, 0, num)

	for {
		rows, err := pr.ReadByNumber(100)
		if err != nil {
			log.Println("Error reading rows:", err)
			break
		}
		if len(rows) == 0 {
			break
		}
		for _, r := range rows {
			record, ok := r.(models.Record)
			if !ok {
				log.Println("Error: row is not a Record:", r)
				continue
			}
			records = append(records, record)
		}
	}

	return records, nil
}
