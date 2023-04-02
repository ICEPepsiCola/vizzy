#!/bin/bash

# 进入 frontend 目录
cd frontend

FRONTEND_PID=
# 检查操作系统类型
if [[ "$OSTYPE" == "msys" ]]; then
  # 如果是 Windows，使用 start 命令打开新窗口并运行 yarn dev
  start cmd /c "yarn dev"
  FRONTEND_PID=$!
else
  # 如果是 macOS，直接在当前窗口运行 yarn dev
  yarn dev &
   FRONTEND_PID=$!
fi


echo $FRONTEND_PID

# 返回上级目录
cd ..

# 运行 Go 程序
go run main.go

# 当 Go 程序结束时，结束 yarn dev 进程
kill $FRONTEND_PID
