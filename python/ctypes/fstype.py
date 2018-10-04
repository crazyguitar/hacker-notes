from __future__ import print_function

from ctypes import *
from sys import platform

'''
Usage
-----

    $ python3 statfs.py
    Is ext4? True

'''

if platform not in ('linux', 'linux2'):
    raise RuntimeError("Not support '{}'".format(platform))


# from Linux/include/uapi/linux/magic.h

EXT_SUPER_MAGIC      = 0x137D
EXT2_OLD_SUPER_MAGIC = 0xEF51
EXT2_SUPER_MAGIC     = 0xEF53
EXT3_SUPER_MAGIC     = 0xEF53
EXT4_SUPER_MAGIC     = 0xEF53
BTRFS_SUPER_MAGIC    = 0x9123683E


class KernelFsid(Structure):
    '''
    From Linux/arch/mips/include/asm/posix_types.h

    typedef struct {
            long    val[2];
    } __kernel_fsid_t;
    '''
    _fields_ = [('val', POINTER(c_long) * 2)]

class Statfs(Structure):
    '''
    From Linux/arch/mips/include/asm/statfs.h

    struct statfs {
            long            f_type;
    #define f_fstyp f_type
            long            f_bsize;
            long            f_frsize;
            long            f_blocks;
            long            f_bfree;
            long            f_files;
            long            f_ffree;
            long            f_bavail;

            /* Linux specials */
            __kernel_fsid_t f_fsid;
            long            f_namelen;
            long            f_flags;
            long            f_spare[5];
    };
    '''
    _fields_ = [('f_type',    c_long),
                ('f_bsize',   c_long),
                ('f_frsize',  c_long),
                ('f_block',   c_long),
                ('f_bfree',   c_long),
                ('f_files',   c_long),
                ('f_ffree',   c_long),
                ('f_fsid',    KernelFsid),
                ('f_namelen', c_long),
                ('f_flags',   c_long),
                ('f_spare',   POINTER(c_long) * 5)]


libc = CDLL('libc.so.6', use_errno=True)
statfs = libc.statfs

path = create_string_buffer(b'/etc')
fst = Statfs()
ret = statfs(path, byref(fst))
assert ret == 0

print('Is ext4? {}'.format(fst.f_type == EXT4_SUPER_MAGIC))
