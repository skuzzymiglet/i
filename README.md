# i

(sorry for the terribly short name - feel free to alias)

Barebones shell command runner. Connects to Neovim via RPC, and reruns your command with your specified input each time your buffer changes. First line is the command, the rest of the lines get piped to stdin

```
go get -v github.com/skuzzymiglet/i
nvim --listen <unix socket address>
```

In another window, connect to Neovim:

```
i <same unix socket address>
```

Try it out. Type this in the Neovim window:

```
wc -l
many
lines
of
text
```

Edit the command and the input lines and see it change, live

# TODOs

- Delay before running command, so we don't risk running partially completed commands that do nasty stuff
- Multi-line commands
