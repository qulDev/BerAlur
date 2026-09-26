param([string]$OutputDirectory = './backups')
$ErrorActionPreference = 'Stop'
Set-Location (Split-Path $PSScriptRoot -Parent)
New-Item -ItemType Directory -Force -Path $OutputDirectory | Out-Null
$destination = Join-Path (Resolve-Path -LiteralPath $OutputDirectory) ('beralur-' + (Get-Date -Format 'yyyyMMdd-HHmmss') + '.dump')
docker compose -f compose.yaml -f compose.local.yaml exec -T db sh -c 'pg_dump --username="$POSTGRES_USER" --dbname="$POSTGRES_DB" --format=custom --file=/tmp/beralur-backup.dump'
if ($LASTEXITCODE -ne 0) { throw 'Database backup failed.' }
docker compose -f compose.yaml -f compose.local.yaml cp db:/tmp/beralur-backup.dump $destination
if ($LASTEXITCODE -ne 0) { throw 'Copying the backup failed.' }
Write-Output $destination
