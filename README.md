# Pinata Go CLI

Welcome to the Pinata Go CLI! This is still in active development so please let us know if you have any questions! :)

## Installation

> [!NOTE]
> If you are on Windows please use WSL when installing. If you get an error that it was not able to resolve the github host run `git config --global --unset http.proxy`

### Install Script

The easiest way to install is to copy and paste this script into your terminal

```bash
curl -fsSL https://cli.pinata.cloud/install-web3 | bash
```

### Building from Source

To build and instal from source make sure you have [Go](https://go.dev/) installed on your computer and the following command returns a version:

```
go version
```

Then paste and run the following into your terminal:

```
git clone https://github.com/PinataCloud/pinata-go-cli && cd pinata-go-cli && go install .
```

### Linux Binary

As versions are released you can visit the [Releases](https://github.com/PinataCloud/pinata-go-cli/releases) page and download the appropriate binary for your system, them move it into your bin folder.

For example, this is how I install the CLI for my Raspberry Pi

```
wget https://github.com/PinataCloud/pinata-go-cli/releases/download/v0.1.2/pinata-go-cli_Linux_arm64.tar.gz

tar -xzf files-cli_Linux_arm64.tar.gz

sudo mv pinata /usr/bin
```

## Usage

The Pinata CLI is equipped with the majortiry of features on the Pinata API.

### `auth` - Authentication

With the CLI installed you will first need to authenticate it with your [Pinata JWT](https://docs.pinata.cloud/docs/api-keys)

```shell
pinata-web3 auth <your-jwt>
```

### `upload` - Uploads

After authentication you can now upload using the `upload` command or `u` for short, then pass in the path to the file or folder you want to upload.

```shell
pinata-web3 upload ~/Pictures/somefolder/image.png
```

The following flags are also available to set the name or CID version of the upload.

```shell
--version value, -v value  Set desired CID version to either 0 or 1. Default is 1. (default: 1)
--name value, -n value     Add a name for the file you are uploading. By default it will use the filename on your system. (default: "nil")
--cid-only                 Use if you only want the CID returned after an upload (default: false)

```

### `list` - List Files

You can list files with the `list` command or the alias `l`. The results are printed in raw JSON to help increase composability.

```shell
pinata-web3 list
```

By default it will retrieve the 10 latest files, but with the flags below you can get more results or fine tune your search.

```shell
--cid value, -c value         Search files by CID (default: "null")
--amount value, -a value      The number of files you would like to return, default 10 max 1000 (default: "10")
--name value, -n value        The name of the file (default: "null")
--status value, -s value      Status of the file. Options are 'pinned', 'unpinned', or 'all'. Default: 'pinned' (default: "pinned")
--pageOffset value, -p value  Allows you to paginate through files. If your file amount is 10, then you could set the pageOffset to '10' to see the next 10 files. (default: "null")
```

### `delete` - Delete Files

If you ever need to you can delete a file by CID using the `delete` command or alias `d` followed by the file CID.

```shell
pinata-web3 delete QmVLwvmGehsrNEvhcCnnsw5RQNseohgEkFNN1848zNzdng
```

### `pin` - Pin by CID

Separate from the `upload` command which uploads files from your machine to Pinata, you can also pin a file already on the IPFS network by using the `pin` command or alias `p` followed by the CID. This will start a pin by CID request which will go into a queue.

```shell
pinata-web3 pin QmVLwvmGehsrNEvhcCnnsw5RQNseohgEkFNN1848zNzdng
```

To check the queue use the `request` command.

### `requests` - Pin by CID Requests

As mentioned in the `pin` command, when you submit an existing CID on IPFS to be pinned to your Pinata account, it goes into a request queue. From here it will go through multiple status'. For more info on these please consult the [documentation](https://docs.pinata.cloud/reference/get_pinning-pinjobs).

```shell
pinata-web3 requests
```

You can use flags to help filter requests as well.

```shell
--cid value, -c value         Search pin by CID requests by CID (default: "null")
--status value, -s value      Search by status for pin by CID requests. See https://docs.pinata.cloud/reference/get_pinning-pinjobs for more info. (default: "null")
--pageOffset value, -p value  Allows you to paginate through requests by number of requests. (default: "null")
```

## Contact

If you have any questions please feel free to reach out to us!

[team@pinata.cloud](mailto:team@pinata.cloud)
