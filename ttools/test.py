#-*- coding: utf-8 -*-
import requests
import re
import time
import hashlib
import base64
import datetime
import struct

URL = "http://172.16.154.172:10085"  #待修改
REQ_TYPE = "login"
USERNAME = "davidclaude111"
PASSWORD = "success0325"
CNT_TYPE = "chksmd5"

def getHeader(timestamp):
    headers = {
        "req_tp":REQ_TYPE,
        "ts":timestamp,
        "usr":USERNAME,
        "cnt_tp":CNT_TYPE
    }
    return headers

def getLoginData(timestamp):
    checksum = PASSWORD + timestamp + USERNAME
    m2 = hashlib.md5()
    m2.update(checksum)
    checksumMd5 = m2.hexdigest()
    return bytes(checksumMd5)

def getRegisterData():
    return bytes(base64.b64encode(PASSWORD))

ts = str(int(time.time()))
resp = requests.post(URL, data=getLoginData(ts), headers=getHeader(ts))
print ("Code: " + resp.headers["code"] + ", desc: " + resp.headers["desc"])