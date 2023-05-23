/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package oncecmd

import (
	"bytes"
	"encoding/gob"
	"io"
	"os"
	"sort"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/provider"
	"github.com/charmbracelet/log"
	"github.com/klauspost/compress/zstd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var lg = internal.BaseLogger

func init() {}

// OnceCmd represents the once command
var OnceCmd = &cobra.Command{
	Use:   "once",
	Short: "This command is intended to quickly implement a one-time task or similar",
	Long:  `A One-time tasks or experiment can easily implemented based the foundation glooomberg provides. Just create a new .go file in the once/ directory for your task or add it to a already available one and call it from this command.`,
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	//
	// initialize some basic resources
	//

	// pool of providers
	pool, err := provider.FromConfig(viper.Get("provider"))
	if err != nil {
		gbl.Log.Fatal("❌ running provider failed, exiting")
	}

	// ethclient.Client
	client := pool.GetProviders()[0].Client

	//
	// if you need other resources besides the ethclient.Client, feel free to initiate them here
	//

	//
	//
	// call your task/experiment/whatever from here on
	//
	//

	// lawless metadata: get the lawless on-chain metadata and save it to a json file
	analyzeLawlessTokenNames(client)
}

//
// helper & utility functions
//

func sortMapByValue(m map[string]uint64, reverse bool) []string {
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

func writeDataToFile(data interface{}, filePath string) {
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

	zstdCompress(&buf, file)
}

//
// zstd compression

func readDataFromFile(filePath string, data interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		log.Errorf("failed to open file: %s", err)

		return err
	}
	defer file.Close()

	var buf bytes.Buffer
	err = gob.NewDecoder(file).Decode(&buf)
	if err != nil {
		return err
	}

	zstdDecompress(bytes.NewReader(buf.Bytes()), data)

	return nil
}

func zstdCompress(in io.Reader, out io.Writer) error {
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

func zstdDecompress(in io.Reader, out interface{}) error {
	dec, err := zstd.NewReader(in, zstd.WithDecoderConcurrency(0))
	if err != nil {
		return err
	}

	err = gob.NewDecoder(dec).Decode(out)
	if err != nil {
		return err
	}

	return nil
}
