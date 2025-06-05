package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("Program started")

	fmt.Print("Enter a string: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	log.Info().Str("input", input).Msg("User input received")

	hash := sha256.Sum256([]byte(input))
	hashStr := hex.EncodeToString(hash[:])
	log.Info().Str("sha256", hashStr).Msg("Hash computed")
	fmt.Println(hashStr)
}
