#!/bin/bash

# 检查是否提供了进程名称参数
if [ -z "$1" ]; then
  echo "Usage: $0 <process_name>"
  exit 1
fi

PROCESS_NAME=$1

# 查找进程ID
PIDS=$(pgrep "$PROCESS_NAME")

# 检查是否找到了任何进程
if [ -z "$PIDS" ]; then
  echo "No processes found with name: $PROCESS_NAME"
fi

# 杀死找到的所有进程
echo "Killing the following processes:"
echo "$PIDS"
kill $PIDS
# 设置要编译的 Go 文件路径，如果没有提供则使用当前目录中的主文件
GO_FILE=./main.go



# 检查 Go 文件是否存在
if [ ! -f "$GO_FILE" ]; then
  echo "Go file $GO_FILE does not exist."
  exit 1
fi

# 执行 go build 命令
echo "Building $GO_FILE..."
go build -o "$PROCESS_NAME" "$GO_FILE"

# 检查 go build 命令是否成功
if [ $? -eq 0 ]; then
  echo "Successfully built $GO_FILE. Output file: $PROCESS_NAME"
else
  echo "Failed to build $GO_FILE."
  exit 1
fi
LOG_DIR="./log_data"
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
NOHUP_FILE="nohup.out"
BACKUP_FILE="$LOG_DIR/nohup_$TIMESTAMP.out"
mv "$NOHUP_FILE" "$BACKUP_FILE"
# 使用 nohup 启动编译后的可执行文件，并将输出日志记录到 nohup.out
echo "Starting $PROCESS_NAME with nohup..."
nohup ./"$PROCESS_NAME" > nohup.out 2>&1 &

# 检查 nohup 命令是否成功
if [ $? -eq 0 ]; then
  echo "$PROCESS_NAME started successfully. Logs are being written to nohup.out."
else
  echo "Failed to start $PROCESS_NAME."
  exit 1
fi

