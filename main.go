package main

import (
	"errors"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

type UploadResponse struct {
	IpfsHash    string `json:"IpfsHash"`
	PinSize     int    `json:"PinSize"`
	Timestamp   string `json:"Timestamp"`
	IsDuplicate bool   `json:"isDuplicate"`
}

type PinByCIDResponse struct {
	Id     string `json:"id"`
	CID    string `json:"ipfsHash"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

type Options struct {
	CidVersion int `json:"cidVersion"`
}

type Metadata struct {
	Name      string                 `json:"name"`
	KeyValues map[string]interface{} `json:"keyvalues"`
}

type Pin struct {
	Id            string   `json:"id"`
	IPFSPinHash   string   `json:"ipfs_pin_hash"`
	Size          int      `json:"size"`
	UserId        string   `json:"user_id"`
	DatePinned    string   `json:"date_pinned"`
	DateUnpinned  *string  `json:"date_unpinned"`
	Metadata      Metadata `json:"metadata"`
	MimeType      string   `json:"mime_type"`
	NumberOfFiles int      `json:"number_of_files"`
}

type ListResponse struct {
	Rows []Pin `json:"rows"`
}

type Request struct {
	Id        string `json:"id"`
	CID       string `json:"ipfs_pin_hash"`
	StartDate string `json:"date_queued"`
	Name      string `json:"name"`
	Status    string `json:"status"`
}

type RequestsResponse struct {
	Rows []Request `json:"rows"`
}

func main() {
  exeName := "pinata"
	app := &cli.App{
		Name:  "pinata",
		Usage: "A CLI for uploading files to Pinata! To get started make an API key at https://app.pinata.cloud/keys, then authorize the CLI with the auth command with your JWT",
		Commands: []*cli.Command{
			{
				Name:      "auth",
				Aliases:   []string{"a"},
				Usage:     "Authorize the CLI with your Pinata JWT",
				ArgsUsage: "[your Pinata JWT]",
				Action: func(ctx *cli.Context) error {
					jwt := ctx.Args().First()
					if jwt == "" {
						return errors.New("no jwt supplied")
					}
					err := SaveJWT(jwt)
					return err
				},
			},
			{
				Name:      "upload",
				Aliases:   []string{"u"},
				Usage:     "Upload a file or folder to Pinata",
				ArgsUsage: "[path to file]",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "version",
						Aliases: []string{"v"},
						Value:   1,
						Usage:   "Set desired CID version to either 0 or 1. Default is 1.",
					},
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Value:   "nil",
						Usage:   "Add a name for the file you are uploading. By default it will use the filename on your system.",
					},
          &cli.BoolFlag{
            Name: "cid-only",
            Usage: "Use if you only want the CID returned after an upload",
          },
				},
				Action: func(ctx *cli.Context) error {
					filePath := ctx.Args().First()
					version := ctx.Int("version")
					name := ctx.String("name")
          cidOnly := ctx.Bool("cid-only")
					if filePath == "" {
						return errors.New("no file path provided")
					}
					_, err := Upload(filePath, version, name, cidOnly)
					return err
				},
			},
			{
				Name:      "pin",
				Aliases:   []string{"p"},
				Usage:     "Pin an existing CID on IPFS to Pinata",
				ArgsUsage: "[CID of file on IPFS]",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Value:   "null",
						Usage:   "Add a name for the file you are trying to pin.",
					},
				},
				Action: func(ctx *cli.Context) error {
					cid := ctx.Args().First()
					name := ctx.String("name")
					if cid == "" {
						return errors.New("no cid provided")
					}
					_, err := PinByCID(cid, name)
					return err
				},
			},
			{
				Name:    "requests",
				Aliases: []string{"r"},
				Usage:   "List pin by CID requests.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "cid",
						Aliases: []string{"c"},
						Value:   "null",
						Usage:   "Search pin by CID requests by CID",
					},
					&cli.StringFlag{
						Name:    "status",
						Aliases: []string{"s"},
						Value:   "null",
						Usage:   "Search by status for pin by CID requests. See https://docs.pinata.cloud/reference/get_pinning-pinjobs for more info.",
					},
					&cli.StringFlag{
						Name:    "pageOffset",
						Aliases: []string{"p"},
						Value:   "null",
						Usage:   "Allows you to paginate through requests by number of requests.",
					},
				},
				Action: func(ctx *cli.Context) error {
					cid := ctx.String("cid")
					status := ctx.String("status")
					offset := ctx.String("pageOffset")
					_, err := Requests(cid, status, offset)
					return err
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"d"},
				Usage:     "Delete a file by CID",
				ArgsUsage: "[CID of file]",
				Action: func(ctx *cli.Context) error {
					cid := ctx.Args().First()
					if cid == "" {
						return errors.New("no CID provided")
					}
					err := Delete(cid)
					return err
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List most recent files",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "cid",
						Aliases: []string{"c"},
						Value:   "null",
						Usage:   "Search files by CID",
					},
					&cli.StringFlag{
						Name:    "amount",
						Aliases: []string{"a"},
						Value:   "10",
						Usage:   "The number of files you would like to return, default 10 max 1000",
					},
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Value:   "null",
						Usage:   "The name of the file",
					},
					&cli.StringFlag{
						Name:    "status",
						Aliases: []string{"s"},
						Value:   "pinned",
						Usage:   "Status of the file. Options are 'pinned', 'unpinned', or 'all'. Default: 'pinned'",
					},
					&cli.StringFlag{
						Name:    "pageOffset",
						Aliases: []string{"p"},
						Value:   "null",
						Usage:   "Allows you to paginate through files. If your file amount is 10, then you could set the pageOffset to '10' to see the next 10 files.",
					},
				},
				Action: func(ctx *cli.Context) error {
					cid := ctx.String("cid")
					amount := ctx.String("amount")
					name := ctx.String("name")
					status := ctx.String("status")
					offset := ctx.String("pageOffset")
					_, err := ListFiles(amount, cid, name, status, offset)
					return err
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
