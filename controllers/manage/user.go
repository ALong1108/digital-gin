package manage

//
//import (
//	"github.com/gin-gonic/gin"
//)
//
///**
//* @api {post} /user/login 用户登录
//* @apiVersion 1.0.0
//* @apiGroup User
//* @apiDescription 用户登录
//*
//* @apiParam {String} username 用户名
//* @apiParam {String} password 密码
//* @apiParam {String} [token] 登录授权码
//*
//* @apiParamExample {json} Request-Example:
//{
//	"username": "xxx",
//	"password": "df81b3f78e9bd3990ba51e28ca6d11fd"
//	"token":"sss"
//}
//* @apiSuccess {Number} err_no 错误码，大于0表示出错
//* @apiSuccess {String} err_msg 错误信息
//* @apiSuccess {String} results 返回的结果
//* @apiSuccessExample {json} Response-Example:
//{
//    "err_no":0,
//    "err_msg":"",
//    "results": {
//		"token": "asfasfqfsfe",
//		"group_id": 1,
//		"user_id": 2,
//		"authority": ["a", "b", "c"]
//	}
//}
//*/
//
//func Login(httpCtx *gin.Context) {
//	req := &user2.LoginRequest{}
//	httpCtx.ThrowCheck(lib.IntegrityError, api.RequestUnmarshal(httpCtx, req))
//	httpCtx.ThrowCheck(lib.LegalityError, lib.ValidParams(req))
//	res, err := req.Login()
//	httpCtx.ThrowCheck(lib.VerifyError(err))
//	httpCtx.Results = res
//}
//
///**
//* @api {post} /user/logout 用户登出
//* @apiVersion 1.0.0
//* @apiGroup User
//* @apiDescription 用户登出
//*
//* @apiParam {String} token 登录授权码
//*
//* @apiParamExample {json} Request-Example:
//{
//	"token":"sss"
//}
//* @apiSuccess {Number} err_no 错误码，大于0表示出错
//* @apiSuccess {String} err_msg 错误信息
//* @apiSuccess {String} results 返回的结果
//* @apiSuccessExample {json} Response-Example:
//{
//    "err_no":0,
//    "err_msg":"",
//    "results": "Success"
//}
//*/
//
//func Logout(httpCtx *gin.Context) {
//	if !user2.Offline(httpCtx.Data["token"].(string)) {
//		httpCtx.ThrowCheck(lib.MarkPlatformDataError, lib.InvalidUsernameOrPassword)
//	}
//	httpCtx.Results = "Success"
//}
//
///**
//* @api {post} /user/update 修改密码
//* @apiVersion 1.0.0
//* @apiGroup User
//* @apiDescription 重置秘密
//*
//* @apiParam {String} password 密码
//* @apiParam {String} token 登录授权码
//*
//* @apiParamExample {json} Request-Example:
//{
//	"password": "df81b3f78e9bd3990ba51e28ca6d11fd",
//	"token":"sss"
//}
//* @apiSuccess {Number} err_no 错误码，大于0表示出错
//* @apiSuccess {String} err_msg 错误信息
//* @apiSuccess {String} results 返回的结果
//* @apiSuccessExample {json} Response-Example:
//{
//    "err_no":0,
//    "err_msg":"",
//    "results": "Success"
//}
//*/
//
//func Update(httpCtx *gin.Context) {
//	req := &user2.ResetPasswordRequest{}
//	httpCtx.ThrowCheck(lib.IntegrityError, api.RequestUnmarshal(httpCtx, req))
//	httpCtx.ThrowCheck(lib.LegalityError, lib.ValidParams(req))
//	err := req.ResetPassword(httpCtx)
//	httpCtx.ThrowCheck(lib.VerifyError(err))
//	httpCtx.Results = "Success"
//}
