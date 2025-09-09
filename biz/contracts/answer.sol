// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

/**
 * @title QuestionAnswerContract
 * @dev 用于记录用户回答问题的合约
 * @notice 每个地址每天每个问题ID只能提交一次答案
 */
contract QuestionAnswerContract {

    // 答案记录结构体
    struct AnswerRecord {
        address user;           // 用户地址
        string answer;          // 答案内容
        uint256 questionId;     // 问题ID
        uint256 timestamp;      // 提交时间戳
        uint256 day;           // 提交日期（天数）
    }

    // 存储所有答案记录
    AnswerRecord[] public answerRecords;

    uint256 public startTime;


    // 记录用户每天每个问题的提交状态
    // mapping(用户地址 => mapping(天数 => mapping(问题ID => 是否已提交)))
    mapping(address => mapping(uint256 => mapping(uint256 => bool))) public dailySubmissions;

    // 记录用户的答案历史
    // mapping(用户地址 => 答案记录ID数组)
    mapping(address => uint256[]) public userAnswerHistory;

    // 记录每个问题的答案数量
    mapping(uint256 => uint256) public questionAnswerCount;

    // 事件定义
    event AnswerSubmitted(
        address indexed user,
        string answer,
        uint256 questionId,
        uint256 timestamp,
        uint256 day
    );

    event DailyLimitReached(
        address indexed user,
        uint256 questionId,
        uint256 day
    );

    // 错误定义
    error AlreadySubmittedToday(address user, uint256 questionId, uint256 day);
    error EmptyAnswer();
    error InvalidQuestionId();

    constructor(uint256 _startTime) {
        startTime = _startTime;
    }

    /**
     * @dev 提交答案函数
     * @param answer 答案内容
     * @param id 问题ID
     */
    function submit_answer(string calldata answer, uint256 id) external {
        // 验证输入参数
        if (bytes(answer).length == 0) {
            revert EmptyAnswer();
        }

        if (id == 0) {
            revert InvalidQuestionId();
        }

        // 获取当前天数（从1970年1月1日开始的天数）
        uint256 currentDay = getCurrentDayOffset();

        // 检查用户今天是否已经提交过这个问题的答案
        if (dailySubmissions[msg.sender][currentDay][id]) {
            emit DailyLimitReached(msg.sender, id, currentDay);
            revert AlreadySubmittedToday(msg.sender, id, currentDay);
        }

        // 记录提交状态
        dailySubmissions[msg.sender][currentDay][id] = true;

        // 创建答案记录
        AnswerRecord memory newRecord = AnswerRecord({
            user: msg.sender,
            answer: answer,
            questionId: id,
            timestamp: block.timestamp,
            day: currentDay
        });

        // 保存到数组中
        answerRecords.push(newRecord);

        // 记录用户答案历史
        uint256 recordIndex = answerRecords.length - 1;
        userAnswerHistory[msg.sender].push(recordIndex);

        // 更新问题答案计数
        questionAnswerCount[id]++;

        // 发送事件
        emit AnswerSubmitted(
            msg.sender,
            answer,
            id,
            block.timestamp,
            currentDay
        );
    }

    /**
     * @dev 获取当前天数
     * @return 从1970年1月1日开始的天数
     */
    function getCurrentDayOffset() public view returns (uint256) {
        return (block.timestamp - startTime) / 86400;
    }

    /**
     * @dev 检查用户今天是否已经提交过某个问题的答案
     * @param user 用户地址
     * @param questionId 问题ID
     * @return 是否已提交
     */
    function hasSubmittedToday(address user, uint256 questionId) external view returns (bool) {
        uint256 currentDay = getCurrentDayOffset();
        return dailySubmissions[user][currentDay][questionId];
    }

    /**
     * @dev 检查用户在指定日期是否已经提交过某个问题的答案
     * @param user 用户地址
     * @param day 指定日期
     * @param questionId 问题ID
     * @return 是否已提交
     */
    function hasSubmittedOnDay(address user, uint256 day, uint256 questionId) external view returns (bool) {
        return dailySubmissions[user][day][questionId];
    }

    /**
     * @dev 获取用户的答案历史记录数量
     * @param user 用户地址
     * @return 答案记录数量
     */
    function getUserAnswerCount(address user) external view returns (uint256) {
        return userAnswerHistory[user].length;
    }

    /**
     * @dev 获取用户的答案历史记录（分页）
     * @param user 用户地址
     * @param offset 偏移量
     * @param limit 限制数量
     * @return records 答案记录数组
     */
    function getUserAnswerHistory(
        address user,
        uint256 offset,
        uint256 limit
    ) external view returns (AnswerRecord[] memory records) {
        uint256[] memory userRecordIds = userAnswerHistory[user];
        uint256 totalRecords = userRecordIds.length;

        if (offset >= totalRecords) {
            return new AnswerRecord[](0);
        }

        uint256 end = offset + limit;
        if (end > totalRecords) {
            end = totalRecords;
        }

        uint256 resultLength = end - offset;
        records = new AnswerRecord[](resultLength);

        for (uint256 i = 0; i < resultLength; i++) {
            records[i] = answerRecords[userRecordIds[offset + i]];
        }
    }

    /**
     * @dev 获取指定问题的答案记录（分页）
     * @param questionId 问题ID
     * @param offset 偏移量
     * @param limit 限制数量
     * @return records 答案记录数组
     * @return total 总记录数
     */
    function getQuestionAnswers(
        uint256 questionId,
        uint256 offset,
        uint256 limit
    ) external view returns (AnswerRecord[] memory records, uint256 total) {
        // 先计算总数和符合条件的记录
        uint256 matchCount = 0;
        uint256 totalRecords = answerRecords.length;

        // 计算匹配的记录数
        for (uint256 i = 0; i < totalRecords; i++) {
            if (answerRecords[i].questionId == questionId) {
                matchCount++;
            }
        }

        total = matchCount;

        if (offset >= matchCount || limit == 0) {
            return (new AnswerRecord[](0), total);
        }

        uint256 end = offset + limit;
        if (end > matchCount) {
            end = matchCount;
        }

        uint256 resultLength = end - offset;
        records = new AnswerRecord[](resultLength);

        uint256 currentIndex = 0;
        uint256 resultIndex = 0;

        for (uint256 i = 0; i < totalRecords && resultIndex < resultLength; i++) {
            if (answerRecords[i].questionId == questionId) {
                if (currentIndex >= offset) {
                    records[resultIndex] = answerRecords[i];
                    resultIndex++;
                }
                currentIndex++;
            }
        }
    }

    /**
     * @dev 获取总答案记录数
     * @return 总记录数
     */
    function getTotalAnswerCount() external view returns (uint256) {
        return answerRecords.length;
    }

    /**
     * @dev 获取指定问题的答案数量
     * @param questionId 问题ID
     * @return 答案数量
     */
    function getQuestionAnswerCount(uint256 questionId) external view returns (uint256) {
        return questionAnswerCount[questionId];
    }

    /**
     * @dev 获取答案记录详情
     * @param recordId 记录ID
     * @return record 答案记录
     */
    function getAnswerRecord(uint256 recordId) external view returns (AnswerRecord memory record) {
        require(recordId < answerRecords.length, "Record does not exist");
        return answerRecords[recordId];
    }

    /**
     * @dev 批量获取答案记录
     * @param offset 偏移量
     * @param limit 限制数量
     * @return records 答案记录数组
     */
    function getAnswerRecords(
        uint256 offset,
        uint256 limit
    ) external view returns (AnswerRecord[] memory records) {
        uint256 totalRecords = answerRecords.length;

        if (offset >= totalRecords) {
            return new AnswerRecord[](0);
        }

        uint256 end = offset + limit;
        if (end > totalRecords) {
            end = totalRecords;
        }

        uint256 resultLength = end - offset;
        records = new AnswerRecord[](resultLength);

        for (uint256 i = 0; i < resultLength; i++) {
            records[i] = answerRecords[offset + i];
        }
    }
}
