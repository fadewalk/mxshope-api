package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mxshop-api/user-web/proto"
	"net/http"
)

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	// 将grpc 的codei 转换成为http 的状态码
	if err == nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message,
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					//"msg": "其他错误" ,
					"msg": e.Code(),
				})
			}
			return
		}
	}

}

func GetUserList(ctx *gin.Context) {
	ip := "127.0.0.1"
	port := 50051

	// 拨号连接用户grpc 服务器
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 【用户服务失败】",
			"msg", err.Error())
	}
	// 生成grpc的 client 并调用接口
	userSrvClient := proto.NewUserClient(userConn)

	rep, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    0,
		PSize: 0,
	})
	if err != nil {
		zap.S().Errorw("[GetUserList] 查询【用户列表】失败")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	result := make([]interface{}, 0)
	for _, value := range rep.Data {
		data := make(map[string]interface{})
		data["id"] = value.Id
		data["name"] = value.NickName
		data["birthday"] = value.BirthDay
		data["gender"] = value.Gender
		data["mobile"] = value.Mobile
		result = append(result, data)
	}
	ctx.JSON(http.StatusOK, result)

	zap.S().Debug("获取用户列表页")
}
