package handler

import (
    "context"

    "github.com/golang/protobuf/ptypes/empty"
    "github.com/xmlking/logger/log"

    myErrors "github.com/xmlking/micro-starter-kit/shared/errors"
    transactionPB "github.com/xmlking/micro-starter-kit/srv/recorder/proto/transaction"
    "github.com/xmlking/micro-starter-kit/srv/recorder/repository"
)

type recorderHandler struct {
    repo repository.TransactionRepository
}

func NewTransactionHandler(repo repository.TransactionRepository) transactionPB.TransactionServiceHandler {
    return &recorderHandler{
        repo: repo,
    }
}

func (h *recorderHandler) Read(ctx context.Context, req *transactionPB.ReadRequest, rsp *transactionPB.TransactionEvent) error {
    if rsp, err := h.repo.Read(ctx, req.GetKey()); err != nil {
        log.Errorw("Received transactionService.Read request error", err)
        return myErrors.AppError(myErrors.DBE, err)
    } else {
        log.Infof("Got transactionService responce %s", rsp.GetReq())
    }
    return nil
}
func (h *recorderHandler) Write(ctx context.Context, req *transactionPB.WriteRequest, rsp *empty.Empty) (err error) {
    if err := h.repo.Write(ctx, req.GetKey(), req.GetEvent()); err != nil {
        log.Errorw("Received TransactionHandler.Write request error", err)
        return myErrors.AppError(myErrors.DBE, err)
    }
    return nil
}
