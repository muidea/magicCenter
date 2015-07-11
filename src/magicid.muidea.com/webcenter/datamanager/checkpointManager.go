package datamanager

import (
	"fmt"
	"log"
	"muidea.com/dao"
)

type CheckpointManager struct {
	checkpointInfo map[int]Checkpoint
	dao           *dao.Dao
}

func (this *CheckpointManager) Load() bool {
	dao, err := dao.Fetch("root", "rootkit", "localhost:3306", "magicid_db")
	if err != nil {
		log.Printf("fetch dao failed, err:%s", err.Error())
		return false
	}

	this.checkpointInfo = make(map[int]Checkpoint)
	this.dao = dao
	return true
}

func (this *CheckpointManager) Unload() {
	this.dao.Release()
	this.dao = nil
	this.checkpointInfo = nil
}

func (this *CheckpointManager) AddCheckpoint(checkpoint Checkpoint) bool {
	sql := fmt.Sprintf("insert into magicid_db.checkpoint value (%d, %s, %d)", checkpoint.Id, checkpoint.Name, checkpoint.Rid)
	if !this.dao.Execute(sql) {
		log.Printf("execute failed, sql:%s", sql)
		return false
	}

	return true
}

func (this *CheckpointManager) ModCheckpoint(checkpoint Checkpoint) bool {
	sql := fmt.Sprintf("update magicid_db.checkpoint set name ='%s', rid=%d where id =%d", checkpoint.Name, checkpoint.Rid, checkpoint.Id)
	if !this.dao.Execute(sql) {
		log.Printf("execute failed, sql:%s", sql)
		return false
	}

	this.checkpointInfo[checkpoint.Id] = checkpoint
	return true
}

func (this *CheckpointManager) DelCheckpoint(id int) {
	delete(this.checkpointInfo, id)

	sql := fmt.Sprintf("delete from magicid_db.checkpoint where id =%d", id)
	if !this.dao.Execute(sql) {
		log.Printf("execute failed, sql:%s", sql)
		return
	}
}

func (this *CheckpointManager) FindCheckpointById(id int) (Checkpoint, bool) {
	checkpoint, found := this.checkpointInfo[id]
	if !found {
		sql := fmt.Sprintf("select * from magicid_db.checkpoint where id=%d", id)
		if !this.dao.Query(sql) {
			log.Printf("query failed, sql:%s", sql)
			return checkpoint, false
		}

		for this.dao.Next() {
			checkpoint := Checkpoint{}
			this.dao.GetField(&checkpoint.Id, &checkpoint.Name, &checkpoint.Rid)
			this.checkpointInfo[checkpoint.Id] = checkpoint
		}
	}
	checkpoint, found = this.checkpointInfo[id]

	return checkpoint, found
}

func (this *CheckpointManager) FindCheckpointByRId(id int) ([]Checkpoint, bool) {
	checkpoints := []Checkpoint{}
	sql := fmt.Sprintf("select * from magicid_db.checkpoint where rid=%d", id)
	if !this.dao.Query(sql) {
		log.Printf("query failed, sql:%s", sql)
		return checkpoints, false
	}

	for this.dao.Next() {
		checkpoint := Checkpoint{}
		this.dao.GetField(&checkpoint.Id, &checkpoint.Name, &checkpoint.Rid)
		this.checkpointInfo[checkpoint.Id] = checkpoint
		checkpoints = append(checkpoints, checkpoint)
	}
	
	return checkpoints, true
}

