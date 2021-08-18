package ratelimiter

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/logging"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

type RateLimiter struct {
	DB            *sql.DB
	GeneralLimit  int
	ResetInterval time.Duration
}

var requestCountsQuery = models.RequestCounts //nolint

func (r *RateLimiter) Allow(identifier string, ctx echo.Context) (bool, error) {
	var requestCount *models.RequestCount

	logger := logging.FromContext(ctx)

	dbCtx := boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(true, logger))
	requestCountSlice, err := requestCountsQuery(qm.Where("user_guid=?", identifier)).All(dbCtx, r.DB)
	if err != nil {
		return false, err
	}

	if len(requestCountSlice) < 1 {
		requestCount = &models.RequestCount{}
		requestCount.UserGUID = null.StringFrom(identifier)
		resetCount(requestCount, r.ResetInterval)
		err = requestCount.Insert(dbCtx, r.DB, boil.Infer())
		if err != nil {
			return false, err
		}
	} else {
		requestCount = requestCountSlice[0]
		requestCount.Count = null.IntFrom(requestCount.Count.Int + 1)
		if requestCount.ValidUntil.Time.UTC().Before(time.Now().UTC()) {
			resetCount(requestCount, r.ResetInterval)
		}
		_, err = requestCount.Update(dbCtx, r.DB, boil.Infer())
		if err != nil {
			return false, err
		}
	}

	var limitRemaining int
	if remaining := r.GeneralLimit - requestCount.Count.Int; remaining < 0 {
		limitRemaining = 0
	} else {
		limitRemaining = remaining
	}

	base := 10
	ctx.Response().Header().Set("X-RateLimit-Limit", strconv.Itoa(r.GeneralLimit))
	ctx.Response().Header().Set("X-RateLimit-Reset", strconv.FormatInt(requestCount.ValidUntil.Time.Unix(), base))
	ctx.Response().Header().Set("X-RateLimit-Remaining", strconv.Itoa(limitRemaining))

	if requestCount.Count.Int > r.GeneralLimit {
		return false, nil
	}
	return true, nil
}

func resetCount(requestCount *models.RequestCount, resetInterval time.Duration) {
	expiryTime := time.Now().UTC().Add(resetInterval)
	requestCount.ValidUntil = null.TimeFrom(expiryTime)
	requestCount.Count = null.IntFrom(1)
}

func CustomRateLimiter(rateLimiter RateLimiter) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			identifier, ok := ctx.Get("username").(string)
			if !ok {
				return fmt.Errorf("something went wrong with casting")
			}
			if allow, err := rateLimiter.Allow(identifier, ctx); !allow {
				err := &echo.HTTPError{
					Code:     http.StatusTooManyRequests,
					Message:  "Too Many Requests",
					Internal: err,
				}
				return err
			}
			return next(ctx)
		}
	}
}
