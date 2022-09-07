// Copyright (c) 2021 Tigera, Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tcdefs

const (
	MarkSeen                         = 0x01000000
	MarkSeenMask                     = MarkSeen
	MarkSeenBypass                   = MarkSeen | 0x02000000
	MarkSeenBypassMask               = MarkSeenMask | MarkSeenBypass
	MarkSeenFallThrough              = MarkSeen | 0x04000000
	MarkSeenFallThroughMask          = MarkSeenMask | MarkSeenFallThrough
	MarkSeenBypassForward            = MarkSeenBypass | 0x00300000
	MarkSeenBypassForwardMask        = MarkSeenBypassMask | 0x00f00000
	MarkSeenBypassForwardSourceFixup = MarkSeenBypass | 0x00500000
	MarkSeenNATOutgoing              = MarkSeenBypass | 0x00800000
	MarkSeenNATOutgoingMask          = MarkSeenBypassMask | 0x00f00000
	MarkSeenMASQ                     = MarkSeenBypass | 0x00600000
	MarkSeenMASQMask                 = MarkSeenBypassMask | 0x00f00000
	MarkSeenSkipFIB                  = MarkSeen | 0x00100000

	MarkLinuxConntrackEstablished     = 0x08000000
	MarkLinuxConntrackEstablishedMask = 0x08000000

	MarkSeenToNatIfaceOut   = 0x41000000
	MarkSeenFromNatIfaceOut = 0x81000000

	MarksMask uint32 = 0x1ff00000
)

const (
	ProgIndexPolicy = iota
	ProgIndexAllowed
	ProgIndexIcmp
	ProgIndexDrop
	ProgIndexHostCtConflict
	ProgIndexV6Prologue
	ProgIndexV6Policy
	ProgIndexV6Icmp
	ProgIndexV6Drop
)

var ProgramNames = []string{
	"calico_tc_norm_pol_tail",
	"calico_tc_skb_accepted_entrypoint",
	"calico_tc_skb_send_icmp_replies",
	"calico_tc_skb_drop",
	"calico_tc_host_ct_conflict",
	"calico_tc_v6",
	"calico_tc_v6_norm_pol_tail",
	"calico_tc_v6_skb_accepted_entrypoint",
	"calico_tc_v6_skb_send_icmp_replies",
	"calico_tc_v6_skb_drop",
}

var JumpMapIndexes = map[string][]int{
	"IPv4": []int{
		ProgIndexPolicy,
		ProgIndexAllowed,
		ProgIndexIcmp,
		ProgIndexDrop,
		ProgIndexHostCtConflict,
	},
	"IPv6": []int{
		ProgIndexV6Prologue,
		ProgIndexV6Policy,
		ProgIndexV6Icmp,
		ProgIndexV6Drop,
	},
}
