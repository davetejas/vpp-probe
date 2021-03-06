package vpp

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"go.ligato.io/vpp-probe/probe"
	"go.ligato.io/vpp-probe/vpp/api"
)

var (
	cliShowVersionVerbose    = regexp.MustCompile(`Version:\s+(\S+)`)
	cliShowVersionVerbosePID = regexp.MustCompile(`PID:\s+([0-9]+)`)
)

// vpp# show version verbose
// Version:                  v20.09-rc0~399-gef80ad6bf~b1658
// Compiled by:              root
// Compile host:             31cb557be35c
// Compile date:             2020-09-09T11:13:09
// Compile location:         /w/workspace/vpp-merge-master-ubuntu1804
// Compiler:                 Clang/LLVM 9.0.0 (tags/RELEASE_900/final)
// Current PID:              170

func GetVersionInfoCLI(cli probe.CliExecutor) (*api.VersionInfo, error) {
	out, err := cli.RunCli("show version verbose")
	if err != nil {
		return nil, err
	}

	info := &api.VersionInfo{}

	matchVersion := cliShowVersionVerbose.FindStringSubmatch(out)
	if len(matchVersion) > 1 {
		info.Version = matchVersion[1]
	}

	matchPid := cliShowVersionVerbosePID.FindStringSubmatch(out)
	if len(matchPid) > 1 {
		info.Pid, _ = strconv.Atoi(matchPid[1])
	}

	return info, nil
}

var (
	cliShowClock = regexp.MustCompile(`Time\s+now\s+([0-9\.]+),\s+(\S+)`)
)

// vpp# show click
// Time now 3180.278756, Tue, 1 Dec 2020 11:52:45 GMT

func GetUptimeCLI(cli probe.CliExecutor) (time.Duration, error) {
	out, err := cli.RunCli("show clock")
	if err != nil {
		return 0, err
	}

	var uptime time.Duration

	matchUptime := cliShowClock.FindStringSubmatch(out)
	if len(matchUptime) > 1 {
		rawUptime := matchUptime[1]
		floatUptime, err := strconv.ParseFloat(rawUptime, 64)
		if err != nil {
			logrus.Debugf("parse float %v error: %v", rawUptime, err)
		} else {
			uptime = time.Duration(floatUptime * float64(time.Second))
		}
	}

	return uptime, nil
}

var (
// TODO: parse log entries from CLI
//cliShowLog = regexp.MustCompile(``)
)

// vpp# show log
// 2020/12/01 10:59:44:837 notice     plugin/load    Loaded plugin: abf_plugin.so (Access Control List (ACL) Based Forwarding)
// 2020/12/01 10:59:44:841 notice     plugin/load    Loaded plugin: acl_plugin.so (Access Control Lists (ACL))
// 2020/12/01 10:59:44:841 notice     plugin/load    Loaded plugin: adl_plugin.so (Allow/deny list plugin)
// 2020/12/01 10:59:44:843 notice     plugin/load    Loaded plugin: avf_plugin.so (Intel Adaptive Virtual Function (AVF) Device
// ...

func DumpLogsCLI(cli probe.CliExecutor) ([]string, error) {
	out, err := cli.RunCli("show log")
	if err != nil {
		return nil, err
	}
	logs := strings.Split(out, "\n")
	return logs, nil
}
