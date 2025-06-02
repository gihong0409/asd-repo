package utils

import (
	"ASD/dmrsapi"
	"fmt"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats/ktformats"
	"git.datau.co.kr/ferrari/ferrari-common/tcrsapi/tcrsformats/sktformats"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

// extract birth year
func ExtBD(data map[string]interface{}, tel int) int {

	print(tel, ": ")
	switch tel {

	//SKT
	case 0:

		body, ok := data["BodyInfo"].(sktformats.UserInfoRsp)

		if !ok {
			logrus.Error("SKT: error: 응답 바디 에러")
			return -1
		}
		if len(body.SSN_BIRTH_DT) < 8 {
			logrus.Error("SKT: error: SSN_BIRTH_DT ")

			return -1
		}
		bd := body.SSN_BIRTH_DT
		bd = bd[:4]
		println(": SKT:", body.SSN_BIRTH_DT, ": 생년월일: ", bd)

		result, _ := strconv.Atoi(bd)
		return result
	//KT
	case 1:
		body, ok := data["Body"].(ktformats.RSPUserInfoAndKways)
		if !ok {
			logrus.Error("KT: error: 응답 바디 에러")

			return -1
		}
		// USER_SSN_FRONT의 길이가 충분한지 확인
		if len(body.USER_SSN_FRONT) < 2 {
			logrus.Error("KT: error: USER_SSN_FRONT too short")

			return -1
		}

		bd := body.USER_SSN_FRONT[:2]
		bd = ktAgeConverter(bd)

		println(": KT: ", body.USER_SSN_FRONT, ": 생년: ", bd)

		result, _ := strconv.Atoi(bd)
		return result
		//
	case 2:
		body, ok := data["Body"].(dmrsapi.LGUPRSPUserInfo)

		if !ok {
			logrus.Error("lGUP: error: 응답 바디 에러")

			return -1
		}
		bd := lgupAgeToBirthYear(body.Age)

		println(": LGUP: ", body.Age, ": 생년월일: ", bd)

		result, _ := strconv.Atoi(bd)
		return result
	}
	logrus.Error("extBD error")

	return -1

}

// AgeToBirthYear는 나이를 받아 생년(YYYY)으로 반환
func lgupAgeToBirthYear(ageStr string) string {
	// 나이를 숫자로 변환
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return fmt.Sprintf("invalid age format: %v", err)
	}

	// 나이 유효성 검사
	if age < 0 || age > 120 {
		return "invalid age range"
	}

	// 현재 연도 가져오기
	currentYear := time.Now().Year() // 2025

	// 생년 계산
	birthYear := currentYear - age

	// 생년 반환
	return fmt.Sprintf("%d", birthYear)
}

// 두 자리 연도를 4자리 연도로 변환
func ktAgeConverter(twoDigitYear string) string {
	// 입력 길이 확인
	if len(twoDigitYear) != 2 {
		return "error: 입력이 두자리가 아닙니다."
	}

	// 두 자리 연도를 숫자로 변환
	year, err := strconv.Atoi(twoDigitYear)
	if err != nil {
		return "error: 유효한 숫자가 아닙니다."
	}

	// 연도 변환 로직
	var fourDigitYear int
	if year >= 0 && year <= 25 {
		fourDigitYear = 2000 + year
	} else if year >= 24 && year <= 99 {
		fourDigitYear = 1900 + year
	} else {
		return "error: 유효하지 않은 연도 범위입니다."
	}

	// 네 자리 연도를 문자열로 반환
	return fmt.Sprintf("%04d", fourDigitYear)
}
