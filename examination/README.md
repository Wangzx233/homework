服务器地址：http://47.106.170.23/



https://github.com/Wangzx233/homework/tree/main/examination

gihub始终进不去，盲push的，不是很确定有没有

## AIP接口:

登录

| GET      | /USER/LOGIN |
| :------- | ----------- |
| username | 用户uid     |
| password | 密码        |

注册

| POST     | /USER/REGISTER |
| :------- | -------------- |
| username | 用户uid        |
| password | 密码           |

进入房间前验证是否满足要求

| POST    | /USER/CHEEK_INTER |
| :------ | ----------------- |
| room_id | 房间号            |

创建房间前验证是否满足要求

| POST    | /USER/CHEEK_create |
| :------ | ------------------ |
| room_id | 房间号             |

创建房间时客户端发送wesocket

| 名称    | 值       |
| :------ | -------- |
| type    | "create" |
| user    | 玩家id   |
| content | ""       |
| room_id | 房间号   |

加入房间时客户端发送wesocket

| 名称    | 值      |
| :------ | ------- |
| type    | "login" |
| user    | 玩家id  |
| content | ""      |
| room_id | 房间号  |

当玩家准备时发送结构体

| 名称    | 值      |
| :------ | ------- |
| type    | "ready" |
| user    | 用户名  |
| content | ""      |
| room_id | 房间号  |

当玩家取消准备时发送结构体

| 名称    | 值        |
| :------ | --------- |
| type    | "unready" |
| user    | 用户名    |
| content | ""        |
| room_id | 房间号    |



都准备完毕时返回

| 名称    | 值                             |
| :------ | ------------------------------ |
| type    | "gameStart"                    |
| user    | 1或者2（1代表黑棋，2代表白旗） |
| content | ""                             |
| room_id | 房间号                         |



玩家发送移动请求时

| 名称    | 值                       |
| :------ | ------------------------ |
| type    | "move"                   |
| user    | 1或者2（当前回合的玩家） |
| content | 坐标，例如"1,2"          |
| room_id | 房间号                   |

玩家发送移动通过验证规则后服务器返回

| 名称    | 值                         |
| :------ | -------------------------- |
| type    | "next_move"                |
| user    | 1或者2（下一个回合的玩家） |
| content | 坐标，例如"1,2"            |
| room_id | 房间号                     |

| 名称    | 值                       |
| :------ | ------------------------ |
| type    | "moved"                  |
| user    | 1或者2（当前回合的玩家） |
| content | 坐标，例如"1,2"          |
| room_id | 房间号                   |

当玩家胜利时返回

| 名称    | 值                             |
| :------ | ------------------------------ |
| type    | "game_over"                    |
| user    | 胜利的玩家（由之前获得的1和2） |
| content | ""                             |
| room_id | 房间号                         |

## 加分项及实现方法：

- 服务拆分，用户中心使用begonia单独部署
- 使用html完成了棋盘的等
- 房间内和游戏中玩家可以聊天

## 代码运行环境：mysql

