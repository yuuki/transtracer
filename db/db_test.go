package db

import (
	"net"
	"testing"
	"time"

	"github.com/lib/pq"
	"github.com/yuuki/lstf/tcpflow"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateSchema(t *testing.T) {
	db, mock := NewTestDB()
	defer db.Close()

	mock.ExpectExec("CREATE TYPE (.+)").WillReturnResult(sqlmock.NewResult(1, 1))

	err := db.CreateSchema()
	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestInsertOrUpdateHostFlows_empty(t *testing.T) {
	db, mock := NewTestDB()
	defer db.Close()

	flows := []*tcpflow.HostFlow{}

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO nodes")
	mock.ExpectPrepare("SELECT node_id FROM nodes")
	mock.ExpectPrepare("INSERT INTO flows")
	mock.ExpectCommit()

	err := db.InsertOrUpdateHostFlows(flows)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestInsertOrUpdateHostFlows(t *testing.T) {
	db, mock := NewTestDB()
	defer db.Close()

	flows := []*tcpflow.HostFlow{
		{
			Direction:   tcpflow.FlowActive,
			Local:       &tcpflow.AddrPort{Addr: "10.0.10.1", Port: "many"},
			Peer:        &tcpflow.AddrPort{Addr: "10.0.10.2", Port: "5432"},
			Process:     &tcpflow.Process{Pgid: 1001, Name: "python"},
			Connections: 10,
		},
		{
			Direction:   tcpflow.FlowPassive,
			Local:       &tcpflow.AddrPort{Addr: "10.0.10.1", Port: "80"},
			Peer:        &tcpflow.AddrPort{Addr: "10.0.10.2", Port: "many"},
			Process:     &tcpflow.Process{Pgid: 1002, Name: "nginx"},
			Connections: 12,
		},
	}

	mock.ExpectBegin()
	stmt1 := mock.ExpectPrepare("INSERT INTO nodes")
	mock.ExpectPrepare("SELECT node_id FROM nodes")
	stmt3 := mock.ExpectPrepare("INSERT INTO flows")

	// first loop
	stmt1.ExpectQuery().WithArgs("10.0.10.1", 0, 1001, "python").WillReturnRows(sqlmock.NewRows([]string{"node_id"}).AddRow(1))
	stmt1.ExpectQuery().WithArgs("10.0.10.2", 5432, 0, "").WillReturnRows(sqlmock.NewRows([]string{"node_id"}).AddRow(2))
	stmt3.ExpectExec().WithArgs("active", 1, 2, 10).WillReturnResult(sqlmock.NewResult(1, 1))

	// second loop
	stmt1.ExpectQuery().WithArgs("10.0.10.1", 80, 1002, "nginx").WillReturnRows(sqlmock.NewRows([]string{"node_id"}).AddRow(3))
	stmt1.ExpectQuery().WithArgs("10.0.10.2", 0, 0, "").WillReturnRows(sqlmock.NewRows([]string{"node_id"}).AddRow(4))
	stmt3.ExpectExec().WithArgs("passive", 4, 3, 12).WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	err := db.InsertOrUpdateHostFlows(flows)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestInsertOrUpdateHostFlows_empty_process(t *testing.T) {
	db, mock := NewTestDB()
	defer db.Close()

	flows := []*tcpflow.HostFlow{
		{
			Direction:   tcpflow.FlowActive,
			Local:       &tcpflow.AddrPort{Addr: "10.0.10.1", Port: "many"},
			Peer:        &tcpflow.AddrPort{Addr: "10.0.10.2", Port: "5432"},
			Process:     nil,
			Connections: 10,
		},
	}

	mock.ExpectBegin()
	stmt1 := mock.ExpectPrepare("INSERT INTO nodes")
	mock.ExpectPrepare("SELECT node_id FROM nodes")
	stmt3 := mock.ExpectPrepare("INSERT INTO flows")

	// first loop
	stmt1.ExpectQuery().WithArgs("10.0.10.1", 0, 0, "").WillReturnRows(sqlmock.NewRows([]string{"node_id"}).AddRow(1))
	stmt1.ExpectQuery().WithArgs("10.0.10.2", 5432, 0, "").WillReturnRows(sqlmock.NewRows([]string{"node_id"}).AddRow(2))
	stmt3.ExpectExec().WithArgs("active", 1, 2, 10).WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	err := db.InsertOrUpdateHostFlows(flows)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFindListeningPortsByAddrs(t *testing.T) {
	db, mock := NewTestDB()
	defer db.Close()

	straddrs := pq.Array([]string{"192.0.2.1", "192.0.2.2"})
	columns := sqlmock.NewRows([]string{"ipv4", "port"})
	mock.ExpectQuery("SELECT ipv4, port FROM nodes").WithArgs(straddrs).WillReturnRows(columns.AddRow("192.0.2.1", 80).AddRow("192.0.2.2", 443))

	addrs := []net.IP{
		net.ParseIP("192.0.2.1"),
		net.ParseIP("192.0.2.2"),
	}
	portsbyaddr, err := db.FindListeningPortsByAddrs(addrs)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(portsbyaddr) != 2 {
		t.Errorf("portsbyaddr should be 2, but %v", len(portsbyaddr))
	}
	if ports, ok := portsbyaddr["192.0.2.1"]; !ok || ports[0] != 80 {
		t.Errorf("portsbyaddr should have '192.0.2.1' as key. value should be 80: %v", ports)
	}
	if ports, ok := portsbyaddr["192.0.2.2"]; !ok || ports[0] != 443 {
		t.Errorf("portsbyaddr should have '192.0.2.2' as key. value should be 443: %v", ports)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFindSourceByDestAddrAndPort(t *testing.T) {
	db, mock := NewTestDB()
	defer db.Close()

	addr := net.ParseIP("192.0.10.1")
	port := int16(8000)

	columns := sqlmock.NewRows([]string{"connections", "updated", "source_ipv4", "source_port"})
	mock.ExpectQuery("SELECT (.+) FROM flows").WithArgs(addr.String(), port).WillReturnRows(
		columns.AddRow(10, time.Now(), "192.0.10.2", 0),
	)

	addrports, err := db.FindSourceByDestAddrAndPort(addr, port)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(addrports) != 1 {
		t.Errorf("addrports should be 1, but %v", len(addrports))
	}

	if addrports[0].IPAddr.String() != "192.0.10.2" {
		t.Errorf("addrports[0].Addr should be '192.0.10.2', but %s", addrports[0].IPAddr)
	}
	if addrports[0].Port != 0 {
		t.Errorf("addrports[0].Port should be '8000', but %v", addrports[0].Port)
	}
	if addrports[0].Connections != 10 {
		t.Errorf("addrports[0].Connections should be '10', but %v", addrports[0].Connections)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
