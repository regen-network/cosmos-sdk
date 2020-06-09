package rest

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/cosmos/cosmos-sdk/x/msg_authorization/types"
)

func registerTxRoutes(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc("/msg_authorization/grant", grantHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/msg_authorization/revoke", revokeHandler(clientCtx)).Methods("POST")
}

type GrantRequest struct {
	BaseReq       rest.BaseReq         `json:"base_req" yaml:"base_req"`
	Granter       sdk.AccAddress       `json:"granter"`
	Grantee       sdk.AccAddress       `json:"grantee"`
	Authorization types.AuthorizationI `json:"authorization"`
	Expiration    time.Time            `json:"expiration"`
}

type RevokeRequest struct {
	BaseReq              rest.BaseReq   `json:"base_req" yaml:"base_req"`
	Granter              sdk.AccAddress `json:"granter"`
	Grantee              sdk.AccAddress `json:"grantee"`
	AuthorizationMsgType string         `json:"authorization_msg_type"`
}

func grantHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req GrantRequest

		if !rest.ReadRESTReq(w, r, clientCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg, err := types.NewMsgGrantAuthorization(req.Granter, req.Grantee, req.Authorization, req.Expiration)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		authclient.WriteGenerateStdTxResponse(w, clientCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func revokeHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RevokeRequest

		if !rest.ReadRESTReq(w, r, clientCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgRevokeAuthorization(req.Granter, req.Grantee, req.AuthorizationMsgType)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		authclient.WriteGenerateStdTxResponse(w, clientCtx, req.BaseReq, []sdk.Msg{msg})
	}
}
