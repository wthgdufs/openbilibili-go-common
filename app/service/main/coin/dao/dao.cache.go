// Code generated by $GOPATH/src/go-common/app/tool/cache/gen. DO NOT EDIT.

/*
  Package dao is a generated cache proxy package.
  It is generated from:
  type _cache interface {
		// cache: -nullcache=-1 -check_null_code=$==-1 -singleflight=true
		UserCoin(c context.Context, mid int64) (count float64, err error)
		// cache: -nullcache=-1 -check_null_code=$==-1 -singleflight=true
		ItemCoin(c context.Context, aid int64, tp int64) (count int64, err error)
	}
*/

package dao

import (
	"context"

	"go-common/library/stat/prom"

	"golang.org/x/sync/singleflight"
)

var _ _cache
var cacheSingleFlights = [2]*singleflight.Group{{}, {}}

// UserCoin get data from cache if miss will call source method, then add to cache.
func (d *Dao) UserCoin(c context.Context, id int64) (res float64, err error) {
	addCache := true
	res, err = d.CacheUserCoin(c, id)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if res == -1 {
			res = 0
		}
	}()
	if res != 0 {
		prom.CacheHit.Incr("UserCoin")
		return
	}
	var rr interface{}
	sf := d.cacheSFUserCoin(id)
	rr, err, _ = cacheSingleFlights[0].Do(sf, func() (r interface{}, e error) {
		prom.CacheMiss.Incr("UserCoin")
		r, e = d.RawUserCoin(c, id)
		return
	})
	res = rr.(float64)
	if err != nil {
		return
	}
	miss := res
	if miss == 0 {
		miss = -1
	}
	if !addCache {
		return
	}
	d.cache.Do(c, func(c context.Context) {
		d.AddCacheUserCoin(c, id, miss)
	})
	return
}

// ItemCoin get data from cache if miss will call source method, then add to cache.
func (d *Dao) ItemCoin(c context.Context, id int64, tp int64) (res int64, err error) {
	addCache := true
	res, err = d.CacheItemCoin(c, id, tp)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if res == -1 {
			res = 0
		}
	}()
	if res != 0 {
		prom.CacheHit.Incr("ItemCoin")
		return
	}
	var rr interface{}
	sf := d.cacheSFItemCoin(id, tp)
	rr, err, _ = cacheSingleFlights[1].Do(sf, func() (r interface{}, e error) {
		prom.CacheMiss.Incr("ItemCoin")
		r, e = d.RawItemCoin(c, id, tp)
		return
	})
	res = rr.(int64)
	if err != nil {
		return
	}
	miss := res
	if miss == 0 {
		miss = -1
	}
	if !addCache {
		return
	}
	d.cache.Do(c, func(c context.Context) {
		d.AddCacheItemCoin(c, id, miss, tp)
	})
	return
}
