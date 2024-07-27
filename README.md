# unphoto

Download the pic of the day from differing sources. This is inspired by another
application I had that would download and set a new wallpaper each day.
Currently it's only working with The Guardian daily photo's.

## Usage

A few example usage examples

```
unphoto --guardian
unphoto --guardian --wallpaper
unphoto --local --local-path /home/trent/Pictures --wallpaper
```

The latter 2 will change your current wallpaper.

You may want to automate this either through a startup script or on recurring timer.
Using a systemd timer:

```bash
cat << EOF > $HOME/.config/systemd/user/unphoto.service
[Unit]
Description="Changes the wallpaper from the given folder"
Wants=unphoto.timer

[Service]
ExecStart=$HOME/go/bin/unphoto --local --local-path $HOME/Pictures/CorporateBackgrounds --wallpaper

[Install]
WantedBy=default.target
EOF
```

And the timer file to run every hour:

```bash
cat <<EOF > $HOME/.config/systemd/user/unphoto.timer
[Unit]
Description=Timer for the unphoto service that changes the desktop wallpaper
Requires=unphoto.service

[Timer]
Unit=unphoto.service
OnCalendar=*-*-* *:00:00

[Install]
WantedBy=timers.target

EOF
```

After that, reload daemon

```bash
systemctl --user daemon-reload
systemctl --user list-timers
```

After creating the service file, reload the aemon

## Author

Trent Schafer, 2019
