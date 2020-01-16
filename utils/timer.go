// Copyright 2017-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not use this file except in compliance with the License. A copy of the License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

package utils

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
	"time"
)

// Timer is the same as time.Timer except that it has jitters.
// A Timer must be created with NewTimer.
type Timer struct {
	t      *time.Timer
	d      time.Duration
	jitter time.Duration
	rand   *rand.Rand
}

// NewTimer creates a new Timer that will send the current time on its channel.
func NewTimer(d, jitter time.Duration) *Timer {
	var seed int64
	var buf [8]byte
	if _, err := crand.Read(buf[:]); err != nil {
		seed = time.Now().UnixNano()
	} else {
		seed = int64(binary.BigEndian.Uint64(buf[:]))
	}
	rnd := rand.New(rand.NewSource(seed))

	t := time.NewTimer(d - time.Duration(rnd.Int63n(int64(jitter))))

	jitteredTimer := Timer{
		t:      t,
		d:      d,
		jitter: jitter,
		rand:   rnd,
	}

	return &jitteredTimer
}

// C is channel.
func (j *Timer) C() <-chan time.Time {
	return j.t.C
}

// Reset resets the timer.
// Reset should be invoked only on stopped or expired timers with drained channels.
func (j *Timer) Reset() {
	j.t.Reset(j.d - time.Duration(j.rand.Int63n(int64(j.jitter))))
}
