#!/usr/bin/bash
#
# Bash completion for ctl command
#
# Author: Mohammed El Bahja
# License: MIT
# See: https://github.com/melbahja/ctl
#

_ctl_completions()
{
	local SUGS=$(ctl complete $COMP_LINE 2>/dev/null)

	if [[ $(echo $SUGS | tr -d '\r\n') = $(echo $COMP_LINE | rev | cut -d' ' -f1 | rev | tr -d '\r\n') ]]; then
		return
	fi


	COMPREPLY=($(compgen -W "$(echo $SUGS | tr '\r\n' ' ')" -- "${COMP_WORDS[$COMP_CWORD]}"))
}

complete -F _ctl_completions ctl
