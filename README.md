# Digital-gin 数科技术平台-gin

## 页面地址

- dev： 192.168.0.217 digitalweb.dev.com
- test：192.168.0.217 digitalweb.test.com
- prod：116.63.203.160:60011

# docs

- dev： 192.168.0.217:30012/digital/docs
- test：192.168.0.217:31012/digital/docs
- prod：116.63.203.160:60007/digital/docs

## 用户组

| **GroupID** | **GroupName** | **Comment** |
|-------------|---------------|-------------|
| `-1`        | `　上帝　`        | `God`       |
| `1`         | `　管理员　`       | `Admin`     |
| `2`         | `　数据支持`       | `Support`   |
| `3`         | `　数据中心`       | `Artery`    |
| `4`         | `　人工智能`       | `AI`        |
| `5`         | `　大数据　`       | `BigData`   |
| `6`         | `　测试　　`       | `Test`      |
| `7`         | `　运维　　`       | `IT`        |
| `8`         | `　游客　　`       | `Guest`     |

# 错误码对照表（初稿）

| **const**                         | **number** | **error**   | **Scenes**                |
|-----------------------------------|------------|-------------|---------------------------|
| `Pending`                         | `1001`     | 请求执行中　　　　　　 | 请求未完成                     |
| `IntegrityError`                  | `4002`     | 请求参数不合规范　　　 | 必填参数缺失                    |
| `LegalityError`                   | `4003`     | 请求参数不合规范　　　 | 参数格式问题                    |
| `MarkPlatformSystemError`         | `5000`     | 数科技术平台系统异常　 | 字符串解析异常、json解析出错、数据库相关异常等 |
| `ArteryResultError`               | `5001`     | 数据流结果异常　　　　 | 接口调用异常、服务异常               |
| `AlgoResultError`                 | `5002`     | 算法结果异常　　　　　 | 接口调用异常、服务异常               |
| `SearchRecommendationResultError` | `5003`     | 搜索结果异常　　　　　 | 接口调用异常、服务异常               |
| `MarkPlatformDataError`           | `5004`     | 数科技术平台数据异常　 | 数据不存在                     |

# 项目注意事项
0.本项目是 *digital* 项目基于 *gin* 框架重构版  
1.在添加“数科内部菜单”组的菜单时，需要手动把“parent_id”关联到对应的一级菜单下  
2.组为“数科内部菜单”的权限，内部所有组都可以使用  
3.目前一些没有组的“数科内部菜单”权限，都统一挂在“域名管理组”