from __future__ import print_function, unicode_literals

import os
import sys
import errno
import platform

from ctypes import *

'''
Example
-------

    $ python sendfile.py /etc/resolv.conf resolve.conf; cat resolve.conf
    nameserver  192.168.1.1
'''

# check os
p = platform.system()
if p != "Linux":
    raise OSError("Not support '{}'".format(p))

# check linux version
ver = platform.release()
if tuple(map(int, ver.split('.'))) < (2,6,33):
    raise OSError("Upgrade kernel after 2.6.33")

# check input arguments
if len(sys.argv) != 3:
    print("Usage: sendfile.py f1 f2", file=sys.stderr)
    exit(1)

libc = CDLL('libc.so.6', use_errno=True)
sendfile = libc.sendfile

src = sys.argv[1]
dst = sys.argv[2]
src_size = os.stat(src).st_size

# clean destination first
try:
    os.remove(dst)
except OSError as e:
    if e.errno != errno.ENOENT: raise

offset = c_int64(0)

with open(src, 'r') as f1:
    with open(dst, 'w') as f2:
        src_fd = c_int(f1.fileno())
        dst_fd = c_int(f2.fileno())
        ret = sendfile(dst_fd, src_fd, byref(offset), src_size)
        if ret < 0:
            errno_ = get_errno()
            errmsg = "sendfile failed. {}".format(os.strerror(errno_))
            raise OSError(errno_, errmsg)
