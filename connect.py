import uiautomator2 as u2
import log


class AdbConnector:
    device: u2.Device
    logger = log.logger

    def __init__(self, ip="", port="", serial="", new_command_timeout=300):
        """
        Args:
            ip:
            port:
            serial:
            new_command_timeout:  配置 accessibility 服务的最大空闲时间，超时将自动释放。默认 3 分钟。
        """
        self.ip = ip
        self.port = port
        self.serial = serial
        self.new_command_timeout = new_command_timeout  #

    def _set_new_command_timeout(self):
        self.device.set_new_command_timeout(self.new_command_timeout)

    def connect_usb(self) -> u2.Device:
        d = u2.connect_usb(self.serial)
        self.logger.info(d.info)
        self.device = d
        return d

    def connect_usb_wifi(self) -> u2.Device:
        d = u2.connect_adb_wifi(f"{self.ip}:{self.port}")
        self.device = d
        return d

    def connect(self) -> u2.Device:
        d = u2.connect(self.ip)
        self.device = d
        return d
