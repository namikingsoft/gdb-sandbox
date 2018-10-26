set disassembly-flavor intel

set print pretty on

# dir /usr/local/src/glibc-2.27/libio
# dir /usr/local/src/glibc-2.27/stdio-common

# dashboard assembly -style function True
# dashboard assembly -style opcodes True
alias -a exit = quit

define hook-quit
  set confirm off
end
