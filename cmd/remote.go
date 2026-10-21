package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/xnet-admin-1/Headhunter/internal/daemon"
)

// Remote connection flags for `headhunter remote`.
var (
	remotePair     string
	remoteHost     string
	remoteList     bool
	remoteForget   string
	remotePairCode string // hidden: fixed code for tests
)

var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Connect to a paired headhunter daemon and run a session on it",
	Long: `Connect to a headhunter daemon and run a session on the remote host.

First-time use pairs this machine with the host:

  headhunter remote --pair A1B2C3D4     # code shown by 'headhunter daemon pair' on the host

After pairing, reconnect by the name you saved — no code needed:

  headhunter remote --host homelab

The session runs entirely on the host; this terminal just renders it.
Ctrl-] d detaches; /quit ends the session.

'headhunter session attach --host <name>' is the same thing with session switching:
Ctrl-] s picks another session, Ctrl-] w lists every paired host.`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		ctx := cmd.Context()
		switch {
		case remoteList:
			hosts, err := daemon.ListHosts()
			if err != nil {
				return err
			}
			if len(hosts) == 0 {
				fmt.Println("No paired hosts. Pair one with: headhunter remote --pair <code>")
				return nil
			}
			fmt.Printf("%-16s %-12s %s\n", "NAME", "PAIRED", "LAST USED")
			for _, h := range hosts {
				fmt.Printf("%-16s %-12s %s\n",
					h.Name,
					h.AddedAt.Format("2006-01-02"),
					h.LastUsed.Format("2006-01-02 15:04"))
			}
			return nil
		case remoteForget != "":
			return daemon.ForgetHost(remoteForget)
		case remotePair != "":
			code := remotePair
			if remotePairCode != "" {
				code = remotePairCode // hidden testing override
			}
			return daemon.RunPair(ctx, daemon.PairOptions{Code: code, Name: remoteHost})
		case remoteHost != "":
			_, _ = fmt.Fprintln(cmd.ErrOrStderr(), "Deprecated: use 'headhunter session attach --host' to attach to remote sessions.")
			return daemon.RunHost(ctx, remoteHost, daemon.AttachOptions{
				Pick: localPicker,
			})
		default:
			return cmd.Help()
		}
	},
}

func init() {
	remoteCmd.Flags().StringVar(&remotePair, "pair", "", "pair with a host using the code from 'headhunter daemon pair'")
	remoteCmd.Flags().StringVar(&remoteHost, "host", "", "connect to a paired host by saved name")
	remoteCmd.Flags().BoolVar(&remoteList, "list", false, "list paired hosts")
	remoteCmd.Flags().StringVar(&remoteForget, "forget", "", "forget a paired host by name")
	remoteCmd.Flags().StringVar(&remotePairCode, "code", "", "use a fixed pairing code (testing)")
	_ = remoteCmd.Flags().MarkHidden("code")
	rootCmd.AddCommand(remoteCmd)
}
