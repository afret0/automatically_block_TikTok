import uiautomator2 as u2
import log
import time
from functools import wraps


def timethis(func):
    @wraps(func)
    def wrapper(*args, **kwargs):
        start = time.perf_counter()
        r = func(*args, **kwargs)
        end = time.perf_counter()
        print("{}.{} : {}".format(func.__module__, func.__name__, end - start))
        return r

    return wrapper


class DouYin:
    device: u2.Device
    app_name: str
    logger = log.logger
    follow_button_resource_id = "com.ss.android.ugc.aweme:id/f1s"
    next_video_delay_time = 1
    auto_stop = 0

    def __init__(
        self, app_name: str, device: u2.Device, next_video_delay_time: int = 0
    ):
        """
        Args:
            app_name:
            device:
            next_video_delay_time:  滑动下一个视频的间隔时间
        """

        if next_video_delay_time:
            self.next_video_delay_time = next_video_delay_time
        self.device = device
        self.app_name = app_name

    def set_implicitly_wait(self, t: float = 5):
        """
        设置元素查找等待时间
        Args:
            t: 设置元素查找等待时间

        Returns:

        """
        self.device.implicitly_wait(t)
        self.logger.info(f"set implicitly wait : {t}")

    def set_wait_time(self, t: float = 3):
        """
        默认控件等待时间
        Args:
            t:

        Returns:

        """
        self.device.settings["wait_timeout"] = t
        self.logger.info(f"set wait timeout: {t}")

    def start_douyin(self):
        self.device.app_start(self.app_name, use_monkey=True)
        self.logger.info(self.device.app_info(self.app_name))

    @timethis
    def is_ad(self):
        if self.device(resourceId="com.ss.android.ugc.aweme:id/h1q").exists(
            timeout=0.3
        ):
            return True
        return False

    @timethis
    def is_live(self) -> (bool, str):
        """

        Returns: (bool, str)  (是否为直播, 进入直播页面资源值)

        """
        r_id_a = "com.ss.android.ugc.aweme:id/szv"
        if self.device(
            resourceId=r_id_a, text="点击进入直播间", className="android.widget.TextView"
        ).exists(timeout=0.3):
            return True, r_id_a
        r_id_a = "com.ss.android.ugc.aweme:id/t3t"
        if self.device(resourceId=r_id_a, className="android.widget.Button").exists(
            timeout=0.3
        ):
            return True, r_id_a
        return False, ""

    @timethis
    def is_follow_when_live(self, rs_id):
        """
        是否关注该直播
        Returns:
        """
        if rs_id == "com.ss.android.ugc.aweme:id/szv" and self.device(
            resourceId="com.ss.android.ugc.aweme:id/jsa"
        ).exists(timeout=0.3):
            return True
        return rs_id == "com.ss.android.ugc.aweme:id/t3t" and not self.device(
            resourceId="com.ss.android.ugc.aweme:id/f1s"
        ).exists(timeout=0.2)

    def enter_to_live_room(self, rs_id: str):
        self.device(resourceId=rs_id).click()

    def next_video(self):
        self.device.swipe_ext("up", scale=0.8)
        if self.is_ad():
            self.device.swipe_ext("up", scale=0.8)

    def last_video(self):
        self.device.swipe_ext("down", scale=0.8)
        self.logger.info("下滑~")

    @timethis
    def check_follow_in_main_page(self) -> bool:
        return not self.device(resourceId="com.ss.android.ugc.aweme:id/f1s").exists(
            timeout=0.5
        )

    def check_follow_in_user_info_page(self) -> bool:
        return bool(
            self.device(resourceId="com.ss.android.ugc.aweme:id/f14").exists(
                timeout=0.5
            )
        )

    def follow(self):
        self.device(resourceId=self.follow_button_resource_id).click_exists()
        self.logger.info("点击关注...")
        self.next_video()

    def enter_user_info_from_live(self):
        """
        从直播间进入用户主页
        Returns:

        """
        # 点击直播间左上角头像
        self.device(resourceId="com.ss.android.ugc.aweme:id/go6").click()
        # 点击右下角头像
        self.device(resourceId="com.ss.android.ugc.aweme:id/lqw").click()
        self.logger.info("从直播间跳转到个人页...")

    def is_avatar_in_main_page(self):
        return bool(
            self.device(resourceId="com.ss.android.ugc.aweme:id/user_avatar").exists(
                timeout=0.3
            )
        )

    def click_avatar_in_main_page(self):
        # 点击头像
        self.device(resourceId="com.ss.android.ugc.aweme:id/user_avatar").click()
        self.logger.info("点击首页头像...")

    def skip_buy_page_in_user_info(self):
        if self.device(resourceId="com.ss.android.ugc.aweme:id/m71").exists(
            timeout=0.1
        ):
            self.device(resourceId="com.ss.android.ugc.aweme:id/back_btn").click()
            self.logger.info("存在购买页面, 点击 x ,关闭...")

    def block(self):
        self.skip_buy_page_in_user_info()
        # 点击 ...  拉黑
        self.device(resourceId="com.ss.android.ugc.aweme:id/rd2").click(timeout=1)
        # 页面弹出有延迟, 等待一秒查询
        self.device(resourceId="com.ss.android.ugc.aweme:id/desc", text="拉黑").click(
            timeout=1
        )
        self.device(resourceId="com.ss.android.ugc.aweme:id/c6a").click()
        self.device(resourceId="com.ss.android.ugc.aweme:id/back_btn").click()
        self.logger.info("点击拉黑...")

    def click_x_in_live_room(self):
        # 从直播间进入用户页拉黑后, 点击返回会返回直播间, 点击 x 返回
        self.device(resourceId="com.ss.android.ugc.aweme:id/cj9").click()

    @timethis
    def block_and_swipe_next_video(self):
        self.logger.info("block_and_swipe_next_video 开始执行 ...")
        live, rs_id = self.is_live()
        if live:
            if self.is_follow_when_live(rs_id):
                self.next_video()
                self.logger.info("当前直播已关注, 划过~\n block_and_swipe_next_video 执行完成...")
                return
            self.logger.info("当前视频为直播且未关注,点击进入直播间...")
            self.enter_to_live_room(rs_id)
            self.enter_user_info_from_live()
            self.block()
            self.click_x_in_live_room()
            self.logger.info("点击 x, 退出直播间...")
        elif self.is_ad():
            self.next_video()
            self.logger.info("该视频为广告,滑过~")
        elif self.check_follow_in_main_page():
            self.next_video()
            self.logger.info("已关注, 下一个视频...")
        else:
            self.click_avatar_in_main_page()
            self.block()
        self.logger.info("block_and_swipe_next_video 执行完成...")
