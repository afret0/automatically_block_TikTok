import logging

import logzero

formatter = logzero.LogFormatter(datefmt=logging.Formatter.default_time_format)
logger = logzero.setup_logger(
    name="auto_block_douyin",
    formatter=formatter,
    logfile="log.txt",
    level=logging.INFO,
    maxBytes=600000,
    backupCount=3,
)
