package main

import (
	"bufio"
	"file_compress/file_compress/huffman"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const version = "1.0.0"

func main() {
	if len(os.Args) > 1 {
		os.Exit(runCLI(os.Args[1:], os.Stdout, os.Stderr))
	}
	runInteractive()
}

func runInteractive() {
	r := bufio.NewReader(os.Stdin)

	fmt.Println("=== Huffman Compression Tool ===")
	fmt.Println()

	fmt.Print("Input file path: ")
	input, _ := r.ReadString('\n')
	input = cleanPath(input)

	if _, err := os.Stat(input); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "error: file not found: %s\n", input)
		pause()
		os.Exit(1)
	}

	fmt.Print("Mode (compress/decompress) [compress]: ")
	mode, _ := r.ReadString('\n')
	mode = strings.TrimSpace(mode)
	if mode == "" {
		mode = "compress"
	}
	mode = strings.ToLower(mode)
	if mode != "compress" && mode != "decompress" {
		fmt.Fprintf(os.Stderr, "error: mode must be 'compress' or 'decompress'\n")
		pause()
		os.Exit(1)
	}

	var output string
	if mode == "compress" {
		ext := filepath.Ext(input)
		output = strings.TrimSuffix(input, ext) + ".huf"
	} else {
		ext := filepath.Ext(input)
		output = strings.TrimSuffix(input, ext) + ".out"
	}

	fmt.Printf("Output file path [%s]: ", output)
	outStr, _ := r.ReadString('\n')
	outStr = cleanPath(outStr)
	if outStr != "" {
		output = outStr
	}

	if _, err := os.Stat(output); err == nil {
		fmt.Print("Output file exists. Overwrite? (y/N): ")
		answer, _ := r.ReadString('\n')
		answer = strings.TrimSpace(strings.ToLower(answer))
		if answer != "y" && answer != "yes" {
			fmt.Println("Cancelled.")
			pause()
			os.Exit(0)
		}
	}

	var err error
	if mode == "compress" {
		fmt.Println("Compressing...")
		err = huffman.Compress(input, output)
	} else {
		fmt.Println("Decompressing...")
		err = huffman.Decompress(input, output)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		pause()
		os.Exit(1)
	}

	fmt.Printf("Done: %s -> %s\n", input, output)
	pause()
}

func pause() {
	fmt.Print("\nPress Enter to exit...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func runCLI(args []string, stdout, stderr io.Writer) int {
	if len(args) < 1 {
		printUsage(stderr)
		return 2
	}

	cmd := args[0]
	switch cmd {
	case "compress", "decompress":
		return runCmd(cmd, args[1:], stdout, stderr)
	case "-h", "--help", "help":
		printUsage(stdout)
		return 0
	case "-v", "--version", "version":
		fmt.Fprintf(stdout, "huffman %s\n", version)
		return 0
	default:
		fmt.Fprintf(stderr, "unknown command: %s\n", cmd)
		printUsage(stderr)
		return 2
	}
}

func parseArgs(args []string) (output, input string, force bool, err error) {
	for i := 0; i < len(args); i++ {
		switch {
		case args[i] == "-o" || args[i] == "--output":
			if i+1 >= len(args) {
				return "", "", false, fmt.Errorf("-%s requires a value", args[i])
			}
			i++
			output = args[i]
		case args[i] == "-f" || args[i] == "--force":
			force = true
		default:
			if input != "" {
				return "", "", false, fmt.Errorf("unexpected argument: %s", args[i])
			}
			input = args[i]
		}
	}
	return output, input, force, nil
}

func runCmd(cmd string, args []string, stdout, stderr io.Writer) int {
	outputArg, input, force, err := parseArgs(args)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 2
	}
	if input == "" {
		fmt.Fprintf(stderr, "usage: huffman %s [-o output] [-f] <input>\n", cmd)
		return 2
	}
	if _, err := os.Stat(input); os.IsNotExist(err) {
		fmt.Fprintf(stderr, "error: input file not found: %s\n", input)
		return 1
	}

	outputPath := outputArg
	if outputPath == "" {
		ext := filepath.Ext(input)
		if cmd == "compress" {
			outputPath = strings.TrimSuffix(input, ext) + ".huf"
		} else {
			outputPath = strings.TrimSuffix(input, ext) + ".out"
		}
	}

	if !force {
		if _, err := os.Stat(outputPath); err == nil {
			fmt.Fprintf(stderr, "error: output file exists: %s (use -f to overwrite)\n", outputPath)
			return 1
		}
	}

	var err2 error
	if cmd == "compress" {
		err2 = huffman.Compress(input, outputPath)
	} else {
		err2 = huffman.Decompress(input, outputPath)
	}
	if err2 != nil {
		fmt.Fprintf(stderr, "error: %v\n", err2)
		return 1
	}

	fmt.Fprintf(stdout, "%s: %s -> %s OK\n", cmd, input, outputPath)
	return 0
}

func cleanPath(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, `"`)
	return s
}

func printUsage(w io.Writer) {
	fmt.Fprintf(w, `huffman - Huffman compression tool %s

Usage (CLI):
  huffman compress [-o output] [-f] <input>
  huffman decompress [-o output] [-f] <input>
  huffman -h | --help
  huffman -v | --version

Usage (Interactive - double-click):
  Just run the exe without arguments.

Commands:
  compress    compress file using Huffman coding
  decompress  decompress Huffman-compressed file

Flags:
  -o <path>   output file path (default: auto-generated)
  -f          overwrite output file if it exists
`, version)
}
