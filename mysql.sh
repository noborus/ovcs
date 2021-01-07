#!/bin/bash

# Set Session Name
SESSION="mysql"
SESSIONEXISTS=$(tmux list-sessions | grep $SESSION)

HELPSOCK="/tmp/ov-mysql-help.sock"
OVSOCK="/tmp/ov-mysql.sock"

# Only create tmux session if it doesn't already exist
if [ "$SESSIONEXISTS" = "" ]
then
    # Start New Session with our name
    tmux new-session -d -s $SESSION
    tmux split-window -h -t $SESSION:0
    tmux resize-pane -L 50
    tmux split-window -v -t $SESSION:0
    tmux resize-pane -U 10
    tmux send-keys -t $SESSION:0.0 "ovcs server -w=f -p $HELPSOCK" C-m
    tmux send-keys -t $SESSION:0.2 "ovcs server -d '|' -H3 -C -p $OVSOCK" C-m
    while [ ! -S $HELPSOCK ]
    do
	sleep 1
    done
    tmux send-keys -t $SESSION:0.1 "mysql $*" C-m
    tmux send-keys -t $SESSION:0.1 "pager ovcs client -p $HELPSOCK" C-m
    tmux send-keys -t $SESSION:0.1 'show tables\;' C-m
    tmux send-keys -t $SESSION:0.1 "pager ovcs client -p $OVSOCK" C-m
    tmux select-pane -t $SESSION:0.1
fi

# Attach Session, on the Main window
tmux attach-session -t $SESSION:0
