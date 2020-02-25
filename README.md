# CTL
Manage and run all systemd commands and tools in one command `ctl`!

`ctl` is running all systemd command and tools as a subcommand, for example: 

`ctl status` instead of `systemctl status` 

`ctl log --system --boot` instead of `journalctl --system --boot` 

`ctl bus status` instead of `busctl status`

`ctl run --scope ls` instead of `systemd-run --scope ls`

## Install
You can check releases page for compiled versions or follow these steps:
```bash
gi clone https://github.com/melbahja/ctl /tmp/ctl
cd /tmp/ctl
make
sudo make install
sudo make completion
```
You should be able to use and execute `ctl` command
