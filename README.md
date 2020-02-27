# CTL
Manage and run all systemd commands and tools as a subcommand of one command `ctl`!

## Usage
```bash
ctl [command] [command options] [arguments...]
```
example:

`ctl status` instead of `systemctl status` 

`ctl log --system --boot` instead of `journalctl --system --boot` 

`ctl mount --help` instead of `systemd-mount --help`

`ctl run --scope ls` instead of `systemd-run --scope ls`

## Installation

```bash
go get github.com/melbahja/ctl
# To install bash completion:
cd $GOPATH/src/github.com/melbahja/ctl
sudo make completion
```
You should be able to use and execute `ctl` command

## License
[MIT](https://github.com/melbahja/ctl/blob/master/LICENSE)
