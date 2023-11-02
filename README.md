# Lua Capture 

Test your lua capture patterns against strings online.

[try online](https://lua-capture.onrender.com/)

## Motivation

While adding a parser to [nvim-lint](https://github.com/mfussenegger/nvim-lint) i realized that there were no websites to test lua capture patterns easily and here we are.

## How it works

Go backend runs an embedded lua function to return captured groups.

## Tech Stack

As minimal as it can be:

- Frontend with [htmx](https://htmx.org/)
- Backend with Go and templating with [templ](https://github.com/a-h/templ)
- Styling with [picocss](https://picocss.com/)
