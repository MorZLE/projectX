$envFile = "msrvs/msrv-bot-tg/config/_secret.env"

if (Test-Path $envFile) {
    $envContent = Get-Content $envFile
    foreach ($line in $envContent) {
        if ($line -match '^(.*?)=(.*)$') {
            $name = $matches[1]
            $value = $matches[2]
            Set-Item -Path "Env:$name" -Value $value
        }
    }
    Write-Host "Переменные окружения из файла .env успешно применены."
} else {
    Write-Host "Файл .env не найден."
}