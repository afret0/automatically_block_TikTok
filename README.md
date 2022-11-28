# 抖音一键拉黑

使用 uiautomator2 实现对手机的操作~

手机需要开启开发者模式, 允许调试, 允许安装应用, 允许使用 usb 调试~

项目第一次启动会自动安装所需要套件, uiautomator2

__miui需要启动 uiautomator2__, 在 ATX(uiautomator2自动安装) 刷新服务状态, 或点击启动 uiautomator2
![atx](https://raw.githubusercontent.com/Afret1/image/master/img/907430100956156311.jpg)

# Requirements

uiautomator2

安装过程请参考项目文档, 避免更新不及时就不抄录了....

> https://github.com/openatx/uiautomator2/blob/master/README.md

adb 

`scoop install adb`

或自行参阅 google 相关文档安装...

# 使用

## 安装 uiautomator2

```
# Since uiautomator2 is still under development, you have to add --pre to install the development version
pip install --upgrade --pre uiautomator2

# Or you can install directly from github source
git clone https://github.com/openatx/uiautomator2
pip install -e uiautomator2
```

## 修改 ip, port, 以及 serial(可使用 adb devices 命令获取)为你自己的设备相关信息...

## python main.py

# 操作说明

使用方向键操作

- ↑ 上滑屏幕, 即下一个视频
- ↓ 下滑屏幕
- ← 点击关注
- → 拉黑

拉黑操作会判断是否为已关注, 若已关注则自动滑过视频, 不进行拉黑操作~

# Q&A

## 操作超时...

在判断是否为直播或者广告时, 查找元素时间过长...

