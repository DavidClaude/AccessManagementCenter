package codex

import (
	"errors"
	"strconv"
	"regexp"
	"generic-op/utils"
)

type Args struct {
	ArgMap map[string]string
}

func NewArgs(osArgs []string) (args *Args) {
	m := make(map[string]string)
	args = &Args{
		ArgMap: m,
	}
	args.Parse(osArgs)
	return args
}

func (args *Args) Parse(osArgs []string) () {
	l := len(osArgs)
	for i, v := range osArgs {
		b, _ := IsOpt(v)
		if b { // if v is option
			b1, _ := IsOpt(osArgs[(i+1)%l]) // the next elem of v is option ?
			if b1 { // if yes, return empty
				args.ArgMap[utils.StringKick(v, "-")] = ""
			} else { //if no, return this elem as v's arg
				args.ArgMap[utils.StringKick(v, "-")] = osArgs[(i+1)%l]
			}
		}
	}
}

func (args *Args) GetBool(opt string) (b bool, err error) {
	if _, ok := args.ArgMap[opt]; !ok {
		return false, errors.New("No such option")
	}
	if args.ArgMap[opt] == "0" {
		return false, nil
	} else {
		return true, nil
	}
}

func (args *Args) GetInt(opt string) (i int, err error) {
	if _, ok := args.ArgMap[opt]; !ok {
		return 0, errors.New("No such option")
	}
	i, _ = strconv.Atoi(args.ArgMap[opt])
	return i, nil
}

func (args *Args) GetString(opt string) (s string, err error) {
	if _, ok := args.ArgMap[opt]; !ok {
		return "", errors.New("No such option")
	}
	return args.ArgMap[opt], nil
}

func IsOpt(s string) (b bool, err error) {
	match, err := regexp.MatchString("^-+(.+)", s)
	if err != nil {
		return false, err
	}
	return match, nil
}
