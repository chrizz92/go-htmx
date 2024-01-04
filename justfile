# use PowerShell instead of sh:
set shell := ["powershell.exe", "-c"]

default:
    start powershell{air -c .air.toml}
    npx tailwindcss -i ./tw-input.css -o ./assets/css/tw-output.css --watch