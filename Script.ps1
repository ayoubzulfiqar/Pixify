Write-Host "Response ImageShack"
Invoke-RestMethod -Uri "https://api.imageshack.com/v2/images/p1KyWfgSj"

Write-Host "`n`nResponse IMGUR"
Invoke-RestMethod -Uri "https://api.imgur.com/3/gallery/image/duhCHQu"

$headers = @{
    Authorization = 'Basic ' + [Convert]::ToBase64String([Text.Encoding]::ASCII.GetBytes('api:47yoyAZ-RSj34OwwRCJ5EImb8nUDLHx8'))
}

$file = Get-Item -Path "..\CFzq6zN.jpg"
$url = "https://api.tinify.com/shrink"

$response = Invoke-RestMethod -Uri $url -Headers $headers -Method Post -InFile $file.FullName

Write-Output $response
