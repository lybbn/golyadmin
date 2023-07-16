package utils

import (
	"gitee.com/lybbn/golyadmin/utils/uuid"
)

// 生成唯一文件名（根据时间戳和随机数字）
func GenerateUniqueTimeStampFileName() string {
	return MD5([]byte(FormatInt642String(GetTimestamp()) + FormatInt2String(GenerateRandomNumsInt(1, 200))))
}

// 生成唯一文件名（根据uuid）
func GenerateUniqueUUIDFileName() string {
	return uuid.MakeUUID()
}
