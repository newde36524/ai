#!/bin/bash

# 获取当前操作系统信息
OS=$(uname -s)
USER=$(whoami)
# 定义文件名和下载链接
ARCH=$(uname -m)
case "$OS" in
Linux)
    BIN_DIR="/usr/local/bin/ai"
    case "$ARCH" in
    x86_64)
        FILE="ai_Linux_x86_64.tar.gz"
        ;;
    i686)
        FILE="ai_Linux_i386.tar.gz"
        ;;
    arm64)
        FILE="ai_Linux_arm64.tar.gz"
        ;;
    aarch64)
        FILE="ai_Linux_arm64.tar.gz"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
    esac
    ;;
Darwin)
    BIN_DIR="/usr/bin/ai"
    case "$ARCH" in
    arm64)
        FILE="ai_Darwin_arm64.tar.gz"
        ;;
    x86_64)
        FILE="ai_Darwin_x86_64.tar.gz"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
    esac
    ;;
*)
    echo "Unsupported operating system: $OS"
    exit 1
    ;;
esac

URL="https://gh-proxy.com/https://github.com/newde36524/ai/releases/latest/download/$FILE"

# 下载文件
# echo "Downloading $FILE from $URL"
curl -L -o "$FILE" "$URL"

# 检查下载是否成功
if [ $? != 0 ]; then
    echo "Download failed: $FILE"
    exit 1
fi

# 创建目标目录
mkdir -p "$BIN_DIR"

# 解压文件
# echo "Extracting $FILE to $BIN_DIR"
tar -xzf "$FILE" -C "$BIN_DIR"

# 删除下载的压缩文件
# echo "Deleting the downloaded tar.gz file"
rm "$FILE"

# 设置可执行权限
AI_BIN="${BIN_DIR}/${FILE%.tar.gz}"
if [ -f "$AI_BIN" ]; then
    # echo "Setting executable permission for $AI_BIN"
    chmod +x "$AI_BIN"
fi

# 设置环境变量
if ! grep -q "$BIN_DIR" <<<"$PATH"; then
    export PATH="$PATH:$BIN_DIR"
    echo 'export PATH="$PATH:'"$BIN_DIR"'"' >>~/.bashrc
    # echo "Updated PATH environment variable."
fi
source ~/.bashrc
# echo "Please restart your terminal or run 'source ~/.bashrc' for the PATH changes to take effect."
echo "ok, usage: ai hi or ai -l hi"
