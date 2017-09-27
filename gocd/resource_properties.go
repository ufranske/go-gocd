package gocd

import (
	"encoding/csv"
	"bytes"
	"strings"
	"io"
)

type Properties struct {
	UnmarshallWithHeader bool
	Header               []string
	DataFrame            [][]string
}

func NewPropertiesFrame(frame [][]string) *Properties {
	p := Properties{}
	for i, line := range frame {
		if i == 0 {
			p.Header = line
		} else {
			p.AddRow(line)
		}
	}
	return &p
}

func (p Properties) Get(row int, column string) string {
	var columnIdx int
	for i, key := range p.Header {
		if key == column {
			columnIdx = i
		}
	}
	return p.DataFrame[row][columnIdx]
}

func (p *Properties) AddRow(r []string) {
	p.SetRow(len(p.DataFrame), r)
}

func (p *Properties) SetRow(row int, r []string) {
	for row >= len(p.DataFrame) {
		p.DataFrame = append(p.DataFrame, []string{})
	}
	p.DataFrame[row] = r
}

func (p Properties) MarshallCSV() (string, error) {
	buf := new(bytes.Buffer)
	w := csv.NewWriter(buf)
	if err := w.Write(p.Header); err != nil {
		return buf.String(), err
	}
	for _, line := range p.DataFrame {
		if err := w.Write(line); err != nil {
			return buf.String(), err
		}
	}
	w.Flush()

	return buf.String(), nil
}

func (p *Properties) UnmarshallCSV(raw string) error {
	r := csv.NewReader(strings.NewReader(raw))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if p.UnmarshallWithHeader && len(p.Header) == 0 && len(p.DataFrame) == 0 {
			p.Header = record
		} else {
			p.AddRow(record)
		}
	}
	return nil
}
