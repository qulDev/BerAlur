param(
  [Parameter(Mandatory)][string]$BackupFile,
  [ValidatePattern('^[a-z][a-z0-9_]{0,62}$')][string]$RestoreDatabase = 'beralur_restore_check'
)
$ErrorActionPreference = 'Stop'
$backup = (Resolve-Path -LiteralPath $BackupFile).Path
Set-Location (Split-Path $PSScriptRoot -Parent)
$applicationDatabase = docker compose -f compose.yaml -f compose.local.yaml exec -T db printenv POSTGRES_DB
if ($LASTEXITCODE -ne 0) { throw 'Cannot read database configuration.' }
if ($RestoreDatabase -eq $applicationDatabase.Trim()) { throw 'Restore must target a new database, never the application database.' }
docker compose -f compose.yaml -f compose.local.yaml exec -T db sh -c 'createdb --username="$POSTGRES_USER" "$1"' -- $RestoreDatabase
if ($LASTEXITCODE -ne 0) { throw 'Cannot create restore database. Choose a new database name.' }
docker compose -f compose.yaml -f compose.local.yaml cp $backup db:/tmp/beralur-restore.dump
if ($LASTEXITCODE -ne 0) { throw 'Copying the backup failed. The new database was left in place for inspection.' }
docker compose -f compose.yaml -f compose.local.yaml exec -T db sh -c 'pg_restore --username="$POSTGRES_USER" --dbname="$1" --no-owner --no-privileges --exit-on-error /tmp/beralur-restore.dump' -- $RestoreDatabase
if ($LASTEXITCODE -ne 0) { throw 'Restore failed. Inspect the new database before retrying with another name.' }
Write-Output "Restored into new database: $RestoreDatabase"
