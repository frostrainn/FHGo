package registry

import (
	"sort"
	"time"
)

func (r *ZRegistry) putServer(path, addr string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	s := r.servers[path]
	if s == nil || len(s) == 0 {

		r.servers[path] = []*ServerItem{}
		r.servers[path] = append(r.servers[path], &ServerItem{Addr: addr, start: time.Now()})
	} else if i := getIndex(r.servers[path], addr); i >= 0 {
		// if exists, update start time to keep alive
		s[i].start = time.Now()
	} else {
		// if not exists, append to slice
		r.servers[path] = append(r.servers[path], &ServerItem{Addr: addr, start: time.Now()})
	}
}

func (r *ZRegistry) aliveServers(path string) []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	var alive []string
	var dead []string
	for index, s := range r.servers[path] {
		if r.timeout == 0 || s.start.Add(r.timeout).After(time.Now()) {
			alive = append(alive, r.servers[path][index].Addr)
		} else {
			// delete the timeout server
			dead = append(dead, r.servers[path][index].Addr)
		}
	}
	sort.Strings(alive)
	r.servers[path] = removeAddrs(r.servers[path], dead)
	return alive
}

func getIndex(items []*ServerItem, item string) int {
	for index, eachItem := range items {
		if eachItem.Addr == item {
			return index
		}
	}
	return -1
}

func removeAddr(items []*ServerItem, item string) []*ServerItem {
	for index, eachItem := range items {
		if eachItem.Addr == item {
			return append(items[:index], items[index+1:]...)
		}
	}
	return items
}

func removeAddrs(items []*ServerItem, addrs []string) []*ServerItem {
	for _, addr := range addrs {
		items = removeAddr(items, addr)
	}
	return items
}
