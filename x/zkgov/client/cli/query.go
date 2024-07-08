package cli

// import (
// 	"fmt"
// 	"strconv"

// 	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/types"

// 	"github.com/cosmos/cosmos-sdk/client"
// 	"github.com/cosmos/cosmos-sdk/client/flags"
// 	"github.com/spf13/cobra"
// )

// func GetQueryCmd() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   types.ModuleName,
// 		Short: "Querying commands for the zkgov module",
// 		RunE:  client.ValidateCmd,
// 	}

// 	cmd.AddCommand(
// 		GetUser(),
// 	)

// 	return cmd
// }

// func GetUser() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "get",
// 		Short: "List all the leaves",
// 		Long: `List all the leaves which are accepted or rejected by the admin,
// 		`,
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			clientCtx, err := client.GetClientQueryContext(cmd)
// 			if err != nil {
// 				return err
// 			}
// 			queryClient := types.NewQueryClient(clientCtx)
// 			userid, _ := strconv.Atoi(args[0])
// 			fmt.Println("user id>>>>>>>>>>>", userid)
// 			params := &types.QueryUserRequset{Userid: uint64(userid)}
// 			res, _ := queryClient.GetUser(cmd.Context(), params)
// 			return clientCtx.PrintProto(res)
// 		},
// 	}
// 	flags.AddQueryFlagsToCmd(cmd)
// 	return cmd
// }
