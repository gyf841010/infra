package errorUtil

import (
	"fmt"
)

const HERROR_PREFIX = "HERROR:"

//ZError的基本code，如果service相关的code，请service自行定义
const (
	ERROR_CODE_OTHER_SERVER_ERROR                 = 111
	ERROR_CODE_PARAMETER_FORMAT_INVALID           = 400
	ERROR_CODE_ACCESS_TOKEN_TIMEOUT               = 401
	ERROR_CODE_NOT_FOUND                          = 404
	ERROR_CODE_PARAMETER_AUTH_INVALID             = 406
	ERROR_INTERNAL_SERVER_ERROR                   = 500
	ERROR_CODE_ACCESS_TOKEN_ERROR                 = 1001
	ERROR_CODE_MOBILE_IS_REGISTERED               = 1002
	ERROR_CODE_MOBILE_IS_NOT_EXISTED              = 1003
	ERROR_CODE_CAPTCHA_TYPE_ERROR                 = 1004
	ERROR_CODE_USER_NOT_FOUND                     = 1005
	ERROR_CODE_MOBILE_IS_INVALID                  = 1006
	ERROR_CODE_REQUEST_SMS_OVERLOAD               = 1007
	ERROR_CODE_CAPTCHA_TOKEN_ERROR                = 1008
	ERROR_CODE_CM_ACTION_NOT_FOUND                = 1009
	ERROR_CODE_AUTO_LOGIN_FAILED                  = 1010
	ERROR_CODE_INVALID_MOBILE                     = 1011
	ERROR_CODE_LOGIN_FAILED                       = 1012
	ERROR_CODE_CM_ACTION_TIME_NOT_FOUND           = 1013
	ERROR_CODE_REPORT_ACTION_TYPE_ERROR           = 1014
	ERROR_CODE_REDIS_KEY_NULL                     = 1015
	ERROR_CODE_REDIS_VALUE_NULL_PTR               = 1016
	ERROR_CODE_REDIS_VALUE_NULL                   = 1017
	ERROR_CODE_REDIS_KEY_NOT_EXIST                = 1018
	ERROR_CODE_REPORT_ACTION_DATA_ERROR           = 1019
	ERROR_CODE_SN_NULL                            = 1020
	ERROR_CODE_MAC_NULL                           = 1021
	ERROR_CODE_FIRMWORD_TYPE_NULL                 = 1022
	ERROR_CODE_MANUFACTURE_FEATURE_ERROR          = 1023
	ERROR_CODE_MANUFACTURE_AUTH_ERROR             = 1024
	ERROR_CODE_DEVICE_NOT_FOUND                   = 1025
	ERROR_CODE_BIND_INFO_NULL                     = 1026
	ERROR_CODE_USER_ID_NULL                       = 1027
	ERROR_CODE_DEVICE_ID_NULL                     = 1028
	ERROR_CODE_DEVICE_BIND_BY_OTHER_USER          = 1029
	ERROR_CODE_USER_BIND_WITH_OTHER_DEVICE        = 1030
	ERROR_CODE_USER_ACCOUNT_PROPS_NOT_FOUND       = 1031
	ERROR_CODE_USER_BIND_INFO_NOT_FOUND           = 1032
	ERROR_CODE_USER_INFO_NULL                     = 1033
	ERROR_CODE_THIRD_PARTY_INFO_NULL              = 1034
	ERROR_CODE_CAPTCHA_ERROR                      = 1035
	ERROR_CODE_USER_FACE_DB_NOT_FOUND             = 1036
	ERROR_CODE_UPLOAD_URL_NULL                    = 1037
	ERROR_CODE_FACE_DB_TYPE_NULL                  = 1038
	ERROR_CODE_ACTION_GROUP_NOT_FOUND             = 1039
	ERROR_CODE_CLOTHES_NOT_FOUND                  = 1040
	ERROR_CODE_TOTAL_ACTION_DATA_ERROR            = 1041
	ERROR_CODE_CM_ACTION_GROUP_NOT_FOUND          = 1042
	ERROR_CODE_CM_CLOTHES_NOT_FOUND               = 1043
	ERROR_CODE_CM_SOUND_NOT_FOUND                 = 1044
	ERROR_CODE_CM_SKILL_TYPE_NOT_FOUND            = 1045
	ERROR_CODE_NO_BIND_DEVICE                     = 1046
	ERROR_CODE_SMS_SERVER_ERROR                   = 1047
	ERROR_CODE_CM_MUSIC_TAG_NOT_FOUND             = 1048
	ERROR_CODE_CM_VIDEO_TAG_NOT_FOUND             = 1049
	ERROR_CODE_SOUND_NOT_FOUND_BY_TYPE            = 1050
	ERROR_CODE_USER_VIP_NOT_FOUND                 = 1051
	ERROR_CODE_JD_OAUTH_CALLBACK_ERROR            = 1052
	ERROR_CODE_CM_CLOTHES_TYPE_NOT_FOUND          = 1053
	ERROR_CODE_USER_WARDROBE_NOT_FOUND            = 1054
	ERROR_CODE_CM_CLOTHES_SHOW_TYPE_ERROR         = 1055
	ERROR_CODE_CM_PROP_NOT_FOUND                  = 1056
	ERROR_CODE_PROP_NOT_FOUND_BY_TYPE             = 1057
	ERROR_CODE_PROP_NOT_FOUND                     = 1058
	ERROR_CODE_ACTION24_PRE_NOT_FOUND             = 1059
	ERROR_CODE_ACTION24_TIME_ERROR                = 1060
	ERROR_CODE_NOT_IN_ACTION24_ERROR              = 1061
	ERROR_CODE_DEVICE_IS_BANDED_BY_USER           = 1062
	ERROR_CODE_DEVICE_NOT_BANDED_BY_USER          = 1063
	ERROR_CODE_GET_USER_VIP_ERROR                 = 1064
	ERROR_CODE_USER_INFO_NOT_FOUND                = 1066
	ERROR_CODE_CM_SOUND_TYPES_NOT_FOUND           = 1067
	ERROR_CODE_CM_SEMANTICS_TYPE_NOT_FOUND        = 1068
	ERROR_CODE_LATEST_SERVICE_NOT_FOUND           = 1069
	ERROR_CODE_DEVICE_LOCATION_IS_NULL            = 1070
	ERROR_CODE_DEVICE_LOCATION_NOT_FOUND          = 1071
	ERROR_CODE_CM_DANCE_TYPE_NOT_FOUND            = 1072
	ERROR_CODE_CM_DANCE_CONFIG_NOT_FOUND          = 1073
	ERROR_CODE_PARAMETER_ERROR                    = 1074
	ERROR_CODE_REDIS_HGET_ERROR                   = 1075
	ERROR_CODE_USER_FACE_IS_REGISTERED            = 1076
	ERROR_CODE_REPORT_SERVICE_NAME_ERROR          = 1077
	ERROR_CODE_DAILY_ACTION_IS_EXPIRED_ERROR      = 1078
	ERROR_CODE_JD_REGISTER_DEVICE_ERROR           = 1088
	ERROR_CODE_LATEST_ACTION_NULL                 = 1089
	ERROR_CODE_GET_BIND_DEVICE_ID_FAILED          = 1099
	ERROR_CODE_QINIU_CONFIG_IS_EMPTY              = 1100
	ERROR_CODE_UPLOAD_FILE_CONTENT_IS_EMPTY       = 1101
	ERROR_CODE_NO_INDEX_FOUND                     = 1102
	ERROR_CODE_INVALID_INDEX                      = 1103
	ERROR_CODE_CAN_NOT_DO_ACTION                  = 1104
	ERROR_CODE_USER_CURRENT_CLOTHES_NULL          = 1105
	ERROR_CODE_BLUETOOTH_STATE_IS_NOT_VALID       = 1106
	ERROR_CODE_USER_NOT_HAVE_THIS_CLOTHES         = 1107
	ERROR_CODE_USER_NOT_HAVE_THIS_KIND_OF_CLOTHES = 1108
	ERROR_CODE_USER_NOT_HAVE_THIS_TAG_OF_CLOTHES  = 1109
	ERROR_CODE_USER_NOT_HAVE_CLOTHES              = 1110
	ERROR_CODE_USER_IS_WEARING_THIS_CLOTHES       = 1111
	ERROR_CODE_ACTION_RECORDS_NOT_FOUND           = 1112
	ERROR_CODE_PROMO_CODE_EXPIRE                  = 1113
	ERROR_CODE_PROMO_CODE_IS_USED                 = 1114
	ERROR_CODE_PROMO_CODE_IS_NOT_EXIST            = 1115
	ERROR_CODE_CM_EFFECT_NOT_FOUND                = 1116
	ERROR_CODE_EFFECT_NOT_FOUND                   = 1117
	ERROR_CODE_EFFECT_NEED_HIGHER_VIP             = 1118
	ERROR_CODE_SKILL_NEED_HIGHER_VIP              = 1119
	ERROR_CODE_NO_SOUND_WITH_USER_VIP             = 1120
	ERROR_CODE_TRANSMIT_JD_OAUTH_CALLBACK_ERROR   = 1121
	ERROR_CODE_CM_ACTION_UNABLE_NOT_FOUND         = 1122
	ERROR_CODE_ACTION_CLOTHES_TYPE_ERROR          = 1123
	ERROR_CODE_ACTION_SOUND_TYPE_ERROR            = 1124
	ERROR_CODE_CM_PROP_TYPE_NOT_FOUND             = 1125
	ERROR_CODE_ACTION_PROP_TYPE_ERROR             = 1126
	ERROR_CODE_IDENTIFIER_NULL                    = 1127
	ERROR_CODE_FACTORY_NULL                       = 1128
	ERROR_CODE_CM_SN_FACTORY_NOT_FOUND            = 1129
	ERROR_CODE_CM_SN_PRODUCT_NOT_FOUND            = 1130
	ERROR_CODE_CM_SN_MODEL_NOT_FOUND              = 1131
	ERROR_CODE_CM_SN_COLOR_NOT_FOUND              = 1132
	ERROR_CODE_PRODUCT_NULL                       = 1133
	ERROR_CODE_MODEL_NULL                         = 1134
	ERROR_CODE_COLOR_NULL                         = 1135
	ERROR_CODE_CAN_NOT_PARSE_SN_TO_INT            = 1136
	ERROR_CODE_NO_AVAILABLE_SN                    = 1137
	ERROR_CODE_SN_EXEED_THE_UPPER_LIMIT           = 1138
	ERROR_CODE_NO_SN_BATCH                        = 1139
	ERROR_CODE_JD_GET_USER_INFO_ERROR             = 1140
	ERROR_CODE_PROMO_CODE_TYPE_IS_NULL            = 1141
	ERROR_CODE_PROMO_EXTEND_DAYS_IS_NULL          = 1142
	ERROR_CODE_PROMO_EFFECTIVE_TIME_IS_NULL       = 1143
	ERROR_CODE_PROMO_CODE_TYPE_LIST_NOT_EXIST     = 1144
	ERROR_CODE_PROMO_CODE_TYPE_NOT_EXIST          = 1145
	ERROR_CODE_FACTORY_EXIST                      = 1146
	ERROR_CODE_PRODUCT_EXIST                      = 1147
	ERROR_CODE_MODEL_EXIST                        = 1148
	ERROR_CODE_COLOR_EXIST                        = 1149
	ERROR_CODE_LATEST_SHOW_RECORD_NOT_FOUND       = 1150
	ERROR_CODE_ACTION_ID_NULL                     = 1151
	ERROR_CODE_UPGRADE_PARAMETER_ERROR            = 1152
	ERROR_CODE_UPGRADE_NO_PHONE_VERSION_ERROR     = 1153
	ERROR_CODE_PHONE_OS_TYPE_ERROR                = 1154
	ERROR_CODE_ACTIVATE_JD_MUSIC_ERROR            = 1155
	ERROR_CODE_CM_HOME_TAG_NOT_FOUND              = 1156
	ERROR_CODE_REFRESH_JD_TOKEN_ERROR             = 1157
	ERROR_CODE_GOWILD_CENTER_ERROR                = 1158
	ERROR_CODE_USE_USER_TAG_ERROR                 = 1159
	ERROR_CODE_USER_TAG_IS_NULL                   = 1160
	ERROR_CODE_USER_LIST_IS_NULL                  = 1161
	ERROR_CODE_UNKNOW_PUSH_PLATFORM               = 1162
	ERROR_CODE_CM_DEBUG_CMD_NOT_FOUND             = 1163
	ERROR_CODE_MACS_IS_NULL                       = 1164
	ERROR_CODE_CLIENT_TYPE_ERROR                  = 1165
	ERROR_CODE_UNKNOW_PUSH_TYPE                   = 1166
	ERROR_CODE_ORDER_NOT_FOUND                    = 1167
	ERROR_CODE_ORDER_ID_IS_NULL                   = 1168
	ERROR_CODE_ORDER_STATUS_IS_NOT_PENDING        = 1169
	ERROR_CODE_ORDER_TYPE_ERROR                   = 1170
	ERROR_CODE_ACTION24_TIME_5MIN_ERROR           = 1171
	ERROR_CODE_NOTIFY_PRODUCT_ERROR               = 1172
	ERROR_CODE_ORDER_STATUS_IS_NOT_SUCCESS        = 1173
	ERROR_CODE_PRODUCT_NOT_FOUND                  = 1174
	ERROR_CODE_GOODS_TYPE_ERROR                   = 1175
	ERROR_CODE_GOWILD_IS_NULL                     = 1176
	ERROR_CODE_VIP_PACKAGE_NOT_FOUND              = 1178
	ERROR_CODE_TRANSMIT_ALI_PAY_NOTIFY_ERROR      = 1179
	ERROR_CODE_DANCE_TYPE_IS_NULL                 = 1180
	ERROR_CODE_REQUEST_URL_IS_NULL                = 1181
	ERROR_CODE_GOWILD_NOT_TAG                     = 1182
	ERROR_CODE_CANT_FIND_JOB_BY_ID                = 1183
	ERROR_CODE_CANT_DEL_JOB_ACC                   = 1184
	ERROR_CODE_SOUND_NOT_FOUND_BY_FUNCTION        = 1185
	ERROR_CODE_CM_COIN_NOT_FOUND                  = 1186
	ERROR_CODE_COIN_COUNT_NOT_ENOUGH              = 1187
	ERROR_CODE_ACTION_CENTER_ERROR                = 1188
	ERROR_CODE_GOODS_REQUEST_TYPE_ERROR           = 1189
	ERROR_CODE_NOT_FRIST_REGISTER_FAIL            = 1284
	ERROR_CODE_SPORT_GET_USER_CONF_FAIL           = 1285
	ERROR_CODE_SPORT_DB_FAIL                      = 1286
	ERROR_CODE_SPORT_WEEKLY_LIST_NOTENOUGH        = 1287
	ERROR_CODE_SPORT_LEVEL_EMPTY                  = 1288
	ERROR_CODE_SPORT_STAGE_EMPTY                  = 1289
	ERROR_CODE_SPORT_NO_WEEKLY_RECORD             = 1290
	ERROR_CODE_SPORT_WEEKLY_RECORD_NOT_ENOUGH     = 1291
	ERROR_CODE_SPORT_IS_NOTMICRO_INT64S           = 1292
	ERROR_CODE_GOODS_ID_IS_NULL                   = 1293
	ERROR_CODE_GOODS_PRICE_AMOUNT_ERROR           = 1294
	ERROR_CODE_GOODS_SALE_TIME_ERROR              = 1295
	ERROR_CODE_GOODS_OFF_SALE_TIME_ERROR          = 1296
	ERROR_CODE_GOODS_S_SUB_TYPE_ERROR             = 1297
	ERROR_CODE_GOODS_VALIDATE_ERROR               = 1298
	ERROR_CODE_CLOTHES_PART_NOT_FOUND             = 1299
	ERROR_CODE_CAPTCHA_IS_NULL                    = 1300
	ERROR_CODE_SOUND_NOT_BOUND                    = 1301
	ERROR_REPORT_STATE_URL_LEN_NOT_ENOUGH         = 1302
	ERROR_EMPTY_GIFT_PACKAGE                      = 1303
	ERROR_DAILY_CHECK_LIMIT                       = 1304
	ERROR_KEY_VALUE_LENGTH_NOT_EQUIL              = 1305
	ERROR_CODE_CM_ACTION_TEMPLATE_NOT_FOUND       = 1306
	ERROR_CODE_CM_STUDY_NOT_FOUND                 = 1307
	ERROR_CODE_CM_WORK_NOT_FOUND                  = 1308
	ERROR_CODE_CM_STUDY_TYPE_NOT_FOUND            = 1309
	ERROR_CODE_USER_PROPERTY_NOT_FOUND            = 1310
	ERROR_CODE_CM_PROPERTY_NOT_FOUND              = 1311
	ERROR_CODE_USER_NO_THIS_PROPERTY              = 1312
	ERROR_CODE_CM_CERTIFICATE_NOT_FOUND           = 1313
	ERROR_CODE_PROPERTY_NOT_FOUND                 = 1314
	ERROR_SYS_MSG_MUST_HAVE_TYPE                  = 1315
	ERROR_CODE_CM_OCCUPATION_NOT_FOUND            = 1316
	ERROR_CODE_UNFOUND_DOING_DAILY_ACTION         = 1317
	ERROR_CODE_CM_PROPERTY_TYPE_NOT_FOUND         = 1318
	ERROR_CODE_STUDY_NOT_FOUND_BY_TYPE            = 1319
	ERROR_CODE_CM_BAG_TYPE_NOT_FOUND              = 1320
	ERROR_CODE_PROPERTY_GRADE_NOT_FOUND           = 1321
	ERROR_CODE_USER_PROPERTY_SNAPSHOT_NOT_FOUND   = 1322
	ERROR_CODE_DAILY_RECORD_NOT_FOUND             = 1323
	ERROR_CODE_NO_CAN_STUDY_SOURCE                = 1324
	ERROR_CODE_SERVICE_NAME_ERROR                 = 1325
	ERROR_CODE_CAN_FIND_SUCH_PLAY_ID              = 1326
	ERROR_CODE_CANT_FIND_CAN_DO_WORK              = 1327
	ERROR_CODE_CANT_FIND_CAN_DO_STUDY             = 1328
	ERROR_CODE_CM_BAG_TYPE_ERROR                  = 1329
	ERROR_CODE_UNLOCK_ERROR                       = 1330
	ERROR_CODE_DB_NO_RECORD                       = 1331
	ERROR_CODE_DB_UNKNOWN                         = 1332
	ERROR_CODE_CM_STUDY_SOURCE_NOT_FOUND          = 1333
	ERROR_CODE_CM_WORK_SOURCE_NOT_FOUND           = 1334
	ERROR_CODE_GOODS_LIST_NOT_FOUND               = 1335
	ERROR_CODE_CM_GOODS_NOT_FOUND                 = 1336
	ERROR_CODE_GOODS_PRICE_ERROR                  = 1337
	ERROR_CODE_CM_DEVELOP_DIRECTION_NOT_FOUND     = 1338
	ERROR_CODE_DATE_NULL                          = 1339
	ERROR_CODE_SERVICE_SWITCHING_LEN_ZERO         = 1340
	ERROR_CODE_UNDEFIND_RANK                      = 1341
	ERROR_CODE_TEMPLATE_ID_NULL                   = 1342
	ERROR_CODE_NO_DEFAULT_DAILY_ACTION            = 1343
	ERROR_CODE_SAVE_DAILY_ACTION_LESS_24          = 1344
	ERROR_CODE_DEFAULT_PROPERTY_IS_NULL           = 1345
	ERROR_CODE_GOODS_S_SUB_TYPE_IS_NULL           = 1346
	ERROR_CODE_EVE_FUND_NOT_ENOUGH                = 1347
	ERROR_CODE_LENGTH_RECEIPTS_ZERO               = 1348
	ERROR_CODE_GET_PLAY_TYPE_BY_ACTION_ID_ERROR   = 1349
	ERROR_CODE_SERVICE_DETAIL_IS_NULL             = 1350
	ERROR_CODE_SAVE_DAILY_ACTION_LIST_IS_NULL     = 1351
	ERROR_CODE_TEMPLATE_SOME_COLUMN_IS_NULL       = 1352
	ERROR_CODE_CM_CAMPAIGN_CLOTHES_NOT_FOUND      = 1353
	ERROR_CODE_FESTIVAL_ONE_CHANCE                = 1354
	ERROR_CODE_FESTIVAL_NOT_TIME                  = 1355
	ERROR_CODE_KICK_TO_SOON                       = 1356
	ERROR_CODE_DONT_HAVE_TYPE_FESTIVAL            = 1357
	ERROR_CODE_CM_FOOD_NOT_FOUND                  = 1358
	ERROR_CODE_CM_CONSUMPTION_ACITION_NOT_FOUND   = 1359
	ERROR_CODE_CM_ARRANGE_TYPE_NOT_FOUND          = 1360
	ERROR_CODE_CM_FOOD_TYPE_NOT_FOUND             = 1361
	ERROR_CODE_PROPERTY_BODY_NOT_FOUND            = 1362
	ERROR_CODE_USER_DONT_HAVE_ENOUGH_GOODS        = 1363
	ERROR_CODE_BEYOUND_PURCHASE_LIMITS            = 1364
	ERROR_CODE_USER_PROPERTY_NOT_ENOUGH           = 1365
	ERROR_CODE_CM_CONSUMPTION_TYPE_ERROR          = 1366
	ERROR_CODE_CMEI_NULL                          = 1369
	ERROR_CODE_CMEI_IS_USED                       = 1370
	ERROR_CODE_COMMIT_IS_NULL                     = 1371
	ERROR_CODE_CM_DRUG_NOT_FOUND                  = 1372
	ERROR_CODE_CM_ILLNESS_NOT_FOUND               = 1373
	ERROR_CODE_CM_PROPERTY_HEALTH_NOT_FOUND       = 1374
	ERROR_CODE_CM_ILLNESS_CONFIG_ERROR            = 1375
	ERROR_CODE_PHONE_DATE_ERROR                   = 1376
	ERROR_CODE_CM_DRUG_TYPE_NOT_FOUND             = 1377
	ERROR_CODE_CM_HEALTH_PROPOSE_NOT_FOUND        = 1378
	ERROR_CODE_HEALTH_GRADE_NOT_FOUND             = 1379
	ERROR_CODE_CM_WEIGHT_PROPOSE_NOT_FOUND        = 1380
	ERROR_CODE_ILLNESS_TYPE_NOT_MIN_ILLNESS       = 1381
	ERROR_CODE_CM_IDOL_TYPE_NOT_FOUND             = 1382
	ERROR_CODE_IDOL_NULL                          = 1383
	ERROR_CODE_USER_IDOL_NULL                     = 1384
	ERROR_CODE_USER_FUND_NOT_FOUND                = 1385
	ERROR_CODE_CM_MATERIAL_NOT_FOUND              = 1386
	ERROR_CODE_CM_WEAPON_FORM_NOT_FOUND           = 1387
	ERROR_CODE_WEAPON_FORM_NOT_FOUND              = 1388
	ERROR_CODE_WEAPON_FORM_NOT_AVAILABLE          = 1389
	ERROR_CODE_NO_STUDY_MEET_REQUIREMENT          = 1390
	ERROR_CODE_CM_THEATRE_NOT_FOUND               = 1391
	ERROR_CODE_THEATRE_NOT_AVAILABLE              = 1392
	ERROR_CODE_IDOL_ERROR                         = 1393
	ERROR_CODE_INVALID_VITALITY                   = 1394
	ERROR_CODE_VATALITY_USE_LIMIT                 = 1395 //活力使用上限
	ERROR_CODE_USER_HAS_NO_DRUG                   = 1396
	ERROR_CODE_USER_HAS_NO_STUDY                  = 1397
	ERROR_CODE_IOS_PRODUCT_ID_ERROR               = 1398
	ERROR_CODE_GOODS_TYPE_NOT_FOUND               = 1399
	ERROR_CODE_CM_PORTRAIT_NOT_FOUND              = 1601
	ERROR_CODE_AVATAR_GROUP_ID_ERROR              = 1602
	ERROR_CODE_USER_CAN_NOT_CHANGE_COMBATS        = 1603
	ERROR_CODE_CM_CONSUMPTION_ACITION_NULL        = 1604
	ERROR_CODE_IDOL_NO_MATCH_MODEL                = 1605
	ERROR_CODE_USER_NAME_SENSITIVE                = 1606
	ERROR_CODE_MOVE_NOT_AVAILABLE                 = 1607
	ERROR_CODE_CM_TRAINING_NOT_FOUND              = 1608
	ERROR_CODE_TRAINING_ID_IS_NULL                = 1609
	ERROR_CODE_TRAINING_NOT_FOUND                 = 1610
	ERROR_CODE_TRAINING_ACTION_ID_NULL            = 1611
	ERROR_CODE_EFFECT_ERROR                       = 1612
	ERROR_CODE_THEATRE_BY_CLOTHES_ERROR           = 1613

	ERROR_DONT_FIND_GOWILD_ID_OR_ACCOUNT = 1400 // unify专用 1400-1499
	ERROR_COIN_COUNT_NOT_ENOUGH          = 1401 // unify专用 1400-1499

	ERROR_CODE_KAFKA_PRODUCE_NIL = 1500 //kafka相关
	ERROR_CODE_KAFKA_TOPIC_EMPTY = 1501 //topic为空
	ERROR_CODE_KAFKA_PUSH_FAIL   = 1502 //push失败
)

var ErrorCodeMessageMap = map[int]string{
	ERROR_CODE_OTHER_SERVER_ERROR:                 "其他服务出错",
	ERROR_CODE_PARAMETER_FORMAT_INVALID:           "参数格式不对",
	ERROR_CODE_ACCESS_TOKEN_TIMEOUT:               "access token失效",
	ERROR_CODE_NOT_FOUND:                          "资源没有找到",
	ERROR_INTERNAL_SERVER_ERROR:                   "参数业务验证失败",
	ERROR_CODE_ACCESS_TOKEN_ERROR:                 "access token错误",
	ERROR_CODE_MOBILE_IS_REGISTERED:               "手机已经登录",
	ERROR_CODE_MOBILE_IS_NOT_EXISTED:              "手机不存在",
	ERROR_CODE_CAPTCHA_TYPE_ERROR:                 "验证码类型错误",
	ERROR_CODE_USER_NOT_FOUND:                     "表中不存在用户",
	ERROR_CODE_MOBILE_IS_INVALID:                  "手机号码无效",
	ERROR_CODE_REQUEST_SMS_OVERLOAD:               "申请验证码超出次数",
	ERROR_CODE_CAPTCHA_TOKEN_ERROR:                "验证码token错误",
	ERROR_CODE_CM_ACTION_NOT_FOUND:                "24小时动作指令为空",
	ERROR_CODE_AUTO_LOGIN_FAILED:                  "自动登录失败",
	ERROR_CODE_INVALID_MOBILE:                     "无效手机",
	ERROR_CODE_LOGIN_FAILED:                       "登录失败",
	ERROR_CODE_CM_ACTION_TIME_NOT_FOUND:           "动作时间为空",
	ERROR_CODE_REPORT_ACTION_TYPE_ERROR:           "上报动作指令请求类型错误，1：提前一天生成的总的数据，2：当前完成的某个动作数据",
	ERROR_CODE_REDIS_KEY_NULL:                     "redis: key值为空",
	ERROR_CODE_REDIS_VALUE_NULL_PTR:               "redis: value值为空指针",
	ERROR_CODE_REDIS_VALUE_NULL:                   "redis: value值为空",
	ERROR_CODE_REDIS_KEY_NOT_EXIST:                "redis: key不存在或过期",
	ERROR_CODE_REPORT_ACTION_DATA_ERROR:           "上报动作指令数据错误",
	ERROR_CODE_SN_NULL:                            "sn为空",
	ERROR_CODE_MAC_NULL:                           "mac为空",
	ERROR_CODE_FIRMWORD_TYPE_NULL:                 "固件类型为空",
	ERROR_CODE_MANUFACTURE_FEATURE_ERROR:          "服务端获取工厂验证信息失败",
	ERROR_CODE_MANUFACTURE_AUTH_ERROR:             "工厂验证信息错误",
	ERROR_CODE_DEVICE_NOT_FOUND:                   "设备不存在",
	ERROR_CODE_BIND_INFO_NULL:                     "绑定信息为空",
	ERROR_CODE_USER_ID_NULL:                       "用户id为空",
	ERROR_CODE_DEVICE_ID_NULL:                     "设备id为空",
	ERROR_CODE_DEVICE_BIND_BY_OTHER_USER:          "设备已被其他用户绑定",
	ERROR_CODE_USER_BIND_WITH_OTHER_DEVICE:        "用户已和其他设备绑定",
	ERROR_CODE_USER_ACCOUNT_PROPS_NOT_FOUND:       "账户机器信息不存在",
	ERROR_CODE_USER_BIND_INFO_NOT_FOUND:           "账户绑定信息不存在",
	ERROR_CODE_USER_INFO_NULL:                     "用户信息为空",
	ERROR_CODE_THIRD_PARTY_INFO_NULL:              "第三方验证信息不存在",
	ERROR_CODE_CAPTCHA_ERROR:                      "验证码错误",
	ERROR_CODE_USER_FACE_DB_NOT_FOUND:             "用户人脸数据不存在",
	ERROR_CODE_UPLOAD_URL_NULL:                    "七牛url为空",
	ERROR_CODE_FACE_DB_TYPE_NULL:                  "人脸数据类型参数为空",
	ERROR_CODE_ACTION_GROUP_NOT_FOUND:             "动作组不存在",
	ERROR_CODE_CLOTHES_NOT_FOUND:                  "服装不存在",
	ERROR_CODE_TOTAL_ACTION_DATA_ERROR:            "24小时完整时间轴行为数据错误",
	ERROR_CODE_CM_ACTION_GROUP_NOT_FOUND:          "动作组为空",
	ERROR_CODE_CM_CLOTHES_NOT_FOUND:               "服装为空",
	ERROR_CODE_CM_SOUND_NOT_FOUND:                 "曲目为空",
	ERROR_CODE_CM_SKILL_TYPE_NOT_FOUND:            "技能类型表为空",
	ERROR_CODE_NO_BIND_DEVICE:                     "没有已绑定的设备",
	ERROR_CODE_SMS_SERVER_ERROR:                   "短信验证码服务器报错",
	ERROR_CODE_CM_MUSIC_TAG_NOT_FOUND:             "音乐标签表为空",
	ERROR_CODE_CM_VIDEO_TAG_NOT_FOUND:             "影视标签表为空",
	ERROR_CODE_SOUND_NOT_FOUND_BY_TYPE:            "该类型音效为空",
	ERROR_CODE_USER_VIP_NOT_FOUND:                 "用户vip信息为空",
	ERROR_CODE_JD_OAUTH_CALLBACK_ERROR:            "京东oauth callback错误",
	ERROR_CODE_CM_CLOTHES_TYPE_NOT_FOUND:          "服装类型表为空",
	ERROR_CODE_USER_WARDROBE_NOT_FOUND:            "用户衣橱为空",
	ERROR_CODE_CM_CLOTHES_SHOW_TYPE_ERROR:         "服装表展示类型为空",
	ERROR_CODE_CM_PROP_NOT_FOUND:                  "道具表为空",
	ERROR_CODE_PROP_NOT_FOUND_BY_TYPE:             "道具类型表为空",
	ERROR_CODE_PROP_NOT_FOUND:                     "道具不存在",
	ERROR_CODE_ACTION24_PRE_NOT_FOUND:             "24小时动作轴数据为空",
	ERROR_CODE_ACTION24_TIME_ERROR:                "24小时动作轴时间不对",
	ERROR_CODE_NOT_IN_ACTION24_ERROR:              "生活动作不在上报24小时动作内",
	ERROR_CODE_DEVICE_IS_BANDED_BY_USER:           "设备已经绑定了该用户",
	ERROR_CODE_DEVICE_NOT_BANDED_BY_USER:          "设备没有绑定用户",
	ERROR_CODE_GET_USER_VIP_ERROR:                 "获取用户会员信息失败",
	ERROR_CODE_USER_INFO_NOT_FOUND:                "用户信息不存在",
	ERROR_CODE_CM_SOUND_TYPES_NOT_FOUND:           "才艺类型列表为空",
	ERROR_CODE_CM_SEMANTICS_TYPE_NOT_FOUND:        "语义才艺分类为空",
	ERROR_CODE_LATEST_SERVICE_NOT_FOUND:           "当前服务为空",
	ERROR_CODE_DEVICE_LOCATION_IS_NULL:            "上报设备位置为空",
	ERROR_CODE_DEVICE_LOCATION_NOT_FOUND:          "设备位置信息为空",
	ERROR_CODE_CM_DANCE_TYPE_NOT_FOUND:            "舞蹈类型为空",
	ERROR_CODE_CM_DANCE_CONFIG_NOT_FOUND:          "舞蹈配置为空",
	ERROR_CODE_PARAMETER_ERROR:                    "请求参数有错误",
	ERROR_CODE_REDIS_HGET_ERROR:                   "redis: hgetall返回列表不成对",
	ERROR_CODE_USER_FACE_IS_REGISTERED:            "用户人脸已经注册",
	ERROR_CODE_REPORT_SERVICE_NAME_ERROR:          "上报服务名称错误，跟数据表已有记录不一致",
	ERROR_CODE_DAILY_ACTION_IS_EXPIRED_ERROR:      "数据库24小时动作数据过期",
	ERROR_CODE_JD_REGISTER_DEVICE_ERROR:           "京东注册设备接口错误",
	ERROR_CODE_LATEST_ACTION_NULL:                 "用户当前动作为空",
	ERROR_CODE_GET_BIND_DEVICE_ID_FAILED:          "获取绑定设备id失败",
	ERROR_CODE_QINIU_CONFIG_IS_EMPTY:              "七牛的配置AccessKey或SecretKey为空",
	ERROR_CODE_UPLOAD_FILE_CONTENT_IS_EMPTY:       "七牛上传文件为空",
	ERROR_CODE_NO_INDEX_FOUND:                     "用户id生成器值为空",
	ERROR_CODE_INVALID_INDEX:                      "用户id生成器值无效",
	ERROR_CODE_CAN_NOT_DO_ACTION:                  "不会做这个动作",
	ERROR_CODE_USER_CURRENT_CLOTHES_NULL:          "用户当前服装为空",
	ERROR_CODE_BLUETOOTH_STATE_IS_NOT_VALID:       "蓝牙状态无效",
	ERROR_CODE_USER_NOT_HAVE_THIS_CLOTHES:         "用户没有这件衣服",
	ERROR_CODE_USER_NOT_HAVE_THIS_KIND_OF_CLOTHES: "用户没有这种类型的衣服",
	ERROR_CODE_USER_NOT_HAVE_THIS_TAG_OF_CLOTHES:  "用户没有这种标签的衣服",
	ERROR_CODE_USER_NOT_HAVE_CLOTHES:              "用户没有可穿的衣服",
	ERROR_CODE_USER_IS_WEARING_THIS_CLOTHES:       "用户正穿着这件衣服",
	ERROR_CODE_ACTION_RECORDS_NOT_FOUND:           "最近动作记录为空",
	ERROR_CODE_PROMO_CODE_EXPIRE:                  "活动码过期",
	ERROR_CODE_PROMO_CODE_IS_USED:                 "活动码已经被使用",
	ERROR_CODE_PROMO_CODE_IS_NOT_EXIST:            "活动码不存在",
	ERROR_CODE_CM_EFFECT_NOT_FOUND:                "特效表为空",
	ERROR_CODE_EFFECT_NOT_FOUND:                   "特效不存在",
	ERROR_CODE_EFFECT_NEED_HIGHER_VIP:             "特效需要更高的vip级别",
	ERROR_CODE_SKILL_NEED_HIGHER_VIP:              "技能需要更高的vip级别",
	ERROR_CODE_NO_SOUND_WITH_USER_VIP:             "没有符合用户vip级别的技能",
	ERROR_CODE_TRANSMIT_JD_OAUTH_CALLBACK_ERROR:   "转发京东oauth callback错误",
	ERROR_CODE_CM_ACTION_UNABLE_NOT_FOUND:         "不会行为表为空",
	ERROR_CODE_ACTION_CLOTHES_TYPE_ERROR:          "动作组衣服类型错误",
	ERROR_CODE_ACTION_SOUND_TYPE_ERROR:            "动作组音效类型错误",
	ERROR_CODE_CM_PROP_TYPE_NOT_FOUND:             "道具类型表为空",
	ERROR_CODE_ACTION_PROP_TYPE_ERROR:             "动作组道具类型错误",
	ERROR_CODE_IDENTIFIER_NULL:                    "标识码为空",
	ERROR_CODE_FACTORY_NULL:                       "工厂名称为空",
	ERROR_CODE_CM_SN_FACTORY_NOT_FOUND:            "工厂列表为空",
	ERROR_CODE_CM_SN_PRODUCT_NOT_FOUND:            "产品列表为空",
	ERROR_CODE_CM_SN_MODEL_NOT_FOUND:              "机型列表为空",
	ERROR_CODE_CM_SN_COLOR_NOT_FOUND:              "颜色列表为空",
	ERROR_CODE_PRODUCT_NULL:                       "产品名称为空",
	ERROR_CODE_MODEL_NULL:                         "机型为空",
	ERROR_CODE_COLOR_NULL:                         "颜色为空",
	ERROR_CODE_CAN_NOT_PARSE_SN_TO_INT:            "不能将sn后五位转换成整形",
	ERROR_CODE_NO_AVAILABLE_SN:                    "没有可用的sn，请先生成足够sn",
	ERROR_CODE_SN_EXEED_THE_UPPER_LIMIT:           "SN总数超过了月上限99999",
	ERROR_CODE_NO_SN_BATCH:                        "没有SN批次信息，请先生成批次",
	ERROR_CODE_JD_GET_USER_INFO_ERROR:             "京东获取用户信息接口错误",
	ERROR_CODE_PROMO_CODE_TYPE_IS_NULL:            "活动类型为空",
	ERROR_CODE_PROMO_EXTEND_DAYS_IS_NULL:          "活动延长vip天数为空",
	ERROR_CODE_PROMO_EFFECTIVE_TIME_IS_NULL:       "活动有效期为空",
	ERROR_CODE_PROMO_CODE_TYPE_LIST_NOT_EXIST:     "活动类型列表不存在",
	ERROR_CODE_PROMO_CODE_TYPE_NOT_EXIST:          "活动码类型不存在",
	ERROR_CODE_FACTORY_EXIST:                      "工厂已经存在",
	ERROR_CODE_PRODUCT_EXIST:                      "产品已经存在",
	ERROR_CODE_MODEL_EXIST:                        "机型已经存在",
	ERROR_CODE_COLOR_EXIST:                        "颜色已经存在",
	ERROR_CODE_LATEST_SHOW_RECORD_NOT_FOUND:       "上一次的才艺展示记录不存在",
	ERROR_CODE_ACTION_ID_NULL:                     "动作id为空",
	ERROR_CODE_UPGRADE_PARAMETER_ERROR:            "检查升级参数错误",
	ERROR_CODE_UPGRADE_NO_PHONE_VERSION_ERROR:     "没有手机操作系统及版本信息",
	ERROR_CODE_PHONE_OS_TYPE_ERROR:                "手机系统类型错误",
	ERROR_CODE_ACTIVATE_JD_MUSIC_ERROR:            "激活京东音乐失败",
	ERROR_CODE_CM_HOME_TAG_NOT_FOUND:              "家庭个性化标签表为空",
	ERROR_CODE_REFRESH_JD_TOKEN_ERROR:             "更新京东token失败",
	ERROR_CODE_GOWILD_CENTER_ERROR:                "统一账号中心错误",
	ERROR_CODE_USE_USER_TAG_ERROR:                 "是否使用用户标签错误",
	ERROR_CODE_USER_TAG_IS_NULL:                   "用户标签为空",
	ERROR_CODE_USER_LIST_IS_NULL:                  "用户列表为空",
	ERROR_CODE_UNKNOW_PUSH_PLATFORM:               "不识别的推送平台",
	ERROR_CODE_CM_DEBUG_CMD_NOT_FOUND:             "调试命令不存在",
	ERROR_CODE_MACS_IS_NULL:                       "白名单列表为空",
	ERROR_CODE_CLIENT_TYPE_ERROR:                  "客户端类型错误，1=device、2=ios、android",
	ERROR_CODE_UNKNOW_PUSH_TYPE:                   "不识别的推送类型",
	ERROR_CODE_ORDER_NOT_FOUND:                    "订单没有找到",
	ERROR_CODE_ORDER_ID_IS_NULL:                   "订单ID为空",
	ERROR_CODE_ORDER_STATUS_IS_NOT_PENDING:        "订单状态没有就绪",
	ERROR_CODE_ORDER_TYPE_ERROR:                   "订单类型错误",
	ERROR_CODE_ACTION24_TIME_5MIN_ERROR:           "离24小时动作轴开始时间还差5分钟",
	ERROR_CODE_NOTIFY_PRODUCT_ERROR:               "支付结果通知产品业务后端失败",
	ERROR_CODE_ORDER_STATUS_IS_NOT_SUCCESS:        "订单还没有支付",
	ERROR_CODE_PRODUCT_NOT_FOUND:                  "商品没有找到",
	ERROR_CODE_GOODS_TYPE_ERROR:                   "商品类型错误",
	ERROR_CODE_GOWILD_IS_NULL:                     "Gowild Id 为空",
	ERROR_CODE_VIP_PACKAGE_NOT_FOUND:              "VIP套餐不存在",
	ERROR_CODE_TRANSMIT_ALI_PAY_NOTIFY_ERROR:      "转发阿里支付回调结果失败",
	ERROR_CODE_DANCE_TYPE_IS_NULL:                 "舞蹈类型列表为空",
	ERROR_CODE_REQUEST_URL_IS_NULL:                "推送时请求的uri不可以为空",
	ERROR_CODE_GOWILD_NOT_TAG:                     "使用gowildid时不能同时传标签",
	ERROR_CODE_CANT_FIND_JOB_BY_ID:                "无法使用JOBID找到对应的JOB",
	ERROR_CODE_CANT_DEL_JOB_ACC:                   "不可删除job，该job已执行",
	ERROR_CODE_SOUND_NOT_FOUND_BY_FUNCTION:        "根据才艺功能获取曲目失败",
	ERROR_CODE_CM_COIN_NOT_FOUND:                  "该数量的瓦歌币商品不存在",
	ERROR_CODE_COIN_COUNT_NOT_ENOUGH:              "瓦歌币数不足",
	ERROR_CODE_ACTION_CENTER_ERROR:                "指令中心错误",
	ERROR_CODE_GOODS_REQUEST_TYPE_ERROR:           "商品请求类型错误",
	ERROR_CODE_SPORT_GET_USER_CONF_FAIL:           "用户未初始化运动管理",
	ERROR_CODE_SPORT_DB_FAIL:                      "运动中心数据库查询失败",
	ERROR_CODE_SPORT_WEEKLY_LIST_NOTENOUGH:        "生成每周运动计划长度不足",
	ERROR_CODE_SPORT_LEVEL_EMPTY:                  "运动等级表为空",
	ERROR_CODE_SPORT_STAGE_EMPTY:                  "运动阶段表为空",
	ERROR_CODE_SPORT_NO_WEEKLY_RECORD:             "未生成一周运动计划",
	ERROR_CODE_SPORT_WEEKLY_RECORD_NOT_ENOUGH:     "一周运动计划不完整",
	ERROR_CODE_SPORT_IS_NOTMICRO_INT64S:           "输入时间不是标准的毫秒时间戳",
	ERROR_CODE_GOODS_ID_IS_NULL:                   "商品ID为空",
	ERROR_CODE_GOODS_PRICE_AMOUNT_ERROR:           "商品总价错误",
	ERROR_CODE_GOODS_SALE_TIME_ERROR:              "商品未上架",
	ERROR_CODE_GOODS_OFF_SALE_TIME_ERROR:          "商品已下架",
	ERROR_CODE_GOODS_S_SUB_TYPE_ERROR:             "商品二级子目录错误",
	ERROR_CODE_GOODS_VALIDATE_ERROR:               "商品基本信息错误",
	ERROR_CODE_CLOTHES_PART_NOT_FOUND:             "服装部件没有找到",
	ERROR_CODE_NOT_FRIST_REGISTER_FAIL:            "非首次开启运动管理，重复配置运动信息",
	ERROR_CODE_CAPTCHA_IS_NULL:                    "验证码为空",
	ERROR_CODE_SOUND_NOT_BOUND:                    "曲目尚未购买",
	ERROR_REPORT_STATE_URL_LEN_NOT_ENOUGH:         "上报状态时URL长度不足",
	ERROR_EMPTY_GIFT_PACKAGE:                      "礼品包为空",
	ERROR_DAILY_CHECK_LIMIT:                       "生活轴查询不可大于30天",
	ERROR_KEY_VALUE_LENGTH_NOT_EQUIL:              "键值对长短不一",
	ERROR_CODE_CM_ACTION_TEMPLATE_NOT_FOUND:       "生活轴模板为空",
	ERROR_CODE_CM_STUDY_NOT_FOUND:                 "学习列表为空",
	ERROR_CODE_CM_WORK_NOT_FOUND:                  "工作列表为空",
	ERROR_CODE_CM_STUDY_TYPE_NOT_FOUND:            "学习类型列表为空",
	ERROR_CODE_USER_PROPERTY_NOT_FOUND:            "用户属性为空",
	ERROR_CODE_CM_PROPERTY_NOT_FOUND:              "公共属性为空",
	ERROR_CODE_USER_NO_THIS_PROPERTY:              "用户没有这个属性",
	ERROR_CODE_CM_CERTIFICATE_NOT_FOUND:           "证书列表为空",
	ERROR_CODE_PROPERTY_NOT_FOUND:                 "属性不存在",
	ERROR_SYS_MSG_MUST_HAVE_TYPE:                  "系统消息必须有类型",
	ERROR_CODE_CM_OCCUPATION_NOT_FOUND:            "职业列表为空",
	ERROR_CODE_UNFOUND_DOING_DAILY_ACTION:         "未找到正在做的生活轴信息",
	ERROR_CODE_CM_PROPERTY_TYPE_NOT_FOUND:         "属性类型列表不存在",
	ERROR_CODE_STUDY_NOT_FOUND_BY_TYPE:            "该类型学习列表为空",
	ERROR_CODE_CM_BAG_TYPE_NOT_FOUND:              "背包类型列表为空",
	ERROR_CODE_PROPERTY_GRADE_NOT_FOUND:           "属性等级不存在",
	ERROR_CODE_DAILY_RECORD_NOT_FOUND:             "未找到生活轴记录",
	ERROR_CODE_NO_CAN_STUDY_SOURCE:                "没有可学习资源",
	ERROR_CODE_SERVICE_NAME_ERROR:                 "服务名称错误",
	ERROR_CODE_CAN_FIND_SUCH_PLAY_ID:              "未找到相应的玩法项目ID",
	ERROR_CODE_CANT_FIND_CAN_DO_WORK:              "未找到可以执行的工作",
	ERROR_CODE_CANT_FIND_CAN_DO_STUDY:             "未找到可以执行的玩法",
	ERROR_CODE_CM_BAG_TYPE_ERROR:                  "背包类型错误",
	ERROR_CODE_UNLOCK_ERROR:                       "解锁证书id错误",
	ERROR_CODE_DB_NO_RECORD:                       "数据库没有相关记录",
	ERROR_CODE_DB_UNKNOWN:                         "未知数据库读错误",
	ERROR_CODE_CM_STUDY_SOURCE_NOT_FOUND:          "学习资源不存在",
	ERROR_CODE_CM_WORK_SOURCE_NOT_FOUND:           "工作不存在",
	ERROR_CODE_GOODS_LIST_NOT_FOUND:               "商城列表为空",
	ERROR_CODE_CM_GOODS_NOT_FOUND:                 "商品不存在",
	ERROR_CODE_GOODS_PRICE_ERROR:                  "商品价格错误",
	ERROR_CODE_CM_DEVELOP_DIRECTION_NOT_FOUND:     "生活轴发展方向为空",
	ERROR_CODE_DATE_NULL:                          "日期为空",
	ERROR_CODE_SERVICE_SWITCHING_LEN_ZERO:         "没有满足要求的服务状态记录",
	ERROR_CODE_UNDEFIND_RANK:                      "未定义的排行榜栏目",
	ERROR_CODE_TEMPLATE_ID_NULL:                   "生活轴模板ID不存在",
	ERROR_CODE_NO_DEFAULT_DAILY_ACTION:            "未找到默认生活轴动作",
	ERROR_CODE_SAVE_DAILY_ACTION_LESS_24:          "上报的生活轴模版小于24小时",
	ERROR_CODE_DEFAULT_PROPERTY_IS_NULL:           "默认属性为空",
	ERROR_CODE_GOODS_S_SUB_TYPE_IS_NULL:           "商品二级子目录为空",
	ERROR_CODE_EVE_FUND_NOT_ENOUGH:                "琥珀账户瓦歌币数不足",
	ERROR_CODE_LENGTH_RECEIPTS_ZERO:               "苹果账单长度为0",
	ERROR_CODE_GET_PLAY_TYPE_BY_ACTION_ID_ERROR:   "根据动作id获取学习类型失败",
	ERROR_CODE_SERVICE_DETAIL_IS_NULL:             "上报生活轴中服务细节为空",
	ERROR_CODE_SAVE_DAILY_ACTION_LIST_IS_NULL:     "上报生活轴或模板列表为空",
	ERROR_CODE_TEMPLATE_SOME_COLUMN_IS_NULL:       "模板中有关键字段为空",
	ERROR_CODE_CM_CAMPAIGN_CLOTHES_NOT_FOUND:      "商城列表为空",
	ERROR_CODE_FESTIVAL_ONE_CHANCE:                "双旦活动一天只能做一次",
	ERROR_CODE_FESTIVAL_NOT_TIME:                  "双旦活动时间未到",
	ERROR_CODE_KICK_TO_SOON:                       "点击太频繁",
	ERROR_CODE_DONT_HAVE_TYPE_FESTIVAL:            "没有相应的兑换类型",
	ERROR_CODE_CM_FOOD_NOT_FOUND:                  "食物列表为空",
	ERROR_CODE_CM_CONSUMPTION_ACITION_NOT_FOUND:   "有消耗动作列表为空",
	ERROR_CODE_CM_ARRANGE_TYPE_NOT_FOUND:          "可安排类型列表为空",
	ERROR_CODE_CM_FOOD_TYPE_NOT_FOUND:             "食物类型为空",
	ERROR_CODE_PROPERTY_BODY_NOT_FOUND:            "身体属性范围表为空",
	ERROR_CODE_USER_DONT_HAVE_ENOUGH_GOODS:        "用户所拥有的消耗品不足",
	ERROR_CODE_BEYOUND_PURCHASE_LIMITS:            "购买商品超过上限",
	ERROR_CODE_USER_PROPERTY_NOT_ENOUGH:           "用户所拥有的属性不足",
	ERROR_CODE_CM_CONSUMPTION_TYPE_ERROR:          "有消耗动作类型错误",
	ERROR_CODE_CMEI_NULL:                          "cmei码为空",
	ERROR_CODE_CMEI_IS_USED:                       "cmei码已经被使用",
	ERROR_CODE_COMMIT_IS_NULL:                     "必须填写commit",
	ERROR_CODE_CM_DRUG_NOT_FOUND:                  "药品为空",
	ERROR_CODE_CM_ILLNESS_NOT_FOUND:               "疾病为空",
	ERROR_CODE_CM_PROPERTY_HEALTH_NOT_FOUND:       "属性值对健康影响规则表为空",
	ERROR_CODE_CM_ILLNESS_CONFIG_ERROR:            "疾病表配置错误",
	ERROR_CODE_PHONE_DATE_ERROR:                   "手机日期错误",
	ERROR_CODE_CM_DRUG_TYPE_NOT_FOUND:             "药品类型为空",
	ERROR_CODE_CM_HEALTH_PROPOSE_NOT_FOUND:        "健康建议为空",
	ERROR_CODE_HEALTH_GRADE_NOT_FOUND:             "健康等级为空",
	ERROR_CODE_CM_WEIGHT_PROPOSE_NOT_FOUND:        "体重建议为空",
	ERROR_CODE_ILLNESS_TYPE_NOT_MIN_ILLNESS:       "病类型不是小病",
	ERROR_CODE_CM_IDOL_TYPE_NOT_FOUND:             "偶像IP不存在",
	ERROR_CODE_IDOL_NULL:                          "偶像IP为空",
	ERROR_CODE_USER_IDOL_NULL:                     "用户偶像IP为空",
	ERROR_CODE_USER_FUND_NOT_FOUND:                "用户资金账号为空",
	ERROR_CODE_CM_MATERIAL_NOT_FOUND:              "材料列表为空",
	ERROR_CODE_CM_WEAPON_FORM_NOT_FOUND:           "武器形态为空",
	ERROR_CODE_WEAPON_FORM_NOT_FOUND:              "武器形态不存在",
	ERROR_CODE_WEAPON_FORM_NOT_AVAILABLE:          "武器形态不可用",
	ERROR_CODE_NO_STUDY_MEET_REQUIREMENT:          "没有符合要求的学习资源",
	ERROR_CODE_CM_THEATRE_NOT_FOUND:               "特典或招式为空",
	ERROR_CODE_THEATRE_NOT_AVAILABLE:              "特典不可用",
	ERROR_CODE_IDOL_ERROR:                         "Idol不支持该接口",
	ERROR_CODE_INVALID_VITALITY:                   "活力不存在",
	ERROR_CODE_VATALITY_USE_LIMIT:                 "活力今日使用次数不足",
	ERROR_CODE_USER_HAS_NO_DRUG:                   "用户没有这种药",
	ERROR_CODE_USER_HAS_NO_STUDY:                  "用户没有这种学习资源",
	ERROR_CODE_IOS_PRODUCT_ID_ERROR:               "IOS内购续费票据中的产品id错误",
	ERROR_CODE_GOODS_TYPE_NOT_FOUND:               "商品类型列表为空",
	ERROR_CODE_CM_PORTRAIT_NOT_FOUND:              "用户头像表为空",
	ERROR_CODE_AVATAR_GROUP_ID_ERROR:              "用户头像组id错误",
	ERROR_CODE_USER_CAN_NOT_CHANGE_COMBATS:        "叶修不能指令换战斗服",
	ERROR_CODE_CM_CONSUMPTION_ACITION_NULL:        "消耗动作没有找到",
	ERROR_CODE_IDOL_NO_MATCH_MODEL:                "idol和机器型号不匹配",
	ERROR_CODE_USER_NAME_SENSITIVE:                "用户名称包含敏感字",
	ERROR_CODE_MOVE_NOT_AVAILABLE:                 "招式不可用",
	ERROR_CODE_CM_TRAINING_NOT_FOUND:              "训练表为空",
	ERROR_CODE_TRAINING_ID_IS_NULL:                "训练id为空",
	ERROR_CODE_TRAINING_NOT_FOUND:                 "训练id没有找到",
	ERROR_CODE_TRAINING_ACTION_ID_NULL:            "训练动作id为空",
	ERROR_CODE_EFFECT_ERROR:                       "影响数据包配置错误",
	ERROR_CODE_THEATRE_BY_CLOTHES_ERROR:           "服装赠送招式错误",


	ERROR_DONT_FIND_GOWILD_ID_OR_ACCOUNT: "未找到对应的账号或GOWILDID",
	ERROR_COIN_COUNT_NOT_ENOUGH:          "余额不足",

	ERROR_CODE_KAFKA_PRODUCE_NIL: "producer对象为空",
	ERROR_CODE_KAFKA_TOPIC_EMPTY: "topic为空",
	ERROR_CODE_KAFKA_PUSH_FAIL:   "推送失败",
}

type HError struct {
	ResCode int    `json:"code" description:"0成功 >1失败"`
	Message string `json:"message" description:"提示消息"`
}

func (ze HError) Error() string {
	return fmt.Sprintf("%s%s", HERROR_PREFIX, ze.Message)
}

func (ze HError) Code() int {
	return ze.ResCode
}

//系统错误返回数据结构
func NewHError(code int, message string) error {
	return &HError{code, message}
}

//自定义错误返回数据结构
func NewHErrorCustom(code int) error {
	return &HError{code, ErrorCodeMessageMap[code]}
}

func GetErrorCode(err error) int {
	var errCode int
	switch errType := err.(type) {
	case *HError:
		errCode = errType.ResCode
	default:
		errCode = 1
	}
	return errCode
}
