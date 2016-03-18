// Copyright 2015 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package interfaces

import ()

// Holds the state information for factomd.  This does imply that we will be
// using accessors to access state information in the consensus algorithm.
// This is a bit tedious, but does provide single choke points where information
// can be logged about the execution of Factom.  Also ensures that we do not
// accidentally
type IState interface {

	// Server
	GetFactomNodeName() string
	Clone(number string) IState
	GetCfg() IFactomConfig
	LoadConfig(filename string)
	Init()
	String() string
	GetCoreChainID() IHash
	GetIdentityChainID() IHash
	Sign([]byte) IFullSignature
	GetDirectoryBlockInSeconds() int
	GetServer() IServer
	SetServer(IServer)
	GetDBHeightComplete() uint32
	SetOut(bool)  // Output is turned on if set to true
	GetOut() bool // Return true if Print or Println write output
	LoadDBState(dbheight uint32) (IMsg, error)
	GetFedServerIndexFor(uint32, IHash) (bool, int)
	GetFedServerIndex(uint32) (bool, int)
	SetString()
	ShortString() string
	
	// This is the highest block signed off and recorded in the Database.
	GetHighestRecordedBlock() uint32
	// This is the block the leader is building
	GetLeaderHeight() uint32
	// The highest block for which we have received a message.  Sometimes the same as
	// BuildingBlock(), but can be different depending or the order messages are recieved.
	GetHighestKnownBlock() uint32

	// Find a Directory Block by height
	GetDirectoryBlockByHeight(dbheight uint32) IDirectoryBlock
	// Channels
	//==========

	// Network Processor
	NetworkOutMsgQueue() chan IMsg
	NetworkInvalidMsgQueue() chan IMsg

	// Consensus
	InMsgQueue() chan IMsg // Read by Validate

	// Lists and Maps
	// =====
	GetAuditHeartBeats() []IMsg   // The checklist of HeartBeats for this period
	GetFedServerFaults() [][]IMsg // Keep a fault list for every server

	GetNewEBlocks(dbheight uint32, hash IHash) IEntryBlock
	PutNewEBlocks(dbheight uint32, hash IHash, eb IEntryBlock)

	GetCommits(hash IHash) IMsg
	PutCommits(hash IHash, msg IMsg)
	// Server Configuration
	// ====================

	//Network MAIN = 0, TEST = 1, LOCAL = 2, CUSTOM = 3
	GetNetworkNumber() int  // Encoded into Directory Blocks
	GetNetworkName() string // Some networks have defined names

	GetMatryoshka(dbheight uint32) IHash // Reverse Hash

	LeaderFor(hash []byte) bool // Tests if this server is the leader for this key

	// Database
	// ========
	GetDB() DBOverlay
	SetDB(DBOverlay)

	GetAnchor() IAnchor

	// Web Services
	// ============
	SetPort(int)
	GetPort() int

	// Factoid State
	// =============
	UpdateState()
	GetFactoidState() IFactoidState

	SetFactoidState(dbheight uint32, fs IFactoidState)
	GetFactoshisPerEC() uint64
	SetFactoshisPerEC(factoshisPerEC uint64)
	// MISC
	// ====

	FollowerExecuteMsg(m IMsg) (bool, error) // Messages that go into the process list
	FollowerExecuteAck(m IMsg) (bool, error) // Ack Msg calls this function.
	FollowerExecuteDBState(IMsg) error       // Add the given DBState to this server
	ProcessAddServer(dbheight uint32, addServerMsg IMsg)
	ProcessCommitChain(dbheight uint32, commitChain IMsg)
	ProcessDBS(dbheight uint32, commitChain IMsg)
	ProcessEOM(dbheight uint32, eom IMsg)

	// For messages that go into the Process List
	LeaderExecute(m IMsg) error
	LeaderExecuteAddServer(m IMsg) error
	LeaderExecuteEOM(m IMsg) error
	LeaderExecuteDBSig(m IMsg) error

	NewEOM(int) IMsg

	GetTimestamp() Timestamp

	PrintType(int) bool // Debugging
	Print(a ...interface{}) (n int, err error) 
	Println(a ...interface{}) (n int, err error) 
}
