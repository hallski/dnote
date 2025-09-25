# dnote

A terminal based notes application using Markdown files and wikilinks to link
notes. It has a TUI for browsing the notes and a CLI for interacting with them
as commands.

dNote does not support editing directly but work together with an external
editor. When used with another terminal editor the suggestion is to use either a
terminal multiplexer (like tmux) or open a new terminal with the editor.

![](./images/screenshot.png)

## Features

- Notes have three digit IDs to make them easy to remember and type in when
opening by ID.
- Links have uppercase shortcuts to make navigation fast.

## Requirements

dNote makes use of [ugrep](https://github.com/Genivia/ugrep) for searching.

## Setup
Copy ./dnote.yaml.example to ~/.config/dnote.yaml and edit it to work for your setup.

Set the environmental variable `DNOTES` to point at the directory of your notes
(adding a default notes directory to settings is planned, [#3]).

