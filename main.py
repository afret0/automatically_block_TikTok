from connect import AdbConnector
from douyin import DouYin
import keyboard
from log import logger


def main():
    logger.info(f"抖音自动拉黑脚本启动~\n")

    ip = "192.168.31.164"  # ip
    port = "40421"  # port
    serial = "f7b9b15"  # 序列号

    app_name = "com.ss.android.ugc.aweme"
    adb_connector = AdbConnector(ip=ip, port=port, serial=serial)
    device = adb_connector.connect()
    # device = adb_connector.connect_usb()
    douyin = DouYin(app_name=app_name, device=device)
    douyin.start_douyin()
    douyin.device.jsonrpc.setConfigurator(
        {"waitForIdleTimeout": 10, "waitForSelectorTimeout": 10}
    )
    douyin.set_implicitly_wait(1)
    douyin.set_wait_time(1)
    keyboard.add_hotkey("up", douyin.next_video)
    keyboard.add_hotkey("down", douyin.last_video)
    keyboard.add_hotkey("left", douyin.follow)
    keyboard.add_hotkey("right", douyin.block_and_swipe_next_video)
    keyboard.wait()


if __name__ == "__main__":
    main()
