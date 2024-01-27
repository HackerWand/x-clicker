# README

鼠标连点器(x-clicker)
基于Wails（Vue-TS）构建的简单应用程序，可以帮助您自动进行鼠标点击操作。它支持设置点击方式、点击频率以及自定义热键来启动和停止自动化。

## 功能

- 可分别设置左键和右键单击和双击模式。
- 可设置点击间隔时间（毫秒）。
- 可定义自定义热键来启动和停止自动化（仅支持F1-F12）。
- 简单直观的用户界面。

## 环境安装

需要先安装node、golang以及wails，go和wails安装参考：[https://wails.io/zh-Hans/docs/gettingstarted/installation](https://wails.io/zh-Hans/docs/gettingstarted/installation)

## 开发

```shell
git clone https://github.com/HackerWand/x-clicker.git
cd x-clicker
wails dev
```

## 编译可执行文件

```shell
wails build
```

## 兼容性

鼠标连点器与Windows、macOS和Linux操作系统兼容。

## 免责声明

请负责任地使用鼠标连点器，并遵守您使用该应用程序自动化的任何应用程序或网站的服务条款。开发者对于使用该应用程序进行的任何滥用或非法活动概不负责。

## 第三方依赖
1. 跨平台的GUI自动化
    [go-vgo/robotgo](https://github.com/go-vgo/robotgo)
2. 
[wails](github.com/wailsapp/wails/v2)
[hotkey](golang.design/x/hotkey)

## 贡献

欢迎贡献！如果您有任何建议、错误报告或功能请求，请在鼠标连点器GitHub仓库上提出问题。

## 许可证

鼠标连点器是根据[MIT许可证](https://opensource.org/licenses/MIT)的开源软件。