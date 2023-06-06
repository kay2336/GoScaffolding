package utils

import (
	"errors"
	"golang.org/x/net/context"
)

// 使用key代替int的原因是：避免冲突（见源码注释）
type key int

var userKey key

type UserInfo struct {
	Id uint `json:"id"`
}

// GetUserInfoFromCtx
// 从上下文中获取用户信息
func GetUserInfoFromCtx(ctx context.Context) (*UserInfo, error) {
	userInfo, ok := FromContext(ctx)
	if !ok {
		return nil, errors.New("获取用户信息错误")
	}

	//log.Println(userInfo.Id)

	return userInfo, nil
}

func NewContext(ctx context.Context, u *UserInfo) context.Context {
	return context.WithValue(ctx, userKey, u)
}

func FromContext(ctx context.Context) (*UserInfo, bool) {
	u, ok := ctx.Value(userKey).(*UserInfo)
	return u, ok
}
