package scv

import (
	"sync"
	"time"
)

// Cached object persisted in Mongo
type Stream struct {
	sync.RWMutex
	streamId     string
	targetId     string
	status       string
	frames       int
	errorCount   int
	creationDate int
	activeStream *ActiveStream
}

func NewStream(streamId, targetId, status string,
	frames, errorCount, creationDate int) *Stream {
	stream := &Stream{
		streamId:     streamId,
		targetId:     targetId,
		status:       status,
		frames:       frames,
		errorCount:   errorCount,
		creationDate: creationDate,
	}
	return stream
}

type ActiveStream struct {
	donorFrames  float64 // number of frames done by this donor (including partial frames)
	bufferFrames int     // number of frames stored in the buffer
	authToken    string  // token of the ActiveStream
	user         string  // donor id
	startTime    int     // time the stream was activated
	frameHash    string  // md5 hash of the last frame
	engine       string  // core engine type the stream is assigned to
}

func NewActiveStream(user, token, engine string) *ActiveStream {
	as := &ActiveStream{
		user:      user,
		engine:    engine,
		authToken: token,
		startTime: int(time.Now().Unix()),
	}
	return as
}
