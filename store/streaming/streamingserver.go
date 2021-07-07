package streaming

import (
	"github.com/cosmos/cosmos-sdk/codec"
	listen "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"os"
	"sync"
)

// Hook interface used to hook into the ABCI message processing of the BaseApp
type Hook interface {
	ListenBeginBlock(ctx sdk.Context, req abci.RequestBeginBlock, res abci.ResponseBeginBlock) // update the streaming service with the latest BeginBlock messages
	ListenEndBlock(ctx sdk.Context, req abci.RequestEndBlock, res abci.ResponseEndBlock)       // update the steaming service with the latest EndBlock messages
	ListenDeliverTx(ctx sdk.Context, req abci.RequestDeliverTx, res abci.ResponseDeliverTx)    // update the steaming service with the latest DeliverTx messages
}

// StreamingService interface for registering WriteListeners with the BaseApp and updating the service with the ABCI messages using the hooks
type StreamingService interface {
	Stream(wg *sync.WaitGroup, quitChan <-chan struct{}) // streaming service loop, awaits kv pairs and writes them to some destination stream or file
	Listeners() map[sdk.StoreKey]listen.WriteListener    // returns the streaming service's listeners for the BaseApp to register
	Hook
}

// intermediateWriter is used so that we do not need to update the underlying io.Writer inside the StoreKVPairWriteListener
// everytime we begin writing to a new file
type intermediateWriter struct {
	outChan chan<- []byte
}

// NewIntermediateWriter create an instance of an intermediateWriter that sends to the provided channel
func NewIntermediateWriter(outChan chan<- []byte) *intermediateWriter {
	return &intermediateWriter{
		outChan: outChan,
	}
}

// Write satisfies io.Writer
func (iw *intermediateWriter) Write(b []byte) (int, error) {
	iw.outChan <- b
	return len(b), nil
}

// FileStreamingService is a concrete implementation of StreamingService that writes state changes out to a file
type FileStreamingService struct {
	listeners  map[sdk.StoreKey]listen.WriteListener // the listeners that will be initialized with BaseApp
	srcChan    <-chan []byte                         // the channel that all of the WriteListeners write their data out to
	filePrefix string                                // optional prefix for each of the generated files
	writeDir   string                                // directory to write files into
	dstFile    *os.File                              // the current write output file
	marshaller codec.BinaryCodec                     // marshaller used for re-marshalling the ABCI messages to write them out to the destination files
	stateCache [][]byte                              // cache the protobuf binary encoded StoreKVPairs in the order they are received
}

func (fss *FileStreamingService) Stream(wg *sync.WaitGroup, quitChan <-chan struct{}) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-quitChan:
				return
			case by := <-fss.srcChan:
				fss.stateCache = append(fss.stateCache, by)
			}
		}
	}()
}

func (fss *FileStreamingService) Listeners() map[sdk.StoreKey]listen.WriteListener {
	return fss.listeners
}

func (fss *FileStreamingService) ListenBeginBlock(ctx sdk.Context, req abci.RequestBeginBlock, res abci.ResponseBeginBlock) {
	panic("implement me")
	// NOTE: this could either be done synchronously or asynchronously
	// create a new file with the req info according to naming schema
	// write req to file
	// write all state changes cached for this stage to file
	// reset cache
	// write res to file
	// close file
}

func (fss *FileStreamingService) ListenEndBlock(ctx sdk.Context, req abci.RequestEndBlock, res abci.ResponseEndBlock) {
	panic("implement me")
	// NOTE: this could either be done synchronously or asynchronously
	// create a new file with the req info according to naming schema
	// write req to file
	// write all state changes cached for this stage to file
	// reset cache
	// write res to file
	// close file
}

func (fss *FileStreamingService) ListenDeliverTx(ctx sdk.Context, req abci.RequestDeliverTx, res abci.ResponseDeliverTx) {
	panic("implement me")
	// NOTE: this could either be done synchronously or asynchronously
	// create a new file with the req info according to naming schema
	// NOTE: if the tx failed, handle accordingly
	// write req to file
	// write all state changes cached for this stage to file
	// reset cache
	// write res to file
	// close file
}

// NewFileStreamingService creates a new FileStreamingService for the provided writeDir, (optional) filePrefix, and storeKeys
func NewFileStreamingService(writeDir, filePrefix string, storeKeys []sdk.StoreKey) StreamingService {
	var m codec.BinaryCodec
	listenChan := make(chan []byte, 0)
	iw := NewIntermediateWriter(listenChan)
	listener := listen.NewStoreKVPairWriteListener(iw, m)
	listeners := make(map[sdk.StoreKey]listen.WriteListener, len(storeKeys))
	// in this case, we are using the same listener for each Store
	for _, key := range storeKeys {
		listeners[key] = listener
	}
	//// check that the writeDir exists and is writeable so that we can catch the error here at initialization if it is not
	//// we don't open a dstFile until we receive our first ABCI message
	//if err := fileutil.IsDirWriteable(writeDir); err != nil {
	//	return nil
	//}
	return &FileStreamingService{
		listeners:  listeners,
		srcChan:    listenChan,
		filePrefix: filePrefix,
		writeDir:   writeDir,
		marshaller: m,
		stateCache: make([][]byte, 0),
	}
}
