package dal

import (
	"Ephyra-genesis-api/storage/mysql/model"
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	var st = time.Date(2025, 9, 9, 0, 0, 0, 0, time.UTC)
	fmt.Println(st.Unix())
}

func TestCalculateTaskReward(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name                string
		userModel           *model.User
		targetDay           int32
		expectedPoints      int32
		expectedExtraPoints int32
		expectedNewBadges   []int64
		expectedIsNewUser   bool
		expectedConsecutive int64
	}{
		{
			name: "新用户第一天参与",
			userModel: &model.User{
				ID:         1,
				Address:    "0x123",
				TaskStatus: strings.Repeat("0", 300), // 全0表示新用户
				Badges:     "[]",
			},
			targetDay:           0,
			expectedPoints:      50,
			expectedExtraPoints: 0,
			expectedNewBadges:   []int64{BadgeFirstParticipation},
			expectedIsNewUser:   true,
			expectedConsecutive: 1,
		},
		{
			name: "用户连续第二天参与",
			userModel: &model.User{
				ID:         1,
				Address:    "0x123",
				TaskStatus: "1" + strings.Repeat("0", 299), // 第0天已完成
				Badges:     "[1]",                          // 已有初次参与徽章
			},
			targetDay:           1,
			expectedPoints:      50,
			expectedExtraPoints: 10,
			expectedNewBadges:   []int64{},
			expectedIsNewUser:   false,
			expectedConsecutive: 2,
		},
		{
			name: "用户连续第三天参与",
			userModel: &model.User{
				ID:         1,
				Address:    "0x123",
				TaskStatus: "11" + strings.Repeat("0", 298), // 前2天已完成
				Badges:     "[1]",
			},
			targetDay:           2,
			expectedPoints:      50,
			expectedExtraPoints: 20,
			expectedNewBadges:   []int64{},
			expectedIsNewUser:   false,
			expectedConsecutive: 3,
		},
		{
			name: "用户连续第七天参与，获得七日先知徽章",
			userModel: &model.User{
				ID:         1,
				Address:    "0x123",
				TaskStatus: "111111" + strings.Repeat("0", 294), // 前6天已完成
				Badges:     "[1]",
			},
			targetDay:           6,
			expectedPoints:      50,
			expectedExtraPoints: 60,
			expectedNewBadges:   []int64{BadgeSevenDayOracle},
			expectedIsNewUser:   false,
			expectedConsecutive: 7,
		},
		{
			name: "用户连续第八天参与，积分保持最高",
			userModel: &model.User{
				ID:         1,
				Address:    "0x123",
				TaskStatus: "1111111" + strings.Repeat("0", 293), // 前7天已完成
				Badges:     "[1,2]",                              // 已有初次参与和七日先知徽章
			},
			targetDay:           7,
			expectedPoints:      50,
			expectedExtraPoints: 60,
			expectedNewBadges:   []int64{},
			expectedIsNewUser:   false,
			expectedConsecutive: 8,
		},
		{
			name: "用户中断后重新参与",
			userModel: &model.User{
				ID:         1,
				Address:    "0x123",
				TaskStatus: "1110" + strings.Repeat("0", 296), // 前3天完成，第4天中断
				Badges:     "[1]",
			},
			targetDay:           4,
			expectedPoints:      50,
			expectedExtraPoints: 0, // 重新开始，只有基础积分
			expectedNewBadges:   []int64{},
			expectedIsNewUser:   false,
			expectedConsecutive: 1,
		},
		{
			name: "用户累计参与100天，获得百问智者徽章",
			userModel: &model.User{
				ID:      1,
				Address: "0x123",
				// 创建一个有99个1的字符串，加上当前这次就是100天
				TaskStatus: strings.Repeat("1", 99) + strings.Repeat("0", 201),
				Badges:     "[1,2]",
			},
			targetDay:           99,
			expectedPoints:      50,
			expectedExtraPoints: 60, // 连续100天，但额外积分最多60
			expectedNewBadges:   []int64{BadgeHundredDayWise},
			expectedIsNewUser:   false,
			expectedConsecutive: 100,
		},
		{
			name: "空的TaskStatus新用户",
			userModel: &model.User{
				ID:         1,
				Address:    "0x123",
				TaskStatus: "",
				Badges:     "[]",
			},
			targetDay:           0,
			expectedPoints:      50,
			expectedExtraPoints: 0,
			expectedNewBadges:   []int64{BadgeFirstParticipation},
			expectedIsNewUser:   true,
			expectedConsecutive: 1,
		},
		{
			name: "targetDay超出当前TaskStatus长度",
			userModel: &model.User{
				ID:         1,
				Address:    "0x123",
				TaskStatus: "111", // 只有3天记录
				Badges:     "[1]",
			},
			targetDay:           10, // 第11天
			expectedPoints:      50,
			expectedExtraPoints: 0, // 中断了，重新开始
			expectedNewBadges:   []int64{},
			expectedIsNewUser:   false,
			expectedConsecutive: 1,
		},
		{
			name: "已有所有徽章的用户",
			userModel: &model.User{
				ID:         1,
				Address:    "0x123",
				TaskStatus: "1111111" + strings.Repeat("0", 293),
				Badges:     "[1,2,3]", // 已有所有徽章
			},
			targetDay:           7,
			expectedPoints:      50,
			expectedExtraPoints: 60,
			expectedNewBadges:   []int64{}, // 不会获得新徽章
			expectedIsNewUser:   false,
			expectedConsecutive: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := CalculateTaskReward(ctx, tt.userModel, tt.targetDay)

			// 断言没有错误
			assert.NoError(t, err)
			assert.NotNil(t, result)

			// 断言积分计算正确
			assert.Equal(t, tt.expectedPoints, result.Points, "基础积分不匹配")
			assert.Equal(t, tt.expectedExtraPoints, result.ExtraPoints, "额外积分不匹配")

			// 断言徽章正确
			assert.Equal(t, tt.expectedNewBadges, result.NewBadges, "新徽章不匹配")

			// 断言新用户标识正确
			assert.Equal(t, tt.expectedIsNewUser, result.IsNewUser, "新用户标识不匹配")

			// 断言连续天数正确
			assert.Equal(t, tt.expectedConsecutive, result.ConsecutiveDays, "连续天数不匹配")
		})
	}
}

func TestCalculateTaskReward_InvalidBadgeJSON(t *testing.T) {
	ctx := context.Background()

	userModel := &model.User{
		ID:         1,
		Address:    "0x123",
		TaskStatus: strings.Repeat("0", 300),
		Badges:     "invalid json", // 无效的JSON
	}

	result, err := CalculateTaskReward(ctx, userModel, 0)

	// 即使JSON无效，函数也应该正常工作，只是不会检查已有徽章
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, int32(50), result.Points)
	assert.Equal(t, int32(0), result.ExtraPoints)
	assert.Equal(t, []int64{BadgeFirstParticipation}, result.NewBadges) // 应该获得初次参与徽章
	assert.True(t, result.IsNewUser)
	assert.Equal(t, int64(1), result.ConsecutiveDays)
}

func TestCalculateConsecutiveDaysFromDay(t *testing.T) {
	tests := []struct {
		name       string
		taskStatus string
		targetDay  int
		expected   int64
	}{
		{
			name:       "单独一天",
			taskStatus: "1",
			targetDay:  0,
			expected:   1,
		},
		{
			name:       "连续两天",
			taskStatus: "11",
			targetDay:  1,
			expected:   2,
		},
		{
			name:       "中断后重新开始",
			taskStatus: "1101",
			targetDay:  3,
			expected:   1,
		},
		{
			name:       "连续七天",
			taskStatus: "1111111",
			targetDay:  6,
			expected:   7,
		},
		{
			name:       "中间有中断",
			taskStatus: "11101111",
			targetDay:  7,
			expected:   4, // 从第7天往前，连续4天
		},
		{
			name:       "targetDay超出范围",
			taskStatus: "111",
			targetDay:  5,
			expected:   0,
		},
		{
			name:       "空字符串",
			taskStatus: "",
			targetDay:  0,
			expected:   0,
		},
		{
			name:       "负数targetDay",
			taskStatus: "111",
			targetDay:  -1,
			expected:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateConsecutiveDaysFromDay(tt.taskStatus, tt.targetDay)
			assert.Equal(t, tt.expected, result, "连续天数计算不正确")
		})
	}
}

func TestContainsBadge(t *testing.T) {
	tests := []struct {
		name     string
		badges   []int64
		badgeID  int64
		expected bool
	}{
		{
			name:     "包含徽章",
			badges:   []int64{1, 2, 3},
			badgeID:  2,
			expected: true,
		},
		{
			name:     "不包含徽章",
			badges:   []int64{1, 2, 3},
			badgeID:  4,
			expected: false,
		},
		{
			name:     "空徽章列表",
			badges:   []int64{},
			badgeID:  1,
			expected: false,
		},
		{
			name:     "单个徽章匹配",
			badges:   []int64{5},
			badgeID:  5,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := containsBadge(tt.badges, tt.badgeID)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// 基准测试
func BenchmarkCalculateTaskReward(b *testing.B) {
	ctx := context.Background()
	userModel := &model.User{
		ID:         1,
		Address:    "0x123",
		TaskStatus: "1111111" + strings.Repeat("0", 293),
		Badges:     "[1,2]",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := CalculateTaskReward(ctx, userModel, 7)
		if err != nil {
			b.Fatal(err)
		}
	}
}
