# Requires the API image built by the local Compose stack. Uses disposable containers only.
param([string]$APIImage = 'beralur-local-api')
$ErrorActionPreference = 'Stop'
$testName = 'beralur-proxy-check-' + [guid]::NewGuid().ToString('N').Substring(0, 8)
$oldAPI = "$testName-old"
$newAPI = "$testName-new"
$proxy = "$testName-nginx"
$config = (Resolve-Path (Join-Path $PSScriptRoot '../../infrastructure/nginx/default.conf')).Path
$containers = @()

function Assert-Health {
  for ($attempt = 0; $attempt -lt 20; $attempt++) {
    $body = docker exec $proxy wget -qO- -T 2 http://127.0.0.1/health 2>$null
    if ($LASTEXITCODE -eq 0 -and $body -eq '{"status":"ok"}') { return }
    Start-Sleep -Seconds 1
  }
  throw 'Proxy did not recover after the API address changed.'
}

docker network create $testName | Out-Null
if ($LASTEXITCODE -ne 0) { throw 'Cannot create test network.' }
try {
  docker run -d --name $oldAPI --network $testName --network-alias api --network-alias web -e DATABASE_URL=postgres://test:test@unused/test $APIImage | Out-Null
  if ($LASTEXITCODE -ne 0) { throw 'Cannot start first API.' }
  $containers += $oldAPI
  docker run -d --name $proxy --network $testName --mount "type=bind,src=$config,dst=/etc/nginx/conf.d/default.conf,readonly" nginx:1.29-alpine | Out-Null
  if ($LASTEXITCODE -ne 0) { throw 'Cannot start test proxy.' }
  $containers += $proxy
  Assert-Health
  Write-Output 'PASS: original API is reachable through Nginx.'
  # Start replacement before removing the old API to guarantee a different IP.
  docker run -d --name $newAPI --network $testName --network-alias api --network-alias web -e DATABASE_URL=postgres://test:test@unused/test $APIImage | Out-Null
  if ($LASTEXITCODE -ne 0) { throw 'Cannot start replacement API.' }
  $containers += $newAPI
  $oldAddress = docker inspect --format '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $oldAPI
  $newAddress = docker inspect --format '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $newAPI
  if ($oldAddress -eq $newAddress) { throw 'Test needs different container addresses.' }
  docker rm -f $oldAPI | Out-Null
  if ($LASTEXITCODE -ne 0) { throw 'Cannot remove first API.' }
  $containers = @($containers | Where-Object { $_ -ne $oldAPI })
  Assert-Health
  Write-Output "PASS: Nginx recovered after API changed from $oldAddress to $newAddress."
} finally {
  if ($containers.Count -gt 0) { docker rm -f @containers | Out-Null }
  docker network rm $testName | Out-Null
}
