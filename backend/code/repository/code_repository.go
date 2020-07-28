package repository

import (
	"fmt"
	"math"

	"github.com/eyalch/pipeit/backend/domain"
	"github.com/gomodule/redigo/redis"
)

const (
	availableSetKey = "codes:available"
	activeSetKey    = "codes:active"
)

// codeCount represents the count of all the possible permutations of the code
var codeCount int = int(math.Pow10(domain.CodeLength))

type codeRepository struct {
	redisConn redis.Conn
}

// New creates a new CodeRepository and creates the codes
func New(redisConn redis.Conn) domain.CodeRepository {
	return &codeRepository{redisConn}
}

// checkIfNeedToCreateCodes checks how many codes are there in total (available + active)
// and returns whether the count is equal to the needed count.
func (r *codeRepository) checkIfNeedToCreateCodes() (bool, error) {
	// Check how many codes are there in total
	r.redisConn.Send("SCARD", availableSetKey)
	r.redisConn.Send("SCARD", activeSetKey)
	reply, err := redis.Values(r.redisConn.Do("EXEC"))
	if err != nil {
		return false, err
	}
	var availableCodesCount, activeCodesCount int
	_, err = redis.Scan(reply, &availableCodesCount, &activeCodesCount)
	if err != nil {
		return false, err
	}
	totalCodesCount := availableCodesCount + activeCodesCount

	// If the total codes count doesn't match the required count, then we need
	// to re-create the sets
	return totalCodesCount != codeCount, nil
}

// CreateCodesIfNeeded deletes the Redis Sets and re-creates them by adding all
// the possible codes.
func (r *codeRepository) CreateCodesIfNeeded() error {
	// Check if we need to (re-)create the codes
	if need, err := r.checkIfNeedToCreateCodes(); !need || err != nil {
		return err
	}

	// Delete the sets
	r.redisConn.Send("DEL", availableSetKey)
	r.redisConn.Send("DEL", activeSetKey)

	// Fill the sets
	for i := 0; i < codeCount; i++ {
		// The format of the code should have a length of code.Length and be
		// padded with zeros. For example, if code.Length is 4, then the
		// format would be: "%04d"
		codeFormat := fmt.Sprintf("%%0%dd", domain.CodeLength)

		code := fmt.Sprintf(codeFormat, i)
		r.redisConn.Send("SADD", availableSetKey, code)
	}

	// Complete the transaction
	_, err := r.redisConn.Do("EXEC")
	return err
}

// GetRandomCode pops a random code from the available codes Redis Set, moves it
// to the active codes Redis Set, and returns it
func (r *codeRepository) GetRandomCode() (string, error) {
	code, err := redis.String(r.redisConn.Do("SPOP", availableSetKey))
	if err != nil {
		return "", err
	}

	_, err = r.redisConn.Do("SADD", activeSetKey, code)

	return code, err
}
