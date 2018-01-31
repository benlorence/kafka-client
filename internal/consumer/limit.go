// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package consumer

import "time"

const (
	noLimit = -2
	// defaultLimit set to -1 so that if you use default limit, no messages will be processed.
	defaultLimit              = -1
	defaultLimitCheckInterval = time.Second
)

type topicPartitionLimitMap struct {
	checkInterval time.Duration
	limits        map[TopicPartition]int64
}

func newTopicLimitMap(limit map[TopicPartition]int64) topicPartitionLimitMap {
	return topicPartitionLimitMap{
		limits:        limit,
		checkInterval: defaultLimitCheckInterval,
	}
}

// HasLimits returns true if there are limits set.
func (m *topicPartitionLimitMap) HasLimits() bool {
	return m.limits != nil
}

// Get returns noLimit if there are no limits set.
// If there are limits set but no limits for this topic partition, then defaultLimit will be used.
// Else, it will return the limit stored in the limits map.
func (m *topicPartitionLimitMap) Get(tp TopicPartition) int64 {
	if m.limits == nil {
		return noLimit
	}

	limit, ok := m.limits[tp]
	if !ok {
		return defaultLimit
	}

	return limit
}
