package keeper_test

import (
	"testing"

	addresstypes "cosmossdk.io/core/address"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	cmttime "github.com/cometbft/cometbft/types/time"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/stretchr/testify/suite"
	zkgov "github.com/vishal-kanna/zk/zk-gov/x/zkgov"
	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/keeper"
	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/types"
)

type TestSuite struct {
	suite.Suite

	ctx          sdk.Context
	addrs        []sdk.AccAddress
	baseApp      *baseapp.BaseApp
	encCfg       moduletestutil.TestEncodingConfig
	queryClient  types.QueryClient
	msgSrvr      types.MsgServer
	keeper       keeper.Keeper
	addressCodec addresstypes.Codec

	addresses []string
}

func (s *TestSuite) SetupTest() {
	key := storetypes.NewKVStoreKey(types.ModuleName)
	storeService := runtime.NewKVStoreService(key)
	testCtx := testutil.DefaultContextWithDB(s.T(), key, storetypes.NewTransientStoreKey("transient_test"))
	s.ctx = testCtx.Ctx.WithBlockHeader(cmtproto.Header{Time: cmttime.Now()})
	s.encCfg = moduletestutil.MakeTestEncodingConfig(zkgov.AppModuleBasic{})

	s.addressCodec = address.NewBech32Codec("cosmos")
	// ctrl := gomock.NewController(s.T())
	s.baseApp = baseapp.NewBaseApp(
		"zk-gov",
		log.NewNopLogger(),
		testCtx.DB,
		s.encCfg.TxConfig.TxDecoder(),
	)
	s.baseApp.SetCMS(testCtx.CMS)
	s.baseApp.SetInterfaceRegistry(s.encCfg.InterfaceRegistry)

	s.addrs = simtestutil.CreateIncrementalAccounts(7)

	s.keeper = keeper.NewKeeper(s.encCfg.Codec, storeService)

	queryHelper := baseapp.NewQueryServerTestHelper(s.ctx, s.encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, s.keeper)
	queryClient := types.NewQueryClient(queryHelper)
	s.queryClient = queryClient

	s.msgSrvr = keeper.NewMsgServerImpl(s.keeper)

	alice, err := s.addressCodec.BytesToString([]byte("alice"))
	s.Require().NoError(err)

	bob, err := s.addressCodec.BytesToString([]byte("bob"))
	s.Require().NoError(err)

	_, err = s.addressCodec.StringToBytes(bob)
	s.Require().NoError(err)

	charlie, err := s.addressCodec.BytesToString([]byte("charlie"))
	s.Require().NoError(err)

	s.addresses = []string{alice, bob, charlie}
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

// func (s *TestSuite) TestGetUserInfo() {
// 	_, err := s.addressCodec.BytesToString([]byte("alice"))
// 	s.Require().NoError(err)

// 	// 	_, err = s.msgSrvr.RegisterUser(s.ctx, &types.RegisterUserRequest{
// 	// 	// 	Sender: alice,
// 	// 	// })
// 	// 	// s.Require().NoError(err)

// 	//		// res, err := s.queryClient.GetUser(s.ctx, &types.QueryUserRequset{Userid: 1})
// 	//		// fmt.Println("the result is>>>>>>>>>>", res)
// 	//		// s.Require().Equal(res.Ust.Userid, 1)
// 	//		// s.Require().NoError(err)
// 	//	}
// }
