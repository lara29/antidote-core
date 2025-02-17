package api

import (
	"context"
	"errors"

	pb "github.com/nre-learning/antidote-core/api/exp/generated"
	ot "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	log "github.com/opentracing/opentracing-go/log"
)

// GetCurriculumInfo is designed to only get high-level information about the loaded Curriculum. Specifics on lessons, collections, and more
// are given their own first-level API endpoint elsewhere, but will be pulled from the loaded Curriculum struct being described here.
func (s *AntidoteAPI) GetCurriculumInfo(ctx context.Context, filter *pb.CurriculumFilter) (*pb.CurriculumInfo, error) {
	span := ot.StartSpan("api_curriculum_getinfo", ext.SpanKindRPCClient)
	defer span.Finish()

	curriculum, err := s.Db.GetCurriculum(span.Context())
	if err != nil {
		span.LogFields(log.Error(err))
		ext.Error.Set(span, true)
		return nil, errors.New("Problem retrieving curriculum details")
	}

	return &pb.CurriculumInfo{
		Name:        curriculum.Name,
		Description: curriculum.Description,
		Website:     curriculum.Website,
		AVer:        curriculum.AVer,
		GitRoot:     curriculum.GitRoot,
	}, nil
}
