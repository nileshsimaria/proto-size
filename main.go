package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	gnmi "github.com/nileshsimaria/proto-size/gnmi"
	kv_pb "github.com/nileshsimaria/proto-size/kv"
	ytop_pb "github.com/nileshsimaria/proto-size/ytop"
	"log"
)

func main() {

	state := &ytop_pb.NetworkInstancesNetworkInstanceListMplsTypeSignalingProtocolsTypeRsvpTeTypeSessionsTypeSessionListRecordRouteObjectsTypeRecordRouteObjectListStateType{
		Index:         1,
		Address:       "64.233.160.3",
		ReportedLabel: "NO_LABEL",
		ReportedFlags: 32,
	}
	rr := make([]*ytop_pb.NetworkInstancesNetworkInstanceListMplsTypeSignalingProtocolsTypeRsvpTeTypeSessionsTypeSessionListRecordRouteObjectsTypeRecordRouteObjectList, 1)
	rr[0] = &ytop_pb.NetworkInstancesNetworkInstanceListMplsTypeSignalingProtocolsTypeRsvpTeTypeSessionsTypeSessionListRecordRouteObjectsTypeRecordRouteObjectList{Index: 1, State: state}
	rro := &ytop_pb.NetworkInstancesNetworkInstanceListMplsTypeSignalingProtocolsTypeRsvpTeTypeSessionsTypeSessionListRecordRouteObjectsType{RecordRouteObject: rr}

	session := make([]*ytop_pb.NetworkInstancesNetworkInstanceListMplsTypeSignalingProtocolsTypeRsvpTeTypeSessionsTypeSessionList, 1)
	session[0] = &ytop_pb.NetworkInstancesNetworkInstanceListMplsTypeSignalingProtocolsTypeRsvpTeTypeSessionsTypeSessionList{LocalIndex: uint64(138996), RecordRouteObjects: rro}
	sessions := &ytop_pb.NetworkInstancesNetworkInstanceListMplsTypeSignalingProtocolsTypeRsvpTeTypeSessionsType{Session: session}

	rsvpte := &ytop_pb.NetworkInstancesNetworkInstanceListMplsTypeSignalingProtocolsTypeRsvpTeType{Sessions: sessions}
	sp := &ytop_pb.NetworkInstancesNetworkInstanceListMplsTypeSignalingProtocolsType{RsvpTe: rsvpte}

	mpls := &ytop_pb.NetworkInstancesNetworkInstanceListMplsType{SignalingProtocols: sp}
	ni := make([]*ytop_pb.NetworkInstancesNetworkInstanceList, 1)
	ni[0] = &ytop_pb.NetworkInstancesNetworkInstanceList{InstanceName: "master", Mpls: mpls}

	ytopData := &ytop_pb.NetworkInstances{}
	ytopData.NetworkInstance = ni

	ytopSData, err := proto.Marshal(ytopData)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Printf("ytop size : %d\n", len(ytopSData))
	//fmt.Printf("%v\n", ytopData)

	dd := []string{
		"__prefix__", "/network-instances/network-instance[instance-name='master']/mpls/signaling-protocols/rsvp-te/sessions/session[local-index='138996']/record-route-objects/record-route-object[index='1']/",
		"index", "1",
		"state/index", "1",
		"state/address", "64.233.160.3",
		"state/reported-label", "NO_LABEL",
		"state/reported-flags", "32",
	}

	//kv := make([]*kv_pb.KeyValue, 0, 0)
	kv := []*kv_pb.KeyValue{
		{Key: dd[0], Value: &kv_pb.KeyValue_StrValue{StrValue: dd[1]}},
		{Key: dd[2], Value: &kv_pb.KeyValue_UintValue{UintValue: 1}},
		{Key: dd[4], Value: &kv_pb.KeyValue_UintValue{UintValue: 1}},
		{Key: dd[6], Value: &kv_pb.KeyValue_StrValue{StrValue: dd[7]}},
		{Key: dd[8], Value: &kv_pb.KeyValue_StrValue{StrValue: dd[9]}},
		{Key: dd[10], Value: &kv_pb.KeyValue_UintValue{UintValue: 32}},
	}

	ocData := &kv_pb.OpenConfigData{
		Timestamp: 1510946604929,
		Kv:        kv,
	}

	data, err := proto.Marshal(ocData)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Printf("Multi-level OC prefix size = %d\n", len(data))

	dd1 := []string{
		"__prefix__", "/network-instances/network-instance[instance-name='master']",
		"mpls/signaling-protocols/rsvp-te/sessions/session[local-index='138996']/record-route-objects/record-route-object[index='1']/index", "1",
		"mpls/signaling-protocols/rsvp-te/sessions/session[local-index='138996']/record-route-objects/record-route-object[index='1']/state/index", "1",
		"mpls/signaling-protocols/rsvp-te/sessions/session[local-index='138996']/record-route-objects/record-route-object[index='1']/state/address", "64.233.160.3",
		"mpls/signaling-protocols/rsvp-te/sessions/session[local-index='138996']/record-route-objects/record-route-object[index='1']/state/reported-label", "NO_LABEL",
		"mpls/signaling-protocols/rsvp-te/sessions/session[local-index='138996']/record-route-objects/record-route-object[index='1']/state/reported-flags", "32",
	}
	kv1 := []*kv_pb.KeyValue{
		{Key: dd1[0], Value: &kv_pb.KeyValue_StrValue{StrValue: dd1[1]}},
		{Key: dd1[2], Value: &kv_pb.KeyValue_UintValue{UintValue: 1}},
		{Key: dd1[4], Value: &kv_pb.KeyValue_UintValue{UintValue: 1}},
		{Key: dd1[6], Value: &kv_pb.KeyValue_StrValue{StrValue: dd1[7]}},
		{Key: dd1[8], Value: &kv_pb.KeyValue_StrValue{StrValue: dd1[9]}},
		{Key: dd1[10], Value: &kv_pb.KeyValue_UintValue{UintValue: 32}},
	}

	ocData1 := &kv_pb.OpenConfigData{
		Timestamp: 1510946604929,
		Kv:        kv1,
	}

	data1, err := proto.Marshal(ocData1)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Printf("single-level prefix compression size = %d\n", len(data1))

	dd2 := []string{
		"/network-instances/network-instance[instance-name='master']/mpls/signaling-protocols/rsvp-te/sessions/session[local-index='138996']/record-route-objects/record-route-object[index='1']/index", "1",
		"/network-instances/network-instance[instance-name='master']/mpls/signaling-protocols/rsvp-te/sessions/session[local-index='138996']/record-route-objects/record-route-object[index='1']/state/index", "1",
		"/network-instances/network-instance[instance-name='master']/mpls/signaling-protocols/rsvp-te/sessions/session[local-index='138996']/record-route-objects/record-route-object[index='1']/state/address", "64.233.160.3",
		"/network-instances/network-instance[instance-name='master']/mpls/signaling-protocols/rsvp-te/sessions/session[local-index='138996']/record-route-objects/record-route-object[index='1']/state/reported-label", "NO_LABEL",
		"/network-instances/network-instance[instance-name='master']/mpls/signaling-protocols/rsvp-te/sessions/session[local-index='138996']/record-route-objects/record-route-object[index='1']/state/reported-flags", "32",
	}
	kv2 := []*kv_pb.KeyValue{
		{Key: dd2[0], Value: &kv_pb.KeyValue_UintValue{UintValue: 1}},
		{Key: dd2[2], Value: &kv_pb.KeyValue_UintValue{UintValue: 1}},
		{Key: dd2[4], Value: &kv_pb.KeyValue_StrValue{StrValue: dd2[5]}},
		{Key: dd2[6], Value: &kv_pb.KeyValue_StrValue{StrValue: dd2[7]}},
		{Key: dd2[8], Value: &kv_pb.KeyValue_UintValue{UintValue: 32}},
	}

	ocData2 := &kv_pb.OpenConfigData{
		Timestamp: 1510946604929,
		Kv:        kv2,
	}

	data2, err := proto.Marshal(ocData2)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Printf("no comparession size = %d\n", len(data2))

	gnmiData := &gnmi.Notification{
		Timestamp: 1510946604929,
		Prefix: &gnmi.Path{
			Origin: "openconfig",
			Elem: []*gnmi.PathElem{
				{Name: "network-instances"},
				{Name: "network-instance",
					Key: map[string]string{
						"key":   "instance-name",
						"value": "master",
					},
				},
				{Name: "mpls"},
				{Name: "signaling-protocols"},
				{Name: "rsvp-te"},
				{Name: "sessions"},
				{Name: "session",
					Key: map[string]string{
						"key":   "local-index",
						"value": "138996",
					},
				},
				{Name: "record-route-objects"},
				{Name: "record-route-object",
					Key: map[string]string{
						"key":   "index",
						"value": "1",
					},
				},
			},
		},
		Update: []*gnmi.Update{
			{
				Path: &gnmi.Path{
					//Origin: "openconfig",
					Elem: []*gnmi.PathElem{
						{Name: "index"},
					},
				},
				Val: &gnmi.TypedValue{Value: &gnmi.TypedValue_UintVal{UintVal: 1}},
			},
			{
				Path: &gnmi.Path{
					//Origin: "openconfig",
					Elem: []*gnmi.PathElem{
						{Name: "state/index"},
					},
				},
				Val: &gnmi.TypedValue{Value: &gnmi.TypedValue_UintVal{UintVal: 1}},
			},
			{
				Path: &gnmi.Path{
					//Origin: "openconfig",
					Elem: []*gnmi.PathElem{
						{Name: "state/address"},
					},
				},
				Val: &gnmi.TypedValue{Value: &gnmi.TypedValue_StringVal{StringVal: "64.233.160.3"}},
			},
			{
				Path: &gnmi.Path{
					//Origin: "openconfig",
					Elem: []*gnmi.PathElem{
						{Name: " state/reported-label"},
					},
				},
				Val: &gnmi.TypedValue{Value: &gnmi.TypedValue_StringVal{StringVal: "NO_LABEL"}},
			},
			{
				Path: &gnmi.Path{
					//Origin: "openconfig",
					Elem: []*gnmi.PathElem{
						{Name: " state/reported-flags"},
					},
				},
				Val: &gnmi.TypedValue{Value: &gnmi.TypedValue_UintVal{UintVal: 32}},
			},
		},
	}

	data3, err := proto.Marshal(gnmiData)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Printf("gnmi size = %d\n", len(data3))
	//fmt.Printf("\n%v\n", gnmiData)

}
