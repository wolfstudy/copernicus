package blockchain

import (
	"encoding/binary"
	"io"

	"github.com/btcboost/copernicus/utils"
)

type BlockFileInfo struct {
	Blocks      uint32
	Size        uint32
	UndoSize    uint32
	HeightFirst uint32
	HeightLast  uint32
	timeFirst   uint64
	timeLast    uint64
}

func (blockFileInfo *BlockFileInfo) Serialize(writer io.Writer) error {
	err := utils.BinarySerializer.PutUint32(writer, binary.LittleEndian, blockFileInfo.Blocks)
	if err != nil {
		return err
	}
	err = utils.BinarySerializer.PutUint32(writer, binary.LittleEndian, blockFileInfo.Size)
	if err != nil {
		return err
	}
	err = utils.BinarySerializer.PutUint32(writer, binary.LittleEndian, blockFileInfo.UndoSize)
	if err != nil {
		return err
	}
	err = utils.BinarySerializer.PutUint64(writer, binary.LittleEndian, blockFileInfo.timeFirst)
	if err != nil {
		return err
	}
	err = utils.BinarySerializer.PutUint32(writer, binary.LittleEndian, blockFileInfo.HeightFirst)
	if err != nil {
		return err
	}
	err = utils.BinarySerializer.PutUint32(writer, binary.LittleEndian, blockFileInfo.HeightLast)
	if err != nil {
		return err
	}
	return utils.BinarySerializer.PutUint64(writer, binary.LittleEndian, blockFileInfo.timeLast)
}

func DeserializeBlockFileInfo(reader io.Reader) (*BlockFileInfo, error) {
	blockFileInfo := new(BlockFileInfo)
	blocks, err := utils.BinarySerializer.Uint32(reader, binary.LittleEndian)
	if err != nil {
		return nil, err
	}
	size, err := utils.BinarySerializer.Uint32(reader, binary.LittleEndian)
	if err != nil {
		return nil, err
	}
	undoSize, err := utils.BinarySerializer.Uint32(reader, binary.LittleEndian)
	if err != nil {
		return nil, err
	}
	heightFirst, err := utils.BinarySerializer.Uint32(reader, binary.LittleEndian)
	if err != nil {
		return nil, err
	}
	heightLast, err := utils.BinarySerializer.Uint32(reader, binary.LittleEndian)
	if err != nil {
		return nil, err
	}

	timeFirst, err := utils.BinarySerializer.Uint64(reader, binary.LittleEndian)
	if err != nil {
		return nil, err
	}
	timeLast, err := utils.BinarySerializer.Uint64(reader, binary.LittleEndian)
	if err != nil {
		return nil, err
	}
	blockFileInfo.Blocks = blocks
	blockFileInfo.Size = size
	blockFileInfo.UndoSize = undoSize
	blockFileInfo.HeightFirst = heightFirst
	blockFileInfo.HeightLast = heightLast
	blockFileInfo.timeFirst = timeFirst
	blockFileInfo.timeLast = timeLast
	return blockFileInfo, nil

}
