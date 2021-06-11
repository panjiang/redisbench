package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/panjiang/redisbench/config"
	"github.com/panjiang/redisbench/internal/datasize"
	"github.com/panjiang/redisbench/models"
	"github.com/panjiang/redisbench/tester"
	"github.com/panjiang/redisbench/utils"
	"github.com/panjiang/redisbench/wares"

	"github.com/go-redis/redis/v8"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

const (
	keyPrefix = "benchmark-set"
)

func executeSet(id int, times int, size int, redisClient redis.UniversalClient) {
	defer tester.Wg.Done()
	val := utils.RandSeq(size)
	var err error
	for i := 0; i < times; i++ {
		key := fmt.Sprintf("%s.%d.%d", keyPrefix, id, i)
		err = redisClient.Set(context.Background(), key, val, 0).Err()
		utils.FatalErr(err)
	}
}

func executeDel(id int, times int, redisClient redis.UniversalClient) {
	defer tester.Wg.Done()
	for i := 0; i < times; i++ {
		key := fmt.Sprintf("%s.%d.%d", keyPrefix, id, i)
		redisClient.Del(context.Background(), key)
	}
}

func main() {
	// Parse config arguments from command-line
	config.Parse()
	if config.MultiAddr != "" {
		tester.RPCRun()
	}

	tester.Wg.Wait()

	// Print test initial information
	totalTimes := int64(config.ClientNum * config.TestTimes)
	totalSize := datasize.ByteSize(config.ClientNum * config.TestTimes * config.DataSize)
	log.Info().Str("addr", config.RedisAddr).Msg("Redis")

	log.Info().
		Int("clientNum", config.ClientNum).
		Int("testTimes", config.TestTimes).
		Stringer("dataSize", datasize.ByteSize(config.DataSize)).
		Msg("Config")

	log.Info().
		Int64("times", totalTimes).
		Stringer("size", totalSize).
		Msg("Total")

	// Create a new redis client
	redisClient, err := wares.NewUniversalRedisClient()
	utils.FatalErr(err)

	// Run certain number clients for testing
	log.Info().Msg("Testing...")
	t1 := time.Now()
	for i := 0; i < config.ClientNum; i++ {
		tester.Wg.Add(1)
		go executeSet(i, config.TestTimes, config.DataSize, redisClient)
	}
	tester.Wg.Wait()
	t2 := time.Now()

	// Calculate the duration
	dur := t2.Sub(t1)
	order := 1
	if tester.Multi != nil {
		order = tester.Multi.Order
	}
	result := &models.NodeResult{Order: order, TotalTimes: totalTimes, TsBeg: t1, TsEnd: t2, TotalDur: dur}
	if tester.Multi == nil {
		log.Info().
			Int64("times", result.TotalTimes).
			Stringer("duration", result.TotalDur).
			Int64("tps", tester.CalTps(result.TotalTimes, result.TotalDur)).
			Msg("* Result")

	} else {
		if !tester.Multi.IsMaster() {
			// Notice master to settle
			tester.Multi.NoticeMasterSettle(result)
			log.Info().Msg("* See summary info on node 1")
		} else {
			tester.Wg.Add(1) // Wait all others nodes settling call
			tester.Multi.NodeSettle(result)

			tester.Wg.Wait()
			time.Sleep(time.Second)
			// Summary all nodes result include self
			summary := tester.Multi.Summary()

			// Print testing result
			log.Info().
				Int64("times", summary.TotalTimes).
				Stringer("duration", summary.TotalDur).
				Int("tps", summary.TPS).
				Msg("* Summary")
		}
	}

	log.Debug().Msg("Deleting testing data...")
	for i := 0; i < config.ClientNum; i++ {
		tester.Wg.Add(1)
		go executeDel(i, config.TestTimes, redisClient)
	}
	tester.Wg.Wait()
	log.Debug().Msg("Over")
}
