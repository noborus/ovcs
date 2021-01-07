#!/bin/bash

SESSION="psql"
HELPSOCK="/tmp/ov-psql-help.sock"
OVSOCK="/tmp/ov-psql.sock"

SESSIONEXISTS=$(tmux list-sessions | grep $SESSION)
if [ "$SESSIONEXISTS" = "" ]
then
    tmux new-session -d -s $SESSION
    tmux split-window -h -t $SESSION:0
    tmux resize-pane -L 50
    tmux split-window -v -t $SESSION:0
    tmux resize-pane -U 10
    tmux send-keys -t $SESSION:0.0 "ovcs server -w=f -p $HELPSOCK" C-m
    tmux send-keys -t $SESSION:0.2 "ovcs server -H2 -C -d'|' -p $OVSOCK" C-m
    while [ ! -S $HELPSOCK ]
    do
	sleep 1
    done
    tmux send-keys -t $SESSION:0.1 "psql $*" C-m
    tmux send-keys -t $SESSION:0.1 "\setenv PAGER 'ovcs client -p $HELPSOCK'" C-m
    tmux send-keys -t $SESSION:0.1 "\?" C-m
    tmux send-keys -t $SESSION:0.1 "\d" C-m
    tmux send-keys -t $SESSION:0.1 "\setenv PAGER 'ovcs client  -p $OVSOCK'" C-m
    tmux select-pane -t $SESSION:0.1
fi

tmux attach-session -t $SESSION:0
