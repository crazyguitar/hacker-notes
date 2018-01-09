package tls

const (
	// Cipher Suites
	TLS_NULL_WITH_NULL_NULL                           = 0x0000
	TLS_RSA_WITH_NULL_MD5                             = 0x0001
	TLS_RSA_WITH_NULL_SHA                             = 0x0002
	TLS_RSA_EXPORT_WITH_RC4_40_MD5                    = 0x0003
	TLS_RSA_WITH_RC4_128_MD5                          = 0x0004
	TLS_RSA_WITH_RC4_128_SHA                          = 0x0005
	TLS_RSA_EXPORT_WITH_RC2_CBC_40_MD5                = 0x0006
	TLS_RSA_WITH_IDEA_CBC_SHA                         = 0x0007
	TLS_RSA_EXPORT_WITH_DES40_CBC_SHA                 = 0x0008
	TLS_RSA_WITH_DES_CBC_SHA                          = 0x0009
	TLS_RSA_WITH_3DES_EDE_CBC_SHA                     = 0x000A
	TLS_DH_DSS_EXPORT_WITH_DES40_CBC_SHA              = 0x000B
	TLS_DH_DSS_WITH_DES_CBC_SHA                       = 0x000C
	TLS_DH_DSS_WITH_3DES_EDE_CBC_SHA                  = 0x000D
	TLS_DH_RSA_EXPORT_WITH_DES40_CBC_SHA              = 0x000E
	TLS_DH_RSA_WITH_DES_CBC_SHA                       = 0x000F
	TLS_DH_RSA_WITH_3DES_EDE_CBC_SHA                  = 0x0010
	TLS_DHE_DSS_EXPORT_WITH_DES40_CBC_SHA             = 0x0011
	TLS_DHE_DSS_WITH_DES_CBC_SHA                      = 0x0012
	TLS_DHE_DSS_WITH_3DES_EDE_CBC_SHA                 = 0x0013
	TLS_DHE_RSA_EXPORT_WITH_DES40_CBC_SHA             = 0x0014
	TLS_DHE_RSA_WITH_DES_CBC_SHA                      = 0x0015
	TLS_DHE_RSA_WITH_3DES_EDE_CBC_SHA                 = 0x0016
	TLS_DH_anon_EXPORT_WITH_RC4_40_MD5                = 0x0017
	TLS_DH_anon_WITH_RC4_128_MD5                      = 0x0018
	TLS_DH_anon_EXPORT_WITH_DES40_CBC_SHA             = 0x0019
	TLS_DH_anon_WITH_DES_CBC_SHA                      = 0x001A
	TLS_DH_anon_WITH_3DES_EDE_CBC_SHA                 = 0x001B
	TLS_KRB5_WITH_DES_CBC_SHA                         = 0x001E
	TLS_KRB5_WITH_3DES_EDE_CBC_SHA                    = 0x001F
	TLS_KRB5_WITH_RC4_128_SHA                         = 0x0020
	TLS_KRB5_WITH_IDEA_CBC_SHA                        = 0x0021
	TLS_KRB5_WITH_DES_CBC_MD5                         = 0x0022
	TLS_KRB5_WITH_3DES_EDE_CBC_MD5                    = 0x0023
	TLS_KRB5_WITH_RC4_128_MD5                         = 0x0024
	TLS_KRB5_WITH_IDEA_CBC_MD5                        = 0x0025
	TLS_KRB5_EXPORT_WITH_DES40_CBC_SHA                = 0x0026
	TLS_KRB5_EXPORT_WITH_RC2_CBC_40_SHA               = 0x0027
	TLS_KRB5_EXPORT_WITH_RC4_40_SHA                   = 0x0028
	TLS_KRB5_EXPORT_WITH_DES40_CBC_MD5                = 0x0029
	TLS_KRB5_EXPORT_WITH_RC2_CBC_40_MD5               = 0x002A
	TLS_KRB5_EXPORT_WITH_RC4_40_MD5                   = 0x002B
	TLS_PSK_WITH_NULL_SHA                             = 0x002C
	TLS_DHE_PSK_WITH_NULL_SHA                         = 0x002D
	TLS_RSA_PSK_WITH_NULL_SHA                         = 0x002E
	TLS_RSA_WITH_AES_128_CBC_SHA                      = 0x002F
	TLS_DH_DSS_WITH_AES_128_CBC_SHA                   = 0x0030
	TLS_DH_RSA_WITH_AES_128_CBC_SHA                   = 0x0031
	TLS_DHE_DSS_WITH_AES_128_CBC_SHA                  = 0x0032
	TLS_DHE_RSA_WITH_AES_128_CBC_SHA                  = 0x0033
	TLS_DH_anon_WITH_AES_128_CBC_SHA                  = 0x0034
	TLS_RSA_WITH_AES_256_CBC_SHA                      = 0x0035
	TLS_DH_DSS_WITH_AES_256_CBC_SHA                   = 0x0036
	TLS_DH_RSA_WITH_AES_256_CBC_SHA                   = 0x0037
	TLS_DHE_DSS_WITH_AES_256_CBC_SHA                  = 0x0038
	TLS_DHE_RSA_WITH_AES_256_CBC_SHA                  = 0x0039
	TLS_DH_anon_WITH_AES_256_CBC_SHA                  = 0x003A
	TLS_RSA_WITH_NULL_SHA256                          = 0x003B
	TLS_RSA_WITH_AES_128_CBC_SHA256                   = 0x003C
	TLS_RSA_WITH_AES_256_CBC_SHA256                   = 0x003D
	TLS_DH_DSS_WITH_AES_128_CBC_SHA256                = 0x003E
	TLS_DH_RSA_WITH_AES_128_CBC_SHA256                = 0x003F
	TLS_DHE_DSS_WITH_AES_128_CBC_SHA256               = 0x0040
	TLS_RSA_WITH_CAMELLIA_128_CBC_SHA                 = 0x0041
	TLS_DH_DSS_WITH_CAMELLIA_128_CBC_SHA              = 0x0042
	TLS_DH_RSA_WITH_CAMELLIA_128_CBC_SHA              = 0x0043
	TLS_DHE_DSS_WITH_CAMELLIA_128_CBC_SHA             = 0x0044
	TLS_DHE_RSA_WITH_CAMELLIA_128_CBC_SHA             = 0x0045
	TLS_DH_anon_WITH_CAMELLIA_128_CBC_SHA             = 0x0046
	TLS_DHE_RSA_WITH_AES_128_CBC_SHA256               = 0x0067
	TLS_DH_DSS_WITH_AES_256_CBC_SHA256                = 0x0068
	TLS_DH_RSA_WITH_AES_256_CBC_SHA256                = 0x0069
	TLS_DHE_DSS_WITH_AES_256_CBC_SHA256               = 0x006A
	TLS_DHE_RSA_WITH_AES_256_CBC_SHA256               = 0x006B
	TLS_DH_anon_WITH_AES_128_CBC_SHA256               = 0x006C
	TLS_DH_anon_WITH_AES_256_CBC_SHA256               = 0x006D
	TLS_RSA_WITH_CAMELLIA_256_CBC_SHA                 = 0x0084
	TLS_DH_DSS_WITH_CAMELLIA_256_CBC_SHA              = 0x0085
	TLS_DH_RSA_WITH_CAMELLIA_256_CBC_SHA              = 0x0086
	TLS_DHE_DSS_WITH_CAMELLIA_256_CBC_SHA             = 0x0087
	TLS_DHE_RSA_WITH_CAMELLIA_256_CBC_SHA             = 0x0088
	TLS_DH_anon_WITH_CAMELLIA_256_CBC_SHA             = 0x0089
	TLS_PSK_WITH_RC4_128_SHA                          = 0x008A
	TLS_PSK_WITH_3DES_EDE_CBC_SHA                     = 0x008B
	TLS_PSK_WITH_AES_128_CBC_SHA                      = 0x008C
	TLS_PSK_WITH_AES_256_CBC_SHA                      = 0x008D
	TLS_DHE_PSK_WITH_RC4_128_SHA                      = 0x008E
	TLS_DHE_PSK_WITH_3DES_EDE_CBC_SHA                 = 0x008F
	TLS_DHE_PSK_WITH_AES_128_CBC_SHA                  = 0x0090
	TLS_DHE_PSK_WITH_AES_256_CBC_SHA                  = 0x0091
	TLS_RSA_PSK_WITH_RC4_128_SHA                      = 0x0092
	TLS_RSA_PSK_WITH_3DES_EDE_CBC_SHA                 = 0x0093
	TLS_RSA_PSK_WITH_AES_128_CBC_SHA                  = 0x0094
	TLS_RSA_PSK_WITH_AES_256_CBC_SHA                  = 0x0095
	TLS_RSA_WITH_SEED_CBC_SHA                         = 0x0096
	TLS_DH_DSS_WITH_SEED_CBC_SHA                      = 0x0097
	TLS_DH_RSA_WITH_SEED_CBC_SHA                      = 0x0098
	TLS_DHE_DSS_WITH_SEED_CBC_SHA                     = 0x0099
	TLS_DHE_RSA_WITH_SEED_CBC_SHA                     = 0x009A
	TLS_DH_anon_WITH_SEED_CBC_SHA                     = 0x009B
	TLS_RSA_WITH_AES_128_GCM_SHA256                   = 0x009C
	TLS_RSA_WITH_AES_256_GCM_SHA384                   = 0x009D
	TLS_DHE_RSA_WITH_AES_128_GCM_SHA256               = 0x009E
	TLS_DHE_RSA_WITH_AES_256_GCM_SHA384               = 0x009F
	TLS_DH_RSA_WITH_AES_128_GCM_SHA256                = 0x00A0
	TLS_DH_RSA_WITH_AES_256_GCM_SHA384                = 0x00A1
	TLS_DHE_DSS_WITH_AES_128_GCM_SHA256               = 0x00A2
	TLS_DHE_DSS_WITH_AES_256_GCM_SHA384               = 0x00A3
	TLS_DH_DSS_WITH_AES_128_GCM_SHA256                = 0x00A4
	TLS_DH_DSS_WITH_AES_256_GCM_SHA384                = 0x00A5
	TLS_DH_anon_WITH_AES_128_GCM_SHA256               = 0x00A6
	TLS_DH_anon_WITH_AES_256_GCM_SHA384               = 0x00A7
	TLS_PSK_WITH_AES_128_GCM_SHA256                   = 0x00A8
	TLS_PSK_WITH_AES_256_GCM_SHA384                   = 0x00A9
	TLS_DHE_PSK_WITH_AES_128_GCM_SHA256               = 0x00AA
	TLS_DHE_PSK_WITH_AES_256_GCM_SHA384               = 0x00AB
	TLS_RSA_PSK_WITH_AES_128_GCM_SHA256               = 0x00AC
	TLS_RSA_PSK_WITH_AES_256_GCM_SHA384               = 0x00AD
	TLS_PSK_WITH_AES_128_CBC_SHA256                   = 0x00AE
	TLS_PSK_WITH_AES_256_CBC_SHA384                   = 0x00AF
	TLS_PSK_WITH_NULL_SHA256                          = 0x00B0
	TLS_PSK_WITH_NULL_SHA384                          = 0x00B1
	TLS_DHE_PSK_WITH_AES_128_CBC_SHA256               = 0x00B2
	TLS_DHE_PSK_WITH_AES_256_CBC_SHA384               = 0x00B3
	TLS_DHE_PSK_WITH_NULL_SHA256                      = 0x00B4
	TLS_DHE_PSK_WITH_NULL_SHA384                      = 0x00B5
	TLS_RSA_PSK_WITH_AES_128_CBC_SHA256               = 0x00B6
	TLS_RSA_PSK_WITH_AES_256_CBC_SHA384               = 0x00B7
	TLS_RSA_PSK_WITH_NULL_SHA256                      = 0x00B8
	TLS_RSA_PSK_WITH_NULL_SHA384                      = 0x00B9
	TLS_RSA_WITH_CAMELLIA_128_CBC_SHA256              = 0x00BA
	TLS_DH_DSS_WITH_CAMELLIA_128_CBC_SHA256           = 0x00BB
	TLS_DH_RSA_WITH_CAMELLIA_128_CBC_SHA256           = 0x00BC
	TLS_DHE_DSS_WITH_CAMELLIA_128_CBC_SHA256          = 0x00BD
	TLS_DHE_RSA_WITH_CAMELLIA_128_CBC_SHA256          = 0x00BE
	TLS_DH_anon_WITH_CAMELLIA_128_CBC_SHA256          = 0x00BF
	TLS_RSA_WITH_CAMELLIA_256_CBC_SHA256              = 0x00C0
	TLS_DH_DSS_WITH_CAMELLIA_256_CBC_SHA256           = 0x00C1
	TLS_DH_RSA_WITH_CAMELLIA_256_CBC_SHA256           = 0x00C2
	TLS_DHE_DSS_WITH_CAMELLIA_256_CBC_SHA256          = 0x00C3
	TLS_DHE_RSA_WITH_CAMELLIA_256_CBC_SHA256          = 0x00C4
	TLS_DH_anon_WITH_CAMELLIA_256_CBC_SHA256          = 0x00C5
	TLS_ECDH_ECDSA_WITH_NULL_SHA                      = 0xC001
	TLS_ECDH_ECDSA_WITH_RC4_128_SHA                   = 0xC002
	TLS_ECDH_ECDSA_WITH_3DES_EDE_CBC_SHA              = 0xC003
	TLS_ECDH_ECDSA_WITH_AES_128_CBC_SHA               = 0xC004
	TLS_ECDH_ECDSA_WITH_AES_256_CBC_SHA               = 0xC005
	TLS_ECDHE_ECDSA_WITH_NULL_SHA                     = 0xC006
	TLS_ECDHE_ECDSA_WITH_RC4_128_SHA                  = 0xC007
	TLS_ECDHE_ECDSA_WITH_3DES_EDE_CBC_SHA             = 0xC008
	TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA              = 0xC009
	TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA              = 0xC00A
	TLS_ECDH_RSA_WITH_NULL_SHA                        = 0xC00B
	TLS_ECDH_RSA_WITH_RC4_128_SHA                     = 0xC00C
	TLS_ECDH_RSA_WITH_3DES_EDE_CBC_SHA                = 0xC00D
	TLS_ECDH_RSA_WITH_AES_128_CBC_SHA                 = 0xC00E
	TLS_ECDH_RSA_WITH_AES_256_CBC_SHA                 = 0xC00F
	TLS_ECDHE_RSA_WITH_NULL_SHA                       = 0xC010
	TLS_ECDHE_RSA_WITH_RC4_128_SHA                    = 0xC011
	TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA               = 0xC012
	TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA                = 0xC013
	TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA                = 0xC014
	TLS_ECDH_anon_WITH_NULL_SHA                       = 0xC015
	TLS_ECDH_anon_WITH_RC4_128_SHA                    = 0xC016
	TLS_ECDH_anon_WITH_3DES_EDE_CBC_SHA               = 0xC017
	TLS_ECDH_anon_WITH_AES_128_CBC_SHA                = 0xC018
	TLS_ECDH_anon_WITH_AES_256_CBC_SHA                = 0xC019
	TLS_SRP_SHA_WITH_3DES_EDE_CBC_SHA                 = 0xC01A
	TLS_SRP_SHA_RSA_WITH_3DES_EDE_CBC_SHA             = 0xC01B
	TLS_SRP_SHA_DSS_WITH_3DES_EDE_CBC_SHA             = 0xC01C
	TLS_SRP_SHA_WITH_AES_128_CBC_SHA                  = 0xC01D
	TLS_SRP_SHA_RSA_WITH_AES_128_CBC_SHA              = 0xC01E
	TLS_SRP_SHA_DSS_WITH_AES_128_CBC_SHA              = 0xC01F
	TLS_SRP_SHA_WITH_AES_256_CBC_SHA                  = 0xC020
	TLS_SRP_SHA_RSA_WITH_AES_256_CBC_SHA              = 0xC021
	TLS_SRP_SHA_DSS_WITH_AES_256_CBC_SHA              = 0xC022
	TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256           = 0xC023
	TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384           = 0xC024
	TLS_ECDH_ECDSA_WITH_AES_128_CBC_SHA256            = 0xC025
	TLS_ECDH_ECDSA_WITH_AES_256_CBC_SHA384            = 0xC026
	TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256             = 0xC027
	TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384             = 0xC028
	TLS_ECDH_RSA_WITH_AES_128_CBC_SHA256              = 0xC029
	TLS_ECDH_RSA_WITH_AES_256_CBC_SHA384              = 0xC02A
	TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256           = 0xC02B
	TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384           = 0xC02C
	TLS_ECDH_ECDSA_WITH_AES_128_GCM_SHA256            = 0xC02D
	TLS_ECDH_ECDSA_WITH_AES_256_GCM_SHA384            = 0xC02E
	TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256             = 0xC02F
	TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384             = 0xC030
	TLS_ECDH_RSA_WITH_AES_128_GCM_SHA256              = 0xC031
	TLS_ECDH_RSA_WITH_AES_256_GCM_SHA384              = 0xC032
	TLS_ECDHE_PSK_WITH_RC4_128_SHA                    = 0xC033
	TLS_ECDHE_PSK_WITH_3DES_EDE_CBC_SHA               = 0xC034
	TLS_ECDHE_PSK_WITH_AES_128_CBC_SHA                = 0xC035
	TLS_ECDHE_PSK_WITH_AES_256_CBC_SHA                = 0xC036
	TLS_ECDHE_PSK_WITH_AES_128_CBC_SHA256             = 0xC037
	TLS_ECDHE_PSK_WITH_AES_256_CBC_SHA384             = 0xC038
	TLS_ECDHE_PSK_WITH_NULL_SHA                       = 0xC039
	TLS_ECDHE_PSK_WITH_NULL_SHA256                    = 0xC03A
	TLS_ECDHE_PSK_WITH_NULL_SHA384                    = 0xC03B
	TLS_ECDHE_ECDSA_WITH_CAMELLIA_128_CBC_SHA256      = 0xC072
	TLS_ECDHE_ECDSA_WITH_CAMELLIA_256_CBC_SHA384      = 0xC073
	TLS_ECDH_ECDSA_WITH_CAMELLIA_128_CBC_SHA256       = 0xC074
	TLS_ECDH_ECDSA_WITH_CAMELLIA_256_CBC_SHA384       = 0xC075
	TLS_ECDHE_RSA_WITH_CAMELLIA_128_CBC_SHA256        = 0xC076
	TLS_ECDHE_RSA_WITH_CAMELLIA_256_CBC_SHA384        = 0xC077
	TLS_ECDH_RSA_WITH_CAMELLIA_128_CBC_SHA256         = 0xC078
	TLS_ECDH_RSA_WITH_CAMELLIA_256_CBC_SHA384         = 0xC079
	TLS_RSA_WITH_CAMELLIA_128_GCM_SHA256              = 0xC07A
	TLS_RSA_WITH_CAMELLIA_256_GCM_SHA384              = 0xC07B
	TLS_DHE_RSA_WITH_CAMELLIA_128_GCM_SHA256          = 0xC07C
	TLS_DHE_RSA_WITH_CAMELLIA_256_GCM_SHA384          = 0xC07D
	TLS_DH_RSA_WITH_CAMELLIA_128_GCM_SHA256           = 0xC07E
	TLS_DH_RSA_WITH_CAMELLIA_256_GCM_SHA384           = 0xC07F
	TLS_DHE_DSS_WITH_CAMELLIA_128_GCM_SHA256          = 0xC080
	TLS_DHE_DSS_WITH_CAMELLIA_256_GCM_SHA384          = 0xC081
	TLS_DH_DSS_WITH_CAMELLIA_128_GCM_SHA256           = 0xC082
	TLS_DH_DSS_WITH_CAMELLIA_256_GCM_SHA384           = 0xC083
	TLS_DH_anon_WITH_CAMELLIA_128_GCM_SHA256          = 0xC084
	TLS_DH_anon_WITH_CAMELLIA_256_GCM_SHA384          = 0xC085
	TLS_ECDHE_ECDSA_WITH_CAMELLIA_128_GCM_SHA256      = 0xC086
	TLS_ECDHE_ECDSA_WITH_CAMELLIA_256_GCM_SHA384      = 0xC087
	TLS_ECDH_ECDSA_WITH_CAMELLIA_128_GCM_SHA256       = 0xC088
	TLS_ECDH_ECDSA_WITH_CAMELLIA_256_GCM_SHA384       = 0xC089
	TLS_ECDHE_RSA_WITH_CAMELLIA_128_GCM_SHA256        = 0xC08A
	TLS_ECDHE_RSA_WITH_CAMELLIA_256_GCM_SHA384        = 0xC08B
	TLS_ECDH_RSA_WITH_CAMELLIA_128_GCM_SHA256         = 0xC08C
	TLS_ECDH_RSA_WITH_CAMELLIA_256_GCM_SHA384         = 0xC08D
	TLS_PSK_WITH_CAMELLIA_128_GCM_SHA256              = 0xC08E
	TLS_PSK_WITH_CAMELLIA_256_GCM_SHA384              = 0xC08F
	TLS_DHE_PSK_WITH_CAMELLIA_128_GCM_SHA256          = 0xC090
	TLS_DHE_PSK_WITH_CAMELLIA_256_GCM_SHA384          = 0xC091
	TLS_RSA_PSK_WITH_CAMELLIA_128_GCM_SHA256          = 0xC092
	TLS_RSA_PSK_WITH_CAMELLIA_256_GCM_SHA384          = 0xC093
	TLS_PSK_WITH_CAMELLIA_128_CBC_SHA256              = 0xC094
	TLS_PSK_WITH_CAMELLIA_256_CBC_SHA384              = 0xC095
	TLS_DHE_PSK_WITH_CAMELLIA_128_CBC_SHA256          = 0xC096
	TLS_DHE_PSK_WITH_CAMELLIA_256_CBC_SHA384          = 0xC097
	TLS_RSA_PSK_WITH_CAMELLIA_128_CBC_SHA256          = 0xC098
	TLS_RSA_PSK_WITH_CAMELLIA_256_CBC_SHA384          = 0xC099
	TLS_ECDHE_PSK_WITH_CAMELLIA_128_CBC_SHA256        = 0xC09A
	TLS_ECDHE_PSK_WITH_CAMELLIA_256_CBC_SHA384        = 0xC09B
	TLS_RSA_WITH_AES_128_CCM                          = 0xC09C
	TLS_RSA_WITH_AES_256_CCM                          = 0xC09D
	TLS_DHE_RSA_WITH_AES_128_CCM                      = 0xC09E
	TLS_DHE_RSA_WITH_AES_256_CCM                      = 0xC09F
	TLS_RSA_WITH_AES_128_CCM_8                        = 0xC0A0
	TLS_RSA_WITH_AES_256_CCM_8                        = 0xC0A1
	TLS_DHE_RSA_WITH_AES_128_CCM_8                    = 0xC0A2
	TLS_DHE_RSA_WITH_AES_256_CCM_8                    = 0xC0A3
	TLS_PSK_WITH_AES_128_CCM                          = 0xC0A4
	TLS_PSK_WITH_AES_256_CCM                          = 0xC0A5
	TLS_DHE_PSK_WITH_AES_128_CCM                      = 0xC0A6
	TLS_DHE_PSK_WITH_AES_256_CCM                      = 0xC0A7
	TLS_PSK_WITH_AES_128_CCM_8                        = 0xC0A8
	TLS_PSK_WITH_AES_256_CCM_8                        = 0xC0A9
	TLS_DHE_PSK_WITH_AES_128_CCM_8                    = 0xC0AA
	TLS_DHE_PSK_WITH_AES_256_CCM_8                    = 0xC0AB
	TLS_ECDHE_ECDSA_WITH_AES_128_CCM                  = 0xC0AC
	TLS_ECDHE_ECDSA_WITH_AES_256_CCM                  = 0xC0AD
	TLS_ECDHE_ECDSA_WITH_AES_128_CCM_8                = 0xC0AE
	TLS_ECDHE_ECDSA_WITH_AES_256_CCM_8                = 0xC0AF
	TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256_OLD   = 0xCC13
	TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256_OLD = 0xCC14
	TLS_DHE_RSA_WITH_CHACHA20_POLY1305_SHA256_OLD     = 0xCC15
	TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256       = 0xCCA8
	TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256     = 0xCCA9
	TLS_DHE_RSA_WITH_CHACHA20_POLY1305_SHA256         = 0xCCAA
	TLS_PSK_WITH_CHACHA20_POLY1305_SHA256             = 0xCCAB
	TLS_ECDHE_PSK_WITH_CHACHA20_POLY1305_SHA256       = 0xCCAC
	TLS_DHE_PSK_WITH_CHACHA20_POLY1305_SHA256         = 0xCCAD
	TLS_RSA_PSK_WITH_CHACHA20_POLY1305_SHA256         = 0xCCAE
	TLS_AES_128_GCM_SHA256                            = 0x1301
	TLS_AES_256_GCM_SHA384                            = 0x1302
	TLS_CHACHA20_POLY1305_SHA256                      = 0x1303
	TLS_AES_128_CCM_SHA256                            = 0x1304
	TLS_AES_128_CCM_8_SHA256                          = 0x1305
	SSL_CK_RC4_128_WITH_MD5                           = 0x010080
	SSL_CK_RC4_128_EXPORT40_WITH_MD5                  = 0x020080
	SSL_CK_RC2_128_CBC_WITH_MD5                       = 0x030080
	SSL_CK_RC2_128_CBC_EXPORT40_WITH_MD5              = 0x040080
	SSL_CK_IDEA_128_CBC_WITH_MD5                      = 0x050080
	SSL_CK_DES_64_CBC_WITH_MD5                        = 0x060040
	SSL_CK_DES_192_EDE3_CBC_WITH_MD5                  = 0x0700C0
)
