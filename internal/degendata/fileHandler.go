package degendata

import (
	"bytes"
	"encoding/gob"
	"io"
	"os"
	"sort"

	"github.com/charmbracelet/log"
	"github.com/klauspost/compress/zstd"
)

//
// helper & utility functions
//

func SortMapByValue(m map[string]int64, reverse bool) []string {
	sorted := make([]string, 0)

	for k := range m {
		sorted = append(sorted, k)
	}

	sort.Slice(sorted, func(i, j int) bool {
		if reverse {
			return m[sorted[i]] < m[sorted[j]]
		}

		return m[sorted[i]] > m[sorted[j]]
	})

	return sorted
}

func WriteDataToFile(data interface{}, filePath string) {
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(data)
	if err != nil {
		log.Errorf("failed to encode metadata: %s", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		log.Errorf("failed to create file: %s", err)
	}
	defer file.Close()

	err = ZstdCompress(&buf, file)
	if err != nil {
		log.Errorf("failed to compress file: %s", err)
	}
}

func ReadMetadataFromFile(filePath string) ([]Metadata, error) {
	metadata := make([]Metadata, 0)

	file, err := os.Open(filePath)
	if err != nil {
		log.Errorf("failed to open file: %s", err)

		return metadata, err
	}
	defer file.Close()

	log.Debugf("file: %+v", file)

	decoder, err := zstd.NewReader(file, zstd.WithDecoderConcurrency(0))
	if err != nil {
		log.Errorf("failed to create zstd decoder: %s", err)
	}
	defer decoder.Close()

	err = gob.NewDecoder(decoder.IOReadCloser()).Decode(&metadata)
	if err != nil {
		return metadata, err
	}

	return metadata, nil
}

func ZstdCompress(in io.Reader, out io.Writer) error {
	enc, err := zstd.NewWriter(out, zstd.WithEncoderLevel(zstd.SpeedFastest))
	if err != nil {
		return err
	}

	_, err = io.Copy(enc, in)
	if err != nil {
		enc.Close()

		return err
	}

	return enc.Close()
}
