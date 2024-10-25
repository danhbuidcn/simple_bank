package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	mockdb "simple_bank/db/mock"
	db "simple_bank/db/sqlc"
	"simple_bank/util"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestListAccountsAPI(t *testing.T) {
	accounts := []db.Account{
		randomAccount(),
		randomAccount(),
		randomAccount(),
	}

	testCases := []struct {
		name          string                                                  // name of the test case
		query         ListAccountsRequest                                     // input query parameters
		buildStubs    func(store *mockdb.MockStore)                           // build stubs
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder) // check the response
	}{
		{
			name: "OK",
			query: ListAccountsRequest{
				PageID:   1,
				PageSize: 5,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListAccounts(gomock.Any(), db.ListAccountsParams{
						Limit:  5,
						Offset: 0,
					}).
					Times(1).
					Return(accounts, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccounts(t, recorder.Body, accounts)
			},
		},
		{
			name: "No Data",
			query: ListAccountsRequest{
				PageID:   1,
				PageSize: 5,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListAccounts(gomock.Any(), db.ListAccountsParams{
						Limit:  5,
						Offset: 0,
					}).
					Times(1).
					Return(nil, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccounts(t, recorder.Body, []db.Account{})
			},
		},
		{
			name: "Internal Error",
			query: ListAccountsRequest{
				PageID:   1,
				PageSize: 5,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListAccounts(gomock.Any(), db.ListAccountsParams{
						Limit:  5,
						Offset: 0,
					}).
					Times(1).
					Return(nil, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "Invalid PageID",
			query: ListAccountsRequest{
				PageID:   0,
				PageSize: 5,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListAccounts(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Invalid PageSize",
			query: ListAccountsRequest{
				PageID:   1,
				PageSize: 11,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListAccounts(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "PageSize Less Than Minimum",
			query: ListAccountsRequest{
				PageID:   1,
				PageSize: 4,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListAccounts(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create a mock store
			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			// Create a new server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			// Create a GET request with query parameters
			url := fmt.Sprintf("/accounts?page_id=%d&page_size=%d", tc.query.PageID, tc.query.PageSize)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			// Send the request
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()

	testCases := []struct {
		name          string                                                  // name of the test case
		accountID     int64                                                   // input
		buildStubs    func(store *mockdb.MockStore)                           // build stubs
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder) // check the response
	}{
		{
			name:      "OK",
			accountID: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(account, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccount(t, recorder.Body, account)
			},
		},
		{
			name:      "Not Found",
			accountID: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(db.Account{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:      "Internal Error",
			accountID: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(db.Account{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:      "Invalid ID",
			accountID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create a mock store
			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			// Create a new server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			// Create a GET request
			url := fmt.Sprintf("/accounts/%d", tc.accountID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			// Send the request
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestCreateAccountAPI(t *testing.T) {
	account := randomAccount()

	testCases := []struct {
		name          string                                                  // name of the test case
		payload       CreateAccountRequest                                    // input
		buildStubs    func(store *mockdb.MockStore)                           // build stubs
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder) // check the response
	}{
		{
			name: "OK",
			payload: CreateAccountRequest{
				Owner:    account.Owner,
				Currency: account.Currency,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateAccount(gomock.Any(), gomock.Eq(db.CreateAccountParams{
						Owner:    account.Owner,
						Balance:  0,
						Currency: account.Currency,
					})).
					Times(1).
					Return(account, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccount(t, recorder.Body, account)
			},
		},
		{
			name: "Internal Error",
			payload: CreateAccountRequest{
				Owner:    account.Owner,
				Currency: account.Currency,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateAccount(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Account{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "Invalid Owner",
			payload: CreateAccountRequest{
				Owner:    "",
				Currency: account.Currency,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Invalid Currency",
			payload: CreateAccountRequest{
				Owner:    account.Owner,
				Currency: "XYZ",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create a mock store
			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			// Create a new server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			// Create a POST request with JSON body
			data, err := json.Marshal(tc.payload)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, "/accounts", bytes.NewReader(data))
			require.NoError(t, err)

			// Send the request
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func requireBodyMatchTransfer(t *testing.T, body *bytes.Buffer, transfer db.Transfer) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotTransfer db.Transfer
	err = json.Unmarshal(data, &gotTransfer)
	require.NoError(t, err)
	require.Equal(t, transfer, gotTransfer)
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotAccount db.Account
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)
	require.Equal(t, account, gotAccount)
}

func requireBodyMatchAccounts(t *testing.T, body *bytes.Buffer, accounts []db.Account) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotAccounts []db.Account
	err = json.Unmarshal(data, &gotAccounts)
	require.NoError(t, err)
	require.Equal(t, accounts, gotAccounts)
}

func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}
