# 获取当前操作系统和架构信息
$OS = [System.Runtime.InteropServices.RuntimeInformation]::IsOSPlatform([System.Runtime.InteropServices.OSPlatform]::Windows)
$ARCH = [System.Environment]::GetEnvironmentVariable("PROCESSOR_ARCHITECTURE")
$CurrentUser = [System.Environment]::ExpandEnvironmentVariables("%USERNAME%")

# 定义文件名和下载链接
switch ($ARCH) {
    "AMD64" { $FILE = "ai_Windows_x86_64.zip" }
    "x86" { $FILE = "ai_Windows_i386.zip" }
    "ARM64" { $FILE = "ai_Windows_arm64.zip" }
    default { Write-Host "Unsupported architecture: $ARCH"; exit 1 }
}

# 定义下载链接
$URL = "https://gh-proxy.com/https://github.com/newde36524/ai/releases/latest/download/$FILE"

# 下载文件
# Write-Host "Downloading $FILE from $URL"
Invoke-WebRequest -Uri $URL -OutFile "$FILE"

# 检查下载是否成功
if ($?) {
    # Write-Host "Download successful: $FILE"
}
else {
    Write-Host "Download failed: $FILE"
    exit 1
}

# 创建目标目录
$TARGET_DIR = "C:\Users\$CurrentUser\AppData\Local\Programs\ai"
if (-Not (Test-Path $TARGET_DIR)) {
    New-Item -Path $TARGET_DIR -ItemType Directory
}

# 解压文件
# Write-Host "Extracting $FILE to $TARGET_DIR"
Expand-Archive -Path "$FILE" -DestinationPath $TARGET_DIR -Force 

# 删除下载的压缩文件
# Write-Host "Deleting the downloaded zip file"
Remove-Item -Path "$FILE" -Force

# 重命名解压出来的文件为ai.exe（如果需要）
# 假设解压后只有一个文件需要重命名
$AI_EXE = Get-ChildItem -Path $TARGET_DIR -Filter *.exe | Select-Object -First 1
if ($AI_EXE) {
    $AI_BIN = "$TARGET_DIR\ai.exe"
    # Write-Host "Renaming executable to ai.exe"
    Rename-Item -Path $AI_EXE.FullName -NewName "ai.exe" -Force
}


# 设置环境变量
$userPath = [Environment]::GetEnvironmentVariable('Path', [System.EnvironmentVariableTarget]::User)
if ($userPath -notlike "*$TARGET_DIR*") {
    $newPath = $userPath + ";$TARGET_DIR"
    [Environment]::SetEnvironmentVariable('Path', $newPath, [System.EnvironmentVariableTarget]::User)
    # Write-Host "Updated PATH environment variable for the current user."
    # 重新加载 PATH 环境变量
    $env:PATH = $newPath
    # Write-Host "Reloaded PATH environment variable in the current session."
}
else {
    # Write-Host "The directory $TARGET_DIR is already in the PATH environment variable."
}


# 重新加载 PATH 环境变量
$env:PATH = $newPath
# Write-Host "Reloaded PATH environment variable in the current session."

# 需要重新打开新的 PowerShell 窗口或命令提示符以使环境变量生效
# Write-Host "Please restart your PowerShell window or command prompt for the PATH changes to take effect, or run this script again to reload the PATH variable in the current session."

Write-Host "ok, usage: ai hi or ai -l hi"
