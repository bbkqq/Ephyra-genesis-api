package jobs

import (
	"Ephyra-genesis-api/biz/dal"
	"Ephyra-genesis-api/pkg/cron_job"
	"context"
	"testing"
)

func TestAnswerEventListenerJob(t *testing.T) {
	//init dal
	dal.Init()
	answerEventListenerJob := cron_job.NewCronJob(NewAnswerEventListenerJob())
	answerEventListenerJob.Run(context.Background())
}
