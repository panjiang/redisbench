package tester

import (
	"errors"
	"log"
	"net/rpc"
	"redisbench/config"
	"redisbench/models"
	"redisbench/utils"
	"strings"
	"time"
)

// MasterNodeOrder : The order of master node's
const MasterNodeOrder int = 1

// MultiTester : Multiple testers class
type MultiTester struct {
	Order   int                        // Current order
	Addr    string                     // Current address
	Addrs   map[int]string             // All addresses
	Nodes   map[int]*rpc.Client        // Registered orders
	Results map[int]*models.NodeResult // Every tester result
}

// IsMaster : master receive all nodes connection
// And notice them to do some actions
func (mt *MultiTester) IsMaster() bool {
	return mt.Order == MasterNodeOrder
}

func (mt *MultiTester) connectToNodes() error {
	// Register to others nodes blocking
	for order, client := range mt.Nodes {
		if client != nil {
			continue
		}
		addr := mt.Addrs[order]
		// log.Printf("connect to node: %s#%d", addr, order)
		client, err := rpc.DialHTTP("tcp", addr)
		if err != nil {
			return err
		}
		mt.Nodes[order] = client
		log.Printf("connected to node: %s#%d", addr, order)
	}
	return nil
}

// MustConnectToNodes : Master connects to all others nodes
func (mt *MultiTester) MustConnectToNodes() {
	for {
		err := mt.connectToNodes()
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}
}

// NoticeNodesToStart : Master notices all nodes to start
func (mt *MultiTester) NoticeNodesToStart() {
	for _, client := range mt.Nodes {
		err := client.Call("RPC.Start", MasterNodeOrder, nil)
		utils.FatalErr(err)
	}
}

// NoticeMasterSettle : Nodes notice master settle
func (mt *MultiTester) NoticeMasterSettle(result *models.NodeResult) {
	client := mt.Nodes[1]
	err := client.Call("RPC.Settle", result, nil)
	utils.FatalErr(err)
}

// NodeSettle : One node settle method
func (mt *MultiTester) NodeSettle(result *models.NodeResult) {
	log.Printf("node settle: %s#%d SUM:%d DUR:%.3fs", mt.Addrs[result.Order], result.Order, result.TotalTimes, float64(result.TotalDur)/1000)
	mt.Results[result.Order] = result

	for _, result := range mt.Results {
		if result == nil {
			return
		}
	}

	Wg.Done()
}

// Summary : Summary all nodes' results after them run over
func (mt *MultiTester) Summary() *models.SummaryResult {
	summary := new(models.SummaryResult)

	tsMin := mt.Results[MasterNodeOrder].TsBeg
	tsMax := mt.Results[MasterNodeOrder].TsEnd
	for _, result := range mt.Results {
		summary.TotalTimes += result.TotalTimes
		if result.TsBeg < tsMin {
			tsMin = result.TsBeg
		}
		if result.TsEnd > tsMax {
			tsMax = result.TsEnd
		}
	}
	summary.TotalDur = tsMax - tsMin
	summary.TPS = int(summary.TotalTimes / (summary.TotalDur / 1000.0))
	return summary
}

// NewMultiTester : Create a new MultiTester pointer
func NewMultiTester() (*MultiTester, error) {
	if config.MultiAddr == "" {
		return nil, errors.New("invalid multi addresses has been set")
	}
	addrsArr := strings.Split(config.MultiAddr, ",")

	if config.MultiOrder <= 0 || config.MultiOrder > len(addrsArr) {
		return nil, errors.New("invalid order while multi test has been set")
	}

	multi := &MultiTester{Order: config.MultiOrder}
	multi.Addr = addrsArr[config.MultiOrder-1]
	multi.Addrs = make(map[int]string)
	multi.Nodes = make(map[int]*rpc.Client)
	multi.Results = make(map[int]*models.NodeResult)
	for i, addr := range addrsArr {
		multi.Addrs[i+1] = addr
		multi.Results[i+1] = nil
		if i+1 == multi.Order {
			continue
		}

		// Others nodes only connect master node
		if !multi.IsMaster() && i+1 != MasterNodeOrder {
			continue
		}

		multi.Nodes[i+1] = nil
	}
	return multi, nil
}
