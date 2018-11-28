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
USERNAME = "davidclaude"
PASSWORD = "success0325"

def getHeader(timestamp):
    headers = {
        "req_type":REQ_TYPE,
        "time_stamp":timestamp,
        "user_name":USERNAME,
    }
    return headers

def getData(timestamp):
    checksum = PASSWORD + timestamp + USERNAME
    m2 = hashlib.md5()
    m2.update(checksum)
    checksumMd5 = m2.hexdigest()
    return bytes(checksumMd5)


ts = str(int(time.time()))
resp = requests.post(URL, data=getData(ts), headers=getHeader(ts))

code = resp.headers["code"]
if code != "0":
    print ("ERROR: " + resp.headers["err"] + ", DESC: " + resp.headers["desc"])
else:
    print ("OK")