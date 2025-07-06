package middleware

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(tokenString string) (*UserClaims, error) {
	// 1. ดึง Secret Key มาจาก Environment Variable
	// ต้องเป็น key เดียวกันกับที่ใช้ใน Next.js
	jwtSecret := []byte(os.Getenv("NEXTAUTH_SECRET"))
	if len(jwtSecret) == 0 {
		return nil, fmt.Errorf("JWT secret key is not configured")
	}

	// 2. Parse และ Verify Token ด้วย claims struct ของเรา
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 2.1 ตรวจสอบให้แน่ใจว่าอัลกอริทึมในการเข้ารหัส (signing method) เป็น HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// 2.2 คืนค่า secret key กลับไปให้ library ใช้ตรวจสอบ
		return jwtSecret, nil
	})

	// 3. จัดการกับ Error ที่เกิดจากการ Parse
	// เช่น token หมดอายุ, signature ไม่ถูกต้อง, หรือรูปแบบ token ผิด
	if err != nil {
		return nil, err
	}

	// 4. ดึงข้อมูล claims ออกมาจาก token ที่ผ่านการตรวจสอบแล้ว
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
