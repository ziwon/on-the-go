package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0], pair[1])
	}
}

/*

Results
===============================================================

FOO: 1
BAR:

CLUTTER_IM_MODULE xim
COLORTERM mate-terminal
DBUS_SESSION_BUS_ADDRESS unix:abstract
DEFAULTS_PATH /usr/share/gconf/mate.default.path
DESKTOP_SESSION mate
DISPLAY :0
EDITOR /usr/bin/nvim
GDMSESSION mate
GDM_XSERVER_LOCATION local
GNOME_KEYRING_CONTROL /run/user/1000/keyring-uQe7Bw
GNOME_KEYRING_PID 28974
GPG_AGENT_INFO /run/user/1000/keyring-uQe7Bw/gpg:0:1
GREP_COLOR 37;45
GREP_COLORS mt
GTK_IM_MODULE uim
HOME /home/luno
LANG en_US.UTF-8
LC_ADDRESS ko_KR.UTF-8
LC_IDENTIFICATION ko_KR.UTF-8
LC_MEASUREMENT ko_KR.UTF-8
LC_MONETARY ko_KR.UTF-8
LC_NAME ko_KR.UTF-8
LC_NUMERIC ko_KR.UTF-8
LC_PAPER ko_KR.UTF-8
LC_TELEPHONE ko_KR.UTF-8
LESS -F -g -i -M -R -S -w -X -z-4
LESSOPEN | /usr/bin/env lesspipe %s 2>&-
LESS_TERMCAP_mb [01;31m
LESS_TERMCAP_md [01;31m
LESS_TERMCAP_me [0m
LESS_TERMCAP_se [0m
LESS_TERMCAP_so [00;47;30m
LESS_TERMCAP_ue [0m
LESS_TERMCAP_us [01;32m
LOGNAME luno
LS_COLORS rs
MANDATORY_PATH /usr/share/gconf/mate.mandatory.path
MATE_DESKTOP_SESSION_ID this-is-deprecated
MATHEMATICA_HOME /usr/local/Wolfram/Mathematica/10.3
MDMSESSION mate
MDM_LANG en_US.UTF-8
MDM_XSERVER_LOCATION local
NODE_PATH /usr/lib/nodejs:/usr/lib/node_modules:/usr/share/javascript
OLDPWD /home/luno/Workspace/On-The-Go/src/github.com/patriz/on-the-go
PAGER less
PATH /usr/local/bin:/usr/local/sbin:/usr/local/rbenv/shims:/usr/local/rbenv/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games
PWD /home/luno/Workspace/On-The-Go/src/github.com/patriz/on-the-go/go-by-example
PYTHONPATH /usr/lib/llvm-3.6/lib/python2.7/site-packages
QT4_IM_MODULE uim
QT_IM_MODULE xim
RBENV_ROOT /usr/local/rbenv
RBENV_SHELL Xsession
SESSION_MANAGER local/optimus:@/tmp/.ICE-unix/28984,unix/optimus:/tmp/.ICE-unix/28984
SHELL /bin/zsh
SHLVL 2
SSH_AGENT_PID 29135
SSH_AUTH_SOCK /run/user/1000/keyring-uQe7Bw/ssh
TERM xterm-256color
TEXTDOMAIN im-config
TEXTDOMAINDIR /usr/share/locale/
TMPDIR /tmp/luno
TMUX /tmp/luno/tmux-1000/default,31729,0
TMUX_PANE %1
TMUX_PLUGIN_MANAGER_PATH /home/luno/.tmux/plugins/
TMUX_POWERLINE_SEG_WEATHER_LOCATION 22722547
USER luno
USERNAME luno
VISUAL /usr/bin/vim
WINDOWID 88080421
WINDOWPATH 8
XAUTHORITY /home/luno/.Xauthority
XDG_CONFIG_DIRS /etc/xdg/xdg-mate:/etc/xdg
XDG_CURRENT_DESKTOP MATE
XDG_DATA_DIRS /usr/share/mate:/usr/local/share/:/usr/share/:/usr/share/mdm/
XDG_RUNTIME_DIR /run/user/1000
XDG_SEAT seat0
XDG_SESSION_COOKIE 9846cd7fab942794a2b7cb24560d2b3c-1450984772.335320-1961712165
XDG_SESSION_DESKTOP mate
XDG_SESSION_ID c4
XDG_VTNR 8
XMODIFIERS @im
_ /usr/bin/go
vcs_info_msg_0_ %F{black}  %f%F{236} 8a8998%f %F{black}%F{black}master%f %F{black}%f%f
vcs_info_msg_1_
GOPATH /home/luno/Workspace/On-The-Go
GOBIN /home/luno/Workspace/On-The-Go/bin
FOO 1
*/
