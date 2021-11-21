package env

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func GetString(n string, def string) string {
	if envstr, ok := os.LookupEnv(n); ok {
		return envstr
	}
	return def
}

func GetUint(n string, def uint) uint {
	if envstr, ok := os.LookupEnv(n); ok {
		ret, err := strconv.ParseUint(envstr, 0, 0)
		if err != nil {
			panic(fmt.Errorf("env %q: failed to parse value %q as an unsigned integer: %w", n, envstr, err))
		}
		return uint(ret)
	}
	return def
}

func GetRegexp(n string, def string) *regexp.Regexp {
	if envstr, ok := os.LookupEnv(n); ok {
		ret, err := regexp.Compile(envstr)
		if err != nil {
			panic(fmt.Errorf("env %q: failed to parse value %q as a regex: %w", n, envstr, err))
		}
		return ret
	}
	return regexp.MustCompile(def)
}

func GetDuration(n string, def time.Duration) time.Duration {
	if envstr, ok := os.LookupEnv(n); ok {
		ret, err := time.ParseDuration(envstr)
		if err != nil {
			panic(fmt.Errorf("env %q: failed to parse value %q as a duration: %w", n, envstr, err))
		}
		return ret
	}
	return def
}
