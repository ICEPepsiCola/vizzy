#!/bin/bash

# 定义操作系统类型
OS_TYPE=""

# 读取用户输入
echo "请输入要编译的操作系统类型（1- Windows，2- macOS）:"
read choice

# 根据用户选择设置操作系统类型
if [[ $choice -eq 1 ]]; then
  OS_TYPE="windows"
elif [[ $choice -eq 2 ]]; then
  OS_TYPE="darwin"
else
  echo "输入错误！"
  exit 1
fi

echo "目标编译系统:" $OS_TYPE

echo "执行前端打包⬇️"

# 进入 frontend 目录
cd frontend

# 检查操作系统类型
if [[ "$OSTYPE" == "msys" ]]; then
  # 如果是 Windows，使用 start 命令打开新窗口并运行 yarn dev
  start cmd /c "yarn build"
else
  # 如果是 macOS，直接在当前窗口运行 yarn dev
  yarn build
fi

cd ../

if [ ! -d "resources" ]; then
  mkdir resources
fi

# 复制 frontend/build 目录到 resources 目录
rm -rf resources/*
cp -R frontend/dist/* resources/

echo "执行Go打包⬇️"

# 根据操作系统类型编译
if [[ $OS_TYPE == "windows" ]]; then
  env GOOS=windows GOARCH=amd64 go build -ldflags "-X 'main.mode=prod'" -o app.exe main.go
elif [[ $OS_TYPE == "darwin" ]]; then
  env GOOS=darwin GOARCH=amd64 go build -ldflags "-X 'main.mode=prod'" -o app main.go
fi

echo "编译完成！"

if [[ $OS_TYPE == "windows" ]]; then
  echo "请运行app.exe"
elif [[ $OS_TYPE == "darwin" ]]; then
  echo "请运行./app"
fi

