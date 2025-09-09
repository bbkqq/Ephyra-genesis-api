# Ephyra Genesis API 接口文档

## 接口返回格式说明

所有接口都遵循统一的返回格式：

```json
{
  "status": true,  // 布尔值，表示请求是否成功
  "data": {},      // 具体的返回数据，结构根据接口而定
  "error": ""      // 当status为false时，包含错误信息
}
```

## 认证说明

需要认证的接口需要在请求头中携带JWT token：
```
Authorization: <jwt_token>
```

---

## 1. Common 接口

### 1.1 获取总贡献数

**接口地址**: `GET /api/common/contributions/`

**请求参数**: 无

**返回数据结构**:
```json
{
  "status": true,
  "data": {
    "contribution_num": 12345  // 所有用户参与任务的总次数
  }
}
```

**接口说明**: 获取 user_daily_task_log 表的总记录数，表示所有用户参与任务的总次数。

---

## 2. Login 接口

### 2.1 用户登录

**接口地址**: `POST /api/login/`

**请求参数**:
```json
{
  "address": "0x1234567890abcdef1234567890abcdef12345678",     // 钱包地址
  "sign_message": "Please sign this message to login",        // 签名消息
  "signature": "0xabcdef...",                                 // 钱包签名
  "chain_name": "ethereum"                                    // 区块链网络名称
}
```

**返回数据结构**:
```json
{
  "status": true,
  "data": {
    "user_id": 12345,                                         // 用户ID
    "auth_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",  // JWT认证token
    "address": "0x1234567890abcdef1234567890abcdef12345678",   // 用户钱包地址
    "is_new": false                                           // 是否为新用户
  }
}
```

**接口说明**: 
- 用户通过钱包签名进行登录
- 验证签名的正确性
- 新用户会自动创建账户
- 返回JWT token用于后续接口认证

---

## 3. Task 接口

### 3.1 获取每日问题

**接口地址**: `GET /api/task/daily-questions/`

**请求参数**: 无

**返回数据结构**:
```json
{
  "status": true,
  "data": {
    "questions": [
      {
        "id": 1,                    // 问题ID
        "title": "今日问题标题",     // 问题标题
        "content": "问题详细内容",   // 问题详细内容
        "type": 1                   // 问题类型：1-单选题，2-多选题，3-填空题，4-问答题，5-判断题
      }
    ]
  }
}
```

**接口说明**: 获取当前可用的每日问题列表。

### 3.2 检查是否可以提交答案

**接口地址**: `POST /api/task/can-submit-answer/`

**认证**: 需要JWT token

**请求参数**: 无

**返回数据结构**:
```json
{
  "status": true,
  "data": {
    "status": true,                    // 是否可以提交答案
    "estimated_points": 60,            // 预计获得的基础积分
    "estimated_extra_points": 10       // 预计获得的额外积分（连续奖励）
  }
}
```

**接口说明**: 
- 检查用户今天是否已经提交过答案
- 如果可以提交，会返回预计获得的积分
- 基于用户当前的连续参与天数计算奖励

### 3.3 获取答案交易状态

**接口地址**: `POST /api/task/get-answer-tx-status/`

**请求参数**:
```json
{
  "tx_hash": "0xabcdef1234567890..."  // 区块链交易哈希
}
```

**返回数据结构**:
```json
{
  "status": true,
  "data": {
    "status": 1,                      // 合约执行状态：0-未处理，1-成功，-1-失败
    "badges": [                       // 本次交互获得的徽章列表
      {
        "id": 1,
        "name": "初次参与",
        "image_url": "https://example.com/badge1.png",
        "unlocked": true,
        "unlock_conditions": "完成第一次每日任务"
      }
    ],
    "points": 50,                     // 本次交互的基础积分奖励
    "extra_points": 10                // 本次交互的额外积分奖励（连击奖励）
  }
}
```

**接口说明**: 
- 根据交易哈希查询链上交易的处理结果
- 返回用户在该次交易中获得的积分和徽章
- 积分分为基础积分和连续参与的额外积分

---

## 4. User 接口

### 4.1 获取用户档案

**接口地址**: `GET /api/user/get-user-profile/`

**认证**: 需要JWT token

**请求参数**: 无

**返回数据结构**:
```json
{
  "status": true,
  "data": {
    "uid": 12345,                     // 用户ID
    "address": "0x1234...",           // 用户钱包地址
    "user_name": "用户昵称",           // 用户昵称
    "avatar_url": "https://...",      // 用户头像URL
    "points": 1500,                   // 用户总积分
    "badges": [                       // 用户所有徽章（包括已解锁和未解锁）
      {
        "id": 1,
        "name": "初次参与",
        "image_url": "https://example.com/badge1.png",
        "unlocked": true,             // 是否已解锁
        "unlock_conditions": "完成第一次每日任务"
      },
      {
        "id": 2,
        "name": "七日先知",
        "image_url": "https://example.com/badge2.png",
        "unlocked": false,
        "unlock_conditions": "连续参与7天活动"
      }
    ],
    "contributions": {                // 用户参与贡献统计
      "total_days": 25,               // 总参与天数
      "consecutive_days": 5,          // 当前连续参与天数
      "max_consecutive_days": 12      // 历史最长连续参与天数
    },
    "points_logs": [                  // 最近5次积分历史
      {
        "total_points": 70,           // 总积分
        "base_points": 50,            // 基础积分
        "extra_points": 20,           // 额外积分
        "type": "daily_task",         // 积分类型
        "create_time": 1625097600     // 创建时间（Unix时间戳）
      }
    ]
  }
}
```

**接口说明**: 获取当前用户的完整档案信息，包括积分、徽章、贡献统计等。

### 4.2 设置用户档案

**接口地址**: `POST /api/user/set-user-profile/`

**认证**: 需要JWT token

**请求参数**:
```json
{
  "user_name": "新昵称",              // 用户昵称
  "avatar_url": "https://..."        // 用户头像URL
}
```

**返回数据结构**:
```json
{
  "status": true,
  "data": {
    "status": true                    // 设置是否成功
  }
}
```

**接口说明**: 更新用户的昵称和头像。

### 4.3 获取用户排名

**接口地址**: `GET /api/user/get-user-ranking/`

**请求参数**:
```json
{
  "type": "daily"                    // 排名类型：daily-日榜，weekly-周榜，all-总榜
}
```

**返回数据结构**:
```json
{
  "status": true,
  "data": {
    "user_ranking": 15,               // 当前用户的排名（未登录时为0）
    "ranking_list": [                 // 排行榜列表（前100名）
      {
        "address": "0x1234...",       // 用户地址
        "user_name": "用户昵称",       // 用户昵称
        "points": "1500"              // 用户积分（字符串格式）
      }
    ]
  }
}
```

**接口说明**: 
- 获取指定类型的用户排名
- 支持日榜、周榜、总榜三种类型
- **日榜(daily)**: 从Redis实时获取当天积分排名，数据最新
- **周榜(weekly)**: 从Redis获取缓存的周积分排名，每12小时更新一次
- **总榜(all)**: 从MySQL获取用户总积分排名
- 返回前100名用户和当前用户的排名
- 如果用户排名超过1000名，返回1000
- 未登录用户也可以查看排行榜，但user_ranking为0

---

## 积分系统说明

### 积分计算规则

| 连续天数 | 基础积分 | 额外积分 | 总积分 |
|---------|---------|---------|--------|
| 第1天   | 50分    | 0分     | 50分   |
| 第2天   | 50分    | 10分    | 60分   |
| 第3天   | 50分    | 20分    | 70分   |
| 第4天   | 50分    | 30分    | 80分   |
| 第5天   | 50分    | 40分    | 90分   |
| 第6天   | 50分    | 50分    | 100分  |
| 第7天   | 50分    | 60分    | 110分  |
| 第8天+  | 50分    | 60分    | 110分  |

**注意**: 如果用户中断参与，连续天数会重置，重新从第1天的50分开始累积。

### 徽章系统

| 徽章ID | 徽章名称 | 获得条件 |
|--------|---------|----------|
| 1 | 初次参与 | 用户第一次参与每日任务 |
| 2 | 七日先知 | 连续参与7天活动（只能获得一次） |
| 3 | 百问智者 | 累计参与100天活动 |

---

## 错误码说明

常见错误返回示例：

### 认证失败
```json
{
  "status": false,
  "error": "Unauthorized: Invalid or expired token"
}
```

### 参数错误
```json
{
  "status": false,
  "error": "Invalid parameters"
}
```

### 数据库错误
```json
{
  "status": false,
  "error": "Database error"
}
```

### 签名验证失败
```json
{
  "status": false,
  "error": "Invalid signature"
}
```

---

## 使用流程示例

### 1. 用户登录流程
```javascript
// 1. 用户连接钱包
// 2. 获取签名消息
// 3. 用户签名
// 4. 调用登录接口
const loginResponse = await fetch('/api/login/', {
  method: 'POST',
  body: JSON.stringify({
    address: userAddress,
    sign_message: signMessage,
    signature: signature,
    chain_name: 'ethereum'
  })
});

const loginData = await loginResponse.json();
if (loginData.status) {
  // 保存token用于后续请求
  const token = loginData.data.auth_token;
}
```

### 2. 获取用户信息流程
```javascript
// 使用登录获得的token
const profileResponse = await fetch('/api/user/get-user-profile/', {
  method: 'GET',
  headers: {
    'Authorization': token
  }
});

const profileData = await profileResponse.json();
if (profileData.status) {
  // 显示用户信息
  console.log('用户积分:', profileData.data.points);
  console.log('连续天数:', profileData.data.contributions.consecutive_days);
}
```

### 3. 每日任务流程
```javascript
// 1. 获取每日问题
const questionsResponse = await fetch('/api/task/daily-questions/');
const questionsData = await questionsResponse.json();

// 2. 检查是否可以提交答案
const canSubmitResponse = await fetch('/api/task/can-submit-answer/', {
  method: 'POST',
  headers: { 'Authorization': token }
});
const canSubmitData = await canSubmitResponse.json();

if (canSubmitData.status && canSubmitData.data.status) {
  console.log('可以提交答案，预计获得积分:', 
    canSubmitData.data.estimated_points + canSubmitData.data.estimated_extra_points);
  
  // 3. 用户在区块链上提交答案
  // 4. 获取交易哈希后查询状态
  const txStatusResponse = await fetch('/api/task/get-answer-tx-status/', {
    method: 'POST',
    headers: { 'Authorization': token },
    body: JSON.stringify({ tx_hash: transactionHash })
  });
  
  const txStatusData = await txStatusResponse.json();
  if (txStatusData.status && txStatusData.data.status === 1) {
    console.log('任务完成！获得积分:', txStatusData.data.points + txStatusData.data.extra_points);
    console.log('获得徽章:', txStatusData.data.badges);
  }
}
```

---

## 数据类型说明

### Badge (徽章)
```typescript
interface Badge {
  id: number;                    // 徽章ID
  name: string;                  // 徽章名称
  image_url: string;             // 徽章图片URL
  unlocked: boolean;             // 是否已解锁
  unlock_conditions: string;     // 解锁条件说明
}
```

### UserContributions (用户贡献)
```typescript
interface UserContributions {
  total_days: number;            // 总参与天数
  consecutive_days: number;      // 当前连续参与天数
  max_consecutive_days: number;  // 历史最长连续参与天数
}
```

### UserPointsLogs (积分历史)
```typescript
interface UserPointsLogs {
  total_points: number;          // 总积分
  base_points: number;           // 基础积分
  extra_points: number;          // 额外积分
  type: string;                  // 积分类型
  create_time: number;           // 创建时间（Unix时间戳）
}
```

### Question (问题)
```typescript
interface Question {
  id: number;                    // 问题ID
  title: string;                 // 问题标题
  content: string;               // 问题内容
  type: number;                  // 问题类型
}
```

### UserRanking (用户排名)
```typescript
interface UserRanking {
  address: string;               // 用户地址
  user_name: string;             // 用户昵称
  points: string;                // 用户积分（字符串格式）
}
```

---

## 开发注意事项

1. **时间格式**: 所有时间字段使用Unix时间戳（秒）
2. **积分字段**: 积分相关字段统一使用数字类型，除了排名接口中的积分使用字符串
3. **徽章数据**: 徽章列表包含所有徽章，通过 `unlocked` 字段区分是否已解锁
4. **排名类型**: 支持 `daily`、`weekly`、`all` 三种排名类型
5. **错误处理**: 所有接口都遵循统一的错误返回格式
6. **认证**: 除登录接口和部分查询接口外，其他接口都需要JWT认证

---

## 更新日志

- v1.0.0: 初始版本，包含用户登录、任务系统、积分系统、排名系统
- 支持Web3钱包登录
- 完整的积分奖励机制
- 三种类型徽章系统
- 多维度排行榜
