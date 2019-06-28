package tester

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"github.com/panjiang/redisbench/models"
	"github.com/panjiang/redisbench/utils"
)

// RPC : RPC class between nodes
type RPC int

// Start : Master notices others nodes start testing
func (t *RPC) Start(order int, reply *int) error {
	if order != 1 {
		log.Fatalln("invalid master order")
	}
	Wg.Done()
	return nil
}

// Settle : Others nodes notice master result for settling
func (t *RPC) Settle(result *models.NodeResult, reply *int) error {
	Multi.NodeSettle(result)
	return nil
}

// RPCRun : Register rpc handler, and connect to related nodes
func RPCRun() {
	var err error
	Multi, err = NewMultiTester()
	utils.FatalErr(err)

	// All nodes handle RPC,
	// for receiving start sign
	rpc.Register(new(RPC))
	rpc.HandleHTTP()

	ln, err := net.Listen("tcp", Multi.Addr)
	utils.FatalErr(err)

	Wg.Add(1) // Wait start sign from master
	go http.Serve(ln, nil)
	log.Printf("node listened %s#%d", Multi.Addr, Multi.Order)

	Multi.MustConnectToNodes()

	if Multi.IsMaster() {
		// Master connect to all nodes,
		// then notice all start testing
		Wg.Done() // Don't need wait rpc connect
		Multi.NoticeNodesToStart()
	}
}
